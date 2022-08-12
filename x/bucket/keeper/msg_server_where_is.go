package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/x/bucket/types"
)

func (k msgServer) CreateWhereIs(goCtx context.Context, msg *types.MsgCreateWhereIs) (*types.MsgCreateWhereIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := msg.ValidateBasic()
	k.Logger(ctx).Info("basic request validation finished")

	if err != nil {
		return nil, err
	}

	accts := msg.GetSigners()
	if len(accts) < 1 {
		k.Logger(ctx).Error("Error while querying account: not found")
		return nil, sdkerrors.ErrNotFound
	}

	creator_did := msg.GetCreatorDid()
	uuid := k.GenerateKeyForDID()

	did := fmt.Sprintf("did:snr:%s", uuid)

	var whereIs = types.WhereIs{
		Creator:    creator_did,
		Did:        did,
		Visibility: msg.Visibility,
		Role:       msg.Role,
		IsActive:   true,
		Content:    msg.Content,
		Timestamp:  time.Now().Unix(),
	}

	id := k.AppendWhereIs(
		ctx,
		whereIs,
	)

	return &types.MsgCreateWhereIsResponse{
		Did: id,
	}, nil
}

func (k msgServer) UpdateWhereIs(goCtx context.Context, msg *types.MsgUpdateWhereIs) (*types.MsgUpdateWhereIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var whereIs = types.WhereIs{
		Creator: msg.Creator,
		Did:     msg.Did,
	}

	// Checks that the element exists
	val, found := k.GetWhereIs(ctx, msg.Did)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Did))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetWhereIs(ctx, whereIs)

	return &types.MsgUpdateWhereIsResponse{}, nil
}

func (k msgServer) DeleteWhereIs(goCtx context.Context, msg *types.MsgDeleteWhereIs) (*types.MsgDeleteWhereIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetWhereIs(ctx, msg.Did)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Did))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveWhereIs(ctx, msg.Did)

	return &types.MsgDeleteWhereIsResponse{}, nil
}
