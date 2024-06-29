package models

import (
	"bytes"
	"encoding/json"
	"errors"

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
	Controller      string
	AttestationType string                            `json:"attestationType"`
	ID              []byte                            `json:"id"`
	PublicKey       []byte                            `json:"publicKey"`
	Transport       []protocol.AuthenticatorTransport `json:"transport"`
	Authenticator   Authenticator                     `json:"authenticator"`
	Flags           CredentialFlags                   `json:"flags"`
}

// MakeNewCredential will return a credential pointer on successful validation of a registration response.
func MakeNewCredential(c *protocol.ParsedCredentialCreationData) (*Credential, error) {
	newCredential := &Credential{
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

	return newCredential, nil
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

// Credentials is a map of credentials
type Credentials map[string][]*Credential

// NewCredentials creates a new Credentials map
func NewCredentials() Credentials {
	return make(Credentials)
}

// LinkCredential will add a credential to the vault.
func (c Credentials) LinkCredential(origin string, credential *Credential) error {
	if origin == "" {
		return errors.New("origin cannot be empty")
	}
	if _, ok := c[origin]; !ok {
		c[origin] = make([]*Credential, 0)
	}

	c[origin] = append(c[origin], credential)
	return nil
}

// GetCredentials will return a list of credentials for a given origin.
func (c Credentials) GetCredentials(origin string) []protocol.CredentialDescriptor {
	cds := make([]protocol.CredentialDescriptor, 0)
	cs, ok := c[origin]
	if !ok {
		return nil
	}
	for _, c := range cs {
		cds = append(cds, c.Descriptor())
	}
	return cds
}

// UnlinkCredential will remove a credential from the vault.
func (c Credentials) UnlinkCredential(origin string, credential *Credential) {
	if _, ok := c[origin]; !ok {
		return
	}
	for i, cred := range c[origin] {
		if bytes.Equal(cred.ID, credential.ID) {
			c[origin] = append(c[origin][:i], c[origin][i+1:]...)
		}
	}
}

// Marshal marshals the Credentials to a byte slice
func (c Credentials) Marshal() ([]byte, error) {
	return json.Marshal(c)
}

// Unmarshal unmarshals the Credentials from a byte slice
func (c *Credentials) Unmarshal(data []byte) error {
	return json.Unmarshal(data, c)
}
