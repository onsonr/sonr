package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sonrhq/sonr/x/identity"
)

var _ identity.QueryServer = queryServer{}

// NewQueryServerImpl returns an implementation of the module QueryServer.
func NewQueryServerImpl(k Keeper) identity.QueryServer {
	return queryServer{k}
}

type queryServer struct {
	k Keeper
}

// Params defines the handler for the Query/Params RPC method.
func (qs queryServer) Params(ctx context.Context, req *identity.QueryParamsRequest) (*identity.QueryParamsResponse, error) {
	params, err := qs.k.Params.Get(ctx)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return &identity.QueryParamsResponse{Params: identity.Params{}}, nil
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &identity.QueryParamsResponse{Params: params}, nil
}
