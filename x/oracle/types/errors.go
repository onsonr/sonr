package types

import sdkerrors "cosmossdk.io/errors"

var (
	ErrInvalidGenesisState = sdkerrors.Register(ModuleName, 1, "invalid genesis state")
)
