package handlers

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/internal/orm"
)

// ╭───────────────────────────────────────────────────────────╮
// │                    Login Handlers                         │
// ╰───────────────────────────────────────────────────────────╯

func (a *authAPI) LoginSubjectCheck(e echo.Context) error {
	return e.JSON(200, "HandleCredentialAssertion")
}

func (a *authAPI) LoginSubjectStart(e echo.Context) error {
	opts := &protocol.PublicKeyCredentialRequestOptions{
		UserVerification: "preferred",
		Challenge:        []byte("challenge"),
	}
	return e.JSON(200, opts)
}

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

func (a *authAPI) RegisterSubjectCheck(e echo.Context) error {
	subject := e.FormValue("subject")
	return e.JSON(200, subject)
}

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

// ╭───────────────────────────────────────────────────────────╮
// │                 Group Structures                          │
// ╰───────────────────────────────────────────────────────────╯

type authAPI struct{}

var Auth = new(authAPI)
