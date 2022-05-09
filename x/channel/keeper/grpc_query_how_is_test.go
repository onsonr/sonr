package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/sonr-io/sonr/testutil/keeper"
	"github.com/sonr-io/sonr/testutil/nullify"
	"github.com/sonr-io/sonr/x/channel/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestHowIsQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.ChannelKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNHowIs(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryHowIsRequest
		response *types.QueryHowIsResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryHowIsRequest{
				Did: msgs[0].Did,
			},
			response: &types.QueryHowIsResponse{HowIs: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryHowIsRequest{
				Did: msgs[1].Did,
			},
			response: &types.QueryHowIsResponse{HowIs: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryHowIsRequest{
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
			response, err := keeper.HowIs(wctx, tc.request)
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

func TestHowIsQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.ChannelKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNHowIs(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllHowIsRequest {
		return &types.QueryAllHowIsRequest{
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
			resp, err := keeper.HowIsAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.HowIs), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.HowIs),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.HowIsAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.HowIs), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.HowIs),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.HowIsAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.HowIs),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.HowIsAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
