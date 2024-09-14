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
	ctx := k.CurrentCtx(goCtx)
	return &types.QueryServiceResponse{Service: ctx.GetServiceInfo(req.GetOrigin())}, nil
}

// ParamsAssets implements types.QueryServer.
func (k Querier) ParamsAssets(goCtx context.Context, req *types.QueryRequest) (*types.QueryParamsAssetsResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryParamsAssetsResponse{}, nil
}

// ParamsByAsset implements types.QueryServer.
func (k Querier) ParamsByAsset(goCtx context.Context, req *types.QueryRequest) (*types.QueryParamsByAssetResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryParamsByAssetResponse{}, nil
}

// ParamsKeys implements types.QueryServer.
func (k Querier) ParamsKeys(goCtx context.Context, req *types.QueryRequest) (*types.QueryParamsKeysResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryParamsKeysResponse{}, nil
}

// ParamsByKey implements types.QueryServer.
func (k Querier) ParamsByKey(goCtx context.Context, req *types.QueryRequest) (*types.QueryParamsByKeyResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryParamsByKeyResponse{}, nil
}
