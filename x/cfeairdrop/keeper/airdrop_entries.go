package keeper

import (
	"fmt"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	"github.com/tendermint/tendermint/libs/log"
	"golang.org/x/exp/slices"
	"strconv"
)

const FeegrantDenom = "uc4e"

func (k Keeper) AddUsersEntries(ctx sdk.Context, owner string, campaignId uint64, claimRecords []*types.ClaimRecord) error {
	logger := ctx.Logger().With("add user entries", "owner", owner, "campaignId", campaignId)

	campaign, validationResult := k.ValidateAddUsersEntries(logger, ctx, owner, campaignId)
	if validationResult != nil {
		return validationResult
	}

	usersEntries, claimRecordsAmountSum, err := k.validateClaimRecords(logger, ctx, campaign, claimRecords)
	if err != nil {
		return err
	}

	feegrantFeesSum := calculateFeegrantFeesSum(campaign.FeegrantAmount, int64(len(claimRecords)))
	feesAndClaimRecordsAmountSum := claimRecordsAmountSum.Add(feegrantFeesSum...)

	ownerAcc, _ := sdk.AccAddressFromBech32(owner)
	allBalances := k.bankKeeper.GetAllBalances(ctx, ownerAcc)
	if !allBalances.IsAllGTE(feesAndClaimRecordsAmountSum) {
		logger.Debug("airdrop entry owner balance is too small")
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, fmt.Sprintf("add campaign entries - owner balance is too small (%s < %s)", allBalances, feesAndClaimRecordsAmountSum))
	}

	if slices.Contains(types.GetWhitelistedVestingAccounts(), owner) {
		if err = k.AddClaimRecordsFromWhitelistedVestingAccount(ctx, owner, feesAndClaimRecordsAmountSum); err != nil {
			return err
		}
	}

	err = k.setupAndSendFeegrant(ctx, ownerAcc, campaign, feegrantFeesSum, claimRecords)
	if err != nil {
		return err
	}

	if err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, ownerAcc, types.ModuleName, claimRecordsAmountSum); err != nil {
		return err
	}

	k.IncrementCampaignTotalAmount(ctx, types.CampaignTotalAmount{
		CampaignId: campaignId,
		Amount:     claimRecordsAmountSum,
	})
	k.IncrementCampaignAmountLeft(ctx, types.CampaignAmountLeft{
		CampaignId: campaignId,
		Amount:     claimRecordsAmountSum,
	})
	for _, userEntry := range usersEntries {
		k.SetUserEntry(ctx, *userEntry)
	}

	return nil
}

func calculateFeegrantFeesSum(feegrantAmount sdk.Int, claimRecordsNumber int64) (feesSum sdk.Coins) {
	if feegrantAmount.GT(sdk.ZeroInt()) {
		feesSum = feesSum.Add(sdk.NewCoin(FeegrantDenom, feegrantAmount.MulRaw(claimRecordsNumber)))
	}
	return
}

func (k Keeper) ValidateAddUsersEntries(logger log.Logger, ctx sdk.Context, owner string, campaignId uint64) (*types.Campaign, error) {
	campaign, err := k.ValidateCampaignExists(logger, campaignId, ctx)
	if err != nil {
		return nil, err
	}

	if err = ValidateOwner(logger, campaign, owner); err != nil {
		return nil, err
	}

	if err = ValidateCampaignEnabled(logger, campaign); err != nil {
		return nil, err
	}

	if err = ValidateCampaignEnded(logger, ctx, campaign); err != nil {
		return nil, err
	}
	return &campaign, nil
}

func validateClaimRecord(logger log.Logger, claimRecord *types.ClaimRecord) error {
	if claimRecord.Address == "" {
		logger.Debug("claim record empty address")
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "add campaign entries - claim record empty address")
	}
	if !claimRecord.Amount.IsAllPositive() {
		logger.Debug("claim record must has at least one coin")
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "add campaign entries - claim record must has at least one coin")
	}
	return nil
}

