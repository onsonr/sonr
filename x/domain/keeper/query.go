package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sonrhq/core/x/domain/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                SLD Record Query                                ||
// ! ||--------------------------------------------------------------------------------||

func (k Keeper) SLDRecordAll(goCtx context.Context, req *types.QueryAllSLDRecordRequest) (*types.QueryAllSLDRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var sLDRecords []types.SLDRecord
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	sLDRecordStore := prefix.NewStore(store, types.KeyPrefix(types.SLDRecordKeyPrefix))

	pageRes, err := query.Paginate(sLDRecordStore, req.Pagination, func(key []byte, value []byte) error {
		var sLDRecord types.SLDRecord
		if err := k.cdc.Unmarshal(value, &sLDRecord); err != nil {
			return err
		}

		sLDRecords = append(sLDRecords, sLDRecord)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSLDRecordResponse{SLDRecord: sLDRecords, Pagination: pageRes}, nil
}

func (k Keeper) SLDRecord(goCtx context.Context, req *types.QueryGetSLDRecordRequest) (*types.QueryGetSLDRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetSLDRecord(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetSLDRecordResponse{SLDRecord: val}, nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                TLD Record Query                                ||
// ! ||--------------------------------------------------------------------------------||

func (k Keeper) TLDRecordAll(goCtx context.Context, req *types.QueryAllTLDRecordRequest) (*types.QueryAllTLDRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tLDRecords []types.TLDRecord
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	tLDRecordStore := prefix.NewStore(store, types.KeyPrefix(types.TLDRecordKeyPrefix))

	pageRes, err := query.Paginate(tLDRecordStore, req.Pagination, func(key []byte, value []byte) error {
		var tLDRecord types.TLDRecord
		if err := k.cdc.Unmarshal(value, &tLDRecord); err != nil {
			return err
		}

		tLDRecords = append(tLDRecords, tLDRecord)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTLDRecordResponse{TLDRecord: tLDRecords, Pagination: pageRes}, nil
}

func (k Keeper) TLDRecord(goCtx context.Context, req *types.QueryGetTLDRecordRequest) (*types.QueryGetTLDRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetTLDRecord(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTLDRecordResponse{TLDRecord: val}, nil
}
