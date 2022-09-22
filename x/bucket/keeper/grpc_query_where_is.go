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

func (k Keeper) WhereIsAll(c context.Context, req *types.QueryAllWhereIsRequest) (*types.QueryAllWhereIsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var whereIss []types.WhereIs
	ctx := sdk.UnwrapSDKContext(c)

	if req.Pagination == nil {
		whereIss = k.GetAllWhereIs(ctx)
		return &types.QueryAllWhereIsResponse{
			WhereIs:    whereIss,
			Pagination: nil,
		}, nil
	}

	store := ctx.KVStore(k.storeKey)
	whereIsStore := prefix.NewStore(store, types.KeyPrefix(types.WhereIsKeyPrefix))

	pageRes, err := query.Paginate(whereIsStore, req.Pagination, func(key []byte, value []byte) error {
		var whereIs types.WhereIs
		if err := k.cdc.Unmarshal(value, &whereIs); err != nil {
			return err
		}

		whereIss = append(whereIss, whereIs)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, "error while paginating response: "+err.Error())
	}

	return &types.QueryAllWhereIsResponse{
		WhereIs:    whereIss,
		Pagination: pageRes,
	}, nil
}

func (k Keeper) WhereIs(c context.Context, req *types.QueryGetWhereIsRequest) (*types.QueryGetWhereIsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	whereIs, found := k.GetWhereIs(ctx, req.Creator, req.Did)
	if !found {
		return nil, fmt.Errorf("error while querying whereIs: %s %s", req.Did, sdkerrors.ErrKeyNotFound)
	}

	return &types.QueryGetWhereIsResponse{WhereIs: whereIs}, nil
}
