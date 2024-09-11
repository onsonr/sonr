package types

import sdkerrors "cosmossdk.io/errors"

var (
	ErrInvalidGenesisState     = sdkerrors.Register(ModuleName, 100, "invalid genesis state")
	ErrInvalidETHAddressFormat = sdkerrors.Register(ModuleName, 200, "invalid ETH address format")
	ErrInvalidBTCAddressFormat = sdkerrors.Register(ModuleName, 201, "invalid BTC address format")
	ErrInvalidIDXAddressFormat = sdkerrors.Register(ModuleName, 202, "invalid IDX address format")
	ErrInvalidOriginFormat     = sdkerrors.Register(ModuleName, 203, "invalid origin format")
	ErrInvalidServiceOrigin    = sdkerrors.Register(ModuleName, 300, "invalid service origin")
	ErrUnrecognizedService     = sdkerrors.Register(ModuleName, 301, "unrecognized service")
	ErrUnsupportedKeyEncoding  = sdkerrors.Register(ModuleName, 400, "unsupported key encoding")
	ErrUnsopportedChainCode    = sdkerrors.Register(ModuleName, 401, "unsupported chain code")
	ErrUnsupportedKeyCurve     = sdkerrors.Register(ModuleName, 402, "unsupported key curve")
	ErrInvalidSignature        = sdkerrors.Register(ModuleName, 403, "invalid signature")
)
