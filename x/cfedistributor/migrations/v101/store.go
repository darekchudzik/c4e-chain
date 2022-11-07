package v101

import (
	v100cfedistributor "github.com/chain4energy/c4e-chain/x/cfedistributor/migrations/v100"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

func UpdateParams(ctx sdk.Context, paramStore *paramtypes.Subspace) error {
	var res []v100cfedistributor.SubDistributor
	rawDistributors := paramStore.GetRaw(ctx, types.KeySubDistributors)
	if err := codec.NewLegacyAmino().UnmarshalJSON(rawDistributors, &res); err != nil {
		panic(err)
	}

	var newSubdistributors []types.SubDistributor

	for _, oldSubdistributor := range res {

		var newShares []*types.DestinationShare
		for _, oldShare := range oldSubdistributor.Destination.Share {
			newShare := types.DestinationShare{
				Share: oldShare.Percent.Quo(sdk.NewDec(100)),
				Destination: types.Account{
					Id:   oldShare.Account.Id,
					Type: oldShare.Account.Type,
				},
				Name: oldShare.Name,
			}
			newShares = append(newShares, &newShare)
		}

		newSubdistributor := types.SubDistributor{
			Name: oldSubdistributor.Name,
			Destinations: types.Destinations{
				Shares:    newShares,
				BurnShare: oldSubdistributor.Destination.BurnShare.Percent.Quo(sdk.NewDec(100)),
				PrimaryShare: types.Account{
					Id:   oldSubdistributor.Destination.Account.Id,
					Type: oldSubdistributor.Destination.Account.Type,
				},
			},
		}
		newSubdistributors = append(newSubdistributors, newSubdistributor)
	}
	err := types.ValidateSubDistributors(newSubdistributors)
	if err != nil {
		return err
	}
	paramStore.Set(ctx, types.KeySubDistributors, newSubdistributors)

	return nil
}
