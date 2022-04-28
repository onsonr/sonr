package keeper

import (
	"github.com/sonr-io/sonr/internal/blockchain/x/bucket/types"
)

var _ types.QueryServer = Keeper{}
