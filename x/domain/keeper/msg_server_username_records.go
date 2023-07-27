package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonrhq/core/x/domain/types"
)

func (k msgServer) CreateUsernameRecord(goCtx context.Context, msg *types.MsgCreateUsernameRecords) (*types.MsgCreateUsernameRecordsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetUsernameRecords(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var UsernameRecord = types.UsernameRecord{
		Address: msg.Creator,
		Index:   msg.Index,
		Method: msg.Method,
	}

	k.SetUsernameRecords(
		ctx,
		UsernameRecord,
	)
	return &types.MsgCreateUsernameRecordsResponse{}, nil
}

func (k msgServer) UpdateUsernameRecord(goCtx context.Context, msg *types.MsgUpdateUsernameRecords) (*types.MsgUpdateUsernameRecordsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetUsernameRecords(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Address {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var UsernameRecord = types.UsernameRecord{
		Address: msg.Creator,
		Index:   msg.Index,
	}

	k.SetUsernameRecords(ctx, UsernameRecord)

	return &types.MsgUpdateUsernameRecordsResponse{}, nil
}

func (k msgServer) DeleteUsernameRecord(goCtx context.Context, msg *types.MsgDeleteUsernameRecords) (*types.MsgDeleteUsernameRecordsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetUsernameRecords(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Address {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveUsernameRecords(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteUsernameRecordsResponse{}, nil
}
