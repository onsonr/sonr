package claims

import (
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/golang-jwt/jwt/v5"
)

type CredentialClaims struct {
	Credentials []protocol.CredentialDescriptor `json:"credentials"`
	jwt.RegisteredClaims
}

// NewCredentialClaims returns the CredentialClaims for the JWS to sign
func NewCredentialClaims() CredentialClaims {
	// Create claims with multiple fields populated
	claims := CredentialClaims{
		Credentials: make([]protocol.CredentialDescriptor, 0),
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
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
