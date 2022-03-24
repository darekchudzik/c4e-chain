package keeper_test

import (
	"strconv"

	"github.com/chain4energy/c4e-chain/app"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"testing"

	"github.com/stretchr/testify/require"
)

const helperModuleAccount = "helperTestAcc"
const denom = "uc4e"

func addHelperModuleAccountPerms() {
	perms := []string{authtypes.Minter}
	app.AddMaccPerms(helperModuleAccount, perms)
}

func addCoinsToAccount(vested uint64, ctx sdk.Context, app *app.App, toAddr sdk.AccAddress) string {
	denom := "uc4e"
	mintedCoin := sdk.NewCoin(denom, sdk.NewIntFromUint64(vested))
	mintedCoins := sdk.NewCoins(mintedCoin)
	app.BankKeeper.MintCoins(ctx, helperModuleAccount, mintedCoins)
	app.BankKeeper.SendCoinsFromModuleToAccount(ctx, helperModuleAccount, toAddr, mintedCoins)
	return denom
}

func createAccountVestings(addr string, vested uint64, withdrawn uint64) (types.AccountVestings, *types.Vesting) {
	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(1, 1, 1)
	accountVestings.Address = addr
	accountVestings.Vestings[0].Vested = sdk.NewIntFromUint64(vested)
	accountVestings.Vestings[0].DelegationAllowed = false
	accountVestings.Vestings[0].Withdrawn = sdk.NewIntFromUint64(withdrawn)
	accountVestings.Vestings[0].LastModificationVested = sdk.NewIntFromUint64(vested)
	accountVestings.Vestings[0].LastModificationWithdrawn = sdk.NewIntFromUint64(withdrawn)
	return accountVestings, accountVestings.Vestings[0]
}

func addCoinsToModuleByName(vested uint64, modulaName string, ctx sdk.Context, app *app.App) string {
	denom := "uc4e"
	mintedCoin := sdk.NewCoin(denom, sdk.NewIntFromUint64(vested))
	mintedCoins := sdk.NewCoins(mintedCoin)
	app.BankKeeper.MintCoins(ctx, helperModuleAccount, mintedCoins)
	app.BankKeeper.SendCoinsFromModuleToModule(ctx, helperModuleAccount, modulaName, mintedCoins)
	return denom
}

func verifyAccountBalance(t *testing.T, app *app.App, ctx sdk.Context, accAddr sdk.AccAddress, expectedAmount sdk.Int) {
	balance := app.BankKeeper.GetBalance(ctx, accAddr, denom)
	require.EqualValues(t, expectedAmount, balance.Amount)
}

func verifyModuleAccountByName(accName string, ctx sdk.Context, app *app.App, t *testing.T, expected sdk.Int) {
	moduleAccAddr := app.AccountKeeper.GetModuleAccount(ctx, accName).GetAddress()
	moduleBalance := app.BankKeeper.GetBalance(ctx, moduleAccAddr, denom)
	require.EqualValues(t, expected, moduleBalance.Amount)
}

func verifyModuleAccount(t *testing.T, ctx sdk.Context, app *app.App, expected sdk.Int) {
	verifyModuleAccountByName(types.ModuleName, ctx, app, t, expected)
}

func createValidator(t *testing.T, ctx sdk.Context, sk stakingkeeper.Keeper, addr sdk.ValAddress, pk cryptotypes.PubKey, coin sdk.Coin, commisions stakingtypes.CommissionRates) {
	msg, err := stakingtypes.NewMsgCreateValidator(addr, pk, coin, stakingtypes.Description{}, commisions, sdk.OneInt())
	msgSrvr := stakingkeeper.NewMsgServerImpl(sk)
	require.NoError(t, err)
	res, err := msgSrvr.CreateValidator(sdk.WrapSDKContext(ctx), msg)
	require.NoError(t, err)
	require.NotNil(t, res)

}

