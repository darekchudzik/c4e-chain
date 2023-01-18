package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/telemetry"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateAirdropCampaign(goCtx context.Context, msg *types.MsgCreateAirdropCampaign) (*types.MsgCreateAirdropCampaignResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "create aidrop campaign message")
	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper

	if err := keeper.CreateAidropCampaign(
		ctx,
		msg.Owner,
		msg.Name,
		msg.Description,
		msg.Denom,
		time.Unix(msg.StartTime, 0),
		time.Unix(msg.EndTime, 0),
		msg.LockupPeriod,
		msg.VestingPeriod,
	); err != nil {
		return nil, err
	}

	return &types.MsgCreateAirdropCampaignResponse{}, nil
}

func (k msgServer) CloseAirdropCampaign(goCtx context.Context, msg *types.MsgCloseAirdropCampaign) (*types.MsgCloseAirdropCampaignResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	keeper := k.Keeper
	if err := keeper.CloseAirdropCampaign(
		ctx,
		msg.Owner,
		msg.CampaignId,
		msg.Burn,
		msg.CommunityPoolSend,
	); err != nil {
		return nil, err
	}

	return &types.MsgCloseAirdropCampaignResponse{}, nil
}

func (k msgServer) StartAirdropCampaign(goCtx context.Context, msg *types.MsgStartAirdropCampaign) (*types.MsgStartAirdropCampaignResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	keeper := k.Keeper
	if err := keeper.StartAirdropCampaign(
		ctx,
		msg.Owner,
		msg.CampaignId,
	); err != nil {
		return nil, err
	}

	return &types.MsgStartAirdropCampaignResponse{}, nil
}
