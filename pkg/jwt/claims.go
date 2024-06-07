package jwt

import (
	"context"
	"encoding/json"
	"time"

	"github.com/di-dao/sonr/internal/local"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/golang-jwt/jwt/v5"
)

// CredentialClaims is the claims for a credential.
type CredentialClaims struct {
	jwt.RegisteredClaims
}

// NewCredentialClaims returns the CredentialClaims for the JWS to sign
func NewCredentialClaims(ctx context.Context, address string, origin string, credentials ...protocol.CredentialDescriptor) CredentialClaims {
	snrCtx := local.UnwrapContext(ctx)

	// Create claims with multiple fields populated
	return CredentialClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    snrCtx.ValidatorAddress,
			Subject:   origin,
			ID:        address,
			Audience:  ConvertCredentialDescriptorsToStringList(credentials),
		},
	}
}

// utility function to convert a list of CredentialDescriptors into a string list using JSON encoding
func ConvertCredentialDescriptorsToStringList(credentials []protocol.CredentialDescriptor) []string {
	credentialDescriptors := make([]string, len(credentials))
	for i, credential := range credentials {
		jsonByz, err := json.Marshal(credential)
		if err != nil {
			panic(err)
		}
		credentialDescriptors[i] = string(jsonByz)
	}
	return credentialDescriptors
}

// utility function to convert a string list into a list of CredentialDescriptors using JSON decoding
func ConvertStringListToCredentialDescriptors(credentialDescriptors []string) []protocol.CredentialDescriptor {
	credentials := make([]protocol.CredentialDescriptor, len(credentialDescriptors))
	for i, credentialDescriptor := range credentialDescriptors {
		jsonByz := []byte(credentialDescriptor)
		err := json.Unmarshal(jsonByz, &credentials[i])
		if err != nil {
			panic(err)
		}
	}
	return credentials
}
