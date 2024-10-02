package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onsonr/sonr/x/macaroon/types"
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

// RefreshToken implements types.QueryServer.
func (k Querier) RefreshToken(goCtx context.Context, req *types.QueryRefreshTokenRequest) (*types.QueryRefreshTokenResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryRefreshTokenResponse{}, nil
}

// ValidateToken implements types.QueryServer.
func (k Querier) ValidateToken(goCtx context.Context, req *types.QueryValidateTokenRequest) (*types.QueryValidateTokenResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryValidateTokenResponse{}, nil
}
