package handlers

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/internal/orm"
)

type authAPI struct{}

var Auth = new(authAPI)

// ╭───────────────────────────────────────────────────────────╮
// │                    Login Handlers                         │
// ╰───────────────────────────────────────────────────────────╯

// LoginSubjectCheck handles the login subject check.
func (a *authAPI) LoginSubjectCheck(e echo.Context) error {
	return e.JSON(200, "HandleCredentialAssertion")
}

// LoginSubjectStart handles the login subject start.
func (a *authAPI) LoginSubjectStart(e echo.Context) error {
	opts := &protocol.PublicKeyCredentialRequestOptions{
		UserVerification: "preferred",
		Challenge:        []byte("challenge"),
	}
	return e.JSON(200, opts)
}

// LoginSubjectFinish handles the login subject finish.
func (a *authAPI) LoginSubjectFinish(e echo.Context) error {
	var crr protocol.CredentialAssertionResponse
	if err := e.Bind(&crr); err != nil {
		return err
	}
	return e.JSON(200, crr)
}

// ╭───────────────────────────────────────────────────────────╮
// │                   Register Handlers                       │
// ╰───────────────────────────────────────────────────────────╯

// RegisterSubjectCheck handles the register subject check.
func (a *authAPI) RegisterSubjectCheck(e echo.Context) error {
	subject := e.FormValue("subject")
	return e.JSON(200, subject)
}

// RegisterSubjectStart handles the register subject start.
func (a *authAPI) RegisterSubjectStart(e echo.Context) error {
	// Get subject and address
	subject := e.FormValue("subject")
	address := e.FormValue("address")

	// Get challenge
	chal, err := protocol.CreateChallenge()
	if err != nil {
		return err
	}
	return e.JSON(201, orm.NewCredentialCreationOptions(subject, address, chal))
}

// RegisterSubjectFinish handles the register subject finish.
func (a *authAPI) RegisterSubjectFinish(e echo.Context) error {
	// Deserialize the JSON into a temporary struct
	var ccr protocol.CredentialCreationResponse
	if err := e.Bind(&ccr); err != nil {
		return err
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
	return e.JSON(201, ccr)
}
