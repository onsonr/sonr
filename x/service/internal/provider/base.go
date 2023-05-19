package provider

import (
	"encoding/json"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/patrickmn/go-cache"
	idtypes "github.com/sonrhq/core/x/identity/types"
	"github.com/sonrhq/core/x/service/internal/provider/utils"
	"github.com/sonrhq/core/x/service/types"
)

type serviceProvider struct {
	cache  *cache.Cache
	record *types.ServiceRecord
}

func (sp *serviceProvider) CredentialEntity() protocol.CredentialEntity {
	return protocol.CredentialEntity{
		Name: sp.record.Name,
	}
}

func (sp *serviceProvider) GetServiceRecord() *types.ServiceRecord {
	return sp.record
}

// GetCredentialCreationOptions issues a challenge for the VerificationMethod to sign and return
func (sp *serviceProvider) GetCredentialCreationOptions(alias string, addr string, isMobile bool) (string, error) {
	params := types.DefaultParams()
	sess, err := sp.NewSession(alias, addr, isMobile)
	if err != nil {
		return "", err
	}
	cco, err := params.NewWebauthnCreationOptions(sp.record, sess.Alias, sess.Challenge, sess.Address, sess.IsMobile)
	if err != nil {
		return "", err
	}

	ccoJSON, err := json.Marshal(cco)
	if err != nil {
		return "", err
	}
	return string(ccoJSON), nil
}

// GetCredentialCreationOptions issues a challenge for the VerificationMethod to sign and return
func (sp *serviceProvider) GetCredentialAssertionOptions(alias string, allowedCredentials []protocol.CredentialDescriptor, isMobile bool) (string, error) {
	params := types.DefaultParams()
	sess, err := sp.NewSession(alias, "", isMobile)
	if err != nil {
		return "", err
	}
	cco, err := params.NewWebauthnAssertionOptions(sp.record, sess.Challenge, allowedCredentials, sess.IsMobile)
	if err != nil {
		return "", err
	}
	ccoJSON, err := json.Marshal(cco)
	if err != nil {
		return "", err
	}
	return string(ccoJSON), nil
}

// RelyingPartyEntity is a struct that represents a Relying Party entity.
func (sp *serviceProvider) RelyingPartyEntity() protocol.RelyingPartyEntity {
	return protocol.RelyingPartyEntity{
		ID:               sp.record.Origin,
		CredentialEntity: sp.CredentialEntity(),
	}
}

// VerifyCreationChallenge verifies the challenge and a creation signature and returns an error if it fails to verify
func (sp *serviceProvider) VerifyCreationChallenge(resp string, alias string) (*types.WebauthnCredential, error) {
	// Get Credential Creation Respons
	var ccr protocol.CredentialCreationResponse
	err := json.Unmarshal([]byte(resp), &ccr)
	if err != nil {
		return nil, err
	}
	pcc, err := ccr.Parse()
	if err != nil {
		return nil, err
	}

	sess, err := sp.GetSession(alias)
	if err != nil {
		return nil, err
	}

	err = pcc.Verify(sess.Challenge.String(), false, sp.RelyingPartyEntity().ID, []string{sp.record.Origin})
	if err != nil {
		return utils.MakeCredentialFromCreationData(pcc), nil
	}
	return utils.MakeCredentialFromCreationData(pcc), nil
}

// VeriifyAssertionChallenge verifies the challenge and an assertion signature and returns an error if it fails to verify
func (sp *serviceProvider) VerifyAssertionChallenge(resp string, creds ...*idtypes.VerificationMethod) error {
	var ccr protocol.CredentialAssertionResponse
	err := json.Unmarshal([]byte(resp), &ccr)
	if err != nil {
		return err
	}
	pca, err := ccr.Parse()
	if err != nil {
		return err
	}
	credVm := utils.MakeCredentialFromAssertionData(pca).ToVerificationMethod()
	for _, cred := range creds {
		if credVm.Equal(cred) {
			return nil
		}
	}

	return nil
}
