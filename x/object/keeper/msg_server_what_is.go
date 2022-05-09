package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/internal/blockchain/x/object/types"
)

func (k msgServer) CreateWhatIs(goCtx context.Context, msg *types.MsgCreateWhatIs) (*types.MsgCreateWhatIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetWhatIs(
		ctx,
		msg.Did,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("WhatIs with DID '%s' already exists", msg.Did))
	}

	var whatIs = types.WhatIs{
		Creator:   msg.Creator,
		Did:       msg.Did,
		ObjectDoc: msg.GetObjectDoc(),
	}

	k.SetWhatIs(
		ctx,
		whatIs,
	)
	return &types.MsgCreateWhatIsResponse{}, nil
}

func (k msgServer) UpdateWhatIs(goCtx context.Context, msg *types.MsgUpdateWhatIs) (*types.MsgUpdateWhatIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetWhatIs(
		ctx,
		msg.Did,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("WhatIs with DID '%s' not found", msg.Did))
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var whatIs = types.WhatIs{
		Creator:   msg.Creator,
		Did:       msg.Did,
		ObjectDoc: msg.GetObjectDoc(),
	}

	k.SetWhatIs(ctx, whatIs)

	return &types.MsgUpdateWhatIsResponse{}, nil
}

func (k msgServer) DeleteWhatIs(goCtx context.Context, msg *types.MsgDeleteWhatIs) (*types.MsgDeleteWhatIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetWhatIs(
		ctx,
		msg.Did,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("WhatIs with DID '%s' not found", msg.Did))
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveWhatIs(
		ctx,
		msg.Did,
	)

	return &types.MsgDeleteWhatIsResponse{}, nil
}
