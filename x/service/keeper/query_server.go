package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/didao-org/sonr/x/service/types"
)

var _ types.QueryServer = queryServer{}

// NewQueryServerImpl returns an implementation of the module QueryServer.
func NewQueryServerImpl(k Keeper) types.QueryServer {
	return queryServer{k}
}

type queryServer struct {
	k Keeper
}

// ServiceRecord defines the handler for the Query/ServiceRecord RPC method.
func (qs queryServer) ServiceRecord(ctx context.Context, req *types.QueryServiceRecordRequest) (*types.QueryServiceRecordResponse, error) {
	cx := sdk.UnwrapSDKContext(ctx)
	record, err := qs.k.RecordsMapping.Get(cx, req.Origin)
	if err != nil {
		return nil, err
	}
	srv := types.Record{
		Origin:      record.Origin,
		Name:        record.Name,
		Description: record.Description,
		Authority:   record.Authority,
	}
	return &types.QueryServiceRecordResponse{ServiceRecord: srv}, nil
}
