package vault

import (
	"context"
	"errors"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/golang-jwt/jwt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sonr-hq/sonr/pkg/node/config"
	bank "github.com/sonr-hq/sonr/pkg/vault/keeper"
	v1 "github.com/sonr-hq/sonr/pkg/vault/types/v1"
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
		// sonrClient: client,
	}
	err := v1.RegisterVaultHandlerServer(context.Background(), mux, srv)
	if err != nil {
		return nil, err
	}
	return srv, nil
}

// Challeng returns a random challenge for the client to sign.
func (v *VaultService) Challenge(ctx context.Context, req *v1.ChallengeRequest) (*v1.ChallengeResponse, error) {
	session, err := bank.NewEntry(req.RpId, req.Username)
	if err != nil {
		return nil, err
	}
	optsJson, eID, err := v.bank.StartRegistration(session)
	if err != nil {
		return nil, err
	}
	jwt.NewWithClaims(jwt.SigningMethodES384, jwt.MapClaims{
		"creationOptions": optsJson,
		"sessionId":       eID,
		"rpName":          v.rpName,
		"rpIcon":          v.rpIcon,
	})
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
	// fs, err := fs.New(wallet.Address(), fs.WithClientContext(v.cctx, true))
	// if err != nil {
	// 	return nil, err
	// }
	// err = fs.StoreOfflineWallet(wallet)
	// if err != nil {
	// 	return nil, err
	// }
	// service, err := fs.Export(v.node)
	// if err != nil {
	// 	return nil, err
	// }
	// didDoc.AddService(service)
	// docReq := types.NewMsgCreateDidDocument(didDoc.Address(), didDoc)
	// res, err := wallet.SendTx("vault", docReq)
	// if err != nil {
	// 	return nil, err
	// }
	// if res.TxResponse.Code != 0 {
	// 	return nil, errors.New(res.TxResponse.RawLog)
	// }

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
