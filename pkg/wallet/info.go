package wallet

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/go-webauthn/webauthn/protocol"
)

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

type InfoFile struct {
	Creds      Credentials `json:"credentials"`
	Properties Properties  `json:"properties"`
}

func (i *InfoFile) Marshal() ([]byte, error) {
	return json.Marshal(i)
}

func (i *InfoFile) Unmarshal(data []byte) error {
	return json.Unmarshal(data, i)
}