func redelegate(t *testing.T, ctx sdk.Context, app *app.App, delegatorAddress sdk.AccAddress, delegableAddress sdk.AccAddress,
	validatorSrcAddress sdk.ValAddress, validatorDstAddress sdk.ValAddress, redelegateAmount uint64, delegatorAmountBefore uint64, delegableAmountBefore uint64, delegatorAmountAfter uint64, delegableAmountAfter uint64) {

	verifyAccountBalance(t, app, ctx, delegatorAddress, sdk.NewIntFromUint64(delegatorAmountBefore))
	verifyAccountBalance(t, app, ctx, delegableAddress, sdk.NewIntFromUint64(delegableAmountBefore))

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)
	coin := sdk.NewCoin(denom, sdk.NewIntFromUint64(redelegateAmount))
	msgRe := types.MsgBeginRedelegate{
		DelegatorAddress:    delegatorAddress.String(),
		ValidatorSrcAddress: validatorSrcAddress.String(),
		ValidatorDstAddress: validatorDstAddress.String(),
		Amount:              coin,
	}
	_, err := msgServer.BeginRedelegate(msgServerCtx, &msgRe)
	require.EqualValues(t, nil, err)

	verifyAccountBalance(t, app, ctx, delegatorAddress, sdk.NewIntFromUint64(delegatorAmountAfter))

	verifyAccountBalance(t, app, ctx, delegableAddress, sdk.NewIntFromUint64(delegableAmountAfter))
}

func setupValidators(t *testing.T, ctx sdk.Context, app *app.App, validators []sdk.ValAddress, delegatePerValidator uint64) {
	PKs := commontestutils.CreateTestPubKeys(len(validators))
	commission := stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(0, 1), sdk.NewDecWithPrec(0, 1), sdk.NewDec(0))
	delCoin := sdk.NewCoin(denom, sdk.NewIntFromUint64(delegatePerValidator))
	for i, valAddr := range validators {
		addCoinsToAccount(delegatePerValidator, ctx, app, valAddr.Bytes())
		createValidator(t, ctx, app.StakingKeeper, valAddr, PKs[i], delCoin, commission)
	}
	require.EqualValues(t, len(validators), len(app.StakingKeeper.GetAllValidators(ctx)))
}

func setupStakingBondDenom(ctx sdk.Context, app *app.App) {
	stakeParams := app.StakingKeeper.GetParams(ctx)
	stakeParams.BondDenom = denom
	app.StakingKeeper.SetParams(ctx, stakeParams)
}

func delegate(t *testing.T, ctx sdk.Context, app *app.App, delegatorAddress sdk.AccAddress, delegableAddress sdk.AccAddress,
	validatorAddress sdk.ValAddress, delegateAmount uint64, delegatorAccountAmountBefore int64, delegableAccountAmountBefore int64, delegatorAccountAmountAfter int64, delegableAccountAmountAfter int64) {
	verifyAccountBalance(t, app, ctx, delegatorAddress, sdk.NewInt(delegatorAccountAmountBefore))
	verifyAccountBalance(t, app, ctx, delegableAddress, sdk.NewInt(delegableAccountAmountBefore))

	coin := sdk.NewCoin(denom, sdk.NewIntFromUint64(delegateAmount))
	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgDelegate{DelegatorAddress: delegatorAddress.String(), ValidatorAddress: validatorAddress.String(), Amount: coin}
	_, err := msgServer.Delegate(msgServerCtx, &msg)
	require.EqualValues(t, nil, err)
	verifyAccountBalance(t, app, ctx, delegatorAddress, sdk.NewInt(delegatorAccountAmountAfter))
	verifyAccountBalance(t, app, ctx, delegableAddress, sdk.NewInt(delegableAccountAmountAfter))
}

func undelegate(t *testing.T, ctx sdk.Context, app *app.App, delegatorAddress sdk.AccAddress, delegableAddress sdk.AccAddress,
	validatorAddress sdk.ValAddress, delegateAmount uint64, delegatorAccountAmountBefore int64, delegableAccountAmountBefore int64, delegatorAccountAmountAfter int64, delegableAccountAmountAfter int64) {
	verifyAccountBalance(t, app, ctx, delegatorAddress, sdk.NewInt(delegatorAccountAmountBefore))
	verifyAccountBalance(t, app, ctx, delegableAddress, sdk.NewInt(delegableAccountAmountBefore))

	coin := sdk.NewCoin(denom, sdk.NewIntFromUint64(delegateAmount))
	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msgUn := types.MsgUndelegate{DelegatorAddress: delegatorAddress.String(), ValidatorAddress: validatorAddress.String(), Amount: coin}
	_, err := msgServer.Undelegate(msgServerCtx, &msgUn)
	require.EqualValues(t, nil, err)
	verifyAccountBalance(t, app, ctx, delegatorAddress, sdk.NewInt(delegatorAccountAmountAfter))
	verifyAccountBalance(t, app, ctx, delegableAddress, sdk.NewInt(delegableAccountAmountAfter))
}

