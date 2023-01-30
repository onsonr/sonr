package crypto

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/sonrhq/core/pkg/crypto/internal/types"
)

// NewWebAuthnCredential creates a new WebauthnCredential from a ParsedCredentialCreationData and contains all needed information about a WebAuthn credential for storage.
// This is then used to create a VerificationMethod for the DID Document.
func NewWebAuthnCredential(c *protocol.ParsedCredentialCreationData) *types.WebauthnCredential {
	transportsStr := []string{}
	for _, t := range c.Transports {
		transportsStr = append(transportsStr, string(t))
	}
	return &types.WebauthnCredential{
		Id:              c.Response.AttestationObject.AuthData.AttData.CredentialID,
		PublicKey:       c.Response.AttestationObject.AuthData.AttData.CredentialPublicKey,
		AttestationType: c.Response.AttestationObject.Format,
		Transport:       transportsStr,
		Authenticator: &types.WebauthnAuthenticator{
			Aaguid:    c.Response.AttestationObject.AuthData.AttData.AAGUID,
			SignCount: c.Response.AttestationObject.AuthData.Counter,
		},
	}
}
