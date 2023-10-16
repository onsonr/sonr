package keeper

import (
	"github.com/sonr-io/core/x/domain/types"
)

var _ types.QueryServer = Keeper{}
