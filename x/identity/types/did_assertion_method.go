// Utility functions for DID Assertion Method - https://w3c.github.io/did-core/#assertion
// I.e. Verification Material for Wallets. This is the default Verification Method for DID Documents. (snr, btc, eth, etc.)
package types

import "github.com/sonrhq/core/pkg/crypto"

// AddBlockchainAccount creates a verification method from a new wallet account
func (d *DidDocument) AddBlockchainAccount(accName string, ct crypto.CoinType, pk *crypto.PubKey, metadata ...*KeyValuePair) (*VerificationMethod, error) {
	accAddress, err := pk.Bech32(ct.AddrPrefix())
	if err != nil {
		return nil, err
	}
	vm := &VerificationMethod{
		Id:                  NewBlockchainID(accAddress, accName),
		Type:                crypto.Secp256k1KeyType.PrettyString(),
		BlockchainAccountId: accAddress,
		Controller:          d.Id,
		PublicKeyMultibase:  pk.Multibase(),
		Metadata:            metadata,
	}
	d.AddAssertion(vm)
	return vm, nil
}

// AssertionMethodCount returns the number of Assertion Methods
func (vm *DidDocument) AssertionMethodCount() int {
	return len(vm.AssertionMethod)
}

// UpdateAssertion updates the AssertionMethod of the DID Document to a PubKey and configured with the given options
func (d *DidDocument) UpdateAssertion(vm *VerificationMethod) {
	for _, a := range d.AssertionMethod {
		if a.VerificationMethod.Id == vm.Id {
			a.VerificationMethod = vm
		}
	}
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
