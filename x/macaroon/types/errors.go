package types

import sdkerrors "cosmossdk.io/errors"

var (
	ErrInvalidGenesisState         = sdkerrors.Register(ModuleName, 100, "invalid genesis state")
	ErrUnauthorizedMacaroonToken   = sdkerrors.Register(ModuleName, 200, "unauthorized macaroon token")
	ErrInvalidMacaroonScopes       = sdkerrors.Register(ModuleName, 201, "invalid macaroon scopes")
	ErrInvalidMacaroonController   = sdkerrors.Register(ModuleName, 202, "invalid macaroon controller")
	ErrInvalidTransactionSignature = sdkerrors.Register(ModuleName, 203, "invalid supplied transaction signature")
)
