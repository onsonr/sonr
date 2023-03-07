package service

import (
	"context"
	"fmt"

	"github.com/sonrhq/core/pkg/common"
	"github.com/sonrhq/core/x/identity/protocol/dispatcher"
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
	// // Get Session
	controller, err := v.Dispatcher.BuildNewDIDController()
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
