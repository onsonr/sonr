package models

import (
	"encoding/json"
	"fmt"
)

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

func (c *CredentialDescriptor) ToDBModel(handle, origin string) *Credential {
	return &Credential{
		Handle:     handle,
		Origin:     origin,
		ID:         c.ID,
		Type:       c.Type,
		Transports: c.Transports,
	}
}

func ExtractCredentialDescriptor(jsonString string) (*CredentialDescriptor, error) {
	cred := &CredentialDescriptor{}
	// Unmarshal the credential JSON
	if err := json.Unmarshal([]byte(jsonString), cred); err != nil {
		return nil, err
	}

	// Validate required fields
	if cred.ID == "" || cred.RawID == "" {
		return nil, fmt.Errorf("missing credential ID")
	}
	if cred.Type != "public-key" {
		return nil, fmt.Errorf("invalid credential type")
	}
	if cred.Response.AttestationObject == "" || cred.Response.ClientDataJSON == "" {
		return nil, fmt.Errorf("missing attestation data")
	}

	// Log detailed credential information
	fmt.Printf("Credential Details:\n"+
		"ID: %s\n"+
		"Raw ID: %s\n"+
		"Type: %s\n"+
		"Authenticator Attachment: %s\n"+
		"Transports: %v\n"+
		cred.ID,
		cred.RawID,
		cred.Type,
		cred.AuthenticatorAttachment,
		cred.Transports,
	)
	return cred, nil
}
