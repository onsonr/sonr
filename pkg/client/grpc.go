package client

import (
	"context"

	didv1 "github.com/onsonr/sonr/api/did/v1"
	macaroonv1 "github.com/onsonr/sonr/api/macaroon/v1"
	servicev1 "github.com/onsonr/sonr/api/service/v1"
	vaultv1 "github.com/onsonr/sonr/api/vault/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/cosmos/cosmos-sdk/codec"
)

const (
	kLocalGrpcWebAddr = "http://localhost:9090"
)

type Client struct {
	DIDClient      didv1.QueryClient
	MacaroonClient macaroonv1.QueryClient
	ServiceClient  servicev1.QueryClient
	VaultClient    vaultv1.QueryClient
}

func UseClient(ctx context.Context) (*Client, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.NewClient(kLocalGrpcWebAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())))
	if err != nil {
		return nil, err
	}
	return getGrpcClients(grpcConn), nil
}

func getGrpcClients(conn *grpc.ClientConn) *Client {
	didClient := didv1.NewQueryClient(conn)
	macaroonClient := macaroonv1.NewQueryClient(conn)
	serviceClient := servicev1.NewQueryClient(conn)
	vaultClient := vaultv1.NewQueryClient(conn)
	return &Client{
		DIDClient:      didClient,
		MacaroonClient: macaroonClient,
		ServiceClient:  serviceClient,
		VaultClient:    vaultClient,
	}
}
