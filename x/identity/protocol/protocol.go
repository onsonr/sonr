package protocol

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sonrhq/core/pkg/common"
	"github.com/sonrhq/core/x/identity/protocol/dispatcher"
	"github.com/sonrhq/core/x/identity/protocol/vault/service"
	v1 "github.com/sonrhq/core/x/identity/types/vault/v1"
)

var (
	vaultService *service.AuthenticationService
)

// It creates a new VaultService and registers it with the gRPC server
func RegisterVaultIPFSService(cctx client.Context, mux *runtime.ServeMux, node common.IPFSNode) error {
	vaultService = &service.AuthenticationService{
		Node:       node,
		Dispatcher: dispatcher.New(),
	}
	return v1.RegisterVaultAuthenticationHandlerServer(context.Background(), mux, vaultService)
}
