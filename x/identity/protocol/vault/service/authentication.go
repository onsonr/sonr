package service

import (
	"context"
	"fmt"

	"github.com/sonrhq/core/pkg/client/chain"
	"github.com/sonrhq/core/pkg/common"
	"github.com/sonrhq/core/x/identity/protocol/dispatcher"
	"github.com/sonrhq/core/x/identity/protocol/vault/handler"
	v1 "github.com/sonrhq/core/x/identity/types/vault/v1"
)

// `AuthenticationService` is a type that implements the `v1.VaultServer` interface, and has a field called
// `highway` of type `*HighwayNode`.
// @property  - `v1.VaultServer`: This is the interface that the Vault service implements.
// @property highway - This is the HighwayNode that the AuthenticationService is running on.
type AuthenticationService struct {
	Node       common.IPFSNode
	Dispatcher *dispatcher.Dispatcher
}

// Register registers a new keypair and returns the public key.
func (v *AuthenticationService) RegisterStart(ctx context.Context, req *v1.RegisterStartRequest) (*v1.RegisterStartResponse, error) {
	// Get service handler
	handler, err := handler.NewServiceHandler(req.Origin, chain.SonrLocalRpcOrigin)
	if err != nil {
		return nil, fmt.Errorf("failed to get service handler: %w", err)
	}

	opts, err := handler.BeginRegistration(req)
	if err != nil {
		return nil, fmt.Errorf("failed to begin registration: %w", err)
	}

	// Return response
	return &v1.RegisterStartResponse{
		CreationOptions: string(opts),
	}, nil
}

// CreateAccount derives a new key from the private key and returns the public key.
func (v *AuthenticationService) RegisterFinish(ctx context.Context, req *v1.RegisterFinishRequest) (*v1.RegisterFinishResponse, error) {
	return nil, fmt.Errorf("Method is unimplemented")
}

// LoginStart returns a challenge to be signed by the user.
func (v *AuthenticationService) LoginStart(ctx context.Context, req *v1.LoginStartRequest) (*v1.LoginStartResponse, error) {
	return nil, fmt.Errorf("Method is unimplemented")
}

// LoginFinish returns a challenge to be signed by the user.
func (v *AuthenticationService) LoginFinish(ctx context.Context, req *v1.LoginFinishRequest) (*v1.LoginFinishResponse, error) {
	return nil, fmt.Errorf("Method is unimplemented")
}
