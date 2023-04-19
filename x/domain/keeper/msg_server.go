package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonrhq/core/x/domain/types"
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

// ! ||--------------------------------------------------------------------------------||
// ! ||                              SLD Record Management                             ||
// ! ||--------------------------------------------------------------------------------||

func (k msgServer) CreateSLDRecord(goCtx context.Context, msg *types.MsgCreateSLDRecord) (*types.MsgCreateSLDRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetSLDRecord(
		ctx,
		msg.SldRecord.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var sLDRecord = types.SLDRecord{
		Creator: msg.Creator,
		Index:   msg.SldRecord.Index,
	}

	k.SetSLDRecord(
		ctx,
		sLDRecord,
	)
	return &types.MsgCreateSLDRecordResponse{}, nil
}

func (k msgServer) UpdateSLDRecord(goCtx context.Context, msg *types.MsgUpdateSLDRecord) (*types.MsgUpdateSLDRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetSLDRecord(
		ctx,
		msg.SldRecord.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var sLDRecord = types.SLDRecord{
		Creator: msg.Creator,
		Index:   msg.SldRecord.Index,
	}

	k.SetSLDRecord(ctx, sLDRecord)

	return &types.MsgUpdateSLDRecordResponse{}, nil
}

func (k msgServer) DeleteSLDRecord(goCtx context.Context, msg *types.MsgDeleteSLDRecord) (*types.MsgDeleteSLDRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetSLDRecord(
		ctx,
		msg.Name,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveSLDRecord(
		ctx,
		msg.Name,
	)

	return &types.MsgDeleteSLDRecordResponse{}, nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                              TLD Record Management                             ||
// ! ||--------------------------------------------------------------------------------||

func (k msgServer) CreateTLDRecord(goCtx context.Context, msg *types.MsgCreateTLDRecord) (*types.MsgCreateTLDRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetTLDRecord(
		ctx,
		msg.TldRecord.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var tLDRecord = types.TLDRecord{
		Creator: msg.Creator,
		Index:   msg.TldRecord.Index,
	}

	k.SetTLDRecord(
		ctx,
		tLDRecord,
	)
	return &types.MsgCreateTLDRecordResponse{}, nil
}

func (k msgServer) UpdateTLDRecord(goCtx context.Context, msg *types.MsgUpdateTLDRecord) (*types.MsgUpdateTLDRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetTLDRecord(
		ctx,
		msg.TldRecord.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var tLDRecord = types.TLDRecord{
		Creator: msg.Creator,
		Index:   msg.TldRecord.Index,
	}

	k.SetTLDRecord(ctx, tLDRecord)

	return &types.MsgUpdateTLDRecordResponse{}, nil
}

func (k msgServer) DeleteTLDRecord(goCtx context.Context, msg *types.MsgDeleteTLDRecord) (*types.MsgDeleteTLDRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetTLDRecord(
		ctx,
		msg.Name,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveTLDRecord(
		ctx,
		msg.Name,
	)

	return &types.MsgDeleteTLDRecordResponse{}, nil
}
