package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sonrhq/core/x/identity/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DomainRecordAll(c context.Context, req *types.QueryAllDomainRecordRequest) (*types.QueryAllDomainRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var DomainRecords []types.DomainRecord
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	DomainRecordStore := prefix.NewStore(store, types.KeyPrefix(types.DomainRecordKeyPrefix))

	pageRes, err := query.Paginate(DomainRecordStore, req.Pagination, func(key []byte, value []byte) error {
		var DomainRecord types.DomainRecord
		if err := k.cdc.Unmarshal(value, &DomainRecord); err != nil {
			return err
		}

		DomainRecords = append(DomainRecords, DomainRecord)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDomainRecordResponse{DomainRecord: DomainRecords, Pagination: pageRes}, nil
}

func (k Keeper) DomainRecord(c context.Context, req *types.QueryGetDomainRecordRequest) (*types.QueryGetDomainRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetDomainRecord(
		ctx,
		req.Domain,
		req.Tld,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDomainRecordResponse{DomainRecord: val}, nil
}
