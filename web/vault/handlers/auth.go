package handlers

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"
)

// ╭───────────────────────────────────────────────────────────╮
// │                    Login Handlers                         │
// ╰───────────────────────────────────────────────────────────╯

// LoginSubjectCheck handles the login subject check.
func LoginSubjectCheck(e echo.Context) error {
	return e.JSON(200, "HandleCredentialAssertion")
}

// LoginSubjectStart handles the login subject start.
func LoginSubjectStart(e echo.Context) error {
	opts := &protocol.PublicKeyCredentialRequestOptions{
		UserVerification: "preferred",
		Challenge:        []byte("challenge"),
	}
	return e.JSON(200, opts)
}

// LoginSubjectFinish handles the login subject finish.
func LoginSubjectFinish(e echo.Context) error {
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
func RegisterSubjectCheck(e echo.Context) error {
	subject := e.FormValue("subject")
	return e.JSON(200, subject)
}

// RegisterSubjectStart handles the register subject start.
func RegisterSubjectStart(e echo.Context) error {
	// Get subject and address
	// subject := e.FormValue("subject")

	// Get challenge

	return nil
}

// RegisterSubjectFinish handles the register subject finish.
func RegisterSubjectFinish(e echo.Context) error {
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
