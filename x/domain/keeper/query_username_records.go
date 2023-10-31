package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sonrhq/sonr/x/domain/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UsernameRecordAll(goCtx context.Context, req *types.QueryAllUsernameRecordsRequest) (*types.QueryAllUsernameRecordsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var usernameRecordss []types.UsernameRecord
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	usernameRecordsStore := prefix.NewStore(store, types.KeyPrefix(types.UsernameRecordsKeyPrefix))

	pageRes, err := query.Paginate(usernameRecordsStore, req.Pagination, func(key []byte, value []byte) error {
		var UsernameRecord types.UsernameRecord
		if err := k.cdc.Unmarshal(value, &UsernameRecord); err != nil {
			return err
		}

		usernameRecordss = append(usernameRecordss, UsernameRecord)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUsernameRecordsResponse{UsernameRecords: usernameRecordss, Pagination: pageRes}, nil
}

func (k Keeper) UsernameRecord(goCtx context.Context, req *types.QueryGetUsernameRecordsRequest) (*types.QueryGetUsernameRecordsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetUsernameRecords(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetUsernameRecordsResponse{UsernameRecords: val}, nil
}
