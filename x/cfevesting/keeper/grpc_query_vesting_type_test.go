package keeper_test

import (
	"testing"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestVestingTypesQueryEmpty(t *testing.T) {
	keeper, ctx := testkeeper.CfevestingKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)

	response, err := keeper.VestingType(wctx, &types.QueryVestingTypeRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryVestingTypeResponse{types.VestingTypes{}}, response)
}

func TestVestingTypesQueryNotEmpty(t *testing.T) {
	keeper, ctx := testkeeper.CfevestingKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	vestingTypes := types.VestingTypes{}
	vestingType1 := types.VestingType{"test1", 2324, 42423, 4243, true}
	vestingType2 := types.VestingType{"test2", 1111, 112233, 445566, false}

	vestingTypesArray := []*types.VestingType{&vestingType1, &vestingType2}
	vestingTypes.VestingTypes = vestingTypesArray

	keeper.SetVestingTypes(ctx, vestingTypes)
	response, err := keeper.VestingType(wctx, &types.QueryVestingTypeRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryVestingTypeResponse{vestingTypes}, response)

}