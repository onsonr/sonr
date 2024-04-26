package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/di-dao/core/x/svc/types"
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

// LoginOptions implements types.QueryServer.
func (k Querier) LoginOptions(goCtx context.Context, req *types.QueryLoginOptionsRequest) (*types.QueryLoginOptionsResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("LoginOptions is unimplemented")
	return &types.QueryLoginOptionsResponse{}, nil
}

// Record implements types.QueryServer.
func (k Querier) Record(goCtx context.Context, req *types.QueryRecordRequest) (*types.QueryRecordResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("Record is unimplemented")
	return &types.QueryRecordResponse{}, nil
}

// RegisterOptions implements types.QueryServer.
func (k Querier) RegisterOptions(goCtx context.Context, req *types.QueryRegisterOptionsRequest) (*types.QueryRegisterOptionsResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("RegisterOptions is unimplemented")
	return &types.QueryRegisterOptionsResponse{}, nil
}
