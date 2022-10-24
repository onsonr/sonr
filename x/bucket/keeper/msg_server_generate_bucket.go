package keeper

import (
	"context"

    "github.com/sonr-io/sonr/x/bucket/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


func (k msgServer) GenerateBucket(goCtx context.Context,  msg *types.MsgGenerateBucket) (*types.MsgGenerateBucketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // TODO: Handling the message
    _ = ctx

	return &types.MsgGenerateBucketResponse{}, nil
}
