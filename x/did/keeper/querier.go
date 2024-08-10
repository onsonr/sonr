package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	// "github.com/onsonr/hway/internal/local"
	"github.com/onsonr/hway/x/did/types"
)

var _ types.QueryServer = Querier{}

type Querier struct {
	Keeper
}

func NewQuerier(keeper Keeper) Querier {
	return Querier{Keeper: keeper}
}

// Params returns the total set of did parameters.
func (k Querier) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	p, err := k.Keeper.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryParamsResponse{Params: &p}, nil
}

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
