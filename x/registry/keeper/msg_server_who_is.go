package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/x/registry/types"
)

func (k msgServer) CreateWhoIs(goCtx context.Context, msg *types.MsgCreateWhoIs) (*types.MsgCreateWhoIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetWhoIs(
		ctx,
		msg.Did,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("WhoIs with DID '%s' already exists", msg.Did))
	}

	var whoIs = types.WhoIs{
		Owner:       msg.Creator,
		Name:        msg.Name,
		Did:         msg.Did,
		Document:    msg.Document,
		Credentials: msg.Credentials,
	}

	k.SetWhoIs(
		ctx,
		whoIs,
	)
	return &types.MsgCreateWhoIsResponse{}, nil
}

func (k msgServer) UpdateWhoIs(goCtx context.Context, msg *types.MsgUpdateWhoIs) (*types.MsgUpdateWhoIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetWhoIs(
		ctx,
		msg.Did,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "UpdateWhoIs: DID not found")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var whoIs = types.WhoIs{
		Owner:       msg.Creator,
		Name:        valFound.Name,
		Did:         msg.Did,
		Document:    msg.Document,
		Credentials: msg.Credentials,
	}

	k.SetWhoIs(ctx, whoIs)

	return &types.MsgUpdateWhoIsResponse{}, nil
}

func (k msgServer) DeleteWhoIs(goCtx context.Context, msg *types.MsgDeleteWhoIs) (*types.MsgDeleteWhoIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetWhoIs(
		ctx,
		msg.Did,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("WhoIs with DID '%s' not found", msg.Did))
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveWhoIs(
		ctx,
		msg.Did,
	)

	return &types.MsgDeleteWhoIsResponse{}, nil
}
