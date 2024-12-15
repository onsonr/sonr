package providers

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/onsonr/sonr/crypto/mpc"
	views "github.com/onsonr/sonr/internal/gateway/models"
	"github.com/onsonr/sonr/pkg/ipfsapi"
	"lukechampine.com/blake3"
)

type VaultProvider interface {
	Spawn(sessionID string, handle string, origin string, challenge string) (views.CreatePasskeyData, error)
	Claim(sessionID string, handle string, origin string) (views.CreatePasskeyData, error)
}

type VaultService struct {
	ipfsClient     ipfsapi.Client
	tokenStore     ipfsapi.IPFSTokenStore
	challengeCache map[string]protocol.URLEncodedBase64
	stagedEnclaves map[string]mpc.Enclave
}

func NewVaultService(ipc ipfsapi.Client) VaultProvider {
	svc := &VaultService{
		ipfsClient:     ipc,
		challengeCache: make(map[string]protocol.URLEncodedBase64),
		stagedEnclaves: make(map[string]mpc.Enclave),
		tokenStore:     ipfsapi.NewUCANStore(ipc),
	}
	return svc
}

func (s *VaultService) Spawn(sessionID string, handle string, origin string, challenge string) (views.CreatePasskeyData, error) {
	nonce, err := computeNonceFromSessionID(sessionID)
	if err != nil {
		return views.CreatePasskeyData{}, err
	}
	encl, err := mpc.GenEnclave(nonce)
	if err != nil {
		return views.CreatePasskeyData{}, err
	}
	s.stagedEnclaves[sessionID] = encl
	return views.CreatePasskeyData{
		Address:       encl.Address(),
		Handle:        handle,
		Name:          origin,
		Challenge:     challenge,
		CreationBlock: "00001",
	}, nil
}

func (s *VaultService) Claim(sessionID string, handle string, origin string) (views.CreatePasskeyData, error) {
	return views.CreatePasskeyData{}, nil
}

// Uses blake3 to hash the sessionID to generate a nonce of length 12 bytes
func computeNonceFromSessionID(sessionID string) ([]byte, error) {
	hash := blake3.New(32, nil)
	_, err := hash.Write([]byte(sessionID))
	if err != nil {
		return nil, err
	}
	// Read the hash into a byte slice
	nonce := make([]byte, 12)
	_, err = hash.Write(nonce)
	if err != nil {
		return nil, err
	}
	return nonce, nil
}
