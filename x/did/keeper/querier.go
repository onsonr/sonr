package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc/peer"

	"github.com/onsonr/sonr/x/did/types"
)

var _ types.QueryServer = Querier{}

type Querier struct {
	Keeper
}

func NewQuerier(keeper Keeper) Querier {
	return Querier{Keeper: keeper}
}

// Params returns the total set of did parameters.
func (k Querier) Params(
	c context.Context,
	req *types.QueryRequest,
) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	p, err := k.Keeper.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	params := p.ActiveParams(k.HasIPFSConnection())
	return &types.QueryParamsResponse{Params: &params}, nil
}

// Resolve implements types.QueryServer.
func (k Querier) Resolve(
	goCtx context.Context,
	req *types.QueryRequest,
) (*types.QueryResolveResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryResolveResponse{}, nil
}

// Service implements types.QueryServer.
func (k Querier) Service(
	goCtx context.Context,
	req *types.QueryRequest,
) (*types.QueryServiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, ok := peer.FromContext(goCtx)
	if !ok {
		return nil, fmt.Errorf("failed to get peer from context")
	}

	rec, err := k.OrmDB.ServiceRecordTable().GetByOriginUri(ctx, req.Origin)
	if err != nil {
		return nil, err
	}
	return &types.QueryServiceResponse{Service: convertServiceRecord(rec)}, nil
}

// HTMX implements types.QueryServer.
func (k Querier) HTMX(goCtx context.Context, req *types.QueryRequest) (*httpbody.HttpBody, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &httpbody.HttpBody{
		ContentType: "text/html",
		Data:        []byte("<html><body>HTMX</body></html>"),
	}, nil
}
