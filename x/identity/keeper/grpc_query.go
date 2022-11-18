package keeper

import (
	"github.com/sonr-io/sonr/x/identity/types"
)

var _ types.QueryServer = Keeper{}
