package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/context/repository"
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

func (c *CredentialDescriptor) convertToDBModel(handle, origin string) *repository.Credential {
	return &repository.Credential{
		Handle:       handle,
		Origin:       origin,
		CredentialID: c.ID,
		Type:         c.Type,
		Transports:   c.Transports,
	}
}

// SubmitPublicKeyCredential submits a public key credential
func SubmitPublicKeyCredential(c echo.Context) error {
	credentialJSON := c.FormValue("credential")
	_, err := extractCredentialDescriptor(credentialJSON)
	if err != nil {
		return err
	}
	return nil
}

func extractCredentialDescriptor(jsonString string) (*CredentialDescriptor, error) {
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
