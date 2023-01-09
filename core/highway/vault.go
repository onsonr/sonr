package highway

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/node/ipfs"
	"github.com/sonr-hq/sonr/pkg/vault"
	"github.com/sonr-hq/sonr/pkg/vault/mpc"
	v1 "github.com/sonr-hq/sonr/third_party/types/highway/vault/v1"
	"github.com/sonr-hq/sonr/x/identity/types"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
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
	highway ipfs.IPFS
	rpName  string
	rpIcon  string
	cache   *gocache.Cache
}

// It creates a new VaultService and registers it with the gRPC server
func NewVaultService(ctx context.Context, mux *runtime.ServeMux, hway ipfs.IPFS, cache *gocache.Cache) (*VaultService, error) {
	srv := &VaultService{
		cache:   cache,
		highway: hway,
		// TODO: Make all Webauthn options configurable through cmd line flags
		rpName: "Sonr",
		rpIcon: "https://raw.githubusercontent.com/sonr-hq/sonr/master/docs/static/favicon.png",
	}
	err := v1.RegisterVaultHandlerServer(ctx, mux, srv)
	if err != nil {
		return nil, err
	}
	return srv, nil
}

// Challeng returns a random challenge for the client to sign.
func (v *VaultService) Challenge(ctx context.Context, req *v1.ChallengeRequest) (*v1.ChallengeResponse, error) {
	// Cache the challenge for 2 minutes
	session, err := v.makeNewSession(req.GetRpId())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to initialization a new session with challenge: %s", err))
	}
	bz, err := session.Marshal()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to marshal session: %s", err))
	}
	v.cache.Set(session.Id, bz, -1)
	return &v1.ChallengeResponse{
		RpName:    v.rpName,
		RpOrigins: session.RpOrigins,
		Challenge: session.Challenge,
		SessionId: session.Id,
		RpIcon:    v.rpIcon,
	}, nil
}

// Register registers a new keypair and returns the public key.
func (v *VaultService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterResponse, error) {
	// Get Raw Session from cache
	value, ok := v.cache.Get(req.SessionId)
	if !ok {
		return nil, errors.New(fmt.Sprintf("Failed to get session from cache for Session %s", req.SessionId))
	}

	// Parse Session
	session := &v1.Session{}
	err := session.Unmarshal(value.([]byte))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to unmarshal session: %s", err))
	}

	// Parse Client Credential Data
	pcc, err := getParsedCredentialCreationData(req.CredentialResponse)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to get parsed creation data: %s", err))
	}

	// Verify the challenge
	err = pcc.Verify(session.Challenge, false, session.RpId, session.RpOrigins)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to verify session with client credential data: %s", err))
	}

	// Get WebauthnCredential
	cred := common.NewWebAuthnCredential(pcc)
	vm, _ := types.NewWebAuthnVM(cred)

	// Return Register Response
	return &v1.RegisterResponse{
		Success:            true,
		VerificationMethod: vm,
	}, nil

}

// Keygen generates a new keypair and returns the public key.
func (v *VaultService) Keygen(ctx context.Context, req *v1.KeygenRequest) (*v1.KeygenResponse, error) {
	// Create a new offline wallet
	wallet, err := vault.NewWallet(ctx, req.Prefix, v.highway)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to create new offline wallet using MPC: %s", err))
	}

	// Fetch public key of resulting wallet shares
	pubKey, err := wallet.PublicKey()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to retreive new wallet pubKey: %s", err))
	}

	// Get raw bytes of public key
	pbBz, err := pubKey.Marshal()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to marshal public key: %s", err))
	}

	// Marshal Vault Recovery Share
	bz, err := wallet.Find("vault").Share().Marshal()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to marshal Vault WalletShare: %s", err))
	}

	// Add Encrypted WalletShare to IPFS
	cid, err := v.highway.AddEncrypted(bz, pbBz)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to add encrypted WalletShare to IPFS: %s", err))
	}

	// Return Configuration Response
	return &v1.KeygenResponse{
		Id:          []byte(uuid.New().String()),
		Address:     wallet.Address(),
		VaultCid:    cid,
		ShareConfig: wallet.Find("current").Share(),
	}, nil
}

// Refresh refreshes the keypair and returns the public key.
func (v *VaultService) Refresh(ctx context.Context, req *v1.RefreshRequest) (*v1.RefreshResponse, error) {
	self, wallet, err := v.assembleWalletFromShares(req.VaultCid, req.ShareConfig)
	if err != nil {
		return nil, err
	}

	newWallet, err := wallet.Refresh(self)
	if err != nil {
		return nil, err
	}
	share := newWallet.Find("vault").Share()
	bz, err := share.Marshal()
	if err != nil {
		return nil, err
	}
	cid, err := v.highway.Add(bz)
	if err != nil {
		return nil, err
	}
	return &v1.RefreshResponse{
		Id:          []byte(uuid.New().String()),
		Address:     newWallet.Address(),
		VaultCid:    cid,
		ShareConfig: newWallet.Find(party.ID(req.ShareConfig.SelfId)).Share(),
	}, nil
}

