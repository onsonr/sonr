package builder

import (
	"encoding/base64"
	"fmt"

	"github.com/go-webauthn/webauthn/protocol/webauthncose"
	didv1 "github.com/onsonr/sonr/api/did/v1"
	"github.com/onsonr/sonr/x/did/types"
)

func FormatEC2PublicKey(key *webauthncose.EC2PublicKeyData) (*types.JWK, error) {
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

func FormatRSAPublicKey(key *webauthncose.RSAPublicKeyData) (*types.JWK, error) {
	jwkMap := map[string]interface{}{
		"kty": "RSA",
		"n":   base64.RawURLEncoding.EncodeToString(key.Modulus),
		"e":   base64.RawURLEncoding.EncodeToString(key.Exponent),
	}

	return MapToJWK(jwkMap)
}

func FormatOKPPublicKey(key *webauthncose.OKPPublicKeyData) (*types.JWK, error) {
	curve, err := getOKPCurveName(key.Curve)
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

func MapToJWK(m map[string]interface{}) (*types.JWK, error) {
	jwk := &types.JWK{}
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

func getOKPCurveName(curveID int64) (string, error) {
	switch curveID {
	case int64(webauthncose.Ed25519):
		return "Ed25519", nil
	default:
		return "", fmt.Errorf("unknown OKP curve ID: %d", curveID)
	}
}

func ModulePubKeyToAPI(pk *types.PubKey) *didv1.PubKey {
	return &didv1.PubKey{
		Role:      ModuleKeyRoleToAPI(pk.GetRole()),
		Algorithm: ModuleKeyAlgorithmToAPI(pk.GetAlgorithm()),
		Encoding:  ModuleKeyEncodingToAPI(pk.GetEncoding()),
		Curve:     ModuleKeyCurveToAPI(pk.GetCurve()),
		KeyType:   ModuleKeyTypeToAPI(pk.GetKeyType()),
		Raw:       pk.GetRaw(),
	}
}

func ModuleKeyRoleToAPI(role types.KeyRole) didv1.KeyRole {
	switch role {
	case types.KeyRole_KEY_ROLE_INVOCATION:
		return didv1.KeyRole_KEY_ROLE_INVOCATION
	case types.KeyRole_KEY_ROLE_ASSERTION:
		return didv1.KeyRole_KEY_ROLE_ASSERTION
	case types.KeyRole_KEY_ROLE_DELEGATION:
		return didv1.KeyRole_KEY_ROLE_DELEGATION
	default:
		return didv1.KeyRole_KEY_ROLE_INVOCATION
	}
}

func ModuleKeyAlgorithmToAPI(algorithm types.KeyAlgorithm) didv1.KeyAlgorithm {
	switch algorithm {
	case types.KeyAlgorithm_KEY_ALGORITHM_ES256K:
		return didv1.KeyAlgorithm_KEY_ALGORITHM_ES256K
	case types.KeyAlgorithm_KEY_ALGORITHM_ES256:
		return didv1.KeyAlgorithm_KEY_ALGORITHM_ES256
	case types.KeyAlgorithm_KEY_ALGORITHM_ES384:
		return didv1.KeyAlgorithm_KEY_ALGORITHM_ES384
	case types.KeyAlgorithm_KEY_ALGORITHM_ES512:
		return didv1.KeyAlgorithm_KEY_ALGORITHM_ES512
	case types.KeyAlgorithm_KEY_ALGORITHM_EDDSA:
		return didv1.KeyAlgorithm_KEY_ALGORITHM_EDDSA
	default:
		return didv1.KeyAlgorithm_KEY_ALGORITHM_ES256K
	}
}

func ModuleKeyCurveToAPI(curve types.KeyCurve) didv1.KeyCurve {
	switch curve {
	case types.KeyCurve_KEY_CURVE_P256:
		return didv1.KeyCurve_KEY_CURVE_P256
	case types.KeyCurve_KEY_CURVE_SECP256K1:
		return didv1.KeyCurve_KEY_CURVE_SECP256K1
	case types.KeyCurve_KEY_CURVE_BLS12381:
		return didv1.KeyCurve_KEY_CURVE_BLS12381
	case types.KeyCurve_KEY_CURVE_KECCAK256:
		return didv1.KeyCurve_KEY_CURVE_KECCAK256
	default:
		return didv1.KeyCurve_KEY_CURVE_P256
	}
}

func ModuleKeyEncodingToAPI(encoding types.KeyEncoding) didv1.KeyEncoding {
	switch encoding {
	case types.KeyEncoding_KEY_ENCODING_RAW:
		return didv1.KeyEncoding_KEY_ENCODING_RAW
	case types.KeyEncoding_KEY_ENCODING_HEX:
		return didv1.KeyEncoding_KEY_ENCODING_HEX
	case types.KeyEncoding_KEY_ENCODING_MULTIBASE:
		return didv1.KeyEncoding_KEY_ENCODING_MULTIBASE
	default:
		return didv1.KeyEncoding_KEY_ENCODING_RAW
	}
}

func ModuleKeyTypeToAPI(keyType types.KeyType) didv1.KeyType {
	switch keyType {
	case types.KeyType_KEY_TYPE_BIP32:
		return didv1.KeyType_KEY_TYPE_BIP32
	case types.KeyType_KEY_TYPE_ZK:
		return didv1.KeyType_KEY_TYPE_ZK
	case types.KeyType_KEY_TYPE_WEBAUTHN:
		return didv1.KeyType_KEY_TYPE_WEBAUTHN
	default:
		return didv1.KeyType_KEY_TYPE_BIP32
	}
}
