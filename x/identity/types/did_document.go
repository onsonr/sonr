// Main DID Document Constructor Methods
// I.e. Document allows for Reconstruction from Storage of DID Document and Wallet
package types

import (
	fmt "fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonrhq/core/pkg/crypto"
)

// NewBlankDocument creates a blank document to begin the WebAuthnProcess
func NewBlankDocument(idStr string) *DidDocument {
	return &DidDocument{
		Id:                   idStr,
		Context:              []string{DefaultParams().DidBaseContext, DefaultParams().DidMethodContext},
		Controller:           []string{},
		VerificationMethod:   make([]*VerificationMethod, 0),
		Authentication:       make([]string, 0),
		AssertionMethod:      make([]string, 0),
		CapabilityInvocation: make([]string, 0),
		CapabilityDelegation: make([]string, 0),
		KeyAgreement:         make([]string, 0),
		Service:              make([]*Service, 0),
		AlsoKnownAs:          make([]string, 0),
	}
}

// NewDocument creates a new DID Document from a wallet public key
func NewDocument(pk *crypto.PubKey, opts ...VerificationMethodOption) *DidDocument {
	doc := NewBlankDocument("")
	vm, err := NewVerificationMethodFromPubKey(pk, DIDMethod_DIDMethod_BLOCKCHAIN)
	if err != nil {
		panic(err)
	}
	doc.VerificationMethod = append(doc.VerificationMethod, vm)
	return doc
}

// Address returns the address of the DID
func (d *DidDocument) Address() string {
	ptrs := strings.Split(d.Id, ":")
	return fmt.Sprintf("%s%s", ptrs[len(ptrs)-2], ptrs[len(ptrs)-1])
}

// AccAddress returns the account address of the DID
func (d *DidDocument) AccAddress() (sdk.AccAddress, error) {
	return ConvertDidToAccAddress(d.Id)
}

// CheckAccAddress checks if the provided sdk.AccAddress or string matches the DID ID
func (d *DidDocument) CheckAccAddress(t interface{}) bool {
	docAccAddr, err := d.AccAddress()
	if err != nil {
		return false
	}

	switch t.(type) {
	case sdk.AccAddress:
		return t.(sdk.AccAddress).Equals(docAccAddr)
	case string:
		addr, err := sdk.AccAddressFromBech32(t.(string))
		if err != nil {
			return false
		}
		return addr.Equals(docAccAddr)
	default:
		return false
	}
}

// GetVerificationMethodByFragment returns the VerificationMethod with the given fragment
func (d *DidDocument) GetVerificationMethodByFragment(fragment string) *VerificationMethod {
	for _, vm := range d.VerificationMethod {
		if vm.Id == fmt.Sprintf("%s#%s", d.Id, fragment) {
			return vm
		}
	}
	return nil
}

// SetMetadata sets the metadata of the document
func (vm *DidDocument) SetMetadata(data map[string]string) {
	vm.Metadata = MapToKeyValueList(data)
}
