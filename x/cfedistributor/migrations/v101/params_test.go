package v101_test

import (
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	v100cfedistributor "github.com/chain4energy/c4e-chain/x/cfedistributor/migrations/v100"
	v101cfedistributor "github.com/chain4energy/c4e-chain/x/cfedistributor/migrations/v101"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/stretchr/testify/require"

	"testing"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMigrationSubDistributorsCorrectOrder(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	oldSubDistributors := []v100cfedistributor.SubDistributor{
		createOldSubDistributor(types.INTERNAL_ACCOUNT, types.BASE_ACCOUNT, types.MODULE_ACCOUNT, "CUSTOM_ID"),
		createOldSubDistributor(types.BASE_ACCOUNT, types.INTERNAL_ACCOUNT, types.MODULE_ACCOUNT, "CUSTOM_ID"),
		createOldSubDistributor(types.BASE_ACCOUNT, types.MAIN, types.MODULE_ACCOUNT, "CUSTOM_ID"),
	}
	setOldSubdistributors(t, ctx, testUtil, oldSubDistributors)
	MigrateParamsV100ToV101(t, ctx, testUtil, false)
}

func TestMigrationSubDistributorsWrongOrder(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	oldSubDistributors := []v100cfedistributor.SubDistributor{
		createOldSubDistributor(types.INTERNAL_ACCOUNT, types.BASE_ACCOUNT, types.MODULE_ACCOUNT, "CUSTOM_ID"),
		createOldSubDistributor(types.BASE_ACCOUNT, types.MAIN, types.MODULE_ACCOUNT, "CUSTOM_ID"),
		createOldSubDistributor(types.BASE_ACCOUNT, types.MAIN, types.MODULE_ACCOUNT, "CUSTOM_ID"),
	}

	setOldSubdistributors(t, ctx, testUtil, oldSubDistributors)
	MigrateParamsV100ToV101(t, ctx, testUtil, true)
}

func TestMigrationSubDistributorsDuplicates(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	oldSubDistributors := []v100cfedistributor.SubDistributor{
		createOldSubDistributor(types.INTERNAL_ACCOUNT, types.BASE_ACCOUNT, types.MODULE_ACCOUNT, "CUSTOM_ID"),
		createOldSubDistributor(types.BASE_ACCOUNT, types.INTERNAL_ACCOUNT, types.MODULE_ACCOUNT, "CUSTOM_ID"),
		createOldSubDistributor(types.BASE_ACCOUNT, types.MAIN, types.BASE_ACCOUNT, "CUSTOM_ID"),
	}
	setOldSubdistributors(t, ctx, testUtil, oldSubDistributors)
	MigrateParamsV100ToV101(t, ctx, testUtil, true)
}

func TestMigrationSubDistributorsWrongAccType(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	oldSubDistributors := []v100cfedistributor.SubDistributor{
		createOldSubDistributor(types.INTERNAL_ACCOUNT, types.BASE_ACCOUNT, types.MODULE_ACCOUNT, "CUSTOM_ID"),
		createOldSubDistributor(types.BASE_ACCOUNT, types.MAIN, types.MODULE_ACCOUNT, "CUSTOM_ID"),
		createOldSubDistributor(types.BASE_ACCOUNT, "WRONG_ACCOUNT_TYPE", types.MODULE_ACCOUNT, "CUSTOM_ID"),
	}

	setOldSubdistributors(t, ctx, testUtil, oldSubDistributors)
	MigrateParamsV100ToV101(t, ctx, testUtil, true)
}

func setOldSubdistributors(t *testing.T, ctx sdk.Context, testUtil *testkeeper.ExtendedC4eDistributorKeeperUtils, subdistributors []v100cfedistributor.SubDistributor) {
	store := newStore(ctx, testUtil)
	bz, err := codec.NewLegacyAmino().MarshalJSON(subdistributors)
	require.NoError(t, err)
	store.Set(types.KeySubDistributors, bz)
}

func newStore(ctx sdk.Context, testUtil *testkeeper.ExtendedC4eDistributorKeeperUtils) prefix.Store {
	return prefix.NewStore(ctx.KVStore(testUtil.StoreKey), append([]byte((testUtil.Subspace.Name())), '/'))
}

func MigrateParamsV100ToV101(
	t *testing.T,
	ctx sdk.Context,
	testUtil *testkeeper.ExtendedC4eDistributorKeeperUtils,
	wantError bool,
) {
	var res []v100cfedistributor.SubDistributor
	store := newStore(ctx, testUtil)
	distributors := store.Get(types.KeySubDistributors)
	err := codec.NewLegacyAmino().UnmarshalJSON(distributors, &res)
	require.NoError(t, err)

	err = v101cfedistributor.MigrateParams(ctx, testUtil.StoreKey, &testUtil.Subspace)
	if wantError {
		require.Error(t, err)
		return
	}

	require.NoError(t, err)
	newParams := testUtil.GetC4eDistributorKeeper().GetParams(ctx)
	newSubdistributors := newParams.SubDistributors

	for i, oldSubdistributor := range res {
		require.EqualValues(t,
			newSubdistributors[i].Name, oldSubdistributor.Name)
		require.EqualValues(t,
			newSubdistributors[i].Destinations.BurnShare, oldSubdistributor.Destination.BurnShare.Percent.Quo(sdk.NewDec(100)))
		require.EqualValues(t,
			newSubdistributors[i].Destinations.PrimaryShare.Id, oldSubdistributor.Destination.Account.Id)
		require.EqualValues(t,
			newSubdistributors[i].Destinations.PrimaryShare.Type, oldSubdistributor.Destination.Account.Type)

		for j, oldShare := range oldSubdistributor.Destination.Share {
			require.EqualValues(t,
				newSubdistributors[i].Destinations.Shares[j].Share, oldShare.Percent.Quo(sdk.NewDec(100)))
			require.EqualValues(t,
				newSubdistributors[i].Destinations.Shares[j].Name, oldShare.Name)
			require.EqualValues(t,
				newSubdistributors[i].Destinations.Shares[j].Destination.Id, oldShare.Account.Id)
			require.EqualValues(t,
				newSubdistributors[i].Destinations.Shares[j].Destination.Type, oldShare.Account.Type)
		}

		for j, oldSource := range oldSubdistributor.Sources {
			require.EqualValues(t,
				newSubdistributors[i].Sources[j].Id, oldSource.Id)
			require.EqualValues(t,
				newSubdistributors[i].Sources[j].Type, oldSource.Type)
		}
	}
}

func createOldSubDistributor(
	destinationType string,
	sourceType string,
	destinationShareType string,
	Id string,
) v100cfedistributor.SubDistributor {
	return v100cfedistributor.SubDistributor{
		Name: helpers.RandStringOfLength(10),
		Destination: v100cfedistributor.Destination{
			Account: v100cfedistributor.Account{
				Id:   Id,
				Type: destinationType,
			},
			BurnShare: &v100cfedistributor.BurnShare{
				Percent: sdk.NewDec(25),
			},
			Share: []*v100cfedistributor.Share{
				{
					Name: helpers.RandStringOfLength(10),
					Account: v100cfedistributor.Account{
						Id:   Id,
						Type: destinationShareType,
					},
					Percent: sdk.NewDec(50),
				},
			},
		},
		Sources: []*v100cfedistributor.Account{
			{
				Id:   Id,
				Type: sourceType,
			},
		},
	}
}
