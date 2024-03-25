package keeper

import (
	"context"

	"github.com/didao-org/sonr/x/service/types"
	service "github.com/didao-org/sonr/x/service/types"
)

type msgServer struct {
	k Keeper
}

var _ service.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface.
func NewMsgServerImpl(keeper Keeper) service.MsgServer {
	return &msgServer{k: keeper}
}

// CreateRecord implements types.MsgServer.
func (ms msgServer) CreateRecord(ctx context.Context, msg *types.MsgCreateRecord) (*types.MsgCreateRecordResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("CreateRecord is unimplemented")
	return &types.MsgCreateRecordResponse{}, nil
}

// UpdateRecord implements types.MsgServer.
func (ms msgServer) UpdateRecord(ctx context.Context, msg *types.MsgUpdateRecord) (*types.MsgUpdateRecordResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("UpdateRecord is unimplemented")
	return &types.MsgUpdateRecordResponse{}, nil
}

// DeleteRecord implements types.MsgServer.
func (ms msgServer) DeleteRecord(ctx context.Context, msg *types.MsgDeleteRecord) (*types.MsgDeleteRecordResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("DeleteRecord is unimplemented")
	return &types.MsgDeleteRecordResponse{}, nil
}
