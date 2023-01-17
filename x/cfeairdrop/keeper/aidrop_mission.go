package keeper

import (
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) CompleteMission(ctx sdk.Context, campaignId uint64, missionId uint64, address string, isHook bool) error {
	_, mission, userAirdropEntries, err := k.missionFirstStep(ctx, "complete mission", campaignId, missionId, address, isHook)
	if err != nil {
		return err
	}
	userAirdropEntries, err = k.completeMission(ctx, false, mission, userAirdropEntries)
	if err != nil {
		return err
	}
	k.SetUserAirdropEntries(ctx, *userAirdropEntries)
	return nil
}

// CompleteMission triggers the completion of the mission and distribute the claimable portion of airdrop to the user
// the method fails if the mission has already been completed
func (k Keeper) completeMission(ctx sdk.Context, isInitialClaim bool, mission *types.Mission, userAirdropEntries *types.UserAirdropEntries) (*types.UserAirdropEntries, error) {
	campaignId := mission.CampaignId
	missionId := mission.Id
	address := userAirdropEntries.Address
	// check if the mission is already complted for the claim record
	if !isInitialClaim {
		if !userAirdropEntries.IsInitialMissionClaimed(campaignId) {
			k.Logger(ctx).Error("complete mission - initial mission not completed", "claimerAddress", address, "campaignId", campaignId, "missionId", missionId)
			return nil, sdkerrors.Wrapf(types.ErrMissionNotCompleted, "initial mission not completed: address %s, campaignId: %d, missionId: %d", address, campaignId, 0)
		}
	}

	if userAirdropEntries.IsMissionCompleted(campaignId, missionId) {
		k.Logger(ctx).Error("complete mission - mission already completed", "address", address, "campaignId", campaignId, "missionId", missionId)
		return nil, sdkerrors.Wrapf(types.ErrMissionCompleted, "mission already completed: address %s, campaignId: %d, missionId: %d", address, campaignId, missionId)
	}

	if err := userAirdropEntries.CompleteMission(campaignId, missionId); err != nil {
		k.Logger(ctx).Error("complete mission - cannot complete", "address", address, "campaignId", campaignId, "missionId", missionId)
		return nil, sdkerrors.Wrapf(types.ErrMissionCompletion, err.Error())
	}

	return userAirdropEntries, nil
}

