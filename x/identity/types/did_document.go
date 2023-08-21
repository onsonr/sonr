package types

import (
	"encoding/base64"
	fmt "fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/go-webauthn/webauthn/protocol"
	crypto "github.com/sonrhq/core/pkg/crypto"
)

// NewDIDDocument creates a new DIDDocument from an Identification and optional VerificationRelationships
func NewDIDDocument(primaryAccount *crypto.AccountData, authentication *VerificationMethod, alias string) *DIDDocument {
	params := DefaultParams()
	didDoc := &DIDDocument{
		Id:                   primaryAccount.Address,
		Context:              []string{params.AccountDidMethodContext, params.DidBaseContext},
		Authentication:       make([]*VerificationRelationship, 0),
		AssertionMethod:      make([]*VerificationRelationship, 0),
		CapabilityInvocation: make([]*VerificationRelationship, 0),
		CapabilityDelegation: make([]*VerificationRelationship, 0),
		Controller:           []string{authentication.Id},
		AlsoKnownAs:          []string{alias},
	}
	didDoc.LinkAuthenticationMethod(authentication)
	return didDoc
}

// SearchRelationshipsByCoinType returns all verification relationships of a given the query options
func (d *DIDDocument) SearchRelationshipsByCoinType(coinType crypto.CoinType) []*VerificationRelationship {
	method := coinType.DIDMethod()
	relationships := make([]*VerificationRelationship, 0)
	for _, relationship := range d.Authentication {
		if strings.Contains(relationship.Reference, method) {
			relationships = append(relationships, relationship)
		}
	}
	for _, relationship := range d.AssertionMethod {
		if strings.Contains(relationship.Reference, method) {
			relationships = append(relationships, relationship)
		}
	}
	for _, relationship := range d.CapabilityInvocation {
		if strings.Contains(relationship.Reference, method) {
			relationships = append(relationships, relationship)
		}
	}
	for _, relationship := range d.CapabilityDelegation {
		if strings.Contains(relationship.Reference, method) {
			relationships = append(relationships, relationship)
		}
	}
	for _, relationship := range d.KeyAgreement {
		if strings.Contains(relationship.Reference, method) {
			relationships = append(relationships, relationship)
		}
	}
	return relationships
}

// ListAuthenticationVerificationMethods returns all the VerificationMethods for the AuthenticationRelationships
func (d *DIDDocument) ListAuthenticationVerificationMethods() []*VerificationMethod {
	vms := make([]*VerificationMethod, 0)
	for _, relationship := range d.Authentication {
		vms = append(vms, relationship.VerificationMethod)
	}
	return vms
}

// Address returns the address of the DIDDocument
func (d *DIDDocument) Address() string {
	return strings.Split(d.Id, ":")[2]
}

// LinkAuthenticationMethod adds a VerificationMethod to the Authentication list of the DID Document and returns the VerificationRelationship
// Returns nil if the VerificationMethod is already in the Authentication list
func (id *DIDDocument) LinkAuthenticationMethod(vm *VerificationMethod) (*VerificationRelationship, bool) {
	for _, auth := range id.Authentication {
		if auth.Reference == vm.Id {
			return nil, false
		}
	}
	vm.Controller = id.Id
	vr := &VerificationRelationship{
		Reference:          vm.Id,
		Type:               AuthenticationRelationshipName,
		VerificationMethod: vm,
		Owner:              id.Id,
	}
	id.Authentication = append(id.Authentication, vr)
	id.Controller = append(id.Controller, vm.Id)
	return vr, true
}

// LinkAssertionMethod adds a VerificationMethod to the AssertionMethod list of the DID Document and returns the VerificationRelationship
// Returns nil if the VerificationMethod is already in the AssertionMethod list
func (id *DIDDocument) LinkAssertionMethod(vm *VerificationMethod) (*VerificationRelationship, bool) {
	for _, assertionMethod := range id.AssertionMethod {
		if assertionMethod.Reference == vm.Id {
			return nil, false
		}
	}

	vr := &VerificationRelationship{
		Reference:          vm.Id,
		Type:               AssertionRelationshipName,
		VerificationMethod: vm,
		Owner:              id.Id,
	}
	id.AssertionMethod = append(id.AssertionMethod, vr)
	return vr, true
}

