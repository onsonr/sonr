package service

import "cosmossdk.io/collections"

// ModuleName is the name of the module
const ModuleName = "service"

var (
	// ParamsKey is the key for the module parameters
	ParamsKey = collections.NewPrefix(0)

	// CounterKey is the key for the module counter
	CounterKey = collections.NewPrefix(1)
)
