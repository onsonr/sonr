package keeper

import (
	"github.com/sonr-io/sonr/internal/blockchain/x/registry/types"
)

var _ types.QueryServer = Keeper{}
