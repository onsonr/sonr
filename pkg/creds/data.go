package creds

import (
	"bytes"
	"encoding/json"
)

// CredentialData is the list of credentials that are stored in the vault, indexed by their origin and list of credentials.
type CredentialData struct {
	// Data is the map of origins to the list of credentials.
	Data map[string][]*Credential `json:"data"`
}

// NewCredentialData will return a new Credentials struct.
func NewCredentialData() *CredentialData {
	return &CredentialData{}
}

// AddCredential will add a credential to the vault.
func (c *CredentialData) AddCredential(origin string, credential *Credential) {
	if c.Data == nil {
		c.Data = make(map[string][]*Credential)
	}

	if _, ok := c.Data[origin]; !ok {
		c.Data[origin] = make([]*Credential, 0)
	}

	c.Data[origin] = append(c.Data[origin], credential)
}

// ListCredentials will return a list of credentials for a given origin.
func (c *CredentialData) ListCredentials(origin string) []*Credential {
	if c.Data == nil {
		return nil
	}
	if _, ok := c.Data[origin]; !ok {
		return nil
	}

	return c.Data[origin]
}

// RemoveCredential will remove a credential from the vault.
func (c *CredentialData) RemoveCredential(origin string, credential *Credential) {
	if c.Data == nil {
		return
	}
	if _, ok := c.Data[origin]; !ok {
		return
	}
	for i, cred := range c.Data[origin] {
		if bytes.Equal(cred.ID, credential.ID) {
			c.Data[origin] = append(c.Data[origin][:i], c.Data[origin][i+1:]...)
		}
	}
}

// Marshal returns the JSON encoding of the Credentials.
func (c *CredentialData) Marshal() ([]byte, error) {
	return json.Marshal(c)
}

// Unmarshal parses the JSON-encoded data and stores the result in the Credentials.
func (c *CredentialData) Unmarshal(data []byte) error {
	return json.Unmarshal(data, c)
}
