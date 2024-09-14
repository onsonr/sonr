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

// ParamsAssets implements types.QueryServer.
func (k Querier) ParamsAssets(goCtx context.Context, req *types.QueryRequest) (*types.QueryResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("ParamsAssets is unimplemented")
	return &types.QueryResponse{}, nil
}

// ParamsByAsset implements types.QueryServer.
func (k Querier) ParamsByAsset(goCtx context.Context, req *types.QueryRequest) (*types.QueryResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("ParamsByAsset is unimplemented")
	return &types.QueryResponse{}, nil
}

// ParamsKeys implements types.QueryServer.
func (k Querier) ParamsKeys(goCtx context.Context, req *types.QueryRequest) (*types.QueryResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("ParamsKeys is unimplemented")
	return &types.QueryResponse{}, nil
}

// ParamsByKey implements types.QueryServer.
func (k Querier) ParamsByKey(goCtx context.Context, req *types.QueryRequest) (*types.QueryResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("ParamsByKey is unimplemented")
	return &types.QueryResponse{}, nil
}

// RegistrationOptionsByKey implements types.QueryServer.
func (k Querier) RegistrationOptionsByKey(goCtx context.Context, req *types.QueryRequest) (*types.QueryResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("RegistrationOptionsByKey is unimplemented")
	return &types.QueryResponse{}, nil
// Accounts implements types.QueryServer.
func (k Querier) Accounts(goCtx context.Context, req *types.QueryAccountsRequest) (*types.QueryAccountsResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryAccountsResponse{}, nil
}

// Credentials implements types.QueryServer.
func (k Querier) Credentials(goCtx context.Context, req *types.QueryCredentialsRequest) (*types.QueryCredentialsResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryCredentialsResponse{}, nil
}

// Identities implements types.QueryServer.
func (k Querier) Identities(goCtx context.Context, req *types.QueryIdentitiesRequest) (*types.QueryIdentitiesResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryIdentitiesResponse{}, nil
}

// Resolve implements types.QueryServer.
func (k Querier) Resolve(goCtx context.Context, req *types.QueryResolveRequest) (*types.QueryResolveResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryResolveResponse{}, nil
}

// Service implements types.QueryServer.
func (k Querier) Service(goCtx context.Context, req *types.QueryServiceRequest) (*types.QueryServiceResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryServiceResponse{}, nil
}
