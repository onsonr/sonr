// Utility functions for DID Assertion Method - https://w3c.github.io/did-core/#assertion
// I.e. Verification Material for Wallets. This is the default Verification Method for DID Documents. (snr, btc, eth, etc.)
package types

import common "github.com/sonrhq/core/pkg/common"

// SetAssertion sets the AssertionMethod of the DID Document to a PubKey and configured with the given options
func (d *DidDocument) SetAssertion(pub common.SNRPubKey, opts ...VerificationMethodOption) error {
	pubKey, err := PubKeyFromCommon(pub)
	if err != nil {
		return err
	}
	vm, err := pubKey.VerificationMethod(opts...)
	if err != nil {
		return err
	}
	d.AddAssertion(vm)
	return nil
}

// AssertionMethodCount returns the number of Assertion Methods
func (vm *DidDocument) AssertionMethodCount() int {
	return len(vm.AssertionMethod)
}

// AddAssertionMethod adds a VerificationMethod as AssertionMethod
// If the controller is not set, it will be set to the documents ID
func (d *DidDocument) AddAssertion(v *VerificationMethod) {
	if v.Controller == "" {
		v.Controller = d.Id
	}
	d.VerificationMethod = append(d.VerificationMethod, v)
	d.AssertionMethod = append(d.AssertionMethod, &VerificationRelationship{VerificationMethod: v, Reference: d.Id})
}

// ListBlockchainAccounts returns a list of Blockchain Accounts by the address prefix
func (d *DidDocument) ListBlockchainAccounts() []*VerificationMethod {
	accs := make([]*VerificationMethod, 0)
	for _, vm := range d.AssertionMethod {
		if vm.VerificationMethod.IsBlockchainAccount() {
			accs = append(accs, vm.VerificationMethod)
		}
	}
	return accs
}
