package keeper

import (
	"github.com/sonrhq/core/x/identity/types"
)

var _ types.QueryServer = Keeper{}
