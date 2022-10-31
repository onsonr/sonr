package keeper

import (
	"context"
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
	var whereIs = types.BucketConfig{
		Creator:  msg.Creator,
		IsActive: true,
		Uuid:     uuid,
		Name:     msg.GetLabel(),
	}

	// fmt.Printf("label: %s, vi: %d, role: %d \n", whereIs.Label, whereIs.Visibility, whereIs.Role)
	k.AppendBucket(
		ctx,
		bucket,
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyCreator, bucket.Creator),
			// sdk.NewAttribute(types.AttributeKeyDID, whereIs.Did),
			// sdk.NewAttribute(types.AttributeKeyLabel, whereIs.Label),
			sdk.NewAttribute(types.AttributeKeyTxType, types.EventTypeCreateWhereIs),
		),
	)

	return &types.MsgDefineBucketResponse{
		Status: http.StatusAccepted,
		Bucket: &bucket,
	}, nil
}
