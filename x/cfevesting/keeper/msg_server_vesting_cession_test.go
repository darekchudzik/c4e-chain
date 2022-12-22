package keeper_test

import (
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"
	"time"
)

func TestCorrectVestingCession(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(3, 0)
	accAddr1 := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]
	accAddr3 := acountsAddresses[2]

	accBalance := sdk.NewInt(100000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accBalance, accAddr1)
	sendAmount := sdk.NewInt(10000)
	coins := sdk.Coins{{Amount: sendAmount, Denom: commontestutils.DefaultTestDenom}}
	startTime := time.Date(2025, 2, 3, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2035, 2, 3, 0, 0, 0, 0, time.UTC)

	testHelper.C4eVestingUtils.MessageCreateVestingAccount(
		accAddr1,
		accAddr2,
		coins,
		startTime,
		endTime,
		accBalance,
	)

	testHelper.C4eVestingUtils.MessageNewVestingCession(
		accAddr2,
		accAddr3,
		sendAmount.QuoRaw(2),
		sendAmount,
		endTime,
	)

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}
