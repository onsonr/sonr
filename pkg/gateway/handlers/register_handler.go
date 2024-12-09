package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/crypto/mpc"
	"github.com/onsonr/sonr/pkg/blocks/forms"
	"github.com/onsonr/sonr/pkg/common/response"
	"github.com/onsonr/sonr/pkg/gateway/config"
	"github.com/onsonr/sonr/pkg/gateway/internal/pages/register"
)

func HandleRegisterView(env config.Env) echo.HandlerFunc {
	return func(c echo.Context) error {
		return response.TemplEcho(c, register.ProfileFormView(env.GetTurnstileSiteKey()))
	}
}

func HandleRegisterStart(c echo.Context) error {
	challenge, _ := protocol.CreateChallenge()
	handle := c.FormValue("handle")
	firstName := c.FormValue("first_name")
	lastName := c.FormValue("last_name")

	ks, err := mpc.NewKeyset()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	dat := forms.RegisterPasskeyData{
		Address:       ks.Address(),
		Handle:        handle,
		Name:          fmt.Sprintf("%s %s", firstName, lastName),
		Challenge:     challenge.String(),
		CreationBlock: "00001",
	}
	return response.TemplEcho(c, register.LinkCredentialView(dat))
}

func HandleRegisterFinish(c echo.Context) error {
	// Get the raw credential JSON string
	credentialJSON := c.FormValue("credential")
	if credentialJSON == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing credential data")
	}

	// Define the credential structure matching our frontend data
	var cred struct {
		ID                     string                 `json:"id"`
		RawID                  string                 `json:"rawId"`
		Type                   string                 `json:"type"`
		AuthenticatorAttachment string                `json:"authenticatorAttachment"`
		Transports            []string               `json:"transports"`
		ClientExtensionResults map[string]interface{} `json:"clientExtensionResults"`
		Response              struct {
			AttestationObject string `json:"attestationObject"`
			ClientDataJSON    string `json:"clientDataJSON"`
		} `json:"response"`
	}

	// Unmarshal the credential JSON
	if err := json.Unmarshal([]byte(credentialJSON), &cred); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid credential format: %v", err))
	}

	// Validate required fields
	if cred.ID == "" || cred.RawID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing credential ID")
	}
	if cred.Type != "public-key" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid credential type")
	}
	if cred.Response.AttestationObject == "" || cred.Response.ClientDataJSON == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing attestation data")
	}

	// Decode attestation object and client data
	attestationObj, err := base64.RawURLEncoding.DecodeString(cred.Response.AttestationObject)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid attestation object encoding")
	}

	clientData, err := base64.RawURLEncoding.DecodeString(cred.Response.ClientDataJSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid client data encoding")
	}

	// Log detailed credential information
	fmt.Printf("Credential Details:\n"+
		"ID: %s\n"+
		"Raw ID: %s\n"+
		"Type: %s\n"+
		"Authenticator Attachment: %s\n"+
		"Transports: %v\n"+
		"Attestation Object Size: %d bytes\n"+
		"Client Data Size: %d bytes\n",
		cred.ID,
		cred.RawID,
		cred.Type,
		cred.AuthenticatorAttachment,
		cred.Transports,
		len(attestationObj),
		len(clientData))

	// TODO: Verify the attestation and store the credential
	// This is where you would:
	// 1. Verify the attestation signature
	// 2. Check the origin in client data
	// 3. Verify the challenge
	// 4. Store the credential for future authentications

	return response.TemplEcho(c, register.LoadingVaultView())
}
