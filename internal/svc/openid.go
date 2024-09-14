package svc

import (
	"fmt"

	"github.com/labstack/echo/v4"
	oidc "github.com/onsonr/sonr/x/did/types/oidc"
)

func GrantAuthorization(e echo.Context) error {
	// Implement authorization endpoint using passkey authentication
	// Store session data in cache
	return nil
}

func GetJWKS(e echo.Context) error {
	// Implement token endpoint
	// Use cached session data for validation
	return nil
}

func GetToken(e echo.Context) error {
	// Implement token endpoint
	// Use cached session data for validation
	return nil
}

func GetDiscovery(e echo.Context) error {
	baseURL := "https://" + e.Request().Host // Ensure this is the correct base URL for your service
	discoveryDoc := &oidc.DiscoveryDocument{
		Issuer:                 baseURL,
		AuthorizationEndpoint:  fmt.Sprintf("%s/auth", baseURL),
		TokenEndpoint:          fmt.Sprintf("%s/token", baseURL),
		UserinfoEndpoint:       fmt.Sprintf("%s/userinfo", baseURL),
		JwksUri:                fmt.Sprintf("%s/jwks", baseURL),
		RegistrationEndpoint:   fmt.Sprintf("%s/register", baseURL),
		ScopesSupported:        []string{"openid", "profile", "email", "web3", "sonr"},
		ResponseTypesSupported: []string{"code"},
		ResponseModesSupported: []string{"query", "form_post"},
		GrantTypesSupported:    []string{"authorization_code", "refresh_token"},
		AcrValuesSupported:     []string{"passkey"},
		SubjectTypesSupported:  []string{"public"},
		ClaimsSupported:        []string{"sub", "iss", "name", "email"},
	}
	return e.JSON(200, discoveryDoc)
}
