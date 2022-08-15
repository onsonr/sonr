package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/bucket/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) WhereIsByCreator(c context.Context, req *types.QueryGetWhereIsByCreatorRequest) (*types.QueryGetWhereIsByCreatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	whereIss := k.GetWhereIsByCreator(ctx, req.Creator)

	return &types.QueryGetWhereIsByCreatorResponse{
		WhereIs: whereIss,
	}, nil
}
