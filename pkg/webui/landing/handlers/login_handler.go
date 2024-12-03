package handlers

import "github.com/labstack/echo/v4"

func HandleLogin(c echo.Context) error {
	return nil
}

//
// func LoginHandler(c echo.Context) error {
// 	options := PublicKeyCredentialRequestOptions{
// 		Challenge: "your-challenge-base64url",
// 		RpID:      "yourdomain.com",
// 		Timeout:   60000,
// 		AllowCredentials: []CredentialDescriptor{
// 			{
// 				Type: "public-key",
// 				ID:   "credential-id-base64url",
// 			},
// 		},
// 	}
//
// 	return GetCredentials(c, options)
// }
