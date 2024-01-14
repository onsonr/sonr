package service

import "cosmossdk.io/errors"

var (
	// ErrDuplicateAddress error if there is a duplicate address
	ErrDuplicateAddress = errors.Register(ModuleName, 2, "duplicate address")
)
