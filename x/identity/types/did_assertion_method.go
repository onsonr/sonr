// Utility functions for DID Assertion Method - https://w3c.github.io/did-core/#assertion
// I.e. Verification Material for Wallets. This is the default Verification Method for DID Documents. (snr, btc, eth, etc.)
package types

import "github.com/sonrhq/core/pkg/crypto"

// AddBlockchainAccount creates a verification method from a new wallet account
func (d *DidDocument) AddBlockchainAccount(accName string, ct crypto.CoinType, pk *crypto.PubKey, metadata ...*KeyValuePair) (*VerificationMethod, error) {
	accAddress := ct.FormatAddress(pk)
	vm := &VerificationMethod{
		Id:                  NewBlockchainID(accAddress, accName),
		Type:                crypto.Secp256k1KeyType.PrettyString(),
		BlockchainAccountId: accAddress,
		Controller:          d.Id,
		PublicKeyMultibase:  pk.Base64(),
		Metadata:            metadata,
	}
	d.AddAssertion(vm)
	return vm, nil
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
	d.AssertionMethod = append(d.AssertionMethod, v.Id)
}
