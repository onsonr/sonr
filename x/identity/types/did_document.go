// Main DID Document Constructor Methods
// I.e. Document allows for Reconstruction from Storage of DID Document and Wallet
package types

import (
	"errors"
	fmt "fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonrhq/core/internal/crypto"
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
		AlsoKnownAs:          make([]string, 0),
	}
}

// NewDocument creates a new DID Document from a wallet public key
func NewDocument(pk *crypto.PubKey, opts ...VerificationMethodOption) *DidDocument {
	doc := NewBlankDocument("")
	vm, err := NewVerificationMethodFromPubKey(pk, DIDMethod_DIDMethod_SONR)
	if err != nil {
		panic(err)
	}
	doc.VerificationMethod = append(doc.VerificationMethod, vm)
	return doc
}

// AccAddress returns the SONR address of the DID
func (d *DidDocument) AccAddress() (sdk.AccAddress, error) {
	for _, vm := range d.VerificationMethod {
		if strings.Contains(vm.Id, "did:sonr") {
			return sdk.AccAddressFromBech32(vm.BlockchainAccountId)
		}
	}
	return nil, errors.New("No SONR address found")
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

// GetAuthenticationMethod returns a VerificationMethod if the did exists in the authentication array
func (d *DidDocument) GetAuthenticationMethod(did string) (*VerificationMethod, error) {
	hasAuthId := false
	for _, id := range d.Authentication {
		if id == did {
			hasAuthId = true
		}
	}

	if !hasAuthId {
		return nil, errors.New("The did doesnt exist in the authentication method")
	}

	for _, vm := range d.VerificationMethod {
		if vm.Id == did {
			return vm, nil
		}
	}
	return nil, errors.New("Could not find VerificationMethod for DID")
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

// ImportVerificationMethods imports the given VerificationMethods into the document
func (d *DidDocument) ImportVerificationMethods(category string, vms ...VerificationMethod) {
	idList := []string{}
	for _, vm := range vms {
		if !d.Contains(vm.Id) {
			d.VerificationMethod = append(d.VerificationMethod, &vm)
			idList = append(idList, vm.Id)
		}
	}
	switch strings.ToLower(category) {
	case "authentication":
		d.Authentication = append(d.Authentication, idList...)
	case "assertionmethod":
		d.AssertionMethod = append(d.AssertionMethod, idList...)
	case "capabilityinvocation":
		d.CapabilityInvocation = append(d.CapabilityInvocation, idList...)
	case "capabilitydelegation":
		d.CapabilityDelegation = append(d.CapabilityDelegation, idList...)
	case "keyagreement":
		d.KeyAgreement = append(d.KeyAgreement, idList...)
	}
}

// Contains is a method which recursively checks if a given did is contained within the document
func (d *DidDocument) Contains(did string) bool {
	if d.Id == did {
		return true
	}
	for _, vm := range d.VerificationMethod {
		if vm.Id == did {
			return true
		}
	}
	return false
}

// IsAuthentication checks if the given VerificationMethod is used for authentication
func (d *DidDocument) IsAuthentication(vm *VerificationMethod) bool {
	for _, auth := range d.Authentication {
		if auth == vm.Id {
			return true
		}
	}
	return false
}

// IsAssertionMethod checks if the given VerificationMethod is used for assertion
func (d *DidDocument) IsAssertionMethod(vm *VerificationMethod) bool {
	for _, auth := range d.AssertionMethod {
		if auth == vm.Id {
			return true
		}
	}
	return false
}

// IsCapabilityInvocation checks if the given VerificationMethod is used for capability invocation
func (d *DidDocument) IsCapabilityInvocation(vm *VerificationMethod) bool {
	for _, auth := range d.CapabilityInvocation {
		if auth == vm.Id {
			return true
		}
	}
	return false
}

// IsCapabilityDelegation checks if the given VerificationMethod is used for capability delegation
func (d *DidDocument) IsCapabilityDelegation(vm *VerificationMethod) bool {
	for _, auth := range d.CapabilityDelegation {
		if auth == vm.Id {
			return true
		}
	}
	return false
}

// IsKeyAgreement checks if the given VerificationMethod is used for key agreement
func (d *DidDocument) IsKeyAgreement(vm *VerificationMethod) bool {
	for _, auth := range d.KeyAgreement {
		if auth == vm.Id {
			return true
		}
	}
	return false
}

// MatchesAddress checks if the given VerificationMethod matches the given address
func (d *DidDocument) MatchesAddress(addr string) bool {
	for _, vm := range d.VerificationMethod {
		if vm.BlockchainAccountId == addr {
			return true
		}
	}
	return false
}

// Method returns the DID method of the document
func (d *DidDocument) DIDMethod() string {
	return strings.Split(d.Id, ":")[1]
}

// Identifier returns the DID identifier of the document
func (d *DidDocument) DIDIdentifier() string {
	return strings.Split(d.Id, ":")[2]
}

// Fragment returns the DID fragment of the document
func (d *DidDocument) DIDFragment() string {
	return strings.Split(d.Id, "#")[1]
}

// Username is the first item in the AKA array
func (d *DidDocument) FindUsername() string {
	if len(d.AlsoKnownAs) > 0 {
		return d.AlsoKnownAs[0]
	}
	return "tmp"
}

// FindPrimaryAddress is the first item in the assertion method without the did: prefix
func (d *DidDocument) FindPrimaryAddress() string {
	ptrs := strings.Split(d.Id, ":")
	return ptrs[len(ptrs)-1]
}

// ListAuthenticationMethods returns a list of all authentication methods
func (d *DidDocument) ListAuthenticationMethods() []*VerificationMethod {
	vms := []*VerificationMethod{}
	for _, vm := range d.VerificationMethod {
		if d.IsAuthentication(vm) {
			vms = append(vms, vm)
		}
	}
	return vms
}

// ListAssertionMethods returns a list of all assertion methods
func (d *DidDocument) ListAssertionMethods() []*VerificationMethod {
	vms := []*VerificationMethod{}
	for _, vm := range d.VerificationMethod {
		if d.IsAssertionMethod(vm) {
			vms = append(vms, vm)
		}
	}
	return vms
}
