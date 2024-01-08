package context

import (
	"context"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/ipfs/kubo/client/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SonrContext struct {

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

func (c *SonrContext) getGrpcConn() (*grpc.ClientConn, error) {
	 // Create a connection to the gRPC server.
    grpcConn, err := grpc.Dial(
        "127.0.0.1:9090", // your gRPC server address.
        grpc.WithTransportCredentials(insecure.NewCredentials()), // The Cosmos SDK doesn't support any transport security mechanism.
        // This instantiates a general gRPC codec which handles proto bytes. We pass in a nil interface registry
        // if the request/response types contain interface instead of 'nil' you should pass the application specific codec.
        grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
    )
    if err != nil {
        return nil, err
    }
	return grpcConn, nil
}
