package keeper

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sonr-io/sonr/x/bucket/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) BucketsAll(c context.Context, req *types.QueryAllBucketsRequest) (*types.QueryAllBucketsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var whereIss []types.Bucket
	ctx := sdk.UnwrapSDKContext(c)

	if req.Pagination == nil {
		whereIss = k.GetAllWhereIs(ctx)
		return &types.QueryAllBucketsResponse{
			Buckets:    whereIss,
			Pagination: nil,
		}, nil
	}

	store := ctx.KVStore(k.storeKey)
	whereIsStore := prefix.NewStore(store, types.KeyPrefix(types.BucketKeyPrefix))

	pageRes, err := query.Paginate(whereIsStore, req.Pagination, func(key []byte, value []byte) error {
		var whereIs types.Bucket
		if err := k.cdc.Unmarshal(value, &whereIs); err != nil {
			return err
		}

		whereIss = append(whereIss, whereIs)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, "error while paginating response: "+err.Error())
	}

	return &types.QueryAllBucketsResponse{
		Buckets:    whereIss,
		Pagination: pageRes,
	}, nil
}

func (k Keeper) Bucket(c context.Context, req *types.QueryGetBucketRequest) (*types.QueryGetBucketResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	whereIs, found := k.GetBucket(ctx, req.Uuid)
	if !found {
		return nil, fmt.Errorf("error while querying whereIs: %s %s", req.Uuid, sdkerrors.ErrKeyNotFound)
	}

	return &types.QueryGetBucketResponse{Bucket: whereIs}, nil
}
