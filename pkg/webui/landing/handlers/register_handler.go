package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common/cookie"
	"github.com/onsonr/sonr/pkg/common/response"
	"github.com/onsonr/sonr/pkg/webui/landing/components/register"
)

func HandleRegister(c echo.Context) error {
	if !hasProfile(c) {
		return response.TemplEcho(c, register.ProfileFormView())
	}
	return nil
}

func HandleRegisterProfile(c echo.Context) error {
	name := c.FormValue("name")
	handle := c.FormValue("handle")
	cookie.Write(c, cookie.UserName, name)
	cookie.Write(c, cookie.UserHandle, handle)
	return response.TemplEcho(c, register.SetPasscodeView())
}

func HandleRegisterPasscode(c echo.Context) error {
	return response.TemplEcho(c, register.ConfirmPasscodeView())
}

func HandleConfirmPasscode(c echo.Context) error {
	return response.TemplEcho(c, register.LinkCredentialView())
}

func HandleCredentialLink(c echo.Context) error {
	return nil
}

//
// // Example usage:
// func RegisterHandler(c echo.Context) error {
// 	options := PublicKeyCredentialCreationOptions{
// 		Challenge:       "your-challenge-base64url",
// 		RpName:          "Your App",
// 		RpID:            "yourdomain.com",
// 		UserID:          "user-id-base64url",
// 		UserName:        "username",
// 		UserDisplayName: "User Display Name",
// 		Timeout:         60000,
// 		AttestationType: "none",
// 	}
//
// 	return CreateCredentials(c, options)
// }
//

// ╭────────────────────────────────────────────────────────╮
// │                  	Utility Functions 	                │
// ╰────────────────────────────────────────────────────────╯

func hasProfile(c echo.Context) bool {
	return false
}

func hasPasscode(c echo.Context) bool {
	return false
}

func hasCredentials(c echo.Context) bool {
	return false
}
