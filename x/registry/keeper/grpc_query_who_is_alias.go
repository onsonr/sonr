package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/registry/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) WhoIsAlias(goCtx context.Context, req *types.QueryWhoIsAliasRequest) (*types.QueryWhoIsAliasResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.FindWhoIsByAlias(
		ctx,
		req.GetAlias(),
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryWhoIsAliasResponse{
		WhoIs: &val,
	}, nil
}
