package handlers

import (
	"github.com/labstack/echo/v4"
)

var OIDC = oidcHandler{}

type oidcHandler struct{}

func (p oidcHandler) HandleAuthorize(e echo.Context) error {
	// Implement authorization endpoint using passkey authentication
	// Store session data in cache
	return nil
}

func (p oidcHandler) HandleToken(e echo.Context) error {
	// Implement token endpoint
	// Use cached session data for validation
	return nil
}

func (p oidcHandler) HandleUserInfo(e echo.Context) error {
	// Implement userinfo endpoint
	// Use cached session data for validation
	return nil
}

func (p oidcHandler) HandleDiscovery(e echo.Context) error {
	baseURL := "https://" + e.Request().Host // Ensure this is the correct base URL for your service

	discoveryDoc := DiscoveryDocument{
		Issuer:                            baseURL,
		AuthorizationEndpoint:             baseURL + "/authorize",
		TokenEndpoint:                     baseURL + "/token",
		UserinfoEndpoint:                  baseURL + "/userinfo",
		JwksURI:                           baseURL + "/jwks", // You'll need to implement this endpoint
		RegistrationEndpoint:              baseURL + "/register",
		ScopesSupported:                   []string{"openid", "profile", "email"},
		ResponseTypesSupported:            []string{"code"},
		SubjectTypesSupported:             []string{"public"},
		IDTokenSigningAlgValuesSupported:  []string{"RS256"},
		ClaimsSupported:                   []string{"sub", "iss", "name", "email"},
		GrantTypesSupported:               []string{"authorization_code", "refresh_token"},
		TokenEndpointAuthMethodsSupported: []string{"client_secret_basic", "client_secret_post"},
	}
	return e.JSON(200, discoveryDoc)
}

type DiscoveryDocument struct {
	Issuer                            string   `json:"issuer"`
	AuthorizationEndpoint             string   `json:"authorization_endpoint"`
	TokenEndpoint                     string   `json:"token_endpoint"`
	UserinfoEndpoint                  string   `json:"userinfo_endpoint"`
	JwksURI                           string   `json:"jwks_uri"`
	RegistrationEndpoint              string   `json:"registration_endpoint"`
	ScopesSupported                   []string `json:"scopes_supported"`
	ResponseTypesSupported            []string `json:"response_types_supported"`
	SubjectTypesSupported             []string `json:"subject_types_supported"`
	IDTokenSigningAlgValuesSupported  []string `json:"id_token_signing_alg_values_supported"`
	ClaimsSupported                   []string `json:"claims_supported"`
	GrantTypesSupported               []string `json:"grant_types_supported"`
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`
}
