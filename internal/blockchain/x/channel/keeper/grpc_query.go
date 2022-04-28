package keeper

import (
	"github.com/sonr-io/sonr/internal/blockchain/x/channel/types"
)

var _ types.QueryServer = Keeper{}
