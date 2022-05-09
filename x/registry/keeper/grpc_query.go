package keeper

import (
	"github.com/sonr-io/sonr/x/registry/types"
)

var _ types.QueryServer = Keeper{}
