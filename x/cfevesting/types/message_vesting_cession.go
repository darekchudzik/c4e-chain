package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgVestingCession = "vesting_cession"

var _ sdk.Msg = &MsgVestingCession{}

func NewMsgVestingCession(fromAddress string, toAddress string, amount sdk.Int) *MsgVestingCession {
	return &MsgVestingCession{
		FromAddress: fromAddress,
		ToAddress:   toAddress,
		Amount:      amount,
	}
}

func (msg *MsgVestingCession) Route() string {
	return RouterKey
}

func (msg *MsgVestingCession) Type() string {
	return TypeMsgVestingCession
}

func (msg *MsgVestingCession) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgVestingCession) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVestingCession) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
