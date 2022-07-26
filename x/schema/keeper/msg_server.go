package keeper

import (
	"github.com/sonr-io/sonr/pkg/protocol"
	"github.com/sonr-io/sonr/x/schema/types"
)

type msgServer struct {
	Keeper

	ipfs protocol.IPFS
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
