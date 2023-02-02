package vault

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sonrhq/core/pkg/common"
	"github.com/sonrhq/core/x/identity/protocol/vault/dispatcher"
	v1 "github.com/sonrhq/core/x/identity/types/vault/v1"
)

// Default Variables
var (
	defaultRpOrigins = []string{
		"https://auth.sonr.io",
		"https://sonr.id",
		"https://sandbox.sonr.network",
		"localhost:3000",
	}
	vaultService *VaultService
)

// `VaultService` is a type that implements the `v1.VaultServer` interface, and has a field called
// `highway` of type `*HighwayNode`.
// @property  - `v1.VaultServer`: This is the interface that the Vault service implements.
// @property highway - This is the HighwayNode that the VaultService is running on.
type VaultService struct {
	node       common.IPFSNode
	dispatcher *dispatcher.Dispatcher
}

// It creates a new VaultService and registers it with the gRPC server
func RegisterVaultIPFSService(cctx client.Context, mux *runtime.ServeMux, node common.IPFSNode) error {
	vaultService = &VaultService{
		node:       node,
		dispatcher: dispatcher.New(),
	}
	return v1.RegisterVaultHandlerServer(context.Background(), mux, vaultService)
}

// Register registers a new keypair and returns the public key.
func (v *VaultService) RegisterStart(ctx context.Context, req *v1.RegisterStartRequest) (*v1.RegisterStartResponse, error) {
	// // Get Session
	controller, err := v.dispatcher.BuildNewDIDController()
	if err != nil {
		return nil, err
	}

	credOpts, err := controller.BeginRegistration(req.Aka)
	if err != nil {
		return nil, err
	}
	// Return response
	return &v1.RegisterStartResponse{
		AccountAddress:  controller.Address(),
		CreationOptions: string(credOpts),
		Aka:             req.Aka,
	}, nil
}

// CreateAccount derives a new key from the private key and returns the public key.
func (v *VaultService) RegisterFinish(ctx context.Context, req *v1.RegisterFinishRequest) (*v1.RegisterFinishResponse, error) {
	return nil, fmt.Errorf("Method is unimplemented")
}

// ListAccounts lists all the accounts derived from the private key.
func (v *VaultService) ListAccounts(ctx context.Context, req *v1.ListAccountsRequest) (*v1.ListAccountsResponse, error) {
	return nil, fmt.Errorf("Method is unimplemented")
}

// DeleteAccount deletes the account with the given address.
func (v *VaultService) DeleteAccount(ctx context.Context, req *v1.DeleteAccountRequest) (*v1.DeleteAccountResponse, error) {
	return nil, fmt.Errorf("Method is unimplemented")
}

// Refresh refreshes the keypair and returns the public key.
func (v *VaultService) Refresh(ctx context.Context, req *v1.RefreshRequest) (*v1.RefreshResponse, error) {
	return nil, fmt.Errorf("Method is unimplemented")
}
