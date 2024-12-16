package handlers

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/middleware"
	"github.com/onsonr/sonr/internal/gateway/models"
)

// SubmitProfileHandle submits a profile handle
func SubmitProfileHandle(c echo.Context) error {
	return nil
}

// SubmitPublicKeyCredential submits a public key credential
func SubmitPublicKeyCredential(c echo.Context) error {
	credentialJSON := c.FormValue("credential")
	cred := &models.CredentialDescriptor{}
	// Unmarshal the credential JSON
	if err := json.Unmarshal([]byte(credentialJSON), cred); err != nil {
		return middleware.RenderError(c, err)
	}
	err := middleware.SubmitCredential(c, cred)
	if err != nil {
		return middleware.RenderError(c, err)
	}
	return nil
}
