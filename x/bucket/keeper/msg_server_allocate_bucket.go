package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/bucket/types"
)

func (k msgServer) AllocateBucket(goCtx context.Context, msg *types.MsgAllocateBucket) (*types.MsgAllocateBucketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAllocateBucketResponse{}, nil
}