func verifyDelegations(t *testing.T, ctx sdk.Context, app *app.App, delegableAddress sdk.AccAddress,
	validators []sdk.ValAddress, delegated []int64) {
	delegations := app.StakingKeeper.GetAllDelegatorDelegations(ctx, delegableAddress)
	require.EqualValues(t, len(validators), len(delegations))
	for i, valAddr := range validators {
		found := false
		for _, delegation := range delegations {
			if delegation.ValidatorAddress == valAddr.String() {
				require.EqualValues(t, sdk.NewDec(delegated[i]), delegation.Shares)
				found = true
			}
		}
		require.True(t, found, "delegation not found. Validator Address: "+valAddr.String())
	}

}

func verifyUnbondingDelegations(t *testing.T, ctx sdk.Context, app *app.App, delegableAddress sdk.AccAddress,
	validators []sdk.ValAddress, unbondingAmount []int64) {

	unbondingDelegations := app.StakingKeeper.GetAllUnbondingDelegations(ctx, delegableAddress)
	require.EqualValues(t, len(validators), len(unbondingDelegations))

	for i, valAddr := range validators {
		found := false
		for _, delegation := range unbondingDelegations {
			if delegation.ValidatorAddress == valAddr.String() {
				require.EqualValues(t, 1, len(unbondingDelegations[0].Entries))
				require.EqualValues(t, sdk.NewInt(unbondingAmount[i]), unbondingDelegations[0].Entries[0].Balance)
				require.EqualValues(t, sdk.NewInt(unbondingAmount[i]), unbondingDelegations[0].Entries[0].InitialBalance)
				found = true
			}
		}
		require.True(t, found, "delegation not found. Validator Address: "+valAddr.String())
	}

}

func setupAccountsVestings(ctx sdk.Context, app *app.App, address string, delegableAddress string, numberOfVestings int, vestingAmount uint64, withdrawnAmount uint64, delegationAllowed bool) types.AccountVestings {
	return setupAccountsVestingsWithModification(ctx, app, func(*types.Vesting) {}, address, delegableAddress, numberOfVestings, vestingAmount, withdrawnAmount, delegationAllowed)
}

func setupAccountsVestingsWithModification(ctx sdk.Context, app *app.App, modifyVesting func(*types.Vesting), address string, delegableAddress string, numberOfVestings int, vestingAmount uint64, withdrawnAmount uint64, delegationAllowed bool) types.AccountVestings {
	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(numberOfVestings, 1, 1)
	accountVestings.Address = address
	accountVestings.DelegableAddress = delegableAddress

	for _, vesting := range accountVestings.Vestings {
		vesting.Vested = sdk.NewIntFromUint64(vestingAmount)
		vesting.DelegationAllowed = delegationAllowed
		vesting.Withdrawn = sdk.NewIntFromUint64(withdrawnAmount)
		vesting.LastModificationVested = sdk.NewIntFromUint64(vestingAmount)
		vesting.LastModificationWithdrawn = sdk.NewIntFromUint64(withdrawnAmount)
		modifyVesting(vesting)
	}
	app.CfevestingKeeper.SetAccountVestings(ctx, accountVestings)
	return accountVestings
}

func allocateRewardsToValidator(ctx sdk.Context, app *app.App, validatorRewards uint64, valAddr sdk.ValAddress) {
	valCons := sdk.NewDecCoin(denom, sdk.NewIntFromUint64(validatorRewards))
	val := app.StakingKeeper.Validator(ctx, valAddr)
	app.DistrKeeper.AllocateTokensToValidator(ctx, val, sdk.NewDecCoins(valCons))
}

func verifyQueryRewards(t *testing.T, ctx sdk.Context, app *app.App, delegableAddr sdk.AccAddress, valAddr sdk.ValAddress, hasRewards bool, rewards uint64) {
	msgServerCtx := sdk.WrapSDKContext(ctx)
	query := distrtypes.QueryDelegationRewardsRequest{DelegatorAddress: delegableAddr.String(), ValidatorAddress: valAddr.String()}
	resp, _ := app.DistrKeeper.DelegationRewards(msgServerCtx, &query)
	if hasRewards {
		require.EqualValues(t, 1, len(resp.Rewards))
		require.EqualValues(t, sdk.NewDecFromInt(sdk.NewIntFromUint64(rewards)), resp.Rewards[0].Amount)
	} else {
		require.EqualValues(t, 0, len(resp.Rewards))
	}

}

