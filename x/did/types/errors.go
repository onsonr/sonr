package types

import sdkerrors "cosmossdk.io/errors"

var (
	ErrInvalidGenesisState     = sdkerrors.Register(ModuleName, 100, "invalid genesis state")
	ErrInvalidETHAddressFormat = sdkerrors.Register(ModuleName, 200, "invalid ETH address format")
	ErrInvalidBTCAddressFormat = sdkerrors.Register(ModuleName, 201, "invalid BTC address format")
	ErrInvalidIDXAddressFormat = sdkerrors.Register(ModuleName, 202, "invalid IDX address format")
)
