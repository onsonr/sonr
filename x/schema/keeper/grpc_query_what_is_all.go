package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sonr-io/sonr/x/registry/types"
	st "github.com/sonr-io/sonr/x/schema/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) WhatIsAll(goCtx context.Context, req *st.QueryAllWhatIsRequest) (*st.QueryAllWhatIsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(k.storeKey)

	var whatIss []st.WhatIs
	whatIsStore := prefix.NewStore(store, types.KeyPrefix(st.SchemaKeyPrefix))

	pageRes, err := query.Paginate(whatIsStore, req.Pagination, func(key []byte, value []byte) error {
		var whatIs st.WhatIs
		if err := k.cdc.Unmarshal(value, &whatIs); err != nil {
			return err
		}

		whatIss = append(whatIss, whatIs)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &st.QueryAllWhatIsResponse{
		WhatIs:     whatIss,
		Pagination: pageRes,
	}, nil
}
