package claims

import (
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/golang-jwt/jwt/v5"
)

type ControllerClaims struct {
	jwt.RegisteredClaims
}

// NewControllerClaims returns the CredentialClaims for the JWS to sign
func NewControllerClaims() CredentialClaims {
	claims := CredentialClaims{
		Credentials: make([]protocol.CredentialDescriptor, 0),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}
	return claims
}
