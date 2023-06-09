package v3

import (
	"github.com/chain4energy/c4e-chain/x/cfeminter/exported"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MigrateParams migrates the x/cfeminter module state from the consensus version 2 to
// version 3. Specifically, it takes the parameters that are currently stored
// and managed by the x/params module and stores them directly into the x/cfeminter
// module state.
// The migration also includes:
// - cfeminter module refactoring
// - delete type field from minterConfig
// - minter config is now of type Any rather than using 2 fields (LinearMinting and ExponentialStepMinting)
// - MinterConfig was deleted and minters and start-time was moved directly to cfeminter params
func MigrateParams(ctx sdk.Context, storeKey storetypes.StoreKey, legacySubspace exported.Subspace, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	var oldParams types.LegacyParams
	legacySubspace.GetParamSet(ctx, &oldParams)
	if err := oldParams.MinterConfig.Validate(); err != nil {
		return err
	}

	var newParams types.Params
	newParams.MintDenom = oldParams.MintDenom
	newParams.StartTime = oldParams.MinterConfig.StartTime
	for _, oldMinter := range oldParams.MinterConfig.Minters {
		var newMinter types.Minter
		newMinter.SequenceId = oldMinter.SequenceId
		newMinter.EndTime = oldMinter.EndTime
		newParams.Minters = append(newParams.Minters, &newMinter)
		var config *codectypes.Any
		var err error
		if oldMinter.Type == types.ExponentialStepMintingType {
			config, err = codectypes.NewAnyWithValue(oldMinter.ExponentialStepMinting)
			if err != nil {
				return err
			}
		} else if oldMinter.Type == types.LinearMintingType {
			config, err = codectypes.NewAnyWithValue(oldMinter.LinearMinting)
			if err != nil {
				return err
			}
		} else {
			config, err = codectypes.NewAnyWithValue(&types.NoMinting{})
			if err != nil {
				return err
			}
		}
		newMinter.Config = config
	}

	if err := newParams.Validate(); err != nil {
		return err
	}

	bz, err := cdc.Marshal(&newParams)
	if err != nil {
		return err
	}

	store.Set(types.ParamsKey, bz)

	return nil
}
