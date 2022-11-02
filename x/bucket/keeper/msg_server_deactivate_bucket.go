package keeper

import (
	"context"
	"fmt"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/x/bucket/types"
)

func (k msgServer) DeactivateBucket(goCtx context.Context, msg *types.MsgDeactivateBucket) (*types.MsgDeactivateBucketResponse, error) {
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

	// Checks that the element exists
	val, found := k.GetBucket(ctx, msg.Bucket.Uuid)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Bucket.Uuid))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	val.IsActive = false

	k.SetBucket(
		ctx,
		val,
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Bucket.Uuid),
			sdk.NewAttribute(types.AttributeKeyTxType, types.EventTypeDeactivateBucket),
		),
	)

	return &types.MsgDeactivateBucketResponse{
		Status: http.StatusAccepted,
		Bucket: &val,
	}, nil
}