// Sign signs the data with the private key and returns the signature.
func (v *VaultService) Sign(ctx context.Context, req *v1.SignRequest) (*v1.SignResponse, error) {
	self, wallet, err := v.assembleWalletFromShares(req.VaultCid, req.ShareConfig)
	if err != nil {
		return nil, err
	}
	sig, err := wallet.Sign(self, req.Data)
	if err != nil {
		return nil, err
	}
	return &v1.SignResponse{
		Id:        []byte(uuid.New().String()),
		Signature: sig,
		Data:      req.Data,
		Creator:   wallet.Address(),
	}, nil
}

// Derive derives a new key from the private key and returns the public key.
func (v *VaultService) Derive(ctx context.Context, req *v1.DeriveRequest) (*v1.DeriveResponse, error) {
	s, err := mpc.LoadWalletShare(req.GetShareConfig())
	if err != nil {
		return nil, err
	}
	ws, err := s.Bip32Derive(req.GetChildIndex())
	if err != nil {
		return nil, err
	}

	share := ws.Share()
	bz, err := share.Marshal()
	if err != nil {
		return nil, err
	}

	cid, err := v.highway.Add(bz)
	if err != nil {
		return nil, err
	}
	return &v1.DeriveResponse{
		Id:          []byte(uuid.New().String()),
		Address:     ws.Address(),
		VaultCid:    cid,
		ShareConfig: ws.Share(),
	}, nil
}

//
// Helper functions
//

// assembleWalletFromShares takes a WalletShareConfig and CID to return a Offline Wallet
func (v *VaultService) assembleWalletFromShares(cid string, current *common.WalletShareConfig) (party.ID, common.Wallet, error) {
	// Initialize provided share
	shares := make([]*common.WalletShareConfig, 0)
	shares = append(shares, current)

	// Fetch Vault share from IPFS
	oldbz, err := v.highway.Get(cid)
	if err != nil {
		return "", nil, err
	}

	// Unmarshal share
	share := &common.WalletShareConfig{}
	err = share.Unmarshal(oldbz)
	if err != nil {
		return "", nil, err
	}

	// Load wallet
	wallet, err := vault.LoadOfflineWallet(shares)
	if err != nil {
		return "", nil, err
	}
	return party.ID(current.SelfId), wallet, nil
}

// makeNewSession builds a default session for the given user.
func (v *VaultService) makeNewSession(rpId string) (*v1.Session, error) {
	sessionID := uuid.New().String()[:8]

	// Generate a challenge
	bz := make([]byte, 32)
	_, err := rand.Read(bz)
	if err != nil {
		return nil, err
	}

	// Base64 encode the challenge
	challenge := base64.StdEncoding.EncodeToString(bz)
	return &v1.Session{
		Id:        sessionID,
		Challenge: challenge,
		RpId:      rpId,
		RpOrigins: defaultRpOrigins,
	}, nil
}

// It takes a JSON string, converts it to a struct, and then converts that struct to a different struct
func getParsedCredentialCreationData(bz []byte) (*protocol.ParsedCredentialCreationData, error) {
	// Get Credential Creation Response
	var ccr protocol.CredentialCreationResponse
	err := json.Unmarshal(bz, &ccr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to unmarshal credential response: %v", err))
	}

	if err != nil {
		return nil, errors.New("Parse error for Registration")
	}

	if ccr.ID == "" {
		return nil, errors.New("Parse error for Registration")
	}
	testB64, err := base64.RawURLEncoding.DecodeString(ccr.ID)
	if err != nil || !(len(testB64) > 0) {
		return nil, errors.New("Parse error for Registration")
	}

	if ccr.PublicKeyCredential.Credential.Type == "" {
		return nil, errors.New("Parse error for Registration")
	}

	if ccr.PublicKeyCredential.Credential.Type != "public-key" {
		return nil, errors.New("Parse error for Registration")
	}

	var pcc protocol.ParsedCredentialCreationData
	pcc.ID, pcc.RawID, pcc.Type, pcc.ClientExtensionResults = ccr.ID, ccr.RawID, ccr.Type, ccr.ClientExtensionResults
	pcc.Raw = ccr

	for _, t := range ccr.Transports {
		pcc.Transports = append(pcc.Transports, protocol.AuthenticatorTransport(t))
	}

	parsedAttestationResponse, err := ccr.AttestationResponse.Parse()
	if err != nil {
		return nil, errors.New("Error parsing attestation response")
	}
	pcc.Response = *parsedAttestationResponse
	return &pcc, nil
}
