package handler

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	authenticationpb "github.com/sonr-io/core/types/highway/authentication/v1"
	databasepb "github.com/sonr-io/core/types/highway/database/v1"
	storagepb "github.com/sonr-io/core/types/highway/storage/v1"
	walletpb "github.com/sonr-io/core/types/highway/wallet/v1"
)

// RegisterHandlers registers the Highway Service Server.
func RegisterHandlers(ctx context.Context, mux *runtime.ServeMux) {
	authenticationpb.RegisterAuthenticationServiceHandlerServer(ctx, mux, &AuthenticationHandler{})
	databasepb.RegisterDatabaseServiceHandlerServer(ctx,mux, &DatabaseHandler{})
	storagepb.RegisterStorageServiceHandlerServer(ctx, mux,&StorageHandler{})
	walletpb.RegisterWalletServiceHandlerServer(ctx, mux,&WalletHandler{})
}
