// Utility functions for DID Assertion Method - https://w3c.github.io/did-core/#assertion
// I.e. Verification Material for Wallets. This is the default Verification Method for DID Documents. (snr, btc, eth, etc.)
package types

import "github.com/sonr-hq/sonr/pkg/common/crypto"

// FindAssertionMethod finds a VerificationMethod by its ID
func (d *DidDocument) FindAssertionMethod(id string) *VerificationMethod {
	return d.AssertionMethod.FindByID(id)
}

// FindAssertionMethodByFragment finds a VerificationMethod by its fragment
func (d *DidDocument) FindAssertionMethodByFragment(fragment string) *VerificationMethod {
	return d.AssertionMethod.FindByFragment(fragment)
}

// AddAssertionMethod adds a VerificationMethod as AssertionMethod
// If the controller is not set, it will be set to the documents ID
func (d *DidDocument) AddAssertion(v *VerificationMethod) {
	if v.Controller == "" {
		v.Controller = d.ID
	}
	d.VerificationMethod.Add(v)
	d.AssertionMethod.Add(v)
}

func (d *DidDocument) AddBlockchainAccount(wallet crypto.WalletShare) error {
	pb, err := wallet.PublicKey()
	if err != nil {
		return err
	}
	vm, err := NewSecp256k1VM(pb, WithBlockchainAccount(wallet.Address()))
	if err != nil {
		return err
	}
	d.VerificationMethod.Add(vm)
	d.AssertionMethod.Add(vm)
	return nil
}
