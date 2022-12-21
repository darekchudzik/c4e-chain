package keeper

import (
	"context"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) VestingCession(goCtx context.Context, msg *types.MsgVestingCession) (*types.MsgVestingCessionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.newVestingAccount(ctx, msg.ReceiverAddress, amount, vt.Free,
		vestingPool.LockEnd, vestingPool.LockEnd)
	if err == nil {
		k.AppendVestingAccount(ctx, types.VestingAccount{Address: msg.ReceiverAddress})

		eventErr := ctx.EventManager().EmitTypedEvent(&types.VestingCession{
			OwnerAddress: msg.Creator,
			Address:      msg.ReceiverAddress,
			Amount:       amount.String() + k.Denom(ctx),
		})
		if eventErr != nil {
			k.Logger(ctx).Error("new vesting account from vesting pool emit event error", "error", err.Error())
		}
	}
	_ = ctx

	return &types.MsgVestingCessionResponse{}, nil
}
