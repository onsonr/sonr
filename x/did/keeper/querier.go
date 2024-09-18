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
	ctx := k.CurrentCtx(goCtx)
	return &types.QueryParamsResponse{Params: k.GetParams(ctx.SDK())}, nil
}

// Resolve implements types.QueryServer.
func (k Querier) Resolve(
	goCtx context.Context,
	req *types.QueryRequest,
) (*types.QueryResolveResponse, error) {
	_ = k.CurrentCtx(goCtx)
	return &types.QueryResolveResponse{}, nil
}

// Service implements types.QueryServer.
func (k Querier) Service(
	goCtx context.Context,
	req *types.QueryRequest,
) (*types.QueryServiceResponse, error) {
	// ctx := k.CurrentCtx(goCtx)
	return &types.QueryServiceResponse{}, nil
}

// Sync implements types.QueryServer.
func (k Querier) Sync(goCtx context.Context, req *types.SyncRequest) (*types.SyncResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("Sync is unimplemented")
	return &types.SyncResponse{}, nil
}
