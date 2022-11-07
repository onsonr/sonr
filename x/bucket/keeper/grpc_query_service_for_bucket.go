package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/x/bucket/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ServiceForBucket(c context.Context, req *types.QueryGetServiceForBucketRequest) (*types.QueryGetServiceForBucketResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	svc, found := k.GetService(ctx, req.Uuid)
	if !found {
		return nil, fmt.Errorf("error while querying service for bucket: %s %s", req.Uuid, sdkerrors.ErrKeyNotFound)
	}
	return &types.QueryGetServiceForBucketResponse{
		Service: svc,
	}, nil
}
