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
	d.Authentication = append(d.Authentication, &VerificationRelationship{
		Reference:          d.Id,
		VerificationMethod: v,
	})
}

// SetAuthentication sets the AuthenticationMethod of the DID Document to a PubKey and configured with the given options
func (d *DidDocument) SetAuthentication(pub *crypto.PubKey, opts ...VerificationMethodOption) (*VerificationMethod, error) {
	vm, err := NewVMFromPubKey(pub, DIDMethod_DIDMethod_KEY, opts...)
	if err != nil {
		return nil, err
	}
	d.AddAuthentication(vm)
	return vm, nil
}

// UpdateAuthentication updates the AuthenticationMethod of the DID Document to a PubKey and configured with the given options
func (d *DidDocument) UpdateAuthentication(vm *VerificationMethod) {
	for _, a := range d.Authentication {
		if a.VerificationMethod.Id == vm.Id {
			a.VerificationMethod = vm
		}
	}
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
