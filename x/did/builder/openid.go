package builder

import (
	"fmt"

	"github.com/onsonr/sonr/x/did/types/oidc"
)

func GetDiscovery(origin string) *oidc.DiscoveryDocument {
	baseURL := "https://" + origin // Ensure this is the correct base URL for your service
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
	return discoveryDoc
}
