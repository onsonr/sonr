package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sonr-io/sonr/internal/blockchain/x/object/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) WhatIsAll(c context.Context, req *types.QueryAllWhatIsRequest) (*types.QueryAllWhatIsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var whatIss []types.WhatIs
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	whatIsStore := prefix.NewStore(store, types.KeyPrefix(types.WhatIsKeyPrefix))

	pageRes, err := query.Paginate(whatIsStore, req.Pagination, func(key []byte, value []byte) error {
		var whatIs types.WhatIs
		if err := k.cdc.Unmarshal(value, &whatIs); err != nil {
			return err
		}

		whatIss = append(whatIss, whatIs)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllWhatIsResponse{WhatIs: whatIss, Pagination: pageRes}, nil
}

func (k Keeper) WhatIs(c context.Context, req *types.QueryWhatIsRequest) (*types.QueryWhatIsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetWhatIs(
		ctx,
		req.Did,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	// Check if Object is IsActive
	if !val.IsActive {
		return nil, types.ErrInactiveObject
	}

	return &types.QueryWhatIsResponse{WhatIs: val}, nil
}
