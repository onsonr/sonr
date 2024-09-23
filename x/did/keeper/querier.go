package keeper

import (
	"context"

	"github.com/onsonr/sonr/x/did/types"
)

var _ types.QueryServer = Querier{}

type Querier struct {
	Keeper
}

func NewQuerier(keeper Keeper) Querier {
	return Querier{Keeper: keeper}
}

// Params returns the total set of did parameters.
func (k Querier) Params(
	goCtx context.Context,
	req *types.QueryRequest,
) (*types.QueryParamsResponse, error) {
	ctx := k.UnwrapCtx(goCtx)
	return &types.QueryParamsResponse{Params: ctx.Params()}, nil
}

// Resolve implements types.QueryServer.
func (k Querier) Resolve(
	goCtx context.Context,
	req *types.QueryRequest,
) (*types.QueryResolveResponse, error) {
	_ = k.UnwrapCtx(goCtx)
	return &types.QueryResolveResponse{}, nil
}

// Sync implements types.QueryServer.
func (k Querier) Sync(goCtx context.Context, req *types.SyncRequest) (*types.SyncResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.SyncResponse{}, nil
}
