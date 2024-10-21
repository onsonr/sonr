package types

import sdkerrors "cosmossdk.io/errors"

var (
	ErrInvalidGenesisState  = sdkerrors.Register(ModuleName, 100, "invalid genesis state")
	ErrInvalidServiceOrigin = sdkerrors.Register(ModuleName, 200, "invalid service origin")
	ErrUnrecognizedService  = sdkerrors.Register(ModuleName, 201, "unrecognized service")
)
