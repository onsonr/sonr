package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	keepertest "github.com/sonrhq/core/testutil/keeper"
	testkeeper "github.com/sonrhq/core/testutil/keeper"
	"github.com/sonrhq/core/testutil/nullify"
	"github.com/sonrhq/core/x/domain/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.DomainKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}

func TestSLDRecordQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.DomainKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNSLDRecord(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetSLDRecordRequest
		response *types.QueryGetSLDRecordResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetSLDRecordRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetSLDRecordResponse{SLDRecord: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetSLDRecordRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetSLDRecordResponse{SLDRecord: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetSLDRecordRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.SLDRecord(wctx, tc.request)
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

func TestSLDRecordQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.DomainKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNSLDRecord(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllSLDRecordRequest {
		return &types.QueryAllSLDRecordRequest{
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
			resp, err := keeper.SLDRecordAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SLDRecord), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.SLDRecord),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.SLDRecordAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SLDRecord), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.SLDRecord),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.SLDRecordAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.SLDRecord),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.SLDRecordAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func TestTLDRecordQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.DomainKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNTLDRecord(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetTLDRecordRequest
		response *types.QueryGetTLDRecordResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetTLDRecordRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetTLDRecordResponse{TLDRecord: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetTLDRecordRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetTLDRecordResponse{TLDRecord: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetTLDRecordRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.TLDRecord(wctx, tc.request)
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

func TestTLDRecordQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.DomainKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNTLDRecord(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllTLDRecordRequest {
		return &types.QueryAllTLDRecordRequest{
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
			resp, err := keeper.TLDRecordAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.TLDRecord), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.TLDRecord),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.TLDRecordAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.TLDRecord), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.TLDRecord),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.TLDRecordAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.TLDRecord),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.TLDRecordAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
