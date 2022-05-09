package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/x/channel/types"
)

func (k msgServer) CreateHowIs(goCtx context.Context, msg *types.MsgCreateHowIs) (*types.MsgCreateHowIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetHowIs(
		ctx,
		msg.Did,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("HowIs with DID '%s' already exists", msg.Did))
	}

	var howIs = types.HowIs{
		Creator: msg.Creator,
		Did:     msg.Did,
		Channel: msg.Channel,
	}

	k.SetHowIs(
		ctx,
		howIs,
	)
	return &types.MsgCreateHowIsResponse{}, nil
}

func (k msgServer) UpdateHowIs(goCtx context.Context, msg *types.MsgUpdateHowIs) (*types.MsgUpdateHowIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetHowIs(
		ctx,
		msg.Did,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("HowIs with DID '%s' not found", msg.Did))
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var howIs = types.HowIs{
		Creator: msg.Creator,
		Did:     msg.Did,
		Channel: msg.Channel,
	}

	k.SetHowIs(ctx, howIs)

	return &types.MsgUpdateHowIsResponse{}, nil
}

func (k msgServer) DeleteHowIs(goCtx context.Context, msg *types.MsgDeleteHowIs) (*types.MsgDeleteHowIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetHowIs(
		ctx,
		msg.Did,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("HowIs with DID '%s' not found", msg.Did))
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveHowIs(
		ctx,
		msg.Did,
	)

	return &types.MsgDeleteHowIsResponse{}, nil
}
