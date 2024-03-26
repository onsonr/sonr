package keeper

import (
	"context"

	"github.com/didao-org/sonr/x/identity/types"
)

type msgServer struct {
	k Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{k: keeper}
}

// CreateRecord implements types.MsgServer.
func (ms msgServer) InitializeIdentity(ctx context.Context, msg *types.MsgInitializeIdentity) (*types.MsgInitializeIdentityResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("CreateRecord is unimplemented")
	return &types.MsgInitializeIdentityResponse{}, nil
}
