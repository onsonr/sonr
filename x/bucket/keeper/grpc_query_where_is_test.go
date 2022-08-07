package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/sonr-io/sonr/testutil/keeper"
	"github.com/sonr-io/sonr/testutil/nullify"
	"github.com/sonr-io/sonr/x/bucket/types"
)

func TestWhereIsQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.BucketKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNWhereIs(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetWhereIsRequest
		response *types.QueryGetWhereIsResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetWhereIsRequest{Did: msgs[0].Did},
			response: &types.QueryGetWhereIsResponse{WhereIs: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetWhereIsRequest{Did: msgs[1].Did},
			response: &types.QueryGetWhereIsResponse{WhereIs: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetWhereIsRequest{Did: "not-found"},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.WhereIs(wctx, tc.request)
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

func TestWhereIsQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.BucketKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNWhereIs(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllWhereIsRequest {
		return &types.QueryAllWhereIsRequest{
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
			resp, err := keeper.WhereIsAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.WhereIs), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.WhereIs),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.WhereIsAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.WhereIs), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.WhereIs),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.WhereIsAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.WhereIs),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.WhereIsAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
