package keeper

import (
	"fmt"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/libs/log"
	"time"
)

func (k Keeper) CreateAidropCampaign(ctx sdk.Context, owner string, name string, description string, campaignType types.CampaignType, feegrantAmount *sdk.Int, initialClaimFreeAmount *sdk.Int, startTime *time.Time,
	endTime *time.Time, lockupPeriod *time.Duration, vestingPeriod *time.Duration) error {
	k.Logger(ctx).Debug("create aidrop campaign", "owner", owner, "name", name, "description", description,
		"startTime", startTime, "endTime", endTime, "lockupPeriod", lockupPeriod, "vestingPeriod", vestingPeriod)
	if name == "" {
		k.Logger(ctx).Debug("create airdrop campaign empty campaign name")
		return sdkerrors.Wrap(c4eerrors.ErrParam, "create airdrop campaign - empty campaign name error")
	}
	if description == "" {
		k.Logger(ctx).Error("create airdrop campaign empty campaign description")
		return sdkerrors.Wrap(c4eerrors.ErrParam, "create airdrop campaign - empty campaign description error")
	}
	if startTime == nil {
		k.Logger(ctx).Error("create airdrop campaign start time is nil")
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "create airdrop campaign - start time is nil error")
	}
	if startTime.Before(ctx.BlockTime()) {
		k.Logger(ctx).Error("create airdrop campaign start time in the past", "startTime", startTime)
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "create airdrop campaign - start time in the past error (%s < %s)", startTime, ctx.BlockTime())
	}
	if endTime == nil {
		k.Logger(ctx).Error("create airdrop campaign end time is nil")
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "create airdrop campaign - end time is nil error")
	}
	if startTime.After(*endTime) {
		k.Logger(ctx).Error("create airdrop campaign start time is after end time", "startTime", startTime, "endTime", endTime)
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "create airdrop campaign - start time is after end time error (%s > %s)", startTime, endTime)
	}
	if initialClaimFreeAmount.IsNil() {
		zeroInt := sdk.ZeroInt()
		initialClaimFreeAmount = &zeroInt
	}
	if initialClaimFreeAmount.IsNegative() {
		k.Logger(ctx).Error("create airdrop campaign initial claim free amount cannot be negative", "initialClaimFreeAmount", initialClaimFreeAmount)
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "create airdrop campaign - initial claim free amount (%s) cannot be negative", initialClaimFreeAmount.String())
	}
	if feegrantAmount.IsNil() {
		zeroInt := sdk.ZeroInt()
		feegrantAmount = &zeroInt
	}
	if feegrantAmount.IsNegative() {
		k.Logger(ctx).Error("create airdrop campaign initial feegrant amount cannot be negative", "initialClaimFreeAmount", feegrantAmount)
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "create airdrop campaign - feegrant amount (%s) cannot be negative", feegrantAmount.String())
	}
	_, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		k.Logger(ctx).Error("create vesting account owner parsing error", "owner", owner, "error", err.Error())
		return sdkerrors.Wrap(c4eerrors.ErrParsing, sdkerrors.Wrapf(err, "create vesting account - owner parsing error: %s", owner).Error())
	}

	campaign := types.Campaign{
		Owner:                  owner,
		Name:                   name,
		Description:            description,
		CampaignType:           campaignType,
		FeegrantAmount:         *feegrantAmount,
		InitialClaimFreeAmount: *initialClaimFreeAmount,
		Enabled:                false,
		StartTime:              *startTime,
		EndTime:                *endTime,
		LockupPeriod:           *lockupPeriod,
		VestingPeriod:          *vestingPeriod,
	}

	campaignId := k.AppendNewCampaign(ctx, campaign)
	missionInitial := types.NewInitialMission(campaignId)
	k.AppendNewMission(ctx, campaignId, *missionInitial)
	return nil
}

