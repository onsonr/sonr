package common

import (
	"errors"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
)

// ConvertWebauthnCredential creates a common.WebauthnCredential from a webauthn.Credential from the go-webauthn package
func ConvertWebauthnCredential(wa *webauthn.Credential) *WebauthnCredential {
	transportsStr := []string{}
	for _, t := range wa.Transport {
		transportsStr = append(transportsStr, string(t))
	}
	return &WebauthnCredential{
		Id:              wa.ID,
		PublicKey:       wa.PublicKey,
		AttestationType: wa.AttestationType,
		Transport:       transportsStr,
		Authenticator: &WebauthnAuthenticator{
			Aaguid:       wa.Authenticator.AAGUID,
			SignCount:    wa.Authenticator.SignCount,
			CloneWarning: wa.Authenticator.CloneWarning,
		},
	}
}

// NewWebAuthnCredential creates a new WebauthnCredential from a ParsedCredentialCreationData and contains all needed information about a WebAuthn credential for storage.
// This is then used to create a VerificationMethod for the DID Document.
func NewWebAuthnCredential(c *protocol.ParsedCredentialCreationData) *WebauthnCredential {
	transportsStr := []string{}
	for _, t := range c.Transports {
		transportsStr = append(transportsStr, string(t))
	}
	return &WebauthnCredential{
		Id:              c.Response.AttestationObject.AuthData.AttData.CredentialID,
		PublicKey:       c.Response.AttestationObject.AuthData.AttData.CredentialPublicKey,
		AttestationType: c.Response.AttestationObject.Format,
		Transport:       transportsStr,
		Authenticator: &WebauthnAuthenticator{
			Aaguid:    c.Response.AttestationObject.AuthData.AttData.AAGUID,
			SignCount: c.Response.AttestationObject.AuthData.Counter,
		},
	}
}

// Validate verifies that this WebauthnCredential is identical to the go-webauthn package credential
func (c *WebauthnCredential) Validate(pc *webauthn.Credential) error {
	if len(c.PublicKey) != len(pc.PublicKey) {
		return errors.New("Credential Public Keys do not match")
	}
	return nil
}
