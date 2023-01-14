// Utility functions for DID Authentication - https://w3c.github.io/did-core/#authentication
// I.e. Verification Material for Webauthn Credentials or KeyPrints. These are used to unlock the Controller Wallet.
package types

import "github.com/sonr-hq/sonr/pkg/common"

// FindAuthenticationMethod finds a VerificationMethod by its ID
func (d *DidDocument) FindAuthenticationMethod(id string) *VerificationMethod {
	return d.Authentication.FindByID(id)
}

// FindAuthenticationMethodByFragment finds a VerificationMethod by its fragment
func (d *DidDocument) FindAuthenticationMethodByFragment(fragment string) *VerificationMethod {
	return d.Authentication.FindByFragment(fragment)
}

// AddAuthenticationMethod adds a VerificationMethod as AuthenticationMethod
// If the controller is not set, it will be set to the document's ID
func (d *DidDocument) AddAuthentication(v *VerificationMethod) {
	if v.Controller == "" {
		v.Controller = d.ID
	}
	d.VerificationMethod.Add(v)
	d.Authentication.Add(v)
}

func (d *DidDocument) AddWebauthnCredential(cred *common.WebauthnCredential, label string) error {
	vm, err := NewWebAuthnVM(cred, WithIDFragmentSuffix(label))
	if err != nil {
		return err
	}
	d.VerificationMethod.Add(vm)
	d.Authentication.Add(vm)
	return nil
}
