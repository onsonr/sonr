package chain

import (
	"context"
	"errors"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/sonrhq/core/app"
	identitytypes "github.com/sonrhq/core/x/identity/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"google.golang.org/grpc"
)

type APIEndpoint string

const (
	// List of known origin api endpoints.
	SonrLocalRpcOrigin  APIEndpoint = "localhost:9090"
	SonrPublicRpcOrigin APIEndpoint = "142.93.116.204:9090"
)

func (e APIEndpoint) TcpAddress() string {
	return fmt.Sprintf("tcp://%s", string(e))
}

type SonrQueryClient struct {
	APIEndpoint string
}

// NewClient creates a new client for the sonr chain
func NewClient(apiEndpoint APIEndpoint) *SonrQueryClient {
	return &SonrQueryClient{APIEndpoint: string(apiEndpoint)}
}

// GetDID returns the DID document with the given id
func (c *SonrQueryClient) GetDID(ctx context.Context, id string) (*identitytypes.ResolvedDidDocument, error) {
	conn, err := grpc.Dial(c.APIEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, errors.New("failed to connect to grpc server: " + err.Error())
	}
	resp, err := identitytypes.NewQueryClient(conn).Did(ctx, &identitytypes.QueryGetDidRequest{Did: id})
	if err != nil {
		return nil, err
	}
	return &resp.DidDocument, nil
}

// GetAllDIDs returns all DID documents
func (c *SonrQueryClient) GetAllDIDs(ctx context.Context) ([]*identitytypes.DidDocument, error) {
	conn, err := grpc.Dial(c.APIEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, errors.New("failed to connect to grpc server: " + err.Error())
	}
	resp, err := identitytypes.NewQueryClient(conn).DidAll(ctx, &identitytypes.QueryAllDidRequest{})
	if err != nil {
		return nil, err
	}
	list := make([]*identitytypes.DidDocument, len(resp.DidDocument))
	for i, d := range resp.DidDocument {
		list[i] = &d
	}
	return list, nil
}

// GetService returns the service with the given id
func (c *SonrQueryClient) GetService(ctx context.Context, origin string) (*identitytypes.Service, error) {
	conn, err := grpc.Dial(c.APIEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, errors.New("failed to connect to grpc server: " + err.Error())
	}
	resp, err := identitytypes.NewQueryClient(conn).Service(ctx, &identitytypes.QueryGetServiceRequest{Origin: origin})
	if err != nil {
		return nil, err
	}
	return &resp.Service, nil
}

// GetAllServices returns all services
func (c *SonrQueryClient) GetAllServices(ctx context.Context) ([]*identitytypes.Service, error) {
	conn, err := grpc.Dial(c.APIEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, errors.New("failed to connect to grpc server: " + err.Error())
	}
	resp, err := identitytypes.NewQueryClient(conn).ServiceAll(ctx, &identitytypes.QueryAllServiceRequest{})
	if err != nil {
		return nil, err
	}
	list := make([]*identitytypes.Service, len(resp.Services))
	for i, s := range resp.Services {
		list[i] = &s
	}
	return list, nil
}

// BroadcastTx broadcasts a transaction to the sonr chain
func (c *SonrQueryClient) BroadcastTx(ctx context.Context, tx []byte) (*ctypes.ResultBroadcastTx, error) {
	client, err := client.NewClientFromNode(c.APIEndpoint)
	if err != nil {
		return nil, err
	}

	res, err := client.BroadcastTxAsync(ctx, tx)
	if err != nil {
		return nil, err
	}

	// Print the transaction hash.
	fmt.Printf("Transaction log: %s\n", res.Log)
	return res, nil
}
