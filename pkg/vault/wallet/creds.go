package wallet

import (
	"bytes"

	"github.com/di-dao/sonr/pkg/auth"
	"github.com/go-webauthn/webauthn/protocol"
)

// LinkCredential will add a credential to the vault.
func (w *Wallet) LinkCredential(origin string, credential *auth.Credential) {
	if w.Credentials == nil {
		w.Credentials = make(map[string][]*auth.Credential)
	}

	if _, ok := w.Credentials[origin]; !ok {
		w.Credentials[origin] = make([]*auth.Credential, 0)
	}

	w.Credentials[origin] = append(w.Credentials[origin], credential)
}

// GetCredentials will return a list of credentials for a given origin.
func (w *Wallet) GetCredentials(origin string) []protocol.CredentialDescriptor {
	if w.Credentials == nil {
		return nil
	}
	cds := make([]protocol.CredentialDescriptor, 0)
	cs, ok := w.Credentials[origin]
	if !ok {
		return nil
	}
	for _, c := range cs {
		cds = append(cds, c.Descriptor())
	}
	return cds
}

// UnlinkCredential will remove a credential from the vault.
func (w *Wallet) UnlinkCredential(origin string, credential *auth.Credential) {
	if w.Credentials == nil {
		return
	}
	if _, ok := w.Credentials[origin]; !ok {
		return
	}
	for i, cred := range w.Credentials[origin] {
		if bytes.Equal(cred.ID, credential.ID) {
			w.Credentials[origin] = append(w.Credentials[origin][:i], w.Credentials[origin][i+1:]...)
		}
	}
}
