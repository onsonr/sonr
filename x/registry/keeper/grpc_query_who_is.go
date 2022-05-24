package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sonr-io/sonr/x/registry/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) WhoIsAll(c context.Context, req *types.QueryAllWhoIsRequest) (*types.QueryAllWhoIsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var whoIss []types.WhoIs
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	whoIsStore := prefix.NewStore(store, types.KeyPrefix(types.WhoIsKeyPrefix))

	pageRes, err := query.Paginate(whoIsStore, req.Pagination, func(key []byte, value []byte) error {
		var whoIs types.WhoIs
		if err := k.cdc.Unmarshal(value, &whoIs); err != nil {
			return err
		}

		whoIss = append(whoIss, whoIs)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllWhoIsResponse{WhoIs: whoIss, Pagination: pageRes}, nil
}

func (k Keeper) WhoIs(c context.Context, req *types.QueryWhoIsRequest) (*types.QueryWhoIsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetWhoIs(
		ctx,
		req.GetDid(),
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryWhoIsResponse{WhoIs: &val}, nil
}
