package keeper_test

import (
	// "context"
	// "fmt"
	"testing"

	// keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"

	// "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/stretchr/testify/require"

	"github.com/chain4energy/c4e-chain/app"
	// "github.com/cosmos/cosmos-sdk/x/auth"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestVestDelegationNotAllowed(t *testing.T) {
	vestDelegation(t, false)
}

func TestVestDelegationAllowed(t *testing.T) {
	vestDelegation(t, true)
}

func vestDelegation(t *testing.T, delegationAllowed bool) {
	// sdk.GetConfig().SetBech32PrefixForAccount("c4e", "c4e")
	const addr = "cosmos1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8245g3w"
	const delagableAddr = "cosmos1dfugyfm087qa3jrdglkeaew0wkn59jk8mgw6x6"

	accAddr, _ := sdk.AccAddressFromBech32(addr)
	delegableAccAddr, _ := sdk.AccAddressFromBech32(delagableAddr)

	const vt1 = "test1"
	const initBlock = 1000
	const vested = 1000
	const accInitBalance = 10000
	vestingTypes := types.VestingTypes{}
	vestingType1 := types.VestingType{vt1, 1000, 5000, 10, delegationAllowed}
	vestingType2 := types.VestingType{"test2", 1111, 112233, 445566, false}

	vestingTypesArray := []*types.VestingType{&vestingType1, &vestingType2}
	vestingTypes.VestingTypes = vestingTypesArray

	addHelperModuleAccountPerms()
	// k, ctx := keepertest.CfevestingKeeperWithBlockHeight(t, initBlock)

	app := app.Setup(false)
	header := tmproto.Header{}
	header.Height = initBlock
	ctx := app.BaseApp.NewContext(false, header)

	bank := app.BankKeeper
	// mint := app.MintKeeper
	auth := app.AccountKeeper

	denom := "uc4e"
	// mintedCoin := sdk.NewCoin(denom, sdk.NewIntFromUint64(accInitBalance))
	// mintedCoins := sdk.NewCoins(mintedCoin)
	// mint.MintCoins(ctx, mintedCoins)
	// bank.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, accAddr, mintedCoins)

	addCoinsToAccount(accInitBalance, helperModuleAccount, ctx, bank, accAddr)

	k := app.CfevestingKeeper

	k.SetVestingTypes(ctx, vestingTypes)
	msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)

	msg := types.MsgVest{addr, vested, "test1"}
	_, error := msgServer.Vest(msgServerCtx, &msg)
	require.EqualValues(t, nil, error)

	accVestings := k.GetAllAccountVestings(ctx)

	_, accFound := k.GetAccountVestings(ctx, addr)
	require.EqualValues(t, true, accFound)

	var expectedDelegableAcc string
	if delegationAllowed {
		expectedDelegableAcc = delagableAddr
	} else {
		expectedDelegableAcc = ""
	}
	require.EqualValues(t, 1, len(accVestings))
	require.EqualValues(t, 1, len(accVestings[0].Vestings))
	require.EqualValues(t, expectedDelegableAcc, accVestings[0].DelegableAddress)

	require.EqualValues(t, addr, accVestings[0].Address)
	vesting := accVestings[0].Vestings[0]
	require.EqualValues(t, vt1, vesting.VestingType)
	require.EqualValues(t, initBlock, vesting.VestingStartBlock)
	require.EqualValues(t, initBlock+vestingType1.LockupPeriod, vesting.LockEndBlock)

	require.EqualValues(t, initBlock+vestingType1.LockupPeriod+vestingType1.VestingPeriod, vesting.VestingEndBlock)

	require.EqualValues(t, vested, vesting.Vested)
	require.EqualValues(t, 0, vesting.Claimable)
	require.EqualValues(t, 0, vesting.LastFreeingBlock)
	require.EqualValues(t, vestingType1.TokenReleasingPeriod, vesting.FreeCoinsBlockPeriod)
	require.EqualValues(t, 2, vesting.FreeCoinsPerPeriod)
	require.EqualValues(t, delegationAllowed, vesting.DelegationAllowed)

	balance := bank.GetBalance(ctx, accAddr, denom)
	require.EqualValues(t, sdk.NewIntFromUint64(accInitBalance-vested), balance.Amount)
	moduleAccAddr := auth.GetModuleAccount(ctx, types.ModuleName).GetAddress()
	moduleBalance := bank.GetBalance(ctx, moduleAccAddr, denom)
	if delegationAllowed {
		require.EqualValues(t, sdk.ZeroInt(), moduleBalance.Amount)
		delegableBalance := bank.GetBalance(ctx, delegableAccAddr, denom)
		require.EqualValues(t, sdk.NewIntFromUint64(vested), delegableBalance.Amount)
	} else {
		require.EqualValues(t, sdk.NewIntFromUint64(vested), moduleBalance.Amount)
	}

	_, error = msgServer.Vest(msgServerCtx, &msg)

	require.EqualValues(t, nil, error)

	accVestings = k.GetAllAccountVestings(ctx)

	_, accFound = k.GetAccountVestings(ctx, addr)
	require.EqualValues(t, true, accFound)

	require.EqualValues(t, 1, len(accVestings))
	require.EqualValues(t, 2, len(accVestings[0].Vestings))

	balance = bank.GetBalance(ctx, accAddr, denom)
	require.EqualValues(t, sdk.NewIntFromUint64(accInitBalance-vested-vested), balance.Amount)

	moduleBalance = bank.GetBalance(ctx, moduleAccAddr, denom)
	if delegationAllowed {
		require.EqualValues(t, sdk.ZeroInt(), moduleBalance.Amount)
		delegableBalance := bank.GetBalance(ctx, delegableAccAddr, denom)
		require.EqualValues(t, sdk.NewIntFromUint64(vested+vested), delegableBalance.Amount)
	} else {
		require.EqualValues(t, sdk.NewIntFromUint64(vested+vested), moduleBalance.Amount)
	}
}