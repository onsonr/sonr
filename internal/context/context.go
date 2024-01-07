package context

import (
	"context"

	"github.com/ipfs/kubo/client/rpc"
)

type SonrContext struct {
	// BlockchainClient is the client for the gRPC blockchain endpoint.

	// IPFSClient is the client for the IPFS endpoint.
	IPFSClient *rpc.HttpApi

	// MatrixClient is the client for the Matrix endpoint.

	// Context is the context for the request.
    Context context.Context
}

func New(ctx context.Context) (*SonrContext, error) {
	ipfsC, err := rpc.NewLocalApi()
	if err != nil {
		return nil, err
	}
	return &SonrContext{
		IPFSClient: ipfsC,
        Context: ctx,
	}, nil
}