func (k Keeper) claimMission(ctx sdk.Context, initialClaim bool, campaign *types.Campaign, mission *types.Mission, userAirdropEntries *types.UserAirdropEntries) (*types.UserAirdropEntries, error) {
	campaignId := mission.CampaignId
	missionId := mission.Id
	address := userAirdropEntries.Address

	if !userAirdropEntries.IsMissionCompleted(campaignId, missionId) {
		k.Logger(ctx).Error("claim mission - mission not completed", "address", address, "campaignId", campaignId, "missionId", missionId)
		return nil, sdkerrors.Wrapf(types.ErrMissionNotCompleted, "mission not completed: address %s, campaignId: %d, missionId: %d", address, campaignId, missionId)
	}

	if userAirdropEntries.IsMissionClaimed(campaignId, missionId) {
		k.Logger(ctx).Error("claim mission - mission already claimed", "address", address, "campaignId", campaignId, "missionId", missionId)
		return nil, sdkerrors.Wrapf(types.ErrMissionClaimed, "mission already claimed: address %s, campaignId: %d, missionId: %d", address, campaignId, missionId)
	}

	if err := userAirdropEntries.ClaimMission(campaignId, missionId); err != nil {
		k.Logger(ctx).Error("claim mission - cannot claime", "address", address, "campaignId", campaignId, "missionId", missionId)
		return nil, sdkerrors.Wrapf(types.ErrMissionClaiming, err.Error())
	}
	claimableAmount := sdk.ZeroInt()
	if mission.MissionType == types.MissionInitialClaim {
		missions, _ := k.AllMissionForCampaign(ctx, campaignId)
		airdropEntry := userAirdropEntries.GetAidropEntry(campaignId)
		amountSum := sdk.ZeroInt()
		for _, mission := range missions {
			amountSum.Add(mission.Weight.Mul(sdk.NewDecFromInt(airdropEntry.Amount)).TruncateInt())
		}
		claimableAmount = airdropEntry.Amount.Sub(amountSum)
	} else {
		claimableAmount = userAirdropEntries.ClaimableFromMission(mission)
	}

	// TODO initial mission claim should have not waight but get whats left from ther missions

	claimable := sdk.NewCoins(sdk.NewCoin(campaign.Denom, claimableAmount))

	// calculate claimable after decay factor
	// decayInfo := k.GetParams(ctx).DecayInformation
	// claimable = decayInfo.ApplyDecayFactor(claimable, ctx.BlockTime())

	// check final claimable non-zero
	// if claimable.Empty() {
	// 	return types.ErrNoClaimable
	// }

	// decrease airdrop supply
	// airdropSupply.Amount = airdropSupply.Amount.Sub(claimable.AmountOf(airdropSupply.Denom))
	// if airdropSupply.Amount.IsNegative() {
	// 	return errors.Critical("airdrop supply is lower than total claimable")
	// }

	// send claimable to the user
	claimer, err := sdk.AccAddressFromBech32(userAirdropEntries.ClaimAddress)
	if err != nil {
		return nil, sdkerrors.Wrapf(c4eerrors.ErrParsing, "wrong claiming address %s: "+err.Error(), userAirdropEntries.ClaimAddress)
	}
	start := ctx.BlockTime().Add(campaign.LockupPeriod)
	end := start.Add(campaign.VestingPeriod)
	if err = k.SendToAirdropAccount(ctx, claimer, claimable, start.Unix(), end.Unix(), initialClaim); err != nil {
		return nil, sdkerrors.Wrapf(c4eerrors.ErrSendCoins, "send to claiming address %s error: "+err.Error(), userAirdropEntries.ClaimAddress)
	}
	return userAirdropEntries, nil

}

func (k Keeper) AddMissionToAirdropCampaign(ctx sdk.Context, owner string, campaignId uint64, name string, description string, missionType types.MissionType,
	weight sdk.Dec) error {
	k.Logger(ctx).Debug("add mission to airdrop campaign", "owner", owner, "campaignId", campaignId, "name", name,
		"description", description, "missionType", missionType, "weight", weight)

	if weight.GT(sdk.NewDec(1)) {
		k.Logger(ctx).Error("add mission to airdrop campaign weight is >= 1", "weight", weight)
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "add mission to airdrop campaign weight is >= 1 (%s > 1)", weight.String())
	}

	if name == "" {
		k.Logger(ctx).Error("add mission to airdrop campaign: empty name ")
		return sdkerrors.Wrap(c4eerrors.ErrParam, "add mission to airdrop campaign empty name")
	}

	if description == "" {
		k.Logger(ctx).Error("add mission to airdrop campaign: empty description ")
		return sdkerrors.Wrap(c4eerrors.ErrParam, "add mission to airdrop campaign empty description")
	}

	_, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		k.Logger(ctx).Error("add mission to airdrop campaign owner parsing error", "owner", owner, "error", err.Error())
		return sdkerrors.Wrap(c4eerrors.ErrParsing, sdkerrors.Wrapf(err, "add mission to airdrop campaign - owner parsing error: %s", owner).Error())
	}

	campaign, found := k.GetCampaign(ctx, campaignId)
	if !found {
		k.Logger(ctx).Error("add mission to airdrop campaign not found", "campaignId", campaignId)
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "add mission to airdrop campaign - campaign with id %d not found", campaignId)
	}

	if campaign.Owner != owner {
		k.Logger(ctx).Error("add mission to airdrop you are not the owner of this campaign", "campaignId", campaignId)
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "add mission to airdrop campaign - you are not the owner of campaign with id %d", campaignId)
	}

	_, found = k.GetMission(ctx, campaignId, 0)
	if !found {
		weight = sdk.ZeroDec()
		if missionType != types.MissionInitialClaim {
			k.Logger(ctx).Error("add mission to airdrop first mission must be of INITIAL_CLAIM type", "campaignId", campaignId, "missionType", missionType)
			return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "add mission to airdrop campaign - first mission must be of INITIAL_CLAIM type, campaign id: %d, mission type: %s", campaignId, missionType)
		}
	}

	_, weightSum := k.AllMissionForCampaign(ctx, campaignId)
	if weightSum.Add(weight).GT(sdk.NewDec(1)) {
		k.Logger(ctx).Error("add mission to airdrop campaign weight is >= 1", "weightSum", weightSum)
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "add mission to airdrop campaign campaign missions weight sum is >= 1 (%s > 1)", weightSum.String())
	}

	mission := types.Mission{
		CampaignId:  campaignId,
		Name:        name,
		Description: description,
		MissionType: missionType,
		Weight:      &weight,
	}

	k.AppendNewMission(ctx, campaignId, mission)
	return nil
}

