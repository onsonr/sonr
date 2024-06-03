package props

import (
	"bytes"
	"encoding/json"

	"github.com/di-dao/sonr/pkg/auth"
	"github.com/go-webauthn/webauthn/protocol"
)

// Credentials is a map of credentials
type Credentials map[string][]*auth.Credential

// NewCredentials creates a new Credentials map
func NewCredentials() Credentials {
	return make(Credentials)
}

// LinkCredential will add a credential to the vault.
func (c Credentials) LinkCredential(origin string, credential *auth.Credential) {
	if _, ok := c[origin]; !ok {
		c[origin] = make([]*auth.Credential, 0)
	}

	c[origin] = append(c[origin], credential)
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
func (c Credentials) UnlinkCredential(origin string, credential *auth.Credential) {
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