func makeVesting(t *testing.T, ctx sdk.Context, app *app.App, address sdk.AccAddress, accountVestingsExistsBefore bool, accountVestingsExistsAfter bool,
	delegableAddressExistsBefore bool, delegableAddressExistsAfter bool,
	vestingType types.VestingType, amountToVest int64, accAmountBefore int64, delegableAmountBefore int64, moduleAmountBefore int64,
	accAmountAfter int64, delegableAmountAfter int64, moduleAmountAfter int64) {

	accVestings, accFound := app.CfevestingKeeper.GetAccountVestings(ctx, address.String())
	require.EqualValues(t, accountVestingsExistsBefore, accFound)

	verifyAccountBalance(t, app, ctx, address, sdk.NewInt(accAmountBefore))
	moduleAccAddr := app.AccountKeeper.GetModuleAccount(ctx, types.ModuleName).GetAddress()
	verifyAccountBalance(t, app, ctx, moduleAccAddr, sdk.NewInt(moduleAmountBefore))
	if delegableAddressExistsBefore {
		delegableAddress, _ := sdk.AccAddressFromBech32(accVestings.DelegableAddress)
		verifyAccountBalance(t, app, ctx, delegableAddress, sdk.NewInt(delegableAmountBefore))
	}

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgVest{Creator: address.String(), Amount: sdk.NewInt(amountToVest), VestingType: vestingType.Name}
	_, error := msgServer.Vest(msgServerCtx, &msg)
	require.EqualValues(t, nil, error)

	accVestings, accFound = app.CfevestingKeeper.GetAccountVestings(ctx, address.String())
	require.EqualValues(t, accountVestingsExistsAfter, accFound)

	verifyAccountBalance(t, app, ctx, address, sdk.NewInt(accAmountAfter))
	verifyAccountBalance(t, app, ctx, moduleAccAddr, sdk.NewInt(moduleAmountAfter))
	if delegableAddressExistsAfter {
		delegableAddress, _ := sdk.AccAddressFromBech32(accVestings.DelegableAddress)
		verifyAccountBalance(t, app, ctx, delegableAddress, sdk.NewInt(delegableAmountAfter))
	}
}

func newInts64Array(n int, v int64) []int64 {
	s := make([]int64, n)
	for i := range s {
		s[i] = v
	}
	return s
}

func verifyAccountVestings(t *testing.T, ctx sdk.Context, app *app.App, address sdk.AccAddress,
	vestingTypes []types.VestingType, vestedAmounts []int64, withdrawnAmounts []int64) {

	verifyAccountVestingsWithModification(t, ctx, app, address, 1, vestingTypes, newInts64Array(len(vestingTypes), ctx.BlockHeight()), vestedAmounts, withdrawnAmounts,
		newInts64Array(len(vestingTypes), 0), newInts64Array(len(vestingTypes), ctx.BlockHeight()), vestedAmounts, withdrawnAmounts)
}

func verifyAccountVestingsWithModification(t *testing.T, ctx sdk.Context, app *app.App, address sdk.AccAddress,
	amountOfAllAccVestings int, vestingTypes []types.VestingType, startHeights []int64, vestedAmounts []int64, withdrawnAmounts []int64,
	sentAmounts []int64, modificationsHeights []int64, modificationsVested []int64, modificationsWithdrawn []int64) {
	allAccVestings := app.CfevestingKeeper.GetAllAccountVestings(ctx)

	accVestings, accFound := app.CfevestingKeeper.GetAccountVestings(ctx, address.String())
	require.EqualValues(t, true, accFound)

	require.EqualValues(t, amountOfAllAccVestings, len(allAccVestings))
	require.EqualValues(t, len(vestingTypes), len(accVestings.Vestings))

	delegationsAllowed := false
	for _, vestingType := range vestingTypes {
		if vestingType.DelegationsAllowed {
			delegationsAllowed = true
			break
		}

	}
	require.EqualValues(t, address.String(), accVestings.Address)
	if delegationsAllowed {
		require.NotEqualValues(t, "", accVestings.DelegableAddress)
	} else {
		require.EqualValues(t, "", accVestings.DelegableAddress)
	}

	for i, vesting := range accVestings.Vestings {
		found := false
		if vesting.Id == int32(i+1) {
			require.EqualValues(t, i+1, vesting.Id)
			require.EqualValues(t, vestingTypes[i].Name, vesting.VestingType)
			require.EqualValues(t, startHeights[i], vesting.VestingStart)
			require.EqualValues(t, ctx.BlockHeight()+vestingTypes[i].LockupPeriod, vesting.LockEnd)
			require.EqualValues(t, ctx.BlockHeight()+vestingTypes[i].LockupPeriod+vestingTypes[i].VestingPeriod, vesting.VestingEnd)
			require.EqualValues(t, sdk.NewInt(vestedAmounts[i]), vesting.Vested)
			require.EqualValues(t, vestingTypes[i].TokenReleasingPeriod, vesting.ReleasePeriod)
			require.EqualValues(t, vestingTypes[i].DelegationsAllowed, vesting.DelegationAllowed)
			require.EqualValues(t, sdk.NewInt(withdrawnAmounts[i]), vesting.Withdrawn)

			require.EqualValues(t, sdk.NewInt(sentAmounts[i]), vesting.Sent)
			require.EqualValues(t, modificationsHeights[i], vesting.LastModification)
			require.EqualValues(t, sdk.NewInt(modificationsVested[i]), vesting.LastModificationVested)
			require.EqualValues(t, sdk.NewInt(modificationsWithdrawn[i]), vesting.LastModificationWithdrawn)
			found = true

		}
		require.True(t, found, "not found vesting id: "+strconv.Itoa(i+1))

	}

}

