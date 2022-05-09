package keeper

import (
	"github.com/sonr-io/sonr/internal/blockchain/x/object/types"
)

var _ types.QueryServer = Keeper{}
