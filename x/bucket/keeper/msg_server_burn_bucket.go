package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/x/bucket/types"
)

func (k msgServer) BurnBucket(goCtx context.Context, msg *types.MsgBurnBucket) (*types.MsgBurnBucketResponse, error) {
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

	k.RemoveWhereIs(
		ctx,
		msg.Bucket.Uuid,
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Bucket.Uuid),
			sdk.NewAttribute(types.AttributeKeyTxType, types.EventTypeBurnBucket),
		),
	)

	return &types.MsgBurnBucketResponse{}, nil
}
