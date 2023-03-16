// Utility functions for DID Authentication - https://w3c.github.io/did-core/#authentication
// I.e. Verification Material for Webauthn Credentials or KeyPrints. These are used to unlock the Controller Wallet.
package types

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/sonrhq/core/pkg/crypto"
)

// AddAuthenticationMethod adds a VerificationMethod as AuthenticationMethod
// If the controller is not set, it will be set to the document's ID
func (d *DidDocument) AddAuthentication(v *VerificationMethod) {
	if v.Controller == "" {
		v.Controller = d.Id
	}
	d.VerificationMethod = append(d.VerificationMethod, v)
	d.Authentication = append(d.Authentication, v.Id)
}

// SetAuthentication sets the AuthenticationMethod of the DID Document to a PubKey and configured with the given options
func (d *DidDocument) SetAuthentication(pub *crypto.PubKey, opts ...VerificationMethodOption) (*VerificationMethod, error) {
	vm, err := NewVerificationMethodFromPubKey(pub, DIDMethod_DIDMethod_KEY, opts...)
	if err != nil {
		return nil, err
	}
	d.AddAuthentication(vm)
	return vm, nil
}

// AllowedWebauthnCredentials returns a list of CredentialDescriptors for Webauthn Credentials
func (d *DidDocument) AllowedWebauthnCredentials() []protocol.CredentialDescriptor {
	allowList := make([]protocol.CredentialDescriptor, 0)
	creds := d.WebAuthnCredentials()
	for _, cred := range creds {
		allowList = append(allowList, cred.Descriptor())
	}
	return allowList
}
