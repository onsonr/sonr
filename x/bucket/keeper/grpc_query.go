package keeper

import (
	"github.com/sonr-io/sonr/x/bucket/types"
)

var _ types.QueryServer = Keeper{}