func (k Keeper) EditCampaign(ctx sdk.Context, owner string, campaignId uint64, name string, description string, startTime *time.Time,
	endTime *time.Time, lockupPeriod *time.Duration, vestingPeriod *time.Duration) error {
	k.Logger(ctx).Debug("edit airdrop campaign", "owner", owner, "name", name, "description", description,
		"startTime", startTime, "endTime", endTime, "lockupPeriod", lockupPeriod, "vestingPeriod", vestingPeriod)
	campaign, found := k.GetCampaign(
		ctx,
		campaignId,
	)
	if !found {
		k.Logger(ctx).Error("edit airdrop campaign campaign doesn't exist", "campaignId", campaignId)
		return sdkerrors.Wrapf(c4eerrors.ErrParsing, "edit airdrop campaign -  campaign with id %d doesn't exist", campaignId)
	}
	if campaign.Enabled == true {
		k.Logger(ctx).Error("edit airdrop campaign campaign doesn't exist", "campaignId", campaignId)
		return sdkerrors.Wrapf(c4eerrors.ErrParsing, "edit airdrop campaign -  campaign with id %d doesn't exist", campaignId)
	}
	if campaign.EndTime.Before(ctx.BlockTime()) {
		k.Logger(ctx).Error("edit airdrop campaign campaign doesn't exist", "campaignId", campaignId)
		return sdkerrors.Wrapf(c4eerrors.ErrParsing, "edit airdrop campaign -  campaign with id %d doesn't exist", campaignId)
	}
	if campaign.Owner != owner {
		k.Logger(ctx).Error("edit airdrop campaign you are not the owner of this campaign", "campaignId", campaignId)
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "edit airdrop campaign - you are not the owner of campaign with id %d", campaignId)
	}
	if name != "" {
		campaign.Name = name
	}
	if description != "" {
		campaign.Description = description
	}
	if startTime != nil {
		if endTime != nil {
			if startTime.After(*endTime) {
				k.Logger(ctx).Error("edit airdrop campaign start time is after end time", "startTime", startTime, "endTime", endTime)
				return sdkerrors.Wrapf(c4eerrors.ErrParam, "edit airdrop campaign - start time is after end time error (%s > %s)", startTime, endTime)
			}
			campaign.EndTime = *endTime
		} else {
			if startTime.After(campaign.EndTime) {
				k.Logger(ctx).Error("edit airdrop campaign start time is after end time", "startTime", startTime, "endTime", endTime)
				return sdkerrors.Wrapf(c4eerrors.ErrParam, "cedit airdrop campaign - start time is after end time error (%s > %s)", startTime, endTime)
			}
		}
		campaign.StartTime = *startTime
	}
	if vestingPeriod != nil {
		campaign.VestingPeriod = *vestingPeriod
	}
	if lockupPeriod != nil {
		campaign.LockupPeriod = *lockupPeriod
	}
	k.SetCampaign(ctx, campaign)
	return nil
}

func (k Keeper) CloseCampaign(ctx sdk.Context, owner string, campaignId uint64, campaignCloseAction types.CampaignCloseAction) error {
	k.Logger(ctx).Debug("close airdrop campaign", "owner", owner, "campaignId", campaignId, "campaignCloseAction", campaignCloseAction)
	campaign, found := k.GetCampaign(ctx, campaignId)
	if !found {
		k.Logger(ctx).Error("close airdrop campaign campaign campaign not found", "campaignId", campaignId)
		return sdkerrors.Wrapf(c4eerrors.ErrNotExists, "close airdrop campaign - campaign with id %d not found error", campaignId)
	}
	if campaign.EndTime.After(ctx.BlockTime()) {
		k.Logger(ctx).Debug("close airdrop campaign campaign is not over yet", "startTime", campaign.StartTime)
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "close airdrop campaign - campaign with id %d campaign is not over yet (endtime - %s < %s)", campaignId, campaign.EndTime, ctx.BlockTime())
	}
	if campaign.Owner != owner {
		k.Logger(ctx).Error("close airdrop campaign you are not the owner of this campaign")
		return sdkerrors.Wrap(sdkerrors.ErrorInvalidSigner, "close airdrop campaign - you are not the owner error")
	}
	if campaign.Enabled == false {
		k.Logger(ctx).Error("close airdrop campaign campaign is already closed")
		return sdkerrors.Wrap(types.ErrCampaignDisabled, fmt.Sprintf("close airdrop campaign - campaign with id %d is already closed or have not started yet error", campaignId))
	}
	campaign.Enabled = false
	k.SetCampaign(ctx, campaign)
	return nil
}

