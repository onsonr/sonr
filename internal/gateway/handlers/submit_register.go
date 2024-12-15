package handlers

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/context/repository"
)

// SubmitProfileHandle submits a profile handle
func SubmitProfileHandle(c echo.Context) error {
	return nil
}

// SubmitPublicKeyCredential submits a public key credential
func SubmitPublicKeyCredential(c echo.Context) error {
	credentialJSON := c.FormValue("credential")
	cred := &CredentialDescriptor{}
	// Unmarshal the credential JSON
	if err := json.Unmarshal([]byte(credentialJSON), cred); err != nil {
		return err
	}
	return nil
}

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
