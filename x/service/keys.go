package service

import "cosmossdk.io/collections"

const ModuleName = "service"

var (
	ParamsKey  = collections.NewPrefix(0)
	CounterKey = collections.NewPrefix(1)
)