func (k Keeper) validateInitialClaimClaimable(logger log.Logger, initialClaimFreeAmount sdk.Int, campaignMissions []types.Mission, claimRecord *types.ClaimRecord) error {
	allMissionsAmountSum := sdk.NewCoins()
	for _, mission := range campaignMissions {
		for _, amount := range claimRecord.Amount {
			allMissionsAmountSum = allMissionsAmountSum.Add(sdk.NewCoin(amount.Denom, mission.Weight.Mul(sdk.NewDecFromInt(amount.Amount)).TruncateInt()))
		}
	}
	initialClaimAmount := claimRecord.Amount.Sub(allMissionsAmountSum)

	for _, coin := range claimRecord.Amount {
		if initialClaimAmount.AmountOf(coin.Denom).LT(initialClaimFreeAmount) {
			logger.Debug("airdrop entry amount < campaign initial claim free amount", "amount", coin.Amount, "initialClaimFreeAmount", initialClaimFreeAmount)
			return sdkerrors.Wrapf(c4eerrors.ErrParam, "add campaign entries - claim amount %s < campaign initial claim free amount (%s)", initialClaimAmount.AmountOf(coin.Denom), initialClaimFreeAmount.String())
		}
	}

	return nil
}

func (k Keeper) validateClaimRecords(logger log.Logger, ctx sdk.Context, campaign *types.Campaign, claimRecords []*types.ClaimRecord) (usersEntries []*types.UserEntry, entriesAmountSum sdk.Coins, err error) {
	allCampaignMissions, _ := k.AllMissionForCampaign(ctx, campaign.Id)

	for i, claimRecord := range claimRecords {
		logger = logger.With("claimRecord index", i)

		if err = validateClaimRecord(logger, claimRecord); err != nil {
			return nil, nil, wrapClaimRecordIndex(err, i)
		}

		if err = k.validateInitialClaimClaimable(logger, campaign.InitialClaimFreeAmount, allCampaignMissions, claimRecord); err != nil {
			return nil, nil, wrapClaimRecordIndex(err, i)
		}

		entriesAmountSum = entriesAmountSum.Add(claimRecord.Amount...)

		userEntry, err := k.addUserEntry(ctx, campaign.Id, claimRecord.Address, claimRecord.Amount)
		if err != nil {
			return nil, nil, wrapClaimRecordIndex(err, i)
		}
		usersEntries = append(usersEntries, userEntry)
	}
	return
}

func wrapClaimRecordIndex(err error, index int) error {
	return sdkerrors.Wrap(err, fmt.Sprintf("claim records index %d", index))
}

func (k Keeper) addUserEntry(ctx sdk.Context, campaignId uint64, address string, allCoins sdk.Coins) (*types.UserEntry, error) {
	userEntry, found := k.GetUserEntry(ctx, address)
	if !found {
		userEntry = types.UserEntry{Address: address}
	}
	if userEntry.HasCampaign(campaignId) {
		return nil, sdkerrors.Wrapf(c4eerrors.ErrAlreadyExists, "campaignId %d already exists for address: %s", campaignId, address)
	}
	userEntry.ClaimRecords = append(userEntry.ClaimRecords, &types.ClaimRecord{CampaignId: campaignId, Amount: allCoins})
	return &userEntry, nil
}

func (k Keeper) DeleteClaimRecord(ctx sdk.Context, owner string, campaignId uint64, userAddress string) error {
	logger := ctx.Logger().With("delete claim record", "owner", owner, "campaignId", campaignId, "userAddress", userAddress)

	userEntry, claimRecordAmount, validationResult := k.ValidateRemoveClaimRecord(logger, ctx, owner, campaignId, userAddress)
	if validationResult != nil {
		return validationResult
	}

	k.SetUserEntry(ctx, userEntry)
	k.DecrementCampaignTotalAmount(ctx, campaignId, claimRecordAmount)

	return nil
}