func (k Keeper) StartCampaign(ctx sdk.Context, owner string, campaignId uint64) error {
	logger := ctx.Logger().With("start airdrop campaign", "owner", owner, "campaignId", campaignId)

	campaign, validationResult := k.ValidateStartCampaign(logger, campaignId, ctx, owner)
	if validationResult != nil {
		return validationResult
	}

	campaign.Enabled = true
	k.SetCampaign(ctx, campaign)
	return nil
}

func (k Keeper) RemoveCampaign(ctx sdk.Context, owner string, campaignId uint64) error {
	k.Logger(ctx).Debug("Remove airdrop campaign", "owner", owner, "campaignId", campaignId)
	validationResult := k.ValidateRemove(ctx, owner, campaignId)
	if validationResult != nil {
		return validationResult
	}

	k.removeCampaign(ctx, campaignId)
	k.RemoveAllMissionForCampaign(ctx, campaignId)
	return nil
}

func (k Keeper) ValidateStartCampaign(logger log.Logger, campaignId uint64, ctx sdk.Context, owner string) (types.Campaign, error) {
	campaign, err := k.ValidateCampaignExists(logger, campaignId, ctx)
	if err != nil {
		return types.Campaign{}, err
	}

	if err = ValidateOwner(logger, campaign, owner); err != nil {
		return types.Campaign{}, err
	}

	if err = ValidateCampaignEnabled(logger, campaign); err != nil {
		return types.Campaign{}, err
	}

	if err = ValidateCampaignStart(ctx, campaign, logger); err != nil {
		return types.Campaign{}, err
	}

	return campaign, nil
}

func (k Keeper) ValidateRemove(ctx sdk.Context, owner string, campaignId uint64) error {
	logger := ctx.Logger().With("Remove campaign validation")

	campaign, err := k.ValidateCampaignExists(logger, campaignId, ctx)
	if err != nil {
		return err
	}

	if err = ValidateOwner(logger, campaign, owner); err != nil {
		return err
	}

	if err = ValidateCampaignEnabled(logger, campaign); err != nil {
		return err
	}

	return nil
}

func (k Keeper) ValidateCampaignExists(log log.Logger, campaignId uint64, ctx sdk.Context) (types.Campaign, error) {
	campaign, found := k.GetCampaign(ctx, campaignId)
	if !found {
		log.Debug("campaign not exist", "campaignId", campaignId)
		return types.Campaign{}, sdkerrors.Wrapf(c4eerrors.ErrNotExists, "campaign with id %d not found", campaignId)
	}
	return campaign, nil
}

func ValidateOwner(log log.Logger, campaign types.Campaign, owner string) error {
	if campaign.Owner != owner {
		log.Debug("you are not campaign owner")
		return sdkerrors.Wrap(sdkerrors.ErrorInvalidSigner, "you are not the campaign owner")
	}
	return nil
}

func ValidateCampaignEnabled(log log.Logger, campaign types.Campaign) error {
	if campaign.Enabled == true {
		log.Error("Campaign has already started")
		return sdkerrors.Wrap(c4eerrors.ErrAlreadyExists, "Campaign has already started")
	}
	return nil
}

func ValidateCampaignStart(ctx sdk.Context, campaign types.Campaign, logger log.Logger) error {
	if campaign.StartTime.Before(ctx.BlockTime()) {
		logger.Debug("Campaign start time in the past", "startTime", campaign.StartTime)
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "Campaign with id %d start time in the past error (%s < %s)", campaign.Id, campaign.StartTime, ctx.BlockTime())
	}
	return nil
}
