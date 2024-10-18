package types

import sdkerrors "cosmossdk.io/errors"

var (
	ErrInvalidGenesisState    = sdkerrors.Register(ModuleName, 100, "invalid genesis state")
	ErrUnsupportedKeyEncoding = sdkerrors.Register(ModuleName, 200, "unsupported key encoding")
	ErrUnsopportedChainCode   = sdkerrors.Register(ModuleName, 201, "unsupported chain code")
	ErrUnsupportedKeyCurve    = sdkerrors.Register(ModuleName, 202, "unsupported key curve")
)
