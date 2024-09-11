package svc

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/db/orm"
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
	discoveryDoc := &orm.DiscoveryDocument{
		Issuer:                 "https://sonr.id",
		AuthorizationEndpoint:  "https://api.sonr.id/auth",
		TokenEndpoint:          "https://api.sonr.id/token",
		UserinfoEndpoint:       "https://api.sonr.id/userinfo",
		JwksUri:                baseURL + "/jwks", // You'll need to implement this endpoint
		RegistrationEndpoint:   baseURL + "/register",
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
