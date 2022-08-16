package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sonr-io/sonr/x/bucket/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) WhereIsByCreator(c context.Context, req *types.QueryGetWhereIsByCreatorRequest) (*types.QueryGetWhereIsByCreatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	var whereIss []types.WhereIs
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	whereIsStore := prefix.NewStore(store, types.WhereIsKey(types.MemStoreKey))

	pageRes, err := query.Paginate(whereIsStore, req.Pagination, func(key []byte, value []byte) error {
		var whereIs types.WhereIs
		if err := k.cdc.Unmarshal(value, &whereIs); err != nil {
			k.Logger(ctx).Error("Error while unMarshaling WhatIs: %s", err.Error())
			return err
		}

		if whereIs.Creator == req.Creator {
			whereIss = append(whereIss, whereIs)
		}

		return nil
	})

	if err != nil {
		k.Logger(ctx).Error("Error while querying whatIs: %s", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetWhereIsByCreatorResponse{
		WhereIs:    whereIss,
		Pagination: pageRes,
	}, nil
}