func setupVestingTypes(ctx sdk.Context, app *app.App, numberOfVestingTypes int, amountOf10BasedVestingTypes int, a10BasedVestingTypesDelegationAllowe bool, startId int) types.VestingTypes {
	return setupVestingTypesWithModification(ctx, app, func(*types.VestingType) {}, numberOfVestingTypes, amountOf10BasedVestingTypes, a10BasedVestingTypesDelegationAllowe, startId)
}

func setupVestingTypesWithModification(ctx sdk.Context, app *app.App, modifyVestingType func(*types.VestingType), numberOfVestingTypes int, amountOf10BasedVestingTypes int, a10BasedVestingTypesDelegationAllowe bool, startId int) types.VestingTypes {
	vestingTypesArray := testutils.Generate10BasedVestingTypes(numberOfVestingTypes, amountOf10BasedVestingTypes, a10BasedVestingTypesDelegationAllowe, startId)
	for _, vestingType := range vestingTypesArray {
		modifyVestingType(vestingType)
	}
	vestingTypes := types.VestingTypes{VestingTypes: vestingTypesArray}
	app.CfevestingKeeper.SetVestingTypes(ctx, vestingTypes)
	return vestingTypes
}

func withdrawAllAvailable(t *testing.T, ctx sdk.Context, app *app.App, address sdk.AccAddress, accountBalanceBefore int64, moduleBalanceBefore int64,
	accountBalanceAfter int64, moduleBalanceAfter int64) {

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	verifyAccountBalance(t, app, ctx, address, sdk.NewInt(accountBalanceBefore))
	verifyModuleAccount(t, ctx, app, sdk.NewInt(moduleBalanceBefore))
	msg := types.MsgWithdrawAllAvailable{Creator: address.String()}
	_, error := msgServer.WithdrawAllAvailable(msgServerCtx, &msg)
	require.EqualValues(t, nil, error)
	verifyAccountBalance(t, app, ctx, address, sdk.NewInt(accountBalanceAfter))
	verifyModuleAccount(t, ctx, app, sdk.NewInt(moduleBalanceAfter))
}

func withdrawAllAvailableDelegable(t *testing.T, ctx sdk.Context, app *app.App, address sdk.AccAddress, delegableAddress sdk.AccAddress, accountBalanceBefore int64, delegableBalanceBefore int64,
	moduleBalanceBefore int64, accountBalanceAfter int64, delegableBalanceAfter int64, moduleBalanceAfter int64) {

	verifyAccountBalance(t, app, ctx, delegableAddress, sdk.NewInt(delegableBalanceBefore))
	withdrawAllAvailable(t, ctx, app, address, accountBalanceBefore, moduleBalanceBefore, accountBalanceAfter, moduleBalanceAfter)
	verifyAccountBalance(t, app, ctx, delegableAddress, sdk.NewInt(delegableBalanceAfter))
}

func compareStoredAcountVestings(t *testing.T, ctx sdk.Context, app *app.App, address sdk.AccAddress, accVestings types.AccountVestings) {
	storedAccVestings, accFound := app.CfevestingKeeper.GetAccountVestings(ctx, address.String())
	require.EqualValues(t, true, accFound)
	require.EqualValues(t, len(accVestings.Vestings), len(storedAccVestings.Vestings))
	require.EqualValues(t, address.String(), storedAccVestings.Address)
	require.EqualValues(t, accVestings.DelegableAddress, storedAccVestings.DelegableAddress)

	require.EqualValues(t, accVestings.Vestings, storedAccVestings.Vestings)

}

