package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/object module sentinel errors
var (
	ErrSample         = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInactiveObject = sdkerrors.Register(ModuleName, 1104, "Requested object has been deactivated")
)
