// Code generated from Pkl module `orm`. DO NOT EDIT.
package orm

type DiscoveryDocument struct {
	Issuer string `pkl:"issuer" json:"issuer,omitempty" param:"issuer"`

	AuthorizationEndpoint string `pkl:"authorization_endpoint" json:"authorization_endpoint,omitempty" param:"authorization_endpoint"`

	TokenEndpoint string `pkl:"token_endpoint" json:"token_endpoint,omitempty" param:"token_endpoint"`

	UserinfoEndpoint string `pkl:"userinfo_endpoint" json:"userinfo_endpoint,omitempty" param:"userinfo_endpoint"`

	JwksUri string `pkl:"jwks_uri" json:"jwks_uri,omitempty" param:"jwks_uri"`

	RegistrationEndpoint string `pkl:"registration_endpoint" json:"registration_endpoint,omitempty" param:"registration_endpoint"`

	ScopesSupported []string `pkl:"scopes_supported" json:"scopes_supported,omitempty" param:"scopes_supported"`

	ResponseTypesSupported []string `pkl:"response_types_supported" json:"response_types_supported,omitempty" param:"response_types_supported"`

	ResponseModesSupported []string `pkl:"response_modes_supported" json:"response_modes_supported,omitempty" param:"response_modes_supported"`

	SubjectTypesSupported []string `pkl:"subject_types_supported" json:"subject_types_supported,omitempty" param:"subject_types_supported"`

	IdTokenSigningAlgValuesSupported []string `pkl:"id_token_signing_alg_values_supported" json:"id_token_signing_alg_values_supported,omitempty" param:"id_token_signing_alg_values_supported"`

	ClaimsSupported []string `pkl:"claims_supported" json:"claims_supported,omitempty" param:"claims_supported"`

	GrantTypesSupported []string `pkl:"grant_types_supported" json:"grant_types_supported,omitempty" param:"grant_types_supported"`

	AcrValuesSupported []string `pkl:"acr_values_supported" json:"acr_values_supported,omitempty" param:"acr_values_supported"`

	TokenEndpointAuthMethodsSupported []string `pkl:"token_endpoint_auth_methods_supported" json:"token_endpoint_auth_methods_supported,omitempty" param:"token_endpoint_auth_methods_supported"`
}
