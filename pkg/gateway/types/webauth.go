package types

import "github.com/onsonr/sonr/internal/database/repository"

// Define the credential structure matching our frontend data
type CredentialDescriptor struct {
	ID                      string            `json:"id"`
	RawID                   string            `json:"rawId"`
	Type                    string            `json:"type"`
	AuthenticatorAttachment string            `json:"authenticatorAttachment"`
	Transports              string            `json:"transports"`
	ClientExtensionResults  map[string]string `json:"clientExtensionResults"`
	Response                struct {
		AttestationObject string `json:"attestationObject"`
		ClientDataJSON    string `json:"clientDataJSON"`
	} `json:"response"`
}

func (c *CredentialDescriptor) ToModel(handle, origin string) *repository.Credential {
	return &repository.Credential{
		Handle:                  handle,
		Origin:                  origin,
		CredentialID:            c.ID,
		Type:                    c.Type,
		Transports:              c.Transports,
		AuthenticatorAttachment: c.AuthenticatorAttachment,
	}
}

func CredentialArrayToDescriptors(credentials []repository.Credential) []*CredentialDescriptor {
	var descriptors []*CredentialDescriptor
	for _, cred := range credentials {
		cd := &CredentialDescriptor{
			ID:                      cred.CredentialID,
			RawID:                   cred.CredentialID,
			Type:                    cred.Type,
			AuthenticatorAttachment: cred.AuthenticatorAttachment,
			Transports:              cred.Transports,
		}
		descriptors = append(descriptors, cd)
	}
	return descriptors
}
