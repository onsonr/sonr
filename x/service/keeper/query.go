package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sonrhq/core/x/service/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) ServiceRecordAll(goCtx context.Context, req *types.QueryAllServiceRecordRequest) (*types.QueryAllServiceRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var serviceRecords []types.ServiceRecord
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	serviceRecordStore := prefix.NewStore(store, types.KeyPrefix(types.ServiceRecordKeyPrefix))

	pageRes, err := query.Paginate(serviceRecordStore, req.Pagination, func(key []byte, value []byte) error {
		var serviceRecord types.ServiceRecord
		if err := k.cdc.Unmarshal(value, &serviceRecord); err != nil {
			return err
		}

		serviceRecords = append(serviceRecords, serviceRecord)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllServiceRecordResponse{ServiceRecord: serviceRecords, Pagination: pageRes}, nil
}

func (k Keeper) ServiceRecord(goCtx context.Context, req *types.QueryGetServiceRecordRequest) (*types.QueryGetServiceRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetServiceRecord(
		ctx,
		req.Id,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetServiceRecordResponse{ServiceRecord: val}, nil
}
