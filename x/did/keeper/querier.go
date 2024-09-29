package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
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
func (k Querier) Params(goCtx context.Context, req *types.QueryRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	p, err := k.Keeper.CurrentParams(ctx)
	if err != nil {
		return nil, err
	}
	return &types.QueryParamsResponse{Params: p}, nil
}

// Resolve implements types.QueryServer.
func (k Querier) Resolve(goCtx context.Context, req *types.QueryRequest) (*types.QueryResolveResponse, error) {
	return &types.QueryResolveResponse{}, nil
}

// Sign implements types.QueryServer.
func (k Querier) Sign(goCtx context.Context, req *types.QuerySignRequest) (*types.QuerySignResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("Sign is unimplemented")
	return &types.QuerySignResponse{}, nil
}

// Verify implements types.QueryServer.
func (k Querier) Verify(goCtx context.Context, req *types.QueryVerifyRequest) (*types.QueryVerifyResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("Verify is unimplemented")
	return &types.QueryVerifyResponse{}, nil
}
