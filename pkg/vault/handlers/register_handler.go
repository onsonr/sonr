package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/onsonr/hway/internal/orm"
	pages "github.com/onsonr/hway/pkg/vault/components"
	"github.com/onsonr/hway/pkg/vault/middleware"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"
)

var Register = registerHandler{}

type registerHandler struct{}

func (h registerHandler) Start(e echo.Context) error {
	return e.JSON(0, nil)
}

func (h registerHandler) Finish(e echo.Context) error {
	// Get the serialized credential data from the form
	credentialDataJSON := e.FormValue("credentialData")

	// Deserialize the JSON into a temporary struct
	var ccr protocol.CredentialCreationResponse
	err := json.Unmarshal([]byte(credentialDataJSON), &ccr)
	if err != nil {
		return e.JSON(500, err.Error())
	}

	// Parse the CredentialCreationResponse
	parsedData, err := ccr.Parse()
	if err != nil {
		return e.JSON(500, err.Error())
	}

	// Create the Credential
	credential := orm.MakeNewCredential(parsedData)

	// Set additional fields
	credential.DisplayName = ccr.ID      // You might want to set this to a more meaningful value
	credential.Origin = e.Request().Host // Set the origin to the current host
	credential.Controller = ""           // Set this to the appropriate controller value
	return e.JSON(200, fmt.Sprintf("REGISTER: %s", string(credential.ID)))
}

// FormPage returns the page for registering a new user
func (h registerHandler) FormPage(e echo.Context) error {
	return middleware.Render(e, pages.Register(middleware.SessionID(e), string(middleware.Cache.GetChallenge(e))))
}
