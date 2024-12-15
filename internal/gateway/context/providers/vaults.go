package providers

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/onsonr/sonr/crypto/mpc"
	"github.com/onsonr/sonr/internal/gateway/models"
	"github.com/onsonr/sonr/pkg/ipfsapi"
	"lukechampine.com/blake3"
)

type VaultProvider interface {
	Spawn(sessionID string, handle string, origin string, challenge string) (models.CreatePasskeyParams, error)
	Claim(sessionID string, handle string, origin string) (models.CreatePasskeyParams, error)
}

type VaultProviderService struct {
	ipfsClient ipfsapi.Client
	tokenStore ipfsapi.IPFSTokenStore

	challengeCache map[string]protocol.URLEncodedBase64
	stagedEnclaves map[string]mpc.Enclave
}

func NewVaultService(ipc ipfsapi.Client) VaultProvider {
	svc := &VaultProviderService{
		ipfsClient:     ipc,
		challengeCache: make(map[string]protocol.URLEncodedBase64),
		stagedEnclaves: make(map[string]mpc.Enclave),
		tokenStore:     ipfsapi.NewUCANStore(ipc),
	}
	return svc
}

func (s *VaultProviderService) Spawn(sessionID string, handle string, origin string, challenge string) (models.CreatePasskeyParams, error) {
	nonce, err := calcNonce(sessionID)
	if err != nil {
		return models.CreatePasskeyParams{
			Address:       "",
			Handle:        handle,
			Name:          origin,
			Challenge:     challenge,
			CreationBlock: "00001",
		}, err
	}
	encl, err := mpc.GenEnclave(nonce)
	if err != nil {
		return models.CreatePasskeyParams{}, err
	}
	s.stagedEnclaves[sessionID] = encl
	return models.CreatePasskeyParams{
		Address:       encl.Address(),
		Handle:        handle,
		Name:          origin,
		Challenge:     challenge,
		CreationBlock: "00001",
	}, nil
}

func (s *VaultProviderService) Claim(sessionID string, handle string, origin string) (models.CreatePasskeyParams, error) {
	return models.CreatePasskeyParams{}, nil
}

// Uses blake3 to hash the sessionID to generate a nonce of length 12 bytes
func calcNonce(sessionID string) ([]byte, error) {
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
