package service

import (
	"context"
	"fmt"

	"github.com/sonrhq/core/pkg/client/chain"
	"github.com/sonrhq/core/pkg/common"
	v1 "github.com/sonrhq/core/types/vault/v1"
	"github.com/sonrhq/core/x/identity/protocol/vault/handler"
)

// `AuthenticationService` is a type that implements the `v1.VaultServer` interface, and has a field called
// `highway` of type `*HighwayNode`.
// @property  - `v1.VaultServer`: This is the interface that the Vault service implements.
// @property highway - This is the HighwayNode that the AuthenticationService is running on.
type AuthenticationService struct {
	Node common.IPFSNode
}

// Register registers a new keypair and returns the public key.
func (v *AuthenticationService) RegisterStart(ctx context.Context, req *v1.RegisterStartRequest) (*v1.RegisterStartResponse, error) {
	// Get service handler
	handler, err := handler.NewServiceHandler(req.Origin, chain.SonrPublicRpcOrigin)
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
		RpId:            handler.GetRPID(),
		RpName:          handler.GetRPName(),
	}, nil
}

// CreateAccount derives a new key from the private key and returns the public key.
func (v *AuthenticationService) RegisterFinish(ctx context.Context, req *v1.RegisterFinishRequest) (*v1.RegisterFinishResponse, error) {
	// Get service handler
	handler, err := handler.NewServiceHandler(req.Origin, chain.SonrPublicRpcOrigin)
	if err != nil {
		return nil, fmt.Errorf("failed to get service handler: %w", err)
	}
	cred, err := handler.FinishRegistration(req)
	if err != nil {
		return nil, fmt.Errorf("failed to finish registration: %w", err)
	}
	return &v1.RegisterFinishResponse{
		Address: cred.Address().String(),
	}, nil
}

// LoginStart returns a challenge to be signed by the user.
func (v *AuthenticationService) LoginStart(ctx context.Context, req *v1.LoginStartRequest) (*v1.LoginStartResponse, error) {
	handler, err := handler.NewServiceHandler(req.Origin, chain.SonrPublicRpcOrigin)
	if err != nil {
		return nil, fmt.Errorf("failed to get service handler: %w", err)
	}

	opts, err := handler.BeginLogin(req)
	if err != nil {
		return nil, fmt.Errorf("failed to begin login: %w", err)
	}

	return &v1.LoginStartResponse{
		AccountAddress:    req.AccountAddress,
		CredentialOptions: string(opts),
		RpId:              handler.GetRPID(),
		RpName:            handler.GetRPName(),
	}, nil
}

// LoginFinish returns a challenge to be signed by the user.
func (v *AuthenticationService) LoginFinish(ctx context.Context, req *v1.LoginFinishRequest) (*v1.LoginFinishResponse, error) {
	handler, err := handler.NewServiceHandler(req.Origin, chain.SonrPublicRpcOrigin)
	if err != nil {
		return nil, fmt.Errorf("failed to get service handler: %w", err)
	}

	_, err = handler.FinishLogin(req)
	if err != nil {
		return nil, fmt.Errorf("failed to finish login: %w", err)
	}
	return nil, fmt.Errorf("Method is unimplemented")
}
