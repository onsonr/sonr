package orm

import (
	"encoding/base64"
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"
)

func FormatEC2PublicKey(key *webauthncose.EC2PublicKeyData) (*JWK, error) {
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

func FormatRSAPublicKey(key *webauthncose.RSAPublicKeyData) (*JWK, error) {
	jwkMap := map[string]interface{}{
		"kty": "RSA",
		"n":   base64.RawURLEncoding.EncodeToString(key.Modulus),
		"e":   base64.RawURLEncoding.EncodeToString(key.Exponent),
	}

	return MapToJWK(jwkMap)
}

func FormatOKPPublicKey(key *webauthncose.OKPPublicKeyData) (*JWK, error) {
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

func MapToJWK(m map[string]interface{}) (*JWK, error) {
	jwk := &JWK{}
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

// ConvertTransports converts the transports from strings to protocol.AuthenticatorTransport
func ConvertTransports(transports []string) []protocol.AuthenticatorTransport {
	tss := make([]protocol.AuthenticatorTransport, len(transports))
	for i, t := range transports {
		tss[i] = protocol.AuthenticatorTransport(t)
	}
	return tss
}

// NormalizeTransports returns the transports as strings
func NormalizeTransports(transports []protocol.AuthenticatorTransport) []string {
	tss := make([]string, len(transports))
	for i, t := range transports {
		tss[i] = string(t)
	}
	return tss
}
