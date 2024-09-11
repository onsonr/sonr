package builder

import (
	"encoding/base64"
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"

	"github.com/go-webauthn/webauthn/protocol/webauthncose"
	didv1 "github.com/onsonr/sonr/api/did/v1"
	"github.com/onsonr/sonr/x/did/types"
)

func APIFormatDIDNamespace(namespace types.DIDNamespace) didv1.DIDNamespace {
	return didv1.DIDNamespace(namespace)
}

func APIFormatDIDNamespaces(namespaces []types.DIDNamespace) []didv1.DIDNamespace {
	var s []didv1.DIDNamespace
	for _, namespace := range namespaces {
		s = append(s, APIFormatDIDNamespace(namespace))
	}
	return s
}

func APIFormatKeyRole(role types.KeyRole) didv1.KeyRole {
	return didv1.KeyRole(role)
}

func APIFormatKeyAlgorithm(algorithm types.KeyAlgorithm) didv1.KeyAlgorithm {
	return didv1.KeyAlgorithm(algorithm)
}

func APIFormatKeyEncoding(encoding types.KeyEncoding) didv1.KeyEncoding {
	return didv1.KeyEncoding(encoding)
}

func APIFormatKeyCurve(curve types.KeyCurve) didv1.KeyCurve {
	return didv1.KeyCurve(curve)
}

func APIFormatKeyType(keyType types.KeyType) didv1.KeyType {
	return didv1.KeyType(keyType)
}

func APIFormatPermissions(permissions *types.Permissions) *didv1.Permissions {
	if permissions == nil {
		return nil
	}
	p := didv1.Permissions{
		Grants: APIFormatDIDNamespaces(permissions.Grants),
		Scopes: APIFormatPermissionScopes(permissions.Scopes),
	}
	return &p
}

func APIFormatPermissionScope(scope types.PermissionScope) didv1.PermissionScope {
	return didv1.PermissionScope(scope)
}

func APIFormatPermissionScopes(scopes []types.PermissionScope) []didv1.PermissionScope {
	var s []didv1.PermissionScope
	for _, scope := range scopes {
		s = append(s, APIFormatPermissionScope(scope))
	}
	return s
}

func APIFormatServiceRecord(service *types.Service) *didv1.ServiceRecord {
	return &didv1.ServiceRecord{
		Id:               service.Id,
		ServiceType:      service.ServiceType,
		Authority:        service.Authority,
		Origin:           service.Origin,
		Description:      service.Description,
		ServiceEndpoints: service.ServiceEndpoints,
		Permissions:      APIFormatPermissions(service.Permissions),
	}
}

func APIFormatPubKeyJWK(jwk *types.PubKey_JWK) *didv1.PubKey_JWK {
	return &didv1.PubKey_JWK{
		Kty: jwk.Kty,
		Crv: jwk.Crv,
		X:   jwk.X,
		Y:   jwk.Y,
		N:   jwk.N,
		E:   jwk.E,
	}
}

func APIFormatPubKey(key *types.PubKey) *didv1.PubKey {
	return &didv1.PubKey{
		Role:      APIFormatKeyRole(key.GetRole()),
		Algorithm: APIFormatKeyAlgorithm(key.GetAlgorithm()),
		Encoding:  APIFormatKeyEncoding(key.GetEncoding()),
		Curve:     APIFormatKeyCurve(key.GetCurve()),
		KeyType:   APIFormatKeyType(key.GetKeyType()),
		Raw:       key.GetRaw(),
	}
}

func FormatEC2PublicKey(key *webauthncose.EC2PublicKeyData) (*types.PubKey_JWK, error) {
	curve, err := GetCOSECurveName(key.Curve)
	if err != nil {
		return nil, err
	}

	jwkMap := map[string]interface{}{
		"kty": "EC",
		"crv": curve,
		"x":   base64.RawURLEncoding.EncodeToString(key.XCoord),
		"y":   base64.RawURLEncoding.EncodeToString(key.YCoord),
	}

	return MapToJWK(jwkMap)
}

