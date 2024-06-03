package wallet

import (
	"context"
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/golang-jwt/jwt/v5"
)

// CredentialClaims is the claims for a credential.
type CredentialClaims struct {
	jwt.RegisteredClaims
	Credentials []protocol.CredentialDescriptor `json:"credentials"`
}

// NewCredentialClaims returns the CredentialClaims for the JWS to sign
func NewCredentialClaims(ctx context.Context) CredentialClaims {
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
