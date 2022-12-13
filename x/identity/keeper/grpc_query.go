package keeper

import (
	"github.com/sonr-hq/sonr/x/identity/types"
)

var _ types.QueryServer = Keeper{}
