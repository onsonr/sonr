package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sonr-io/sonr/testutil/nullify"
	"github.com/sonr-io/sonr/x/identity/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func (suite *KeeperTestSuite) TestDidDocumentQuerySingle() {
	keeper := suite.keeper
	ctx := suite.ctx
	wctx := suite.wCtx
	msgs := createNDidDocument(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetDidRequest
		response *types.QueryGetDidResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetDidRequest{
				Did: msgs[0].ID,
			},
			response: &types.QueryGetDidResponse{DidDocument: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetDidRequest{
				Did: msgs[1].ID,
			},
			response: &types.QueryGetDidResponse{DidDocument: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetDidRequest{
				Did: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		suite.T().Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Did(wctx, tc.request)
			if tc.err != nil {
				suite.Assert().ErrorIs(err, tc.err)
			} else {
				suite.Assert().NoError(err)
				suite.Assert().Equal(
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestDidDocumentQueryPaginated() {
	keeper := suite.keeper
	ctx := suite.ctx
	wctx := suite.wCtx
	msgs := createNDidDocument(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllDidRequest {
		return &types.QueryAllDidRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	suite.T().Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.DidAll(wctx, request(nil, uint64(i), uint64(step), false))
			suite.Assert().NoError(err)
			suite.Assert().LessOrEqual(len(resp.DidDocument), step)
			suite.Assert().Subset(
				nullify.Fill(msgs),
				nullify.Fill(resp.DidDocument),
			)
		}
	})
	suite.T().Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.DidAll(wctx, request(next, 0, uint64(step), false))
			suite.Assert().NoError(err)
			suite.Assert().LessOrEqual(len(resp.DidDocument), step)
			suite.Assert().Subset(
				nullify.Fill(msgs),
				nullify.Fill(resp.DidDocument),
			)
			next = resp.Pagination.NextKey
		}
	})
	suite.T().Run("Total", func(t *testing.T) {
		resp, err := keeper.DidAll(wctx, request(nil, 0, 0, true))
		suite.Assert().NoError(err)
		suite.Assert().Equal(len(msgs), int(resp.Pagination.Total))
		suite.Assert().ElementsMatch(
			nullify.Fill(msgs),
			nullify.Fill(resp.DidDocument),
		)
	})
	suite.T().Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.DidAll(wctx, nil)
		suite.Assert().ErrorIs(err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func (suite *KeeperTestSuite) TestQueryByService() {
	keeper := suite.keeper
	ctx := suite.ctx
	wctx := suite.wCtx
	msgs := createDidDocumentsWithPrefix(keeper, ctx, "SVC", 2)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryByServiceRequest
		response *types.QueryByServiceResponse
		err      error
	}{
		{
			desc: "FirstDocFirstService",
			request: &types.QueryByServiceRequest{
				ServiceId: msgs[0].Service.Data[0].ID,
			},
			response: &types.QueryByServiceResponse{DidDocument: msgs[0]},
		},
		{
			desc: "SecondDocSecondService",
			request: &types.QueryByServiceRequest{
				ServiceId: msgs[1].Service.Data[1].ID,
			},
			response: &types.QueryByServiceResponse{DidDocument: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryByServiceRequest{
				ServiceId: "did:snr:id#NON_EXISTENT_SVC_ID",
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		suite.T().Run(tc.desc, func(t *testing.T) {
			response, err := keeper.QueryByService(wctx, tc.request)
			if tc.err != nil {
				suite.Assert().ErrorIs(err, tc.err)
			} else {
				suite.Assert().NoError(err)
				suite.Assert().Equal(
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryByKeyID() {
	keeper := suite.keeper
	ctx := suite.ctx
	wctx := suite.wCtx
	msgs := createDidDocumentsWithPrefix(keeper, ctx, "Key", 1)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryByKeyIDRequest
		response *types.QueryByKeyIDResponse
		err      error
	}{
		{
			desc: "Doc",
			request: &types.QueryByKeyIDRequest{
				KeyId: msgs[0].VerificationMethod.Data[0].ID,
			},
			response: &types.QueryByKeyIDResponse{DidDocument: msgs[0]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryByKeyIDRequest{
				KeyId: "did:snr:id#NON_EXISTENT_KEY_ID",
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		suite.T().Run(tc.desc, func(t *testing.T) {
			response, err := keeper.QueryByKeyID(wctx, tc.request)
			if tc.err != nil {
				suite.Assert().ErrorIs(err, tc.err)
			} else {
				suite.Assert().NoError(err)
				suite.Assert().Equal(
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}

}

func (suite *KeeperTestSuite) TestQueryByAKA() {
	keeper := suite.keeper
	ctx := suite.ctx
	wctx := suite.wCtx
	msgs := createDidDocumentsWithPrefix(keeper, ctx, "AKA", 1)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryByAlsoKnownAsRequest
		response *types.QueryByAlsoKnownAsResponse
		err      error
	}{
		{
			desc: "Doc",
			request: &types.QueryByAlsoKnownAsRequest{
				AkaId: msgs[0].AlsoKnownAs[0],
			},
			response: &types.QueryByAlsoKnownAsResponse{DidDocument: msgs[0]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryByAlsoKnownAsRequest{
				AkaId: "NON_EXISTENT_AKA",
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		suite.T().Run(tc.desc, func(t *testing.T) {
			response, err := keeper.QueryByAlsoKnownAs(wctx, tc.request)
			if tc.err != nil {
				suite.Assert().ErrorIs(err, tc.err)
			} else {
				suite.Assert().NoError(err)
				suite.Assert().Equal(
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}

}
