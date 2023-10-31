package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/sonrhq/sonr/testutil/keeper"
	"github.com/sonrhq/sonr/testutil/nullify"
	"github.com/sonrhq/sonr/x/domain/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestUsernameRecordsQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.DomainKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNUsernameRecords(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetUsernameRecordsRequest
		response *types.QueryGetUsernameRecordsResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetUsernameRecordsRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetUsernameRecordsResponse{UsernameRecords: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetUsernameRecordsRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetUsernameRecordsResponse{UsernameRecords: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetUsernameRecordsRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.UsernameRecord(wctx, tc.request)
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

func TestUsernameRecordsQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.DomainKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNUsernameRecords(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllUsernameRecordsRequest {
		return &types.QueryAllUsernameRecordsRequest{
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
			resp, err := keeper.UsernameRecordAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.UsernameRecords), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.UsernameRecords),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.UsernameRecordAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.UsernameRecords), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.UsernameRecords),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.UsernameRecordAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.UsernameRecords),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.UsernameRecordAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
