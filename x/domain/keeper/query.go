package keeper

import (
	"sonr.io/core/x/domain/types"
)

var _ types.QueryServer = Keeper{}
