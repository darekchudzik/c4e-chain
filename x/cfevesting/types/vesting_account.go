// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cfevesting/vesting_account.proto

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (v VestingAccountTrace) Validate() error {
	_, err := sdk.AccAddressFromBech32(v.Address)
	return err
}
