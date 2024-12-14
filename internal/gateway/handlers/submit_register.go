package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/models"
)

// SubmitPublicKeyCredential submits a public key credential
func SubmitPublicKeyCredential(c echo.Context) error {
	credentialJSON := c.FormValue("credential")
	_, err := models.ExtractCredentialDescriptor(credentialJSON)
	if err != nil {
		return err
	}
	return nil
}
