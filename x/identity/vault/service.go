package vault

import (
	"context"
	"errors"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sonr-hq/sonr/pkg/node"
	"github.com/sonr-hq/sonr/pkg/node/config"
	v1 "github.com/sonr-hq/sonr/x/identity/types/vault/v1"
	"github.com/sonr-hq/sonr/x/identity/vault/store"
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
	bank   *store.VaultBank
	node   config.IPFSNode
	rpName string
	rpIcon string
	ctx    context.Context
	cctx   client.Context
	cache  *gocache.Cache
}

// It creates a new VaultService and registers it with the gRPC server
func RegisterGRPCGatewayRoutes(cctx client.Context, mux *runtime.ServeMux) error {
	ctx := context.Background()
	node, err := node.NewIPFS(ctx, config.WithClientContext(cctx, true))
	if err != nil {
		return err
	}
	cache := gocache.New(time.Minute*2, time.Minute*10)
	vaultService = &VaultService{
		ctx:    ctx,
		cctx:   cctx,
		bank:   store.InitBank(node, cache),
		node:   node,
		rpName: "Sonr",
		rpIcon: "https://raw.githubusercontent.com/sonr-hq/sonr/master/docs/static/favicon.png",
		cache:  cache,
	}
	err = v1.RegisterVaultHandlerServer(context.Background(), mux, vaultService)
	if err != nil {
		return err
	}
	return nil
}

// Challeng returns a random challenge for the client to sign.
func (v *VaultService) Challenge(ctx context.Context, req *v1.ChallengeRequest) (*v1.ChallengeResponse, error) {
	session, err := store.NewEntry(req.RpId, req.Username)
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
func (v *VaultService) NewWallet(ctx context.Context, req *v1.NewWalletRequest) (*v1.NewWalletResponse, error) {
	// Get Session
	didDoc, wallet, err := v.bank.FinishRegistration(req.SessionId, req.CredentialResponse)
	if err != nil {
		return nil, err
	}
	return &v1.NewWalletResponse{
		Success:     true,
		Address:     wallet.Address,
		DidDocument: didDoc,
	}, nil
}

// CreateAccount derives a new key from the private key and returns the public key.
func (v *VaultService) CreateAccount(ctx context.Context, req *v1.CreateAccountRequest) (*v1.CreateAccountResponse, error) {
	return nil, errors.New("Method is unimplemented")
}

// ListAccounts lists all the accounts derived from the private key.
func (v *VaultService) ListAccounts(ctx context.Context, req *v1.ListAccountsRequest) (*v1.ListAccountsResponse, error) {
	return nil, errors.New("Method is unimplemented")
}

// DeleteAccount deletes the account with the given address.
func (v *VaultService) DeleteAccount(ctx context.Context, req *v1.DeleteAccountRequest) (*v1.DeleteAccountResponse, error) {
	return nil, errors.New("Method is unimplemented")
}

// Refresh refreshes the keypair and returns the public key.
func (v *VaultService) Refresh(ctx context.Context, req *v1.RefreshRequest) (*v1.RefreshResponse, error) {
	return nil, errors.New("Method is unimplemented")
}

// Sign signs the data with the private key and returns the signature.
func (v *VaultService) SignTransaction(ctx context.Context, req *v1.SignTransactionRequest) (*v1.SignTransactionResponse, error) {
	return nil, errors.New("Method is unimplemented")
}
