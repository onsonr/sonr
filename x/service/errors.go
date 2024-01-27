package service

import "cosmossdk.io/errors"

// ErrDuplicateAddress error if there is a duplicate address
var ErrDuplicateAddress = errors.Register(ModuleName, 2, "duplicate address")
