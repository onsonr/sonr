package orm

import (
	"encoding/base64"

	"github.com/go-webauthn/webauthn/protocol"
)

// NewCredential will return a credential pointer on successful validation of a registration response.
func NewCredential(c *protocol.ParsedCredentialCreationData, origin, handle string) *Credential {
	return &Credential{
		Subject:         handle,
		Origin:          origin,
		AttestationType: c.Response.AttestationObject.Format,
		CredentialId:    BytesToBase64(c.Response.AttestationObject.AuthData.AttData.CredentialID),
		PublicKey:       BytesToBase64(c.Response.AttestationObject.AuthData.AttData.CredentialPublicKey),
		Transport:       NormalizeTransports(c.Response.Transports),
		SignCount:       uint(c.Response.AttestationObject.AuthData.Counter),
		UserPresent:     c.Response.AttestationObject.AuthData.Flags.HasUserPresent(),
		UserVerified:    c.Response.AttestationObject.AuthData.Flags.HasUserVerified(),
		BackupEligible:  c.Response.AttestationObject.AuthData.Flags.HasBackupEligible(),
		BackupState:     c.Response.AttestationObject.AuthData.Flags.HasAttestedCredentialData(),
	}
}

func BytesToBase64(b []byte) string {
	return base64.RawURLEncoding.EncodeToString(b)
}

func Base64ToBytes(b string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(b)
}

// Descriptor converts a Credential into a protocol.CredentialDescriptor.
func (c *Credential) Descriptor() protocol.CredentialDescriptor {
	id, err := base64.RawURLEncoding.DecodeString(c.CredentialId)
	if err != nil {
		panic(err)
	}
	return protocol.CredentialDescriptor{
		Type:            protocol.PublicKeyCredentialType,
		CredentialID:    id,
		Transport:       ConvertTransports(c.Transport),
		AttestationType: c.AttestationType,
	}
}

// This is a signal that the authenticator may be cloned, see CloneWarning above for more information.
func (a *Credential) UpdateCounter(authDataCount uint) {
	if authDataCount <= a.SignCount && (authDataCount != 0 || a.SignCount != 0) {
		a.CloneWarning = true
		return
	}

	a.SignCount = authDataCount
}
