package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onsonr/sonr/x/dwn/types"
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

// Schema implements types.QueryServer.
func (k Querier) Schema(goCtx context.Context, req *types.QuerySchemaRequest) (*types.QuerySchemaResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("Schema is unimplemented")
	return &types.QuerySchemaResponse{}, nil
}

// Allocate implements types.QueryServer.
func (k Querier) Allocate(goCtx context.Context, req *types.QueryAllocateRequest) (*types.QueryAllocateResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("Allocate is unimplemented")
	return &types.QueryAllocateResponse{}, nil
}

// Sync implements types.QueryServer.
func (k Querier) Sync(goCtx context.Context, req *types.QuerySyncRequest) (*types.QuerySyncResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("Sync is unimplemented")
	return &types.QuerySyncResponse{}, nil
}
