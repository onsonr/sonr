package types

import (
	"strings"

	"github.com/sonrhq/core/types/crypto"
	vaulttypes "github.com/sonrhq/core/x/vault/types"
)

// NewDIDDocument creates a new DIDDocument from an Identification and optional VerificationRelationships
func NewDIDDocument(identification *Identification, relationships ...*VerificationRelationship) *DIDDocument {
	params := DefaultParams()
	didDoc := &DIDDocument{
		Id:                   identification.Id,
		Context:              []string{params.AccountDidMethodContext, params.DidBaseContext},
		Authentication:       make([]*VerificationRelationship, 0),
		AssertionMethod:      make([]*VerificationRelationship, 0),
		CapabilityInvocation: make([]*VerificationRelationship, 0),
		CapabilityDelegation: make([]*VerificationRelationship, 0),
		Controller:           []string{identification.Owner},
		AlsoKnownAs:          []string{identification.PrimaryAlias},
		Metadata:             "",
	}

	for _, relationship := range relationships {
		switch relationship.Type {
		case AuthenticationRelationshipName:
			didDoc.Authentication = append(didDoc.Authentication, relationship)
		case AssertionRelationshipName:
			didDoc.AssertionMethod = append(didDoc.AssertionMethod, relationship)
		case CapabilityInvocationRelationshipName:
			didDoc.CapabilityInvocation = append(didDoc.CapabilityInvocation, relationship)
		case CapabilityDelegationRelationshipName:
			didDoc.CapabilityDelegation = append(didDoc.CapabilityDelegation, relationship)
		case KeyAgreementRelationshipName:
			didDoc.KeyAgreement = append(didDoc.KeyAgreement, relationship)
		default:
			continue
		}
	}
	return didDoc
}

// ToIdentification converts all the VerificationRelationships in the DIDDocument to an Identification string arrays for ID
func (d *DIDDocument) ToIdentification() *Identification {
	id := &Identification{
		Id:                   d.Id,
		Owner:                d.Controller[0],
		PrimaryAlias:         d.AlsoKnownAs[0],
		Authentication:       make([]string, 0),
		AssertionMethod:      make([]string, 0),
		CapabilityInvocation: make([]string, 0),
		CapabilityDelegation: make([]string, 0),
		KeyAgreement:         make([]string, 0),
	}

	for _, relationship := range d.Authentication {
		id.Authentication = append(id.Authentication, relationship.Reference)
	}
	for _, relationship := range d.AssertionMethod {
		id.AssertionMethod = append(id.AssertionMethod, relationship.Reference)
	}
	for _, relationship := range d.CapabilityInvocation {
		id.CapabilityInvocation = append(id.CapabilityInvocation, relationship.Reference)
	}
	for _, relationship := range d.CapabilityDelegation {
		id.CapabilityDelegation = append(id.CapabilityDelegation, relationship.Reference)
	}
	for _, relationship := range d.KeyAgreement {
		id.KeyAgreement = append(id.KeyAgreement, relationship.Reference)
	}
	return id
}

// SearchRelationshipsByCoinType returns all verification relationships of a given the query options
func (d *DIDDocument) SearchRelationshipsByCoinType(coinType crypto.CoinType) []*VerificationRelationship {
	method := coinType.DidMethod()
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

// AddCapabilityInvocationForAccount adds a new VerificationRelationship to the DIDDocument
func (d *DIDDocument) AddCapabilityInvocationForAccount(account vaulttypes.Account) *VerificationRelationship {
	vm := NewVerificationMethodFromVaultAccount(account, d.Id)
	vr := NewCapabilityInvocationRelationship(vm)
	d.CapabilityInvocation = append(d.CapabilityInvocation, &vr)
	return &vr
}
