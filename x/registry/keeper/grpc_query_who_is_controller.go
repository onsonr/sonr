package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/registry/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) WhoIsController(goCtx context.Context, req *types.QueryWhoIsControllerRequest) (*types.QueryWhoIsControllerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.FindWhoIsByController(
		ctx,
		req.GetController(),
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}
	return &types.QueryWhoIsControllerResponse{
		WhoIs: &val,
	}, nil
}
