package types

import sdkerrors "cosmossdk.io/errors"

var (
	ErrInvalidGenesisState    = sdkerrors.Register(ModuleName, 100, "invalid genesis state")
	ErrInvalidSchema          = sdkerrors.Register(ModuleName, 200, "invalid schema")
	ErrVaultAssembly          = sdkerrors.Register(ModuleName, 201, "vault assembly")
	ErrControllerCreation     = sdkerrors.Register(ModuleName, 202, "failed to create controller")
	ErrUnsupportedKeyEncoding = sdkerrors.Register(ModuleName, 300, "unsupported key encoding")
	ErrUnsopportedChainCode   = sdkerrors.Register(ModuleName, 301, "unsupported chain code")
	ErrUnsupportedKeyCurve    = sdkerrors.Register(ModuleName, 302, "unsupported key curve")
)
