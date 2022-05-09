package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/internal/blockchain/x/bucket/types"
)

func (k msgServer) CreateWhichIs(goCtx context.Context, msg *types.MsgCreateWhichIs) (*types.MsgCreateWhichIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetWhichIs(
		ctx,
		msg.Did,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("WhichIs already exists with DID '%s'", msg.Did))
	}

	var whichIs = types.WhichIs{
		Creator: msg.Creator,
		Did:     msg.Did,
		Bucket:  msg.Bucket,
	}

	k.SetWhichIs(
		ctx,
		whichIs,
	)
	return &types.MsgCreateWhichIsResponse{}, nil
}

func (k msgServer) UpdateWhichIs(goCtx context.Context, msg *types.MsgUpdateWhichIs) (*types.MsgUpdateWhichIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetWhichIs(
		ctx,
		msg.Did,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("WhichIs with DID '%s' not found", msg.Did))
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var whichIs = types.WhichIs{
		Creator: msg.Creator,
		Did:     msg.Did,
		Bucket:  msg.Bucket,
	}

	k.SetWhichIs(ctx, whichIs)

	return &types.MsgUpdateWhichIsResponse{}, nil
}

func (k msgServer) DeleteWhichIs(goCtx context.Context, msg *types.MsgDeleteWhichIs) (*types.MsgDeleteWhichIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetWhichIs(
		ctx,
		msg.Did,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("WhichIs with DID '%s' not found", msg.Did))
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveWhichIs(
		ctx,
		msg.Did,
	)

	return &types.MsgDeleteWhichIsResponse{}, nil
}
