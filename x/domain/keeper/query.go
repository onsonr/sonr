package keeper

import (
	"github.com/sonrhq/core/x/domain/types"
)

var _ types.QueryServer = Keeper{}
