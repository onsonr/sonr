package orm

import (
	"fmt"

	"github.com/go-webauthn/webauthn/protocol/webauthncose"
)

// ExtractWebAuthnPublicKey parses the raw public key bytes and returns a JWK representation
func ExtractWebAuthnPublicKey(keyBytes []byte) (*JWK, error) {
	key, err := webauthncose.ParsePublicKey(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	switch k := key.(type) {
	case *webauthncose.EC2PublicKeyData:
		return FormatEC2PublicKey(k)
	case *webauthncose.RSAPublicKeyData:
		return FormatRSAPublicKey(k)
	case *webauthncose.OKPPublicKeyData:
		return FormatOKPPublicKey(k)
	default:
		return nil, fmt.Errorf("unsupported key type")
	}
}
