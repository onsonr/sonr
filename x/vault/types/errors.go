package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/vault module sentinel errors
var (
	ErrSample     = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInboxWrite = sdkerrors.Register(ModuleName, 1101, "error writing to inbox")
	ErrInboxRead  = sdkerrors.Register(ModuleName, 1102, "error reading from inbox")
)
