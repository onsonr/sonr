package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/registry/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) WhoIsAlias(goCtx context.Context, req *types.QueryWhoIsAliasRequest) (*types.QueryWhoIsAliasResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoIsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.WhoIs
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.ContainsAlias(req.GetAlias()) {
			return &types.QueryWhoIsAliasResponse{
				WhoIs: &val,
			}, nil
		}
	}

	return &types.QueryWhoIsAliasResponse{}, nil
}
