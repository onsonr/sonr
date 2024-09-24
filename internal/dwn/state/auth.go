package state

import (
	"encoding/json"
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"
)

func checkSubjectIsValid(e echo.Context) error {
	credentialID := e.FormValue("credentialID")
	return e.JSON(200, credentialID)
}

func handleCredentialAssertion(e echo.Context) error {
	return e.JSON(200, "HandleCredentialAssertion")
}

func handleCredentialCreation(e echo.Context) error {
	// Get the serialized credential data from the form
	credentialDataJSON := e.FormValue("credentialData")

	// Deserialize the JSON into a temporary struct
	var ccr protocol.CredentialCreationResponse
	err := json.Unmarshal([]byte(credentialDataJSON), &ccr)
	if err != nil {
		return e.JSON(500, err.Error())
	}
	//
	// // Parse the CredentialCreationResponse
	// parsedData, err := ccr.Parse()
	// if err != nil {
	// 	return e.JSON(500, err.Error())
	// }
	//
	// // Create the Credential
	// // credential := orm.NewCredential(parsedData, e.Request().Host, "")
	//
	// // Set additional fields
	// credential.Controller = "" // Set this to the appropriate controller value
	return e.JSON(200, fmt.Sprintf("REGISTER: %s", string(ccr.ID)))
}
