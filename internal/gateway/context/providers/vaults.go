package providers

import (
	"context"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/onsonr/sonr/crypto/mpc"
	"github.com/onsonr/sonr/internal/gateway/views"
	"github.com/onsonr/sonr/pkg/ipfsapi"
)

type VaultProvider interface {
	Spawn(ctx context.Context, handle string, origin string, challenge string) (views.CreatePasskeyData, error)
	Claim(ctx context.Context, handle string, origin string) (views.CreatePasskeyData, error)
}

type VaultService struct {
	tokenStore     ipfsapi.IPFSTokenStore
	challengeCache map[string]protocol.URLEncodedBase64
	stagedEnclaves map[string]*mpc.Enclave
}

func NewVaultService(ipc ipfsapi.Client) VaultProvider {
	svc := &VaultService{
		tokenStore: ipfsapi.NewUCANStore(ipc),
	}
	return svc
}

func (s *VaultService) Spawn(ctx context.Context, handle string, origin string, challenge string) (views.CreatePasskeyData, error) {
	return views.CreatePasskeyData{
		Address:       "",
		Handle:        handle,
		Name:          origin,
		Challenge:     challenge,
		CreationBlock: "00001",
	}, nil
}

func (s *VaultService) Claim(ctx context.Context, handle string, origin string) (views.CreatePasskeyData, error) {
	return views.CreatePasskeyData{}, nil
}