func FormatRSAPublicKey(key *webauthncose.RSAPublicKeyData) (*types.PubKey_JWK, error) {
	jwkMap := map[string]interface{}{
		"kty": "RSA",
		"n":   base64.RawURLEncoding.EncodeToString(key.Modulus),
		"e":   base64.RawURLEncoding.EncodeToString(key.Exponent),
	}

	return MapToJWK(jwkMap)
}

func FormatOKPPublicKey(key *webauthncose.OKPPublicKeyData) (*types.PubKey_JWK, error) {
	curve, err := GetOKPCurveName(key.Curve)
	if err != nil {
		return nil, err
	}

	jwkMap := map[string]interface{}{
		"kty": "OKP",
		"crv": curve,
		"x":   base64.RawURLEncoding.EncodeToString(key.XCoord),
	}

	return MapToJWK(jwkMap)
}

func MapToJWK(m map[string]interface{}) (*types.PubKey_JWK, error) {
	jwk := &types.PubKey_JWK{}
	for k, v := range m {
		switch k {
		case "kty":
			jwk.Kty = v.(string)
		case "crv":
			jwk.Crv = v.(string)
		case "x":
			jwk.X = v.(string)
		case "y":
			jwk.Y = v.(string)
		case "n":
			jwk.N = v.(string)
		case "e":
			jwk.E = v.(string)
		}
	}
	return jwk, nil
}

func GetCOSECurveName(curveID int64) (string, error) {
	switch curveID {
	case int64(webauthncose.P256):
		return "P-256", nil
	case int64(webauthncose.P384):
		return "P-384", nil
	case int64(webauthncose.P521):
		return "P-521", nil
	default:
		return "", fmt.Errorf("unknown curve ID: %d", curveID)
	}
}

func GetOKPCurveName(curveID int64) (string, error) {
	switch curveID {
	case int64(webauthncose.Ed25519):
		return "Ed25519", nil
	default:
		return "", fmt.Errorf("unknown OKP curve ID: %d", curveID)
	}
}

// NormalizeTransports returns the transports as strings
func NormalizeTransports(transports []protocol.AuthenticatorTransport) []string {
	tss := make([]string, len(transports))
	for i, t := range transports {
		tss[i] = string(t)
	}
	return tss
}

// GetTransports returns the protocol.AuthenticatorTransport
func ModuleTransportsToProtocol(transport []string) []protocol.AuthenticatorTransport {
	tss := make([]protocol.AuthenticatorTransport, len(transport))
	for i, t := range transport {
		tss[i] = protocol.AuthenticatorTransport(t)
	}
	return tss
}

// ModuleFormatAPIServiceRecord formats a service record for the module
func ModuleFormatAPIServiceRecord(service *didv1.ServiceRecord) *types.Service {
	return &types.Service{
		Id:               service.Id,
		ServiceType:      service.ServiceType,
		Authority:        service.Authority,
		Origin:           service.Origin,
		Description:      service.Description,
		ServiceEndpoints: service.ServiceEndpoints,
		Permissions:      ModuleFormatAPIPermissions(service.Permissions),
	}
}

func ModuleFormatAPIPermissions(permissions *didv1.Permissions) *types.Permissions {
	if permissions == nil {
		return nil
	}
	p := types.Permissions{
		Grants: ModuleFormatAPIDIDNamespaces(permissions.Grants),
		Scopes: ModuleFormatAPIPermissionScopes(permissions.Scopes),
	}
	return &p
}

func ModuleFormatAPIPermissionScope(scope didv1.PermissionScope) types.PermissionScope {
	return types.PermissionScope(scope)
}

func ModuleFormatAPIPermissionScopes(scopes []didv1.PermissionScope) []types.PermissionScope {
	var s []types.PermissionScope
	for _, scope := range scopes {
		s = append(s, ModuleFormatAPIPermissionScope(scope))
	}
	return s
}

func ModuleFormatAPIDIDNamespace(namespace didv1.DIDNamespace) types.DIDNamespace {
	return types.DIDNamespace(namespace)
}

func ModuleFormatAPIDIDNamespaces(namespaces []didv1.DIDNamespace) []types.DIDNamespace {
	var s []types.DIDNamespace
	for _, namespace := range namespaces {
		s = append(s, ModuleFormatAPIDIDNamespace(namespace))
	}
	return s
}
