package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
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

// ServiceRecord defines the handler for the Query/ServiceRecord RPC method.
func (qs queryServer) ServiceRecord(ctx context.Context, req *service.QueryServiceRecordRequest) (*service.QueryServiceRecordResponse, error) {
	record, err := qs.k.db.ServiceTable().GetByOrigin(ctx, req.Origin)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return &service.QueryServiceRecordResponse{}, nil
		}

		return nil, status.Error(codes.Internal, err.Error())
	}
	srv := service.ServiceRecord{
		Origin:      record.Origin,
		Name:        record.Name,
		Description: record.Description,
		Permissions: uint32(record.Permissions.Number()),
		Authority:   record.Authority,
	}
	return &service.QueryServiceRecordResponse{ServiceRecord: srv}, nil
}
