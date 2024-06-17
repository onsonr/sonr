package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/di-dao/sonr/x/did/types"
)

var _ types.QueryServer = Querier{}

type Querier struct {
	Keeper
}

func NewQuerier(keeper Keeper) Querier {
	return Querier{Keeper: keeper}
}

func (k Querier) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	p, err := k.Keeper.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryParamsResponse{Params: &p}, nil
}

// Resolve implements types.QueryServer.
func (k Querier) Resolve(goCtx context.Context, req *types.QueryResolveRequest) (*types.QueryResolveResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("Resolve is unimplemented")
	return &types.QueryResolveResponse{}, nil
}

// LoginOptions implements types.QueryServer.
func (k Querier) LoginOptions(goCtx context.Context, req *types.QueryLoginOptionsRequest) (*types.QueryLoginOptionsResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("LoginOptions is unimplemented")
	return &types.QueryLoginOptionsResponse{}, nil
}

// RegisterOptions implements types.QueryServer.
func (k Querier) RegisterOptions(goCtx context.Context, req *types.QueryRegisterOptionsRequest) (*types.QueryRegisterOptionsResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("RegisterOptions is unimplemented")
	return &types.QueryRegisterOptionsResponse{}, nil
}

// PropertyExists implements types.QueryServer.
func (k Querier) PropertyExists(goCtx context.Context, req *types.QueryExistsRequest) (*types.QueryExistsResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("PropertyExists is unimplemented")
	return &types.QueryExistsResponse{}, nil
}

// ResolveIdentifier implements types.QueryServer.
func (k Querier) ResolveIdentifier(goCtx context.Context, req *types.QueryResolveRequest) (*types.QueryResolveResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("ResolveIdentifier is unimplemented")
	return &types.QueryResolveResponse{}, nil
}
