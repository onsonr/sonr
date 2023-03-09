package chain

import (
	"context"
	"errors"

	"github.com/sonrhq/core/x/identity/types"
	"google.golang.org/grpc"
)

type SonrQueryClient struct {
	APIEndpoint string
}

// NewClient creates a new client for the sonr chain
func NewClient(apiEndpoint APIEndpoint) *SonrQueryClient {
	return &SonrQueryClient{APIEndpoint: string(apiEndpoint)}
}

// GetDID returns the DID document with the given id
func (c *SonrQueryClient) GetDID(ctx context.Context, id string) (*types.DidDocument, error) {
	conn, err := grpc.Dial(c.APIEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, errors.New("failed to connect to grpc server: " + err.Error())
	}
	resp, err := types.NewQueryClient(conn).Did(ctx, &types.QueryGetDidRequest{Did: id})
	if err != nil {
		return nil, err
	}
	return &resp.DidDocument, nil
}

// GetAllDIDs returns all DID documents
func (c *SonrQueryClient) GetAllDIDs(ctx context.Context) ([]*types.DidDocument, error) {
	conn, err := grpc.Dial(c.APIEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, errors.New("failed to connect to grpc server: " + err.Error())
	}
	resp, err := types.NewQueryClient(conn).DidAll(ctx, &types.QueryAllDidRequest{})
	if err != nil {
		return nil, err
	}
	list := make([]*types.DidDocument, len(resp.DidDocument))
	for i, d := range resp.DidDocument {
		list[i] = &d
	}
	return list, nil
}

// GetService returns the service with the given id
func (c *SonrQueryClient) GetService(ctx context.Context, origin string) (*types.Service, error) {
	conn, err := grpc.Dial(c.APIEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, errors.New("failed to connect to grpc server: " + err.Error())
	}
	resp, err := types.NewQueryClient(conn).Service(ctx, &types.QueryGetServiceRequest{Origin: origin})
	if err != nil {
		return nil, err
	}
	return &resp.Service, nil
}

// GetAllServices returns all services
func (c *SonrQueryClient) GetAllServices(ctx context.Context) ([]*types.Service, error) {
	conn, err := grpc.Dial(c.APIEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, errors.New("failed to connect to grpc server: " + err.Error())
	}
	resp, err := types.NewQueryClient(conn).ServiceAll(ctx, &types.QueryAllServiceRequest{})
	if err != nil {
		return nil, err
	}
	list := make([]*types.Service, len(resp.Services))
	for i, s := range resp.Services {
		list[i] = &s
	}
	return list, nil
}
