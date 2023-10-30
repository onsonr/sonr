package keeper

import (
	"context"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sonr-io/sonr/x/service/types"
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
	serviceRecordStore := prefix.NewStore(store, types.KeyPrefix(types.ServiceRecordKey))

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
	serviceRecord, found := k.GetServiceRecord(ctx, req.Origin)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetServiceRecordResponse{ServiceRecord: serviceRecord}, nil
}

func (k Keeper) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

// Removes www. and https:// from the origin url
func cleanOriginUrl(url string) string {
	url = strings.Replace(url, "www.", "", 1)
	url = strings.Replace(url, "https://", "", 1)
	return url
}
