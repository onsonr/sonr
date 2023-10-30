// Utility functions for DID Service - https://w3c.github.io/did-core/#services
// I.e. Service Endpoints for IPFS Cluster
package types

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/sonr-io/sonr/types/webauthn"
)

const (
	VaultServiceType            = "EncryptedVault"
	LinkedDomainServiceType     = "LinkedDomains"
	DIDCommMessagingServiceType = "DIDCommMessaging"
)

func (s *ServiceRecord) GetUserEntity(id string) protocol.UserEntity {
	return protocol.UserEntity{
		ID:          []byte(id),
		DisplayName: id,
		CredentialEntity: protocol.CredentialEntity{
			Name: s.Name,
		},
	}
}

// GetCredentialCreationOptions issues a challenge for the VerificationMethod to sign and return
func (vm *ServiceRecord) GetCredentialCreationOptions(username string, chal protocol.URLEncodedBase64) (string, error) {
	params := DefaultParams()
	rp := vm.RelyingPartyEntity()
	cco, err := params.NewWebauthnCreationOptions(rp, username, chal)
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
func (vm *ServiceRecord) GetCredentialAssertionOptions(allowedCredentials []protocol.CredentialDescriptor, chal protocol.URLEncodedBase64) (string, error) {
	if len(allowedCredentials) == 0 {
		return "", errors.New("No allowed credentials")
	}
	params := DefaultParams()
	cco, err := params.NewWebauthnAssertionOptions(vm, chal, allowedCredentials)
	if err != nil {
		return "", err
	}
	ccoJSON, err := json.Marshal(cco)
	if err != nil {
		return "", err
	}
	return string(ccoJSON), nil
}

// NewServiceRelationship creates a new service relationship for an Identification authenticated by
// a ServiceRecord
func (s *ServiceRecord) NewServiceRelationship(id string) *ServiceRelationship {
	return &ServiceRelationship{
		Reference: s.Id,
		Did:       id,
		Count:     0,
	}
}

// RelyingPartyEntity is a struct that represents a Relying Party entity.
func (s *ServiceRecord) RelyingPartyEntity() protocol.RelyingPartyEntity {
	return protocol.RelyingPartyEntity{
		ID: s.Id,
		CredentialEntity: protocol.CredentialEntity{
			Name: s.Name,
		},
	}
}

// VerifyCreationChallenge verifies the challenge and a creation signature and returns an error if it fails to verify
func (vm *ServiceRecord) VerifyCreationChallenge(resp string, chal string) (*webauthn.Credential, error) {
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

	err = pcc.Verify(chal, false, vm.RelyingPartyEntity().ID, vm.Origins)
	if err != nil {
		return makeCredentialFromCreationData(pcc), err
	}
	return makeCredentialFromCreationData(pcc), nil
}

// VeriifyAssertionChallenge verifies the challenge and an assertion signature and returns an error if it fails to verify
func (vm *ServiceRecord) VerifyAssertionChallenge(resp string) (*webauthn.Credential, error) {
	var ccr protocol.CredentialAssertionResponse
	err := json.Unmarshal([]byte(resp), &ccr)
	if err != nil {
		return nil, err
	}
	pca, err := ccr.Parse()
	if err != nil {
		return nil, err
	}
	cred := makeCredentialFromAssertionData(pca)
	return cred, nil
}

// GetBaseOrigin returns the origin url without a subdomain
func (vm *ServiceRecord) GetBaseOrigin() string {
	for _, origin := range vm.Origins {
		u, err := url.Parse(origin)
		if err != nil {
			continue // skip invalid URLs
		}

		hostparts := strings.Split(u.Hostname(), ".")
		if len(hostparts) > 2 {
			// Remove subdomain
			baseHost := strings.Join(hostparts[len(hostparts)-2:], ".")
			u.Host = baseHost
		}

		return u.String()
	}

	return "" // return empty string if no valid URLs
}
