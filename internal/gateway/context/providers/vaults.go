package providers

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/onsonr/sonr/crypto/mpc"
	"github.com/onsonr/sonr/internal/gateway/views"
	"github.com/onsonr/sonr/pkg/ipfsapi"
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
	encl, err := mpc.GenEnclave()
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
