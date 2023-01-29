// Main DID Document Constructor Methods
// I.e. Document allows for Reconstruction from Storage of DID Document and Wallet
package types

import (
	fmt "fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	common "github.com/sonrhq/core/pkg/common"
)

// NewDocument takes a SNRPubKey and returns a new DID Document
func NewDocument(pubKey common.SNRPubKey) (*DidDocument, error) {
	pk, err := PubKeyFromCommon(pubKey)
	if err != nil {
		return nil, err
	}
	addr, err := pk.Bech32("snr")
	if err != nil {
		return nil, err
	}
	doc := NewBlankDocument(pk.DID())
	vm, err := pk.VerificationMethod(WithBlockchainAccount(addr))
	if err != nil {
		return nil, err
	}
	doc.AddAssertion(vm)
	return doc, nil
}

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

// SetMetadata sets the metadata of the document
func (vm *DidDocument) SetMetadata(data map[string]string) {
	vm.Metadata = MapToKeyValueList(data)
}
