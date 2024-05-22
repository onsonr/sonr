package vault

import (
	"bytes"

	"github.com/di-dao/sonr/pkg/auth"
)

// AddCredential will add a credential to the vault.
func (c *vault) AddCredential(origin string, credential *auth.Credential) {
	if c.wallet.Credentials == nil {
		c.wallet.Credentials = make(map[string][]*auth.Credential)
	}

	if _, ok := c.wallet.Credentials[origin]; !ok {
		c.wallet.Credentials[origin] = make([]*auth.Credential, 0)
	}

	c.wallet.Credentials[origin] = append(c.wallet.Credentials[origin], credential)
}

// ListCredentials will return a list of credentials for a given origin.
func (c *vault) ListCredentials(origin string) []*auth.Credential {
	if c.wallet.Credentials == nil {
		return nil
	}
	if _, ok := c.wallet.Credentials[origin]; !ok {
		return nil
	}

	return c.wallet.Credentials[origin]
}

// RemoveCredential will remove a credential from the vault.
func (c *vault) RemoveCredential(origin string, credential *auth.Credential) {
	if c.wallet.Credentials == nil {
		return
	}
	if _, ok := c.wallet.Credentials[origin]; !ok {
		return
	}
	for i, cred := range c.wallet.Credentials[origin] {
		if bytes.Equal(cred.ID, credential.ID) {
			c.wallet.Credentials[origin] = append(c.wallet.Credentials[origin][:i], c.wallet.Credentials[origin][i+1:]...)
		}
	}
}
