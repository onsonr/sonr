package types

import (
	"fmt"
	"strings"

	"github.com/sonrhq/core/internal/crypto"
	vaulttypes "github.com/sonrhq/core/x/vault/types"
)

// BlankIdentity returns a blank Identity
func BlankIdentity() *Identification {
	return &Identification{
		Id:                   "",
		Owner:                "",
		PrimaryAlias:         "",
		Authentication:       make([]string, 0),
		AssertionMethod:      make([]string, 0),
		CapabilityDelegation: make([]string, 0),
		CapabilityInvocation: make([]string, 0),
		AlsoKnownAs:          make([]string, 0),
		Metadata:             "",
	}
}

// NewIdentityFromVaultAccount returns a new Identity from a VaultAccount
func NewIdentityFromVaultAccount(va vaulttypes.Account, controller string) (*Identification, *VerificationRelationship, bool) {
	vm := NewVerificationMethodFromVaultAccount(va, controller)
	wi := NewWalletIdentity(controller, vm.BlockchainAccountId, va.CoinType())
	wi.LinkAuthenticationMethod(vm)
	vr, ok := wi.LinkCapabilityDelegation(vm)
	return wi, vr, ok
}

// NewSonrIdentity returns a new Identity with the given owner address and constructs
// the DID from the owner address
func NewSonrIdentity(ownerAddress string, primaryAlias string) *Identification {
	did := fmt.Sprintf("did:sonr:%s", ownerAddress)
	identity := BlankIdentity()
	identity.Id = did
	identity.Owner = ownerAddress
	identity.SetPrimaryAlias(primaryAlias)
	return identity
}

// NewWalletIdentity takes an ownerAddress, walletAddress, and CoinType and returns a new Identity
// with the given owner address and constructs the DID from the wallet address
func NewWalletIdentity(ownerAddress, walletAddress string, coinType crypto.CoinType) *Identification {
	did := fmt.Sprintf("did:%s:%s", coinType.DidMethod(), walletAddress)
	identity := BlankIdentity()
	identity.Id = did
	identity.Owner = ownerAddress
	return identity
}

// LinkAccountFromVault adds a VerificationMethod to the CapabilityInvocation list of the DID Document and returns the VerificationRelationship
// Returns nil if the VerificationMethod is already in the CapabilityInvocation list
func (id *Identification) LinkAccountFromVault(account vaulttypes.Account) (*VerificationRelationship, bool) {
	vm := NewVerificationMethodFromVaultAccount(account, id.Owner)
	for _, auth := range id.CapabilityInvocation {
		if auth == vm.Id {
			return nil, false
		}
	}
	id.CapabilityInvocation = append(id.CapabilityInvocation, vm.Id)
	vr := &VerificationRelationship{
		Reference:          vm.Id,
		Type:               CapabilityInvocationRelationshipName,
		VerificationMethod: vm,
		Owner:              id.Owner,
	}
	return vr, true
}

// LinkAuthenticationMethod adds a VerificationMethod to the Authentication list of the DID Document and returns the VerificationRelationship
// Returns nil if the VerificationMethod is already in the Authentication list
func (id *Identification) LinkAuthenticationMethod(vm *VerificationMethod) (*VerificationRelationship, bool) {
	for _, auth := range id.Authentication {
		if auth == vm.Id {
			return nil, false
		}
	}
	id.Authentication = append(id.Authentication, vm.Id)
	vr := &VerificationRelationship{
		Reference:          vm.Id,
		Type:               AuthenticationRelationshipName,
		VerificationMethod: vm,
		Owner:              id.Owner,
	}
	return vr, true
}

// LinkAssertionMethod adds a VerificationMethod to the AssertionMethod list of the DID Document and returns the VerificationRelationship
// Returns nil if the VerificationMethod is already in the AssertionMethod list
func (id *Identification) LinkAssertionMethod(vm *VerificationMethod) (*VerificationRelationship, bool) {
	for _, auth := range id.AssertionMethod {
		if auth == vm.Id {
			return nil, false
		}
	}
	id.AssertionMethod = append(id.AssertionMethod, vm.Id)
	vr := &VerificationRelationship{
		Reference:          vm.Id,
		Type:               AssertionRelationshipName,
		VerificationMethod: vm,
		Owner:              id.Owner,
	}
	return vr, true
}

// LinkCapabilityDelegation adds a VerificationMethod to the CapabilityDelegation list of the DID Document and returns the VerificationRelationship
// Returns nil if the VerificationMethod is already in the CapabilityDelegation list
func (id *Identification) LinkCapabilityDelegation(vm *VerificationMethod) (*VerificationRelationship, bool) {
	for _, auth := range id.CapabilityDelegation {
		if auth == vm.Id {
			return nil, false
		}
	}
	id.CapabilityDelegation = append(id.CapabilityDelegation, vm.Id)
	vr := &VerificationRelationship{
		Reference:          vm.Id,
		Type:              CapabilityDelegationRelationshipName,
		VerificationMethod: vm,
		Owner:              id.Owner,
	}
	return vr, true
}



// LinkCapabilityInvocation adds a VerificationMethod to the CapabilityInvocation list of the DID Document and returns the VerificationRelationship
// Returns nil if the VerificationMethod is already in the CapabilityInvocation list
func (id *Identification) LinkCapabilityInvocation(vm *VerificationMethod) (*VerificationRelationship, bool) {
	for _, auth := range id.CapabilityInvocation {
		if auth == vm.Id {
			return nil, false
		}
	}
	id.CapabilityInvocation = append(id.CapabilityInvocation, vm.Id)
	vr := &VerificationRelationship{
		Reference:          vm.Id,
		Type:              CapabilityInvocationRelationshipName,
		VerificationMethod: vm,
		Owner:              id.Owner,
	}
	return vr, true
}

// LinkKeyAgreement adds a VerificationMethod to the KeyAgreement list of the DID Document and returns the VerificationRelationship
// Returns nil if the VerificationMethod is already in the KeyAgreement list
func (id *Identification) LinkKeyAgreement(vm *VerificationMethod) (*VerificationRelationship, bool) {
	for _, auth := range id.KeyAgreement {
		if auth == vm.Id {
			return nil, false
		}
	}
	id.KeyAgreement = append(id.KeyAgreement, vm.Id)
	vr := &VerificationRelationship{
		Reference:          vm.Id,
		Type:               KeyAgreementRelationshipName,
		VerificationMethod: vm,
		Owner:              id.Owner,
	}
	return vr, true
}

// SetPrimaryAlias sets the PrimaryAlias of the DID Document to the given alias and appends the alias to the AlsoKnownAs list
// Returns false if the alias is already the AlsoKnownAs list.
func (id *Identification) SetPrimaryAlias(alias string) bool {
	for _, aka := range id.AlsoKnownAs {
		if aka == alias {
			id.PrimaryAlias = alias
			return false
		}
	}
	id.AlsoKnownAs = append(id.AlsoKnownAs, alias)
	id.PrimaryAlias = alias
	return true
}

// KnownCredentials returns

// ! ||--------------------------------------------------------------------------------||
// ! ||             Blockchain Identities are intended for Wallet Accounts             ||

func ConvertAccAddressToDid(accAddress string) string {
	return strings.ToLower("did:sonr:" + accAddress)
}
