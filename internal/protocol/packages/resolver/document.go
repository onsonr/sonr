package resolver

import (
	"context"
	"errors"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/getsentry/sentry-go"
	// "github.com/sonrhq/core/app"
	"github.com/sonrhq/core/internal/local"
	identitytypes "github.com/sonrhq/core/x/identity/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"google.golang.org/grpc"
)

// GetDID returns the DID document with the given id
func GetDID(ctx context.Context, id string) (*identitytypes.ResolvedDidDocument, error) {
	snrctx := local.NewContext()
	conn, err := grpc.Dial(snrctx.GrpcEndpoint(), grpc.WithInsecure())
	if err != nil {
		sentry.CaptureException(err)
		return nil, errors.New("failed to connect to grpc server: " + err.Error())
	}
	resp, err := identitytypes.NewQueryClient(conn).Did(ctx, &identitytypes.QueryGetDidRequest{Did: id})
	if err != nil {
		return nil, err
	}
	return &resp.DidDocument, nil
}

// GetAllDIDs returns all DID documents
func GetAllDIDs(ctx context.Context) ([]*identitytypes.DidDocument, error) {
	snrctx := local.NewContext()
	conn, err := grpc.Dial(snrctx.GrpcEndpoint(), grpc.WithInsecure())
	if err != nil {
		sentry.CaptureException(err)
		return nil, errors.New("failed to connect to grpc server: " + err.Error())
	}
	resp, err := identitytypes.NewQueryClient(conn).DidAll(ctx, &identitytypes.QueryAllDidRequest{})
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}
	list := make([]*identitytypes.DidDocument, len(resp.DidDocument))
	for i, d := range resp.DidDocument {
		list[i] = &d
	}
	return list, nil
}

// GetService returns the service with the given id
func GetService(ctx context.Context, origin string) (*identitytypes.Service, error) {
	snrctx := local.NewContext()
	conn, err := grpc.Dial(snrctx.GrpcEndpoint(), grpc.WithInsecure())
	if err != nil {
		sentry.CaptureException(err)
		return nil, errors.New("failed to connect to grpc server: " + err.Error())
	}
	resp, err := identitytypes.NewQueryClient(conn).Service(ctx, &identitytypes.QueryGetServiceRequest{Origin: origin})
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}
	return &resp.Service, nil
}

// GetAllServices returns all services
func GetAllServices(ctx context.Context) ([]*identitytypes.Service, error) {
	snrctx := local.NewContext()
	conn, err := grpc.Dial(snrctx.GrpcEndpoint(), grpc.WithInsecure())
	if err != nil {
		sentry.CaptureException(err)
		return nil, errors.New("failed to connect to grpc server: " + err.Error())
	}
	resp, err := identitytypes.NewQueryClient(conn).ServiceAll(ctx, &identitytypes.QueryAllServiceRequest{})
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}
	list := make([]*identitytypes.Service, len(resp.Services))
	for i, s := range resp.Services {
		list[i] = &s
	}
	return list, nil
}

// BroadcastTx broadcasts a transaction to the sonr chain
func BroadcastTx(ctx context.Context, tx []byte) (*ctypes.ResultBroadcastTx, error) {
	snrctx := local.NewContext()
	// endpoint := currEndpoint()
	client, err := client.NewClientFromNode(snrctx.RpcEndpoint())
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	res, err := client.BroadcastTxAsync(ctx, tx)
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}
	return res, nil
}
