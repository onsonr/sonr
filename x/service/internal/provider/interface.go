package provider

import (
	"fmt"
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/patrickmn/go-cache"
	idtypes "github.com/sonrhq/core/x/identity/types"
	"github.com/sonrhq/core/x/service/internal/provider/utils"
	"github.com/sonrhq/core/x/service/types"
)

const (
	defaultExpiration = 5 * time.Minute
	purgeTime         = 10 * time.Minute
)

type Session struct {
	Alias     string                    `json:"alias"`
	Address   string                    `json:"address"`
	Challenge protocol.URLEncodedBase64 `json:"challenge"`
	IsMobile  bool                      `json:"isMobile"`
}

type ServiceProvider interface {
	CredentialEntity() protocol.CredentialEntity
	GetServiceRecord() *types.ServiceRecord
	GetCredentialCreationOptions(alias string, addr string, isMobile bool) (string, error)
	GetCredentialAssertionOptions(alias string, allowedCredentials []protocol.CredentialDescriptor, isMobile bool) (string, error)
	RelyingPartyEntity() protocol.RelyingPartyEntity
	VerifyCreationChallenge(resp string, alias string) (*types.WebauthnCredential, error)
	VerifyAssertionChallenge(resp string, creds ...*idtypes.VerificationMethod) error
}

func NewServiceProvider(record *types.ServiceRecord) ServiceProvider {
	return &serviceProvider{
		cache:  cache.New(defaultExpiration, purgeTime),
		record: record,
	}
}

func (sp *serviceProvider) NewSession(alias, address string, isMobile bool) (*Session, error) {
	chal, err := utils.CreateChallenge()
	if err != nil {
		return nil, fmt.Errorf("failed to create challenge: %w", err)
	}
	s := Session{
		Alias:     alias,
		Address:   address,
		IsMobile:  isMobile,
		Challenge: chal,
	}
	sp.cache.Set(alias, &s, defaultExpiration)
	return &s, nil
}

func (sp *serviceProvider) GetSession(alias string) (*Session, error) {
	s, ok := sp.cache.Get(alias)
	if !ok {
		return nil, fmt.Errorf("session not found")
	}
	return s.(*Session), nil
}
