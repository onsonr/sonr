package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/bucket/types"
)

func (k msgServer) DeactivateBucket(goCtx context.Context, msg *types.MsgDeactivateBucket) (*types.MsgDeactivateBucketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDeactivateBucketResponse{}, nil
}
