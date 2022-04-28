package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sonr-io/sonr/internal/blockchain/x/channel/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) HowIsAll(c context.Context, req *types.QueryAllHowIsRequest) (*types.QueryAllHowIsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var howIss []types.HowIs
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	howIsStore := prefix.NewStore(store, types.KeyPrefix(types.HowIsKeyPrefix))

	pageRes, err := query.Paginate(howIsStore, req.Pagination, func(key []byte, value []byte) error {
		var howIs types.HowIs
		if err := k.cdc.Unmarshal(value, &howIs); err != nil {
			return err
		}

		howIss = append(howIss, howIs)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllHowIsResponse{HowIs: howIss, Pagination: pageRes}, nil
}

func (k Keeper) HowIs(c context.Context, req *types.QueryHowIsRequest) (*types.QueryHowIsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetHowIs(
		ctx,
		req.Did,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	// Check if Channel is IsActive
	if !val.IsActive {
		return nil, types.ErrInactiveChannel
	}

	return &types.QueryHowIsResponse{HowIs: val}, nil
}
