package types

import (
	fmt "fmt"
	"github.com/cosmos/cosmos-sdk/codec/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

// this line is used by starport scaffolding # genesis/types/import

var _ types.UnpackInterfacesMessage = GenesisState{}

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
		MinterState: MinterState{
			SequenceId:                  1,
			AmountMinted:                sdk.ZeroInt(),
			RemainderToMint:             sdk.ZeroDec(),
			LastMintBlockTime:           time.Now(),
			RemainderFromPreviousMinter: sdk.ZeroDec(),
		},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/Validate
	if err := gs.Params.Validate(); err != nil {
		return err
	}
	if err := gs.MinterState.Validate(); err != nil {
		return err
	}
	if !gs.Params.ContainsMinter(gs.MinterState.SequenceId) {
		return fmt.Errorf("cfeminter genesis validation error: minter state sequence id %d not found in minters", gs.MinterState.SequenceId)
	}
	return nil
}

func (gs GenesisState) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	return gs.Params.UnpackInterfaces(unpacker)
}
