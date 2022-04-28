package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/sonr-io/sonr/internal/blockchain/testutil/keeper"
	"github.com/sonr-io/sonr/internal/blockchain/testutil/nullify"
	"github.com/sonr-io/sonr/internal/blockchain/x/object/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestWhatIsQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.ObjectKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNWhatIs(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryWhatIsRequest
		response *types.QueryWhatIsResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryWhatIsRequest{
				Did: msgs[0].Did,
			},
			response: &types.QueryWhatIsResponse{WhatIs: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryWhatIsRequest{
				Did: msgs[1].Did,
			},
			response: &types.QueryWhatIsResponse{WhatIs: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryWhatIsRequest{
				Did: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.WhatIs(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestWhatIsQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.ObjectKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNWhatIs(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllWhatIsRequest {
		return &types.QueryAllWhatIsRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.WhatIsAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.WhatIs), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.WhatIs),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.WhatIsAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.WhatIs), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.WhatIs),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.WhatIsAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.WhatIs),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.WhatIsAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
