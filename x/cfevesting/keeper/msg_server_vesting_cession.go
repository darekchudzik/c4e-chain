package keeper

import (
	"context"
	"github.com/armon/go-metrics"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) VestingCession(goCtx context.Context, msg *types.MsgVestingCession) (*types.MsgVestingCessionResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "vesting cession message")
	ctx := sdk.UnwrapSDKContext(goCtx)

	defer func() {
		if msg.Amount.IsInt64() {
			telemetry.SetGaugeWithLabels(
				[]string{"tx", "msg", types.ModuleName, msg.Type()},
				float32(msg.Amount.Int64()),
				[]metrics.Label{telemetry.NewLabel("denom", k.Keeper.Denom(ctx))},
			)
		}
	}()

	err := k.Keeper.VestingCesion(ctx, msg.FromAddress, msg.ToAddress, msg.Amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgVestingCessionResponse{}, nil
}
