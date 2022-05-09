package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/x/bucket/types"
)

func (k msgServer) DeactivateBucket(goCtx context.Context, msg *types.MsgDeactivateBucket) (*types.MsgDeactivateBucketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if Object exists
	howis, found := k.GetWhichIs(ctx, msg.GetDid())
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Bucket was not found in this Application")
	}
	howis.IsActive = false
	k.SetWhichIs(ctx, howis)
	return &types.MsgDeactivateBucketResponse{
		Code:    100,
		Message: fmt.Sprintf("Bucket %s has been deactivated", howis.Did),
	}, nil
}
