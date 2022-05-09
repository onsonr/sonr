package keeper

import (
	"github.com/sonr-io/sonr/x/channel/types"
)

var _ types.QueryServer = Keeper{}
