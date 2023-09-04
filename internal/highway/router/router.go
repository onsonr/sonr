package router

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	authenticationpb "github.com/sonrhq/core/types/highway/authentication/v1"
	databasepb "github.com/sonrhq/core/types/highway/database/v1"
	storagepb "github.com/sonrhq/core/types/highway/storage/v1"
	walletpb "github.com/sonrhq/core/types/highway/wallet/v1"
	"google.golang.org/grpc"
)

// RegisterRouter registers the Highway Services to the root mux.
func RegisterRouter(mux *runtime.ServeMux, baseURL string) {
	authenticationpb.RegisterAuthenticationServiceHandlerFromEndpoint(context.Background(), mux, endpointAuth(baseURL), getDialOpts())
	databasepb.RegisterDatabaseServiceHandlerFromEndpoint(context.Background(), mux, endpointDB(baseURL), getDialOpts())
	storagepb.RegisterStorageServiceHandlerFromEndpoint(context.Background(), mux, endpointStore(baseURL), getDialOpts())
	walletpb.RegisterWalletServiceHandlerFromEndpoint(context.Background(), mux, endpointWallet(baseURL), getDialOpts())
}

func getDialOpts() []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithInsecure(),
	}
}
