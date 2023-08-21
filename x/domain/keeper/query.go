package keeper

import (
	"github.com/sonr-io/sonr/x/domain/types"
)

var _ types.QueryServer = Keeper{}
