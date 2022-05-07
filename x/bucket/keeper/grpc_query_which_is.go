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

func (k Keeper) WhichIsAll(c context.Context, req *types.QueryAllWhichIsRequest) (*types.QueryAllWhichIsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var whichIss []types.WhichIs
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	whichIsStore := prefix.NewStore(store, types.KeyPrefix(types.WhichIsKeyPrefix))

	pageRes, err := query.Paginate(whichIsStore, req.Pagination, func(key []byte, value []byte) error {
		var whichIs types.WhichIs
		if err := k.cdc.Unmarshal(value, &whichIs); err != nil {
			return err
		}

		whichIss = append(whichIss, whichIs)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllWhichIsResponse{WhichIs: whichIss, Pagination: pageRes}, nil
}

func (k Keeper) WhichIs(c context.Context, req *types.QueryWhichIsRequest) (*types.QueryWhichIsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetWhichIs(
		ctx,
		req.Did,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	// Check if Bucket is IsActive
	if !val.IsActive {
		return nil, types.ErrInactiveBucket
	}

	return &types.QueryWhichIsResponse{WhichIs: val}, nil
}
