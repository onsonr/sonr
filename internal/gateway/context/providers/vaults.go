package providers

import (
	"context"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/onsonr/sonr/internal/gateway/views"
	"github.com/onsonr/sonr/pkg/ipfsapi"
)

type VaultProvider interface {
	New(ctx context.Context, handle string, origin string) (views.CreatePasskeyData, error)
}

type VaultService struct {
	tokenStore     ipfsapi.IPFSTokenStore
	challengeCache map[string]protocol.URLEncodedBase64
}

func NewVaultService(ipc ipfsapi.Client) VaultProvider {
	svc := &VaultService{
		tokenStore: ipfsapi.NewUCANStore(ipc),
	}
	return svc
}

func (s *VaultService) New(ctx context.Context, handle string, origin string) (views.CreatePasskeyData, error) {
	return views.CreatePasskeyData{}, nil
}
