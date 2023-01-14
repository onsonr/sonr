package vault

import (
	"context"
	"errors"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sonr-hq/sonr/pkg/node/config"
	"github.com/sonr-hq/sonr/pkg/vault/bank"
	"github.com/sonr-hq/sonr/pkg/vault/internal/session"

	v1 "github.com/sonr-hq/sonr/third_party/types/highway/vault/v1"
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

// `VaultService` is a type that implements the `v1.VaultServer` interface, and has a field called
// `highway` of type `*HighwayNode`.
// @property  - `v1.VaultServer`: This is the interface that the Vault service implements.
// @property highway - This is the HighwayNode that the VaultService is running on.
type VaultService struct {
	bank   *bank.VaultBank
	node   config.IPFSNode
	rpName string
	rpIcon string
	cctx   client.Context
}

// It creates a new VaultService and registers it with the gRPC server
func NewService(ctx client.Context, mux *runtime.ServeMux, hway config.IPFSNode, cache *gocache.Cache) (*VaultService, error) {
	vaultBank := bank.CreateBank(hway, cache)
	srv := &VaultService{
		cctx:   ctx,
		bank:   vaultBank,
		node:   hway,
		rpName: "Sonr",
		rpIcon: "https://raw.githubusercontent.com/sonr-hq/sonr/master/docs/static/favicon.png",
	}
	err := v1.RegisterVaultHandlerServer(context.Background(), mux, srv)
	if err != nil {
		return nil, err
	}
	return srv, nil
}

// Challeng returns a random challenge for the client to sign.
func (v *VaultService) Challenge(ctx context.Context, req *v1.ChallengeRequest) (*v1.ChallengeResponse, error) {
	session, err := session.NewEntry(req.RpId, req.Username)
	if err != nil {
		return nil, err
	}
	optsJson, eID, err := v.bank.StartRegistration(session)
	if err != nil {
		return nil, err
	}
	return &v1.ChallengeResponse{
		RpName:          v.rpName,
		CreationOptions: optsJson,
		SessionId:       eID,
		RpIcon:          v.rpIcon,
	}, nil
}

// Register registers a new keypair and returns the public key.
func (v *VaultService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterResponse, error) {
	// Get Session
	didDoc, wallet, err := v.bank.FinishRegistration(req.SessionId, req.CredentialResponse)
	if err != nil {
		return nil, err
	}
	err = didDoc.SetRootWallet(wallet)
	if err != nil {
		return nil, err
	}
	didDoc.WebAuthnCredentials()
	return &v1.RegisterResponse{
		Success:     true,
		DidDocument: didDoc,
		Address:     wallet.Address(),
	}, nil

}

// Refresh refreshes the keypair and returns the public key.
func (v *VaultService) Refresh(ctx context.Context, req *v1.RefreshRequest) (*v1.RefreshResponse, error) {
	return nil, errors.New("Method is unimplemented")

}

// Sign signs the data with the private key and returns the signature.
func (v *VaultService) Sign(ctx context.Context, req *v1.SignRequest) (*v1.SignResponse, error) {
	return nil, errors.New("Method is unimplemented")
}

// Derive derives a new key from the private key and returns the public key.
func (v *VaultService) Derive(ctx context.Context, req *v1.DeriveRequest) (*v1.DeriveResponse, error) {
	return nil, errors.New("Method is unimplemented")
}
