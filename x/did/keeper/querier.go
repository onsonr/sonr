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
) (*types.QueryResponse, error) {
	ctx := k.CurrentCtx(goCtx)
	return &types.QueryResponse{Params: k.GetParams(ctx.SDK())}, nil
}

// Resolve implements types.QueryServer.
func (k Querier) Resolve(
	goCtx context.Context,
	req *types.QueryRequest,
) (*types.QueryResponse, error) {
	ctx := k.CurrentCtx(goCtx)
	return &types.QueryResponse{Params: k.GetParams(ctx.SDK())}, nil
}

// Service implements types.QueryServer.
func (k Querier) Service(
	goCtx context.Context,
	req *types.QueryRequest,
) (*types.QueryResponse, error) {
	ctx := k.CurrentCtx(goCtx)
	return &types.QueryResponse{Service: ctx.GetServiceInfo(req.GetOrigin()), Params: ctx.Params()}, nil
}
