package credentials

import (
	"github.com/labstack/echo/v4"
)

// CreateCredentials initiates the credential creation process
func CreateCredentials(c echo.Context, options PublicKeyCredentialCreationOptions) error {
	return CreateCredential(options).Render(c.Request().Context(), c.Response().Writer)
}

// GetCredentials initiates the credential retrieval process
func GetCredentials(c echo.Context, options PublicKeyCredentialRequestOptions) error {
	return GetCredential(options).Render(c.Request().Context(), c.Response().Writer)
}

// Example usage:
func RegisterHandler(c echo.Context) error {
	options := PublicKeyCredentialCreationOptions{
		Challenge:       "your-challenge-base64url",
		RpName:          "Your App",
		RpID:            "yourdomain.com",
		UserID:          "user-id-base64url",
		UserName:        "username",
		UserDisplayName: "User Display Name",
		Timeout:         60000,
		AttestationType: "none",
	}

	return CreateCredentials(c, options)
}

func LoginHandler(c echo.Context) error {
	options := PublicKeyCredentialRequestOptions{
		Challenge: "your-challenge-base64url",
		RpID:      "yourdomain.com",
		Timeout:   60000,
		AllowCredentials: []CredentialDescriptor{
			{
				Type: "public-key",
				ID:   "credential-id-base64url",
			},
		},
	}

	return GetCredentials(c, options)
}
