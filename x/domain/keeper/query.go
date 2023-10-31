package keeper

import (
	"github.com/sonrhq/sonr/x/domain/types"
)

var _ types.QueryServer = Keeper{}
