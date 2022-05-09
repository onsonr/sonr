package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/registry module sentinel errors
var (
	ErrSample               = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidPacketTimeout = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion       = sdkerrors.Register(ModuleName, 1501, "invalid version")
	ErrControllerNotFound   = sdkerrors.Register(ModuleName, 4001, "controller not found")
	ErrAliasNotFound        = sdkerrors.Register(ModuleName, 4002, "alias not found")
)
