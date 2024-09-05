package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	// "github.com/onsonr/sonr/internal/local"
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
func (k Querier) Params(c context.Context, req *types.QueryRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	p, err := k.Keeper.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryParamsResponse{Params: &p}, nil
}

// Accounts implements types.QueryServer.
func (k Querier) Accounts(goCtx context.Context, req *types.QueryRequest) (*types.QueryAccountsResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryAccountsResponse{}, nil
}

// Credentials implements types.QueryServer.
func (k Querier) Credentials(goCtx context.Context, req *types.QueryRequest) (*types.QueryCredentialsResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryCredentialsResponse{}, nil
}

// Resolve implements types.QueryServer.
func (k Querier) Resolve(goCtx context.Context, req *types.QueryRequest) (*types.QueryResolveResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryResolveResponse{}, nil
}

// Service implements types.QueryServer.
func (k Querier) Service(goCtx context.Context, req *types.QueryRequest) (*types.QueryServiceResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryServiceResponse{}, nil
}

// Token implements types.QueryServer.
func (k Querier) Token(goCtx context.Context, req *types.QueryRequest) (*types.QueryTokenResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryTokenResponse{}, nil
}
