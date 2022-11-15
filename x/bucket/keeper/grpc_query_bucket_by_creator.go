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

func (k Keeper) BucketByCreator(c context.Context, req *types.QueryGetBucketByCreatorRequest) (*types.QueryGetBucketByCreatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	if req.Pagination == nil {
		whereIss := k.GetWhereIsByCreator(ctx, req.Creator)

		return &types.QueryGetBucketByCreatorResponse{
			Buckets:    whereIss,
			Pagination: nil,
		}, nil
	}

	store := ctx.KVStore(k.storeKey)
	whereIsStore := prefix.NewStore(store, types.KeyPrefix(types.BucketKeyPrefix))
	var whereIss []types.BucketConfig
	pageRes, err := query.Paginate(whereIsStore, req.Pagination, func(key []byte, value []byte) error {
		var whereIs types.BucketConfig
		if err := k.cdc.Unmarshal(value, &whereIs); err != nil {
			return err
		}

		if whereIs.Creator == req.Creator {
			whereIss = append(whereIss, whereIs)
		}

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, "error while panginating response"+err.Error())
	}

	return &types.QueryGetBucketByCreatorResponse{
		Buckets:    whereIss,
		Pagination: pageRes,
	}, nil
}
