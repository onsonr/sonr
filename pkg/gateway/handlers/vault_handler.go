package handlers

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/gateway/context"
)

// SubmitProfileHandle submits a profile handle
func SubmitProfileHandle(c echo.Context) error {
	return nil
}

// SubmitPublicKeyCredential submits a public key credential
func SubmitPublicKeyCredential(c echo.Context) error {
	credentialJSON := c.FormValue("credential")
	cred := &context.CredentialDescriptor{}
	// Unmarshal the credential JSON
	if err := json.Unmarshal([]byte(credentialJSON), cred); err != nil {
		return context.RenderError(c, err)
	}
	err := context.SubmitCredential(c, cred)
	if err != nil {
		return context.RenderError(c, err)
	}
	return nil
}
