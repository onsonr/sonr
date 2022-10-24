package keeper

import (
	"context"
	"fmt"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/x/bucket/types"
)

func (k msgServer) DefineBucket(goCtx context.Context, msg *types.MsgDefineBucket) (*types.MsgDefineBucketResponse, error) {
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

	uuid := k.GenerateKeyForDID()
	var whereIs = types.Bucket{
		Creator:  msg.Creator,
		IsActive: true,
		Uuid:     uuid,
	}
	// fmt.Printf("label: %s, vi: %d, role: %d \n", whereIs.Label, whereIs.Visibility, whereIs.Role)
	k.AppendWhereIs(
		ctx,
		whereIs,
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyCreator, whereIs.Creator),
			// sdk.NewAttribute(types.AttributeKeyDID, whereIs.Did),
			// sdk.NewAttribute(types.AttributeKeyLabel, whereIs.Label),
			sdk.NewAttribute(types.AttributeKeyTxType, types.EventTypeCreateWhereIs),
		),
	)

	return &types.MsgDefineBucketResponse{
		Status:  http.StatusAccepted,
		WhereIs: &whereIs,
	}, nil
}

func (k msgServer) UpdateBucket(goCtx context.Context, msg *types.MsgUpdateBucket) (*types.MsgUpdateBucketResponse, error) {
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

	var whereIs = types.Bucket{
		// Label:      msg.Label,
		Creator: msg.Creator,
		// Did:        msg.Did,
		// Visibility: msg.Visibility,
		// Role:       msg.Role,
		IsActive: true,
		// Content:    msg.Content,
		// Timestamp:  time.Now().Unix(),
	}

	// Checks that the element exists
	val, found := k.GetBucket(ctx, msg.Did)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Did))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetWhereIs(ctx, whereIs)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyCreator, whereIs.Creator),
			// sdk.NewAttribute(types.AttributeKeyDID, whereIs.Did),
			// sdk.NewAttribute(types.AttributeKeyLabel, whereIs.Label),
			sdk.NewAttribute(types.AttributeKeyTxType, types.EventTypeUpdateWhereIs),
		),
	)
	return &types.MsgUpdateBucketResponse{
		Status:  http.StatusAccepted,
		WhereIs: &whereIs,
	}, nil
}

func (k msgServer) DeleteBucket(goCtx context.Context, msg *types.MsgDeleteBucket) (*types.MsgDeleteBucketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetBucket(ctx, msg.Did)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Did))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	val.IsActive = false
	k.SetWhereIs(ctx, val)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyCreator, val.Creator),
			// sdk.NewAttribute(types.AttributeKeyDID, val.Did),
			// sdk.NewAttribute(types.AttributeKeyLabel, val.Label),
			sdk.NewAttribute(types.AttributeKeyTxType, types.EventTypeDeleteWhereIs),
		),
	)
	return &types.MsgDeleteBucketResponse{}, nil
}
