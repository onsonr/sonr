package keeper

import (
	"github.com/sonr-io/sonr/x/schema/types"
)

var _ types.QueryServer = Keeper{}
