package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/app/gateway/internal/database"
	"github.com/onsonr/sonr/app/gateway/internal/pages/register"
	"github.com/onsonr/sonr/crypto/mpc"
	"github.com/onsonr/sonr/pkg/common/response"
	"github.com/onsonr/sonr/pkg/common/styles/forms"
)

func HandleRegisterView(c echo.Context) error {
	dat := forms.CreateProfileData{
		FirstNumber: 1,
		LastNumber:  2,
	}
	return response.TemplEcho(c, register.ProfileFormView(dat))
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
	cred := database.Credential{}
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

	return response.TemplEcho(c, register.LoadingVaultView())
}
