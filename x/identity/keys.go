package identity

import "cosmossdk.io/collections"

const ModuleName = "identity"

var (
	ParamsKey  = collections.NewPrefix(0)
	CounterKey = collections.NewPrefix(1)
)
