package crypto

import (
	fmt "fmt"
	"strings"
)

// FormatString returns the string representation of the key type
func (kt KeyType) FormatString() string {
	str := kt.String()
	ptrs := strings.Split(str, "_")
	result := ""
	for i, ptr := range ptrs {
		if i > 0 {
			result += strings.Title(strings.ToLower(ptr))
		}
	}
	return result
}

// FormatString returns the string representation of the key type
func (kt ProofType) FormatString() string {
	str := kt.String()
	ptrs := strings.Split(str, "_")
	result := ""
	for i, ptr := range ptrs {
		if i > 0 {
			result += strings.Title(strings.ToLower(ptr))
		}
	}
	return result
}

// MulticodecType returns the multicodec code for the key type
// -- We represent those as raw public key bytes prefixed with public key
// -- multiformat code.
// | secp256k1  "0xe7"
// | Ed25519    "0xed"
// | P256       "0x1200"
// | P384       "0x1201"
// | P512       "0x1202"
// | RSA        "0x1205"
func (kt KeyType) MulticodecType() uint64 {
	switch kt {
	case KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019:
		return 0xe7
	case KeyType_KeyType_ED25519_VERIFICATION_KEY_2018:
		return 0xed
	case KeyType_KeyType_JSON_WEB_KEY_2020:
		return 0x1200
	case KeyType_KeyType_RSA_VERIFICATION_KEY_2018:
		return 0x1205
	default:
		return 0
	}
}

// PrettyString returns the string representation of the key type
func (kt KeyType) PrettyString() string {
	switch kt {
	case KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019:
		return "secp256k1"
	case KeyType_KeyType_ED25519_VERIFICATION_KEY_2018:
		return "Ed25519"
	case KeyType_KeyType_JSON_WEB_KEY_2020:
		return "JWK"
	case KeyType_KeyType_RSA_VERIFICATION_KEY_2018:
		return "RSA"
	default:
		return "unknown"
	}
}

// KeyTypeFromMulticodec returns the key type
func KeyTypeFromMulticodec(code uint64) (KeyType, error) {
	switch code {
	case 0xe7:
		return KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019, nil
	case 0xed:
		return KeyType_KeyType_ED25519_VERIFICATION_KEY_2018, nil
	case 0x1200:
		return KeyType_KeyType_JSON_WEB_KEY_2020, nil
	case 0x1205:
		return KeyType_KeyType_RSA_VERIFICATION_KEY_2018, nil
	default:
		return KeyType_KeyType_UNSPECIFIED, fmt.Errorf("unknown key type code: %d", code)
	}
}

// KeyTypeFromPrettyString returns the key type
func KeyTypeFromPrettyString(s string) (KeyType, error) {
	switch s {
	case "secp256k1":
		return KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019, nil
	case "Ed25519":
		return KeyType_KeyType_ED25519_VERIFICATION_KEY_2018, nil
	case "JWK":
		return KeyType_KeyType_JSON_WEB_KEY_2020, nil
	case "RSA":
		return KeyType_KeyType_RSA_VERIFICATION_KEY_2018, nil
	default:
		return KeyType_KeyType_UNSPECIFIED, fmt.Errorf("unknown key type: %s", s)
	}
}

// IsBlockchainKey returns true if the key is a edsa or secp256k1 key
func (kt KeyType) IsBlockchainKey() bool {
	return kt == KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019 || kt == KeyType_KeyType_ED25519_VERIFICATION_KEY_2018
}

// IsWebAuthnKey returns true if the key is a webauthn key
func (kt KeyType) IsWebAuthnKey() bool {
	return kt == KeyType_KeyType_WEB_AUTHN_AUTHENTICATION_2018
}