func setupApp(initBlock int64) (*app.App, sdk.Context) {
	app := app.Setup(false)
	header := tmproto.Header{}
	header.Height = initBlock
	ctx := app.BaseApp.NewContext(false, header)
	return app, ctx
}

func addCoinsToCfevestingModule(vested uint64, ctx sdk.Context, app *app.App) string {
	return addCoinsToModuleByName(vested, types.ModuleName, ctx, app)
}

func withdrawDelegatorReward(t *testing.T, ctx sdk.Context, app *app.App, delegatorAddress sdk.AccAddress, delegableAddress sdk.AccAddress,
	validatorAddress sdk.ValAddress, delegatorAccountAmountBefore int64, delegableAccountAmountBefore int64, delegatorAccountAmountAfter int64, delegableAccountAmountAfter int64) {

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	verifyAccountBalance(t, app, ctx, delegatorAddress, sdk.NewInt(delegatorAccountAmountBefore))
	verifyAccountBalance(t, app, ctx, delegableAddress, sdk.NewInt(delegableAccountAmountBefore))

	msgUn := types.MsgWithdrawDelegatorReward{DelegatorAddress: delegatorAddress.String(), ValidatorAddress: validatorAddress.String()}
	_, error := msgServer.WithdrawDelegatorReward(msgServerCtx, &msgUn)
	require.EqualValues(t, nil, error)

	verifyAccountBalance(t, app, ctx, delegatorAddress, sdk.NewInt(delegatorAccountAmountAfter))
	verifyAccountBalance(t, app, ctx, delegableAddress, sdk.NewInt(delegableAccountAmountAfter))
}

func getVestings(t *testing.T, ctx sdk.Context, app *app.App, address sdk.AccAddress) *types.QueryVestingResponse {
	msgServerCtx := sdk.WrapSDKContext(ctx)
	vestingData, error := app.CfevestingKeeper.Vesting(msgServerCtx, &types.QueryVestingRequest{Address: address.String()})
	require.EqualValues(t, nil, error)
	return vestingData
}

func verifyVestingResponseWithStoredAccountVestings(t *testing.T, ctx sdk.Context, app *app.App, response *types.QueryVestingResponse, address sdk.AccAddress, currentHeight int64, delegationAllowed bool) {
	accVests, found := app.CfevestingKeeper.GetAccountVestings(ctx, address.String())
	require.EqualValues(t, true, found)
	verifyVestingResponse(t, response, accVests, currentHeight, delegationAllowed)
}

func verifyVestingResponse(t *testing.T, response *types.QueryVestingResponse, accVestings types.AccountVestings, currentHeight int64, delegationAllowed bool) {
	require.EqualValues(t, len(accVestings.Vestings), len(response.Vestings))
	require.EqualValues(t, accVestings.DelegableAddress, response.DelegableAddress)

	for _, vesting := range accVestings.Vestings {
		found := false
		for _, vestingInfo := range response.Vestings {
			if vesting.Id == vestingInfo.Id {
				require.EqualValues(t, vesting.VestingType, vestingInfo.VestingType)
				require.EqualValues(t, testutils.GetExpectedWithdrawableForVesting(*vesting, currentHeight).String(), response.Vestings[0].Withdrawable)
				require.EqualValues(t, vesting.VestingStart, vestingInfo.VestingStart)
				require.EqualValues(t, vesting.LockEnd, vestingInfo.LockEnd)
				require.EqualValues(t, vesting.VestingEnd, vestingInfo.VestingEnd)
				require.EqualValues(t, "uc4e", response.Vestings[0].Vested.Denom)
				require.EqualValues(t, vesting.Vested, response.Vestings[0].Vested.Amount)
				require.EqualValues(t, vesting.LastModificationVested.Sub(vesting.LastModificationWithdrawn).String(), response.Vestings[0].CurrentVestedAmount)
				require.EqualValues(t, delegationAllowed, response.Vestings[0].DelegationAllowed)
				require.EqualValues(t, vesting.Sent.String(), response.Vestings[0].SentAmount)

				found = true
			}
		}
		require.True(t, found, "not found vesting nfo with Id: "+strconv.FormatInt(int64(vesting.Id), 10))
	}
}
