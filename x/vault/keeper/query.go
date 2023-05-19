package keeper

import (
	"github.com/sonrhq/core/x/vault/types"
)

var _ types.QueryServer = Keeper{}
