package msg

import (
	"github.com/sonr-io/sonr/internal/blockchain/x/bucket/keeper"
	"github.com/sonr-io/sonr/internal/blockchain/x/bucket/types"
)

type msgServer struct {
	keeper.Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper keeper.Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
