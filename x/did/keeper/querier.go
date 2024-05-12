package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/di-dao/core/x/did/types"
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

// Account implements types.QueryServer.
func (k Querier) Account(goCtx context.Context, req *types.QueryAccountRequest) (*types.QueryAccountResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("Account is unimplemented")
	return &types.QueryAccountResponse{}, nil
}

// Exists implements types.QueryServer.
func (k Querier) Exists(goCtx context.Context, req *types.QueryExistsRequest) (*types.QueryExistsResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("Exists is unimplemented")
	return &types.QueryExistsResponse{}, nil
}

// Resolve implements types.QueryServer.
func (k Querier) Resolve(goCtx context.Context, req *types.QueryResolveRequest) (*types.QueryResolveResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("Resolve is unimplemented")
	return &types.QueryResolveResponse{}, nil
}

// StartLogin implements types.QueryServer.
func (k Querier) StartLogin(goCtx context.Context, req *types.QueryStartLoginRequest) (*types.QueryStartLoginResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("StartLogin is unimplemented")
	return &types.QueryStartLoginResponse{}, nil
}

// StartRegister implements types.QueryServer.
func (k Querier) StartRegister(goCtx context.Context, req *types.QueryStartRegisterRequest) (*types.QueryStartRegisterResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("StartRegister is unimplemented")
	return &types.QueryStartRegisterResponse{}, nil
}