func (k Keeper) missionFirstStep(ctx sdk.Context, log string, campaignId uint64, missionId uint64, claimerAddress string, isHook bool) (*types.Campaign, *types.Mission, *types.UserAirdropEntries, error) {
	campaignConfig, campaignFound := k.GetCampaign(ctx, campaignId)
	userAirdropEntries, found := k.GetUserAirdropEntries(ctx, claimerAddress)
	if !found {
		if isHook {
			k.Logger(ctx).Debug(log+" - claim record not found", "claimerAddress", claimerAddress)
			return nil, nil, nil, nil
		}
		k.Logger(ctx).Error(log+" - claim record not found", "claimerAddress", claimerAddress)
		return nil, nil, nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "claim record not found for claimerAddress %s", claimerAddress)
	}
	if !campaignFound {
		k.Logger(ctx).Error(log+" - camapign not found", "campaignId", campaignId)
		return nil, nil, nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "camapign not found: campaignId %d", campaignId)
	}
	if err := campaignConfig.IsEnabled(ctx.BlockTime()); err != nil {
		if isHook {
			k.Logger(ctx).Debug(log+" - camapign disabled", "campaignId", campaignId, "err", err)
			return nil, nil, nil, nil
		}
		k.Logger(ctx).Error(log+" - camapign disabled", "campaignId", campaignId, "err", err)
		return nil, nil, nil, sdkerrors.Wrapf(err, "campaign disabled - campaignId %d", campaignId)
	}
	k.Logger(ctx).Debug(log, "campaignId", campaignId, "missionId", missionId, "blockTime", ctx.BlockTime(), "campaigh start", campaignConfig.StartTime, "campaigh end", campaignConfig.EndTime)

	mission, missionFound := k.GetMission(ctx, campaignId, missionId)
	if !missionFound {
		k.Logger(ctx).Error(log+" - mission not found", "campaignId", campaignId, "missionId", missionId)
		return nil, nil, nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "mission not found - campaignId %d, missionId %d", campaignId, missionId)
	}
	k.Logger(ctx).Debug(log, "mission", mission)

	if !userAirdropEntries.HasCampaign(campaignId) {
		if isHook {
			k.Logger(ctx).Error(log+" - campaign record not found", "claimerAddress", claimerAddress, "campaignId", campaignId)
			return nil, nil, nil, nil
		}
		k.Logger(ctx).Error(log+" - campaign record not found", "claimerAddress", claimerAddress, "campaignId", campaignId)
		return nil, nil, nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "campaign record with id: %d not found for claimerAddress %s", campaignId, claimerAddress)
	}

	return &campaignConfig, &mission, &userAirdropEntries, nil
}

func (k Keeper) missionsWeightGreaterThan1(missions []types.Mission, newMissionWeight sdk.Dec) (bool, sdk.Dec) {
	weightSum := newMissionWeight
	for _, mission := range missions {
		weightSum = weightSum.Add(*mission.Weight)
	}
	if weightSum.GT(sdk.NewDec(1)) {
		return true, weightSum
	}
	return false, weightSum
}
