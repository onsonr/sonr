// Utility functions for DID Assertion Method - https://w3c.github.io/did-core/#assertion
// I.e. Verification Material for Wallets. This is the default Verification Method for DID Documents. (snr, btc, eth, etc.)
package types

import (
	"fmt"

	"github.com/sonr-hq/sonr/pkg/common/crypto"
)

// KnownWalletPrefixes is an enum of known wallet prefixes
type ChainWalletPrefix int

const (
	// KnownWalletPrefixes is an enum of known wallet prefixes
	ChainWalletPrefixNone ChainWalletPrefix = iota
	ChainWalletPrefixSNR
	ChainWalletPrefixBTC
	ChainWalletPrefixETH
)

func NewWalletPrefix(prefix string) ChainWalletPrefix {
	switch prefix {
	case "snr":
		return ChainWalletPrefixSNR
	case "btc":
		return ChainWalletPrefixBTC
	case "eth":
		return ChainWalletPrefixETH
	case "0x":
		return ChainWalletPrefixETH
	default:
		return ChainWalletPrefixNone
	}
}

func (k ChainWalletPrefix) String() string {
	return [...]string{"account", "snr", "btc", "eth"}[k]
}

// Prefix returns the prefix of the wallet
func (k ChainWalletPrefix) Prefix() string {
	if k == ChainWalletPrefixETH {
		return "0x"
	}
	return k.String()
}

func createFragment(wallet crypto.WalletShare, didDoc *DidDocument) string {
	count := didDoc.GetBlockchainAccountCount(wallet.Prefix())
	cwpfx := NewWalletPrefix(wallet.Prefix())
	return fmt.Sprintf("%s-%d", cwpfx.String(), count)
}

// FindAssertionMethod finds a VerificationMethod by its ID
func (d *DidDocument) FindAssertionMethod(id string) *VerificationMethod {
	return d.AssertionMethod.FindByID(id)
}

// FindAssertionMethodByFragment finds a VerificationMethod by its fragment
func (d *DidDocument) FindAssertionMethodByFragment(fragment string) *VerificationMethod {
	return d.AssertionMethod.FindByFragment(fragment)[0]
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
	frag := fmt.Sprintf("%s-%d", wallet.Prefix(), d.GetBlockchainAccountCount(wallet.Prefix())+1)
	vm, err := NewSecp256k1VM(pb, WithBlockchainAccount(wallet.Address()), WithIDFragmentSuffix(frag))
	if err != nil {
		return err
	}
	data := map[string]string{
		"index":      fmt.Sprint(wallet.Index()),
		"prefix":     wallet.Prefix(),
		"party_id":   string(wallet.SelfID()),
		"blockchain": ConvertBoolToString(true),
	}
	vm.SetMetadata(data)
	d.VerificationMethod.Add(vm)
	d.AssertionMethod.Add(vm)
	return nil
}

// GetBlockchainAccountCount returns the number of Blockchain Accounts by the address prefix
func (d *DidDocument) GetBlockchainAccountCount(prefix string) int {
	return len(d.AssertionMethod.FindByFragment(prefix))
}

// ListBlockchainAccounts returns a list of Blockchain Accounts by the address prefix
func (d *DidDocument) ListBlockchainAccounts() []*VerificationMethod {
	accs := make([]*VerificationMethod, 0)
	for _, vm := range d.AssertionMethod.Data {
		if vm.VerificationMethod.IsBlockchainAccount() {
			accs = append(accs, vm.VerificationMethod)
		}
	}
	return accs
}
