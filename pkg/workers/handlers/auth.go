package handlers

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/internal/ctx"
	"github.com/onsonr/sonr/internal/orm"
)

// ╭───────────────────────────────────────────────────────────╮
// │                    Login Handlers                         │
// ╰───────────────────────────────────────────────────────────╯

func LoginSubjectCheck(e echo.Context) error {
	return e.JSON(200, "HandleCredentialAssertion")
}

func LoginSubjectStart(e echo.Context) error {
	opts := &protocol.PublicKeyCredentialRequestOptions{
		UserVerification: "preferred",
		Challenge:        []byte("challenge"),
	}
	return e.JSON(200, opts)
}

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

func RegisterSubjectCheck(e echo.Context) error {
	subject := e.FormValue("subject")
	return e.JSON(200, subject)
}

func RegisterSubjectStart(e echo.Context) error {
	// Get subject and address
	subject := e.FormValue("subject")
	address := e.FormValue("address")

	// Set address in session
	s := ctx.GetSession(e)
	s = ctx.SetAddress(e, address)

	// Get challenge
	chal, err := s.GetChallenge(subject)
	if err != nil {
		return err
	}
	return e.JSON(201, orm.NewCredentialCreationOptions(subject, address, chal))
}

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
