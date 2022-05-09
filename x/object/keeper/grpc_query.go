package keeper

import (
	"github.com/sonr-io/sonr/x/object/types"
)

var _ types.QueryServer = Keeper{}
