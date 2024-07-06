package orm

import (
	"github.com/go-webauthn/webauthn/protocol"
	"gorm.io/gorm"
)

type CredentialFlags struct {
	// Flag UP indicates the users presence.
	UserPresent bool `json:"userPresent"`

	// Flag UV indicates the user performed verification.
	UserVerified bool `json:"userVerified"`

	// Flag BE indicates the credential is able to be backed up and/or sync'd between devices. This should NEVER change.
	BackupEligible bool `json:"backupEligible"`

	// Flag BS indicates the credential has been backed up and/or sync'd. This value can change but it's recommended
	// that RP's keep track of this value.
	BackupState bool `json:"backupState"`
}

// Credential contains all needed information about a WebAuthn credential for storage.
type Credential struct {
	gorm.Model
	DisplayName     string
	Origin          string
	Controller      string
	AttestationType string `json:"attestationType"`
	DID             string
	ID              []byte                            `json:"id"`
	PublicKey       []byte                            `json:"publicKey"`
	Transport       []protocol.AuthenticatorTransport `json:"transport"`
	Authenticator   Authenticator                     `json:"authenticator"`
	Flags           CredentialFlags                   `json:"flags"`
}

// MakeNewCredential will return a credential pointer on successful validation of a registration response.
func MakeNewCredential(c *protocol.ParsedCredentialCreationData) *Credential {
	return &Credential{
		ID:              c.Response.AttestationObject.AuthData.AttData.CredentialID,
		PublicKey:       c.Response.AttestationObject.AuthData.AttData.CredentialPublicKey,
		AttestationType: c.Response.AttestationObject.Format,
		Transport:       c.Response.Transports,
		Flags: CredentialFlags{
			UserPresent:    c.Response.AttestationObject.AuthData.Flags.HasUserPresent(),
			UserVerified:   c.Response.AttestationObject.AuthData.Flags.HasUserVerified(),
			BackupEligible: c.Response.AttestationObject.AuthData.Flags.HasBackupEligible(),
			BackupState:    c.Response.AttestationObject.AuthData.Flags.HasBackupState(),
		},
		Authenticator: Authenticator{
			AAGUID:     c.Response.AttestationObject.AuthData.AttData.AAGUID,
			SignCount:  c.Response.AttestationObject.AuthData.Counter,
			Attachment: c.AuthenticatorAttachment,
		},
	}
}

// Descriptor converts a Credential into a protocol.CredentialDescriptor.
func (c *Credential) Descriptor() protocol.CredentialDescriptor {
	return protocol.CredentialDescriptor{
		Type:            protocol.PublicKeyCredentialType,
		CredentialID:    c.ID,
		Transport:       c.Transport,
		AttestationType: c.AttestationType,
	}
}
