package services

import (
	"context"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/onsonr/sonr/internal/gateway/models"
	"github.com/onsonr/sonr/pkg/ipfsapi"
)

type Vaults interface {
	New(ctx context.Context, handle string, origin string) (models.CreatePasskeyData, error)
}

type VaultService struct {
	tokenStore     ipfsapi.IPFSTokenStore
	challengeCache map[string]protocol.URLEncodedBase64
}

func NewVaultService(ipc ipfsapi.Client) Vaults {
	svc := &VaultService{
		tokenStore: ipfsapi.NewUCANStore(ipc),
	}
	return svc
}

func (s *VaultService) New(ctx context.Context, handle string, origin string) (models.CreatePasskeyData, error) {
	return models.CreatePasskeyData{}, nil
}
