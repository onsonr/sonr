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

// WitnessCredential implements types.QueryServer.
func (k Querier) WitnessCredential(goCtx context.Context, req *types.QueryWitnessCredentialRequest) (*types.QueryWitnessCredentialResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("WitnessCredential is unimplemented")
	return &types.QueryWitnessCredentialResponse{}, nil
}
