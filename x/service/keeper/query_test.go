package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/sonrhq/core/testutil/keeper"
	testkeeper "github.com/sonrhq/core/testutil/keeper"
	"github.com/sonrhq/core/testutil/nullify"
	"github.com/sonrhq/core/x/service/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestServiceRecordQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.ServiceKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNServiceRecord(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryServiceRecordRequest
		response *types.QueryServiceRecordResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryServiceRecordRequest{
				Origin: msgs[0].Id,
			},
			response: &types.QueryServiceRecordResponse{ServiceRecord: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryServiceRecordRequest{
				Origin: msgs[1].Id,
			},
			response: &types.QueryServiceRecordResponse{ServiceRecord: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryServiceRecordRequest{
				Origin: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ServiceRecord(wctx, tc.request)
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

func TestServiceRecordQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.ServiceKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNServiceRecord(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.ListServiceRecordsRequest {
		return &types.ListServiceRecordsRequest{
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
			resp, err := keeper.ListServiceRecords(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ServiceRecord), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ServiceRecord),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ListServiceRecords(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ServiceRecord), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ServiceRecord),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ListServiceRecords(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ServiceRecord),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ListServiceRecords(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.ServiceKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
