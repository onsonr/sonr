package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/x/bucket/types"
)

func (k msgServer) AllocateBucket(goCtx context.Context, msg *types.MsgAllocateBucket) (*types.MsgAllocateBucketResponse, error) {
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

	bucket, found := k.GetBucket(ctx, msg.Bucket.Uuid)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("bucket with id %s not defined", msg.Bucket.Uuid))
	}

	didService := bucket.GetDidService(bucket.GetCreator(), msg.Record)

	k.AddService(ctx, bucket.GetUuid(), didService)

	err = k.UpdateWhoIsService(ctx, bucket, didService)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyCreator, bucket.GetCreator()),
			sdk.NewAttribute(types.AttributeKeyServiceId, didService.GetId()),
			sdk.NewAttribute(types.AttributeKeyLabel, bucket.GetName()),
			sdk.NewAttribute(types.AttributeKeyTxType, types.EventTypeGenerateBucket),
		),
	)

	return &types.MsgAllocateBucketResponse{}, nil
}
