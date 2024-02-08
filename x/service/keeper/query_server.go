package keeper

import (
	"context"
	"encoding/json"
	"errors"

	"cosmossdk.io/collections"
	"github.com/go-webauthn/webauthn/protocol"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sonrhq/sonr/x/service"
)

var _ service.QueryServer = queryServer{}

// NewQueryServerImpl returns an implementation of the module QueryServer.
func NewQueryServerImpl(k Keeper) service.QueryServer {
	return queryServer{k}
}

type queryServer struct {
	k Keeper
}

// Params defines the handler for the Query/Params RPC method.
func (qs queryServer) Params(ctx context.Context, req *service.QueryParamsRequest) (*service.QueryParamsResponse, error) {
	params, err := qs.k.Params.Get(ctx)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return &service.QueryParamsResponse{Params: service.Params{}}, nil
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &service.QueryParamsResponse{Params: params}, nil
}

// Credentials defines the handler for the Query/Credentials RPC method.
func (qs queryServer) Credentials(ctx context.Context, req *service.QueryCredentialsRequest) (*service.QueryCredentialsResponse, error) {
	rec, err := qs.k.db.ServiceRecordTable().GetByOrigin(ctx, req.Origin)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if rec == nil {
		return nil, status.Error(codes.NotFound, "record not found")
	}

	if req.ParamsType == service.ParamsType_PARAMS_TYPE_ATTESTATION {
		opts := service.GetPublicKeyCredentialCreationOptions(rec, protocol.UserEntity{})
		creationOptsBz, err := json.Marshal(opts)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return &service.QueryCredentialsResponse{
			AttestationOptions: string(creationOptsBz),
			Origin:             rec.Origin,
		}, nil
	}

	if req.ParamsType == service.ParamsType_PARAMS_TYPE_ASSERTION {
		opts := service.GetPublicKeyCredentialRequestOptions(rec, []protocol.CredentialDescriptor{})
		requestOptsBz, err := json.Marshal(opts)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return &service.QueryCredentialsResponse{
			AssertionOptions: string(requestOptsBz),
			Origin:           rec.Origin,
		}, nil
	}
	return nil, status.Error(codes.InvalidArgument, "invalid params type")
}
