package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
)

// SetClaimRecordXX set a specific claimRecordXX in the store from its index
func (k Keeper) SetClaimRecord(ctx sdk.Context, claimRecord types.ClaimRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimRecordKeyPrefix))
	b := k.cdc.MustMarshal(&claimRecord)
	store.Set(types.ClaimRecordKey(
		claimRecord.Address,
	), b)
}

// GetClaimRecordXX returns a claimRecordXX from its index
func (k Keeper) GetClaimRecord(
	ctx sdk.Context,
	address string,

) (val types.ClaimRecord, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimRecordKeyPrefix))

	b := store.Get(types.ClaimRecordKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveClaimRecordXX removes a claimRecordXX from the store
func (k Keeper) RemoveClaimRecord(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimRecordKeyPrefix))
	store.Delete(types.ClaimRecordKey(
		index,
	))
}

// GetAllClaimRecordXX returns all claimRecordXX
func (k Keeper) GetAllClaimRecord(ctx sdk.Context) (list []types.ClaimRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimRecordKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ClaimRecord
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) addClaimRecord(ctx sdk.Context, address string, campaignId uint64, claimable sdk.Int) (*types.ClaimRecord, error) {
	claimRecord, found := k.GetClaimRecord(ctx, address)
	if !found {
		claimRecord = types.ClaimRecord{Address: address}
		// k.grantFeeAllowance(ctx, address)
	}
	if claimRecord.HasCampaign(campaignId) {
		return nil, sdkerrors.Wrapf(types.ErrAlreadyExists, "campaignId %d already exists for address: %s", campaignId, address)
	}
	claimRecord.CampaignRecords = append(claimRecord.CampaignRecords, &types.CampaignRecord{CampaignId: campaignId, Claimable: claimable})
	return &claimRecord, nil
}

type CampaignRecordData struct {
	Address   string
	Claimable sdk.Int
}

func (k Keeper) AddCampaignRecords(ctx sdk.Context, campaignId uint64, campaignRecord []*CampaignRecordData) error {
	records := []*types.ClaimRecord{}

	for _, recordData := range campaignRecord {
		record, err := k.addClaimRecord(ctx, recordData.Address, campaignId, recordData.Claimable)
		if err != nil {
			return err
		}
		records = append(records, record)
	}

	for _, record := range records {
		k.SetClaimRecord(ctx, *record)
	}
	return nil
}

func (k Keeper) grantFeeAllowance(ctx sdk.Context, grantee string) error {
	allowance := feegranttypes.BasicAllowance{}
	address, err := sdk.AccAddressFromBech32(grantee)
	if err != nil {
		return nil // TODO
	}
	modAcc := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
	k.feeGrantKeeper.GrantAllowance(ctx, modAcc.GetAddress(), address, &allowance)
	return nil // TODO error handling
}

// func (k Keeper) revokeFeeAllowance(ctx sdk.Context, grantee sdk.AccAddress) error  {
// 	allowance := feegranttypes.BasicAllowance{}

// 	modAcc := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
// 	k.feeGrantKeeper.GrantAllowance(ctx, modAcc.GetAddress(), grantee, &allowance)
// 	return nil // TODO error handling
// }