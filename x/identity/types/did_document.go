// Main DID Document Constructor Methods
// I.e. Document allows for Reconstruction from Storage of DID Document and Wallet
package types

import (
	fmt "fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewBlankDocument creates a blank document to begin the WebAuthnProcess
func NewBlankDocument(idStr string) *DidDocument {
	return &DidDocument{
		Id:                   idStr,
		Context:              []string{DefaultParams().DidBaseContext, DefaultParams().DidMethodContext},
		Controller:           []string{},
		VerificationMethod:   make([]*VerificationMethod, 0),
		Authentication:       make([]*VerificationRelationship, 0),
		AssertionMethod:      make([]*VerificationRelationship, 0),
		CapabilityInvocation: make([]*VerificationRelationship, 0),
		CapabilityDelegation: make([]*VerificationRelationship, 0),
		KeyAgreement:         make([]*VerificationRelationship, 0),
		Service:              make([]*Service, 0),
		AlsoKnownAs:          make([]string, 0),
	}
}

// BlankDocument creates a blank document to begin the WebAuthnProcess
func NewDocumentFromAKA(akaStr string) *DidDocument {
	return &DidDocument{
		Id:                   fmt.Sprintf("did:tmp:%s", akaStr),
		Context:              []string{DefaultParams().DidBaseContext, DefaultParams().DidMethodContext},
		Controller:           []string{},
		VerificationMethod:   make([]*VerificationMethod, 0),
		Authentication:       make([]*VerificationRelationship, 0),
		AssertionMethod:      make([]*VerificationRelationship, 0),
		CapabilityInvocation: make([]*VerificationRelationship, 0),
		CapabilityDelegation: make([]*VerificationRelationship, 0),
		KeyAgreement:         make([]*VerificationRelationship, 0),
		Service:              make([]*Service, 0),
		AlsoKnownAs: []string{
			akaStr,
		},
	}
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

// GetVerificationMethodByBlockchainAccountID returns the VerificationMethod with the given blockchain account
func (d *DidDocument) GetVerificationMethodByBlockchainAccountID(account string) *VerificationMethod {
	for _, vr := range d.AssertionMethod {
		if vr.VerificationMethod.BlockchainAccountId == account {
			return vr.VerificationMethod
		}
	}
	return nil
}

// SetMetadata sets the metadata of the document
func (vm *DidDocument) SetMetadata(data map[string]string) {
	vm.Metadata = MapToKeyValueList(data)
}
