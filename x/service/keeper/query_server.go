package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/didao-org/sonr/x/service/types"
)

var _ types.QueryServer = Querier{}

// NewQueryServerImpl returns an implementation of the module QueryServer.
func NewQueryServerImpl(k Keeper) types.QueryServer {
	return Querier{k}
}

type Querier struct {
	k Keeper
}

// ServiceRecord defines the handler for the Query/ServiceRecord RPC method.
func (qs Querier) ServiceRecord(goCtx context.Context, req *types.QueryServiceRecordRequest) (*types.QueryServiceRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	record, err := qs.k.RecordsMapping.Get(ctx, req.Origin)
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
