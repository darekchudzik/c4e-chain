package cli

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdVestingCession() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vesting-cession [receiver-address] [amount]",
		Short: "Broadcast message vesting-cession",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			fmt.Println("VESTING CESSION START")
			argReceiverAddress := args[0]
			argAmount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount must be a positive integer")
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgVestingCession(
				clientCtx.GetFromAddress().String(),
				argReceiverAddress,
				argAmount,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
