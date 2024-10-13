package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onsonr/sonr/x/vault/types"
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
	ctx := sdk.UnwrapSDKContext(goCtx)

	p, err := k.Keeper.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	return &types.QuerySchemaResponse{
		Schema: p.Schema,
	}, nil
}

// SyncInitial implements types.QueryServer.
func (k Querier) SyncInitial(goCtx context.Context, req *types.SyncInitialRequest) (*types.SyncInitialResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("SyncInitial is unimplemented")
	return &types.SyncInitialResponse{}, nil
}

// SyncCurrent implements types.QueryServer.
func (k Querier) SyncCurrent(goCtx context.Context, req *types.SyncCurrentRequest) (*types.SyncCurrentResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("SyncCurrent is unimplemented")
	return &types.SyncCurrentResponse{}, nil
}
