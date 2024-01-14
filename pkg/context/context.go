package context

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/ipfs/kubo/client/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SonrContext interface {
	// IPFSClient() *rpc.HttpApi
	GrpcConn() *grpc.ClientConn
}

type context struct{}

func Get() SonrContext {
	return &context{}
}

func (c *context) GrpcConn() *grpc.ClientConn {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		"127.0.0.1:9090", // your gRPC server address.
		grpc.WithTransportCredentials(insecure.NewCredentials()), // The Cosmos SDK doesn't support any transport security mechanism.
		// This instantiates a general gRPC codec which handles proto bytes. We pass in a nil interface registry
		// if the request/response types contain interface instead of 'nil' you should pass the application specific codec.
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
	)
	if err != nil {
		panic(err)
	}
	return grpcConn
}

func (c *context) IPFSClient() *rpc.HttpApi {
	// The `IPFSClient()` function is a method of the `context` struct that returns an instance of the `rpc.HttpApi` type.
	ipfsC, err := rpc.NewLocalApi()
	if err != nil {
		panic(err)
	}
	return ipfsC
}
