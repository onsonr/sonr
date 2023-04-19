package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/sonrhq/core/x/service/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateServiceRecord(goCtx context.Context, msg *types.MsgCreateServiceRecord) (*types.MsgCreateServiceRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetServiceRecord(
		ctx,
		msg.Record.Id,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Id already set")
	}
	k.SetServiceRecord(
		ctx,
		*msg.Record,
	)
	return &types.MsgCreateServiceRecordResponse{}, nil
}

func (k msgServer) UpdateServiceRecord(goCtx context.Context, msg *types.MsgUpdateServiceRecord) (*types.MsgUpdateServiceRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetServiceRecord(
		ctx,
		msg.Record.Id,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Id not set")
	}

	// Checks if the the msg Controller is the same as the current owner
	if msg.Controller != valFound.Controller {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	k.SetServiceRecord(ctx, *msg.Record)
	return &types.MsgUpdateServiceRecordResponse{}, nil
}

func (k msgServer) DeleteServiceRecord(goCtx context.Context, msg *types.MsgDeleteServiceRecord) (*types.MsgDeleteServiceRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetServiceRecord(
		ctx,
		msg.Id,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Id not set")
	}

	// Checks if the the msg Controller is the same as the current owner
	if msg.Controller != valFound.Controller {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveServiceRecord(
		ctx,
		msg.Id,
	)

	return &types.MsgDeleteServiceRecordResponse{}, nil
}