// LinkCapabilityDelegation adds a VerificationMethod to the CapabilityDelegation list of the DID Document and returns the VerificationRelationship
// Returns nil if the VerificationMethod is already in the CapabilityDelegation list
func (id *DIDDocument) LinkCapabilityDelegation(vm *VerificationMethod) (*VerificationRelationship, bool) {
	for _, capabilityDelegation := range id.CapabilityDelegation {
		if capabilityDelegation.Reference == vm.Id {
			return nil, false
		}
	}

	vr := &VerificationRelationship{
		Reference:          vm.Id,
		Type:               CapabilityDelegationRelationshipName,
		VerificationMethod: vm,
		Owner:              id.Id,
	}
	id.CapabilityDelegation = append(id.CapabilityDelegation, vr)
	return vr, true
}

// LinkCapabilityInvocation adds a VerificationMethod to the CapabilityInvocation list of the DID Document and returns the VerificationRelationship
// Returns nil if the VerificationMethod is already in the CapabilityInvocation list
func (id *DIDDocument) LinkCapabilityInvocation(vm *VerificationMethod) (*VerificationRelationship, bool) {
	for _, capabilityInvocation := range id.CapabilityInvocation {
		if capabilityInvocation.Reference == vm.Id {
			return nil, false
		}
	}

	vr := &VerificationRelationship{
		Reference:          vm.Id,
		Type:               CapabilityInvocationRelationshipName,
		VerificationMethod: vm,
		Owner:              id.Id,
	}
	id.CapabilityInvocation = append(id.CapabilityInvocation, vr)
	return vr, true
}

// LinkKeyAgreement adds a VerificationMethod to the KeyAgreement list of the DID Document and returns the VerificationRelationship
// Returns nil if the VerificationMethod is already in the KeyAgreement list
func (id *DIDDocument) LinkKeyAgreement(vm *VerificationMethod) (*VerificationRelationship, bool) {
	for _, keyAgreement := range id.KeyAgreement {
		if keyAgreement.Reference == vm.Id {
			return nil, false
		}
	}

	vr := &VerificationRelationship{
		Reference:          vm.Id,
		Type:               KeyAgreementRelationshipName,
		VerificationMethod: vm,
		Owner:              id.Id,
	}
	id.KeyAgreement = append(id.KeyAgreement, vr)
	return vr, true
}

// GetWalletVerificationMethods returns all the VerificationMethods for the CapabilityInvocationRelationships
func (d *DIDDocument) GetWalletVerificationMethods() []*VerificationMethod {
	vms := make([]*VerificationMethod, 0)
	for _, relationship := range d.CapabilityInvocation {
		vms = append(vms, relationship.VerificationMethod)
	}
	return vms
}

// SDKAddress returns the address of the DIDDocument as an sdk.AccAddress
func (id *DIDDocument) SDKAddress() (sdk.AccAddress, error) {
	addr, err := sdk.AccAddressFromBech32(id.Address())
	if err != nil {
		return nil, err
	}
	return addr, nil
}

// ToCredentialDescriptor converts a VerificationMethod to a CredentialDescriptor if the VerificationMethod uses the `did:webauthn` method
// returns an error if the VerificationMethod does not use the `did:webauthn` method
func (vm *VerificationMethod) ToCredentialDescriptor() (protocol.CredentialDescriptor, error) {
	// extract the method from the id
	idParts := strings.Split(vm.Id, ":")
	if len(idParts) < 3 {
		return protocol.CredentialDescriptor{}, fmt.Errorf("malformed ID, expected at least 3 parts separated by colons")
	}
	method := idParts[1]
	credentialIDStr := idParts[2]
	if method != "webauthn" {
		return protocol.CredentialDescriptor{}, fmt.Errorf("verification method is not a webauthn method")
	}
	// extract CredentialID from the id
	credentialID, err := base64.RawURLEncoding.DecodeString(credentialIDStr)
	if err != nil {
		return protocol.CredentialDescriptor{}, fmt.Errorf("error decoding credential id: %w", err)
	}

	transport := make([]protocol.AuthenticatorTransport, 0)
	for _, t := range vm.Transports {
		transport = append(transport, protocol.AuthenticatorTransport(t))
	}

	return protocol.CredentialDescriptor{
		CredentialID:    protocol.URLEncodedBase64(credentialID),
		Type:            protocol.PublicKeyCredentialType,
		Transport:       transport,
		AttestationType: vm.AttestationType,
	}, nil
}