func (k Keeper) setupAndSendFeegrant(ctx sdk.Context, ownerAcc sdk.AccAddress, campaign *types.Campaign, feegrantFeesSum sdk.Coins, claimRecords []*types.ClaimRecord) error {
	if campaign.FeegrantAmount.GT(sdk.ZeroInt()) {
		acc := k.NewModuleAccountSet(ctx, campaign.Id)
		if err := k.bankKeeper.SendCoins(ctx, ownerAcc, acc.GetAddress(), feegrantFeesSum); err != nil {
			return err
		}
		if err := k.grantFeeAllowanceToAllClaimRecords(ctx, acc.GetAddress(), claimRecords, sdk.NewCoins(sdk.NewCoin(FeegrantDenom, campaign.FeegrantAmount))); err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) grantFeeAllowanceToAllClaimRecords(ctx sdk.Context, moduleAddress sdk.AccAddress, airdropEntries []*types.ClaimRecord, grantAmount sdk.Coins) error {
	basicAllowance, err := codectypes.NewAnyWithValue(&feegranttypes.BasicAllowance{
		SpendLimit: grantAmount,
	})
	if err != nil {
		return err
	}

	allowedMsgAllowance := feegranttypes.AllowedMsgAllowance{
		Allowance:       basicAllowance,
		AllowedMessages: []string{"/chain4energy.c4echain.cfeairdrop.MsgInitialClaim"},
	}

	for _, claimRecord := range airdropEntries {
		granteeAddress, _ := sdk.AccAddressFromBech32(claimRecord.Address)
		existingFeeAllowance, _ := k.feeGrantKeeper.GetAllowance(ctx, moduleAddress, granteeAddress)
		if existingFeeAllowance == nil {
			if err = k.feeGrantKeeper.GrantAllowance(ctx, moduleAddress, granteeAddress, &allowedMsgAllowance); err != nil {
				return err
			}
		}
	}

	return nil
}

func (k Keeper) revokeFeeAllowance(ctx sdk.Context, granter sdk.Address, grantee sdk.AccAddress) error {
	keeper, _ := (k.feeGrantKeeper).(feegrantkeeper.Keeper)
	feegrantMsgServer := feegrantkeeper.NewMsgServerImpl(keeper)
	msg := feegranttypes.MsgRevokeAllowance{
		Granter: granter.String(),
		Grantee: grantee.String(),
	}
	_, err := feegrantMsgServer.RevokeAllowance(sdk.WrapSDKContext(ctx), &msg)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) NewModuleAccountSet(ctx sdk.Context, campaignId uint64) *authtypes.ModuleAccount {
	moduleAddressName, accountAddr := FeegrantAccountAddress(campaignId)
	macc := &authtypes.ModuleAccount{
		BaseAccount: &authtypes.BaseAccount{
			Address: accountAddr.String(),
		},
		Name: moduleAddressName,
	}
	k.accountKeeper.SetAccount(ctx, k.accountKeeper.NewAccount(ctx, macc))
	return macc
}

func FeegrantAccountAddress(campaignId uint64) (string, sdk.AccAddress) {
	moduleAddressName := "fee-grant-" + strconv.FormatUint(campaignId, 10)
	return moduleAddressName, authtypes.NewModuleAddress(moduleAddressName)
}

func (k Keeper) AddClaimRecordsFromWhitelistedVestingAccount(ctx sdk.Context, from string, amount sdk.Coins) error {
	fromAddress, err := sdk.AccAddressFromBech32(from)
	if err != nil {
		k.Logger(ctx).Error("delete user airdrop entry owner parsing error", "from", from, "error", err.Error())
		return sdkerrors.Wrap(c4eerrors.ErrParsing, sdkerrors.Wrapf(err, "delete user airdrop entry - owner parsing error: %s", from).Error())
	}
	ak := k.accountKeeper
	bk := k.bankKeeper
	fromAcc := ak.GetAccount(ctx, fromAddress)
	if fromAcc == nil {
		return sdkerrors.Wrapf(c4eerrors.ErrNotExists, "account %s doesn't exist", fromAddress)
	}

	spendableCoins := bk.SpendableCoins(ctx, fromAddress)
	if spendableCoins.IsAllGTE(amount) {
		return nil
	}

	vestingAcc := fromAcc.(*vestingtypes.ContinuousVestingAccount)

	balance := bk.GetAllBalances(ctx, fromAddress)
	if balance.IsAllGT(vestingAcc.OriginalVesting) {
		balanceWithoutOriginalVesting := balance.Sub(vestingAcc.OriginalVesting)
		amount = amount.Sub(balanceWithoutOriginalVesting)
	}

	lockedCoins := vestingAcc.LockedCoins(ctx.BlockTime()).Add(vestingAcc.DelegatedVesting...)
	spendableFromVesting := vestingAcc.OriginalVesting.Sub(lockedCoins)
	amountDiffFree := amount.Sub(spendableFromVesting)

	for _, coin := range amount {
		lockedPercentage := vestingAcc.OriginalVesting.AmountOf(coin.Denom).ToDec().Quo(lockedCoins.AmountOf(coin.Denom).ToDec())
		originalVestingDiff := amountDiffFree.AmountOf(coin.Denom).ToDec().Mul(lockedPercentage).TruncateInt()
		vestingAcc.OriginalVesting = vestingAcc.OriginalVesting.Sub(sdk.NewCoins(sdk.NewCoin(coin.Denom, originalVestingDiff)))
	}

	ak.SetAccount(ctx, vestingAcc)
	spendableCoins = bk.SpendableCoins(ctx, fromAddress)
	return nil
}

func (k Keeper) ValidateRemoveClaimRecord(logger log.Logger, ctx sdk.Context, owner string, campaignId uint64, userAddress string) (types.UserEntry, sdk.Coins, error) {
	campaign, err := k.ValidateCampaignExists(logger, campaignId, ctx)
	if err != nil {
		return types.UserEntry{}, nil, err
	}

	if err = ValidateCampaignTypeIsTeamdrop(logger, campaign); err != nil {
		return types.UserEntry{}, nil, err
	}

	if err = ValidateOwner(logger, campaign, owner); err != nil {
		return types.UserEntry{}, nil, err
	}

	userEntry, err := k.ValidateUserEntry(logger, ctx, userAddress)
	if err != nil {
		return types.UserEntry{}, nil, err
	}

	claimRecordAmount, err := ValidateClaimRecordExists(logger, userEntry, campaignId)
	if err != nil {
		return types.UserEntry{}, nil, err
	}

	return userEntry, claimRecordAmount, nil
}

func (k Keeper) ValidateUserEntry(log log.Logger, ctx sdk.Context, userAddress string) (types.UserEntry, error) {
	userEntry, found := k.GetUserEntry(
		ctx,
		userAddress,
	)
	if !found {
		log.Debug("delete user airdrop entry userEntry doesn't exist")
		return types.UserEntry{}, sdkerrors.Wrapf(c4eerrors.ErrParsing, "delete user airdrop entry - userEntry doesn't exist")
	}

	return userEntry, nil
}

func ValidateClaimRecordExists(log log.Logger, userEntry types.UserEntry, campaignId uint64) (claimRecordAmount sdk.Coins, err error) {
	claimRecordFound := false
	for i, claimRecord := range userEntry.ClaimRecords {
		if claimRecord.CampaignId == campaignId {
			claimRecordFound = true
			claimRecordAmount = claimRecord.Amount
			userEntry.ClaimRecords = append(userEntry.ClaimRecords[:i], userEntry.ClaimRecords[i+1:]...)
			break
		}
	}
	if !claimRecordFound {
		log.Debug("delete user airdrop entry airdrop entry doesn't exist", "campaignId", campaignId)
		return nil, sdkerrors.Wrapf(c4eerrors.ErrParsing, "delete user airdrop entry -  campaign id %d airdrop entry doesn't exist", campaignId)
	}

	return claimRecordAmount, nil
}

func ValidateCampaignTypeIsTeamdrop(log log.Logger, campaign types.Campaign) error {
	if campaign.CampaignType != types.CampaignTeamdrop {
		log.Debug("campaign must be of TEAMDROP type to be able to delete its entries")
		return sdkerrors.Wrap(sdkerrors.ErrInvalidType, "ampaign must be of TEAMDROP type to be able to delete its entries")
	}
	return nil
}
