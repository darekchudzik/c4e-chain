package cfeairdrop

import (
	"testing"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	cfeairdropmodulekeeper "github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
	cfeairdroptypes "github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/stretchr/testify/require"
)

type C4eAirdropUtils struct {
	C4eAirdropKeeperUtils
	helperAccountKeeper *authkeeper.AccountKeeper
	bankUtils           *commontestutils.BankUtils
}

func NewC4eAirdropUtils(t *testing.T, helpeCfeairdropmodulekeeper *cfeairdropmodulekeeper.Keeper,
	helperAccountKeeper *authkeeper.AccountKeeper,
	bankUtils *commontestutils.BankUtils) C4eAirdropUtils {
	return C4eAirdropUtils{C4eAirdropKeeperUtils: NewC4eAirdropKeeperUtils(t, helpeCfeairdropmodulekeeper), helperAccountKeeper: helperAccountKeeper, bankUtils: bankUtils}
}

func (h *C4eAirdropUtils) SendToAirdropAccount(ctx sdk.Context, toAddress sdk.AccAddress,
	amount sdk.Int, startTime int64, endTime int64, createAccount bool) {
	coins := sdk.NewCoins(sdk.NewCoin(commontestutils.DefaultTestDenom, amount))
	moduleBalance := h.bankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName)
	accBalance := h.bankUtils.GetAccountDefultDenomBalance(ctx, toAddress)

	accountBefore := h.helperAccountKeeper.GetAccount(ctx, toAddress)
	periodsAmount := 0
	previousOriginalVesting := sdk.NewCoins()
	if accountBefore != nil {
		if airdropAccount, ok := accountBefore.(*cfeairdroptypes.AirdropVestingAccount); ok {
			periodsAmount = len(airdropAccount.VestingPeriods)
			previousOriginalVesting = previousOriginalVesting.Add(airdropAccount.OriginalVesting...)
		}
	}

	require.NoError(h.t, h.helpeCfeairdropmodulekeeper.SendToAirdropAccount(ctx,
		toAddress,
		coins,
		startTime,
		endTime, createAccount,
	))

	h.bankUtils.VerifyAccountDefultDenomBalance(ctx, toAddress, accBalance.Add(amount))
	h.bankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName, moduleBalance.Sub(amount))

	airdropAccount, _ := h.helperAccountKeeper.GetAccount(ctx, toAddress).(*cfeairdroptypes.AirdropVestingAccount)
	periodPosition := len(airdropAccount.VestingPeriods) - 1
	require.EqualValues(h.t, periodsAmount+1, len(airdropAccount.VestingPeriods))
	require.EqualValues(h.t, startTime, airdropAccount.StartTime)
	require.EqualValues(h.t, endTime, airdropAccount.EndTime)
	require.EqualValues(h.t, previousOriginalVesting.Add(coins...), airdropAccount.OriginalVesting)
	require.EqualValues(h.t, startTime, airdropAccount.VestingPeriods[periodPosition].StartTime)
	require.EqualValues(h.t, endTime, airdropAccount.VestingPeriods[periodPosition].EndTime)
	require.EqualValues(h.t, coins, airdropAccount.VestingPeriods[periodPosition].Amount)
	require.NoError(h.t, airdropAccount.Validate())

}

func (h *C4eAirdropUtils) SendToAirdropAccountError(ctx sdk.Context, toAddress sdk.AccAddress,
	amount sdk.Int, startTime int64, endTime int64, createAccount bool, errorMessage string, expectNewAccount bool) {
	coins := sdk.NewCoins(sdk.NewCoin(commontestutils.DefaultTestDenom, amount))
	moduleBalance := h.bankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName)
	accBalance := h.bankUtils.GetAccountDefultDenomBalance(ctx, toAddress)

	accountBefore := h.helperAccountKeeper.GetAccount(ctx, toAddress)
	wasAirdropAccount := false
	if accountBefore != nil {
		_, wasAirdropAccount = accountBefore.(*cfeairdroptypes.AirdropVestingAccount)
	}

	require.EqualError(h.t, h.helpeCfeairdropmodulekeeper.SendToAirdropAccount(ctx,
		toAddress,
		coins,
		startTime,
		endTime, createAccount,
	), errorMessage)

	h.bankUtils.VerifyAccountDefultDenomBalance(ctx, toAddress, accBalance)
	h.bankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName, moduleBalance)

	accountAfter := h.helperAccountKeeper.GetAccount(ctx, toAddress)
	airdropAccount, isAirdropAccount := h.helperAccountKeeper.GetAccount(ctx, toAddress).(*cfeairdroptypes.AirdropVestingAccount)

	if accountBefore == nil && expectNewAccount {
		require.EqualValues(h.t, true, isAirdropAccount)
		require.EqualValues(h.t, 0, len(airdropAccount.VestingPeriods))
		require.EqualValues(h.t, startTime, airdropAccount.StartTime)
		require.EqualValues(h.t, endTime, airdropAccount.EndTime)
		require.True(h.t, sdk.NewCoins().IsEqual(airdropAccount.OriginalVesting))
		require.NoError(h.t, airdropAccount.Validate())
	} else {
		require.EqualValues(h.t, wasAirdropAccount, isAirdropAccount)
		require.EqualValues(h.t, accountBefore, accountAfter)
	}

}