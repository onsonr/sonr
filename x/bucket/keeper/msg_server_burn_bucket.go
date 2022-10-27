package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/bucket/types"
)

func (k msgServer) BurnBucket(goCtx context.Context, msg *types.MsgBurnBucket) (*types.MsgBurnBucketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgBurnBucketResponse{}, nil
}
