package handler

import (
	"github.com/cosmos/cosmos-sdk/client"
	authenticationpb "github.com/sonrhq/core/types/highway/authentication/v1"
	databasepb "github.com/sonrhq/core/types/highway/database/v1"
	storagepb "github.com/sonrhq/core/types/highway/storage/v1"
	walletpb "github.com/sonrhq/core/types/highway/wallet/v1"
	grpc "google.golang.org/grpc"
)

// RegisterHandlers registers the Highway Service Server.
func RegisterHandlers(cctx client.Context, grpcServer *grpc.Server) {
	authenticationpb.RegisterAuthenticationServiceServer(grpcServer, &AuthenticationHandler{cctx: cctx})
	databasepb.RegisterDatabaseServiceServer(grpcServer, &DatabaseHandler{cctx: cctx})
	storagepb.RegisterStorageServiceServer(grpcServer, &StorageHandler{cctx: cctx})
	walletpb.RegisterWalletServiceServer(grpcServer, &WalletHandler{cctx: cctx})
}
