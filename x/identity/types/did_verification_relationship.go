package types

import (
	"encoding/json"
	"errors"
	fmt "fmt"
	"strings"
)

func resolveVerificationRelationships(relationships []*VerificationRelationship, methods []*VerificationMethod) error {
	for i, relationship := range relationships {
		if relationship.Reference != "" {
			continue
		}
		if resolved := resolveVerificationRelationship(relationship.Reference, methods); resolved == nil {
			return fmt.Errorf("unable to resolve %s: %s", verificationMethodKey, relationship.Reference)
		} else {
			relationships[i] = resolved
			relationships[i].Reference = relationship.Reference
		}
	}
	return nil
}

func resolveVerificationRelationship(reference string, methods []*VerificationMethod) *VerificationRelationship {
	for _, method := range methods {
		if method.Id == reference {
			return &VerificationRelationship{VerificationMethod: method}
		}
	}
	return nil
}

func (v VerificationRelationship) MarshalJSON() ([]byte, error) {
	if v.Reference != "" {
		return json.Marshal(*v.VerificationMethod)
	} else {
		return json.Marshal(v.Reference)
	}
}

func (v *VerificationRelationship) UnmarshalJSON(b []byte) error {
	// try to figure out if the item is an object of a string
	type Alias VerificationRelationship
	switch b[0] {
	case '{':
		tmp := Alias{VerificationMethod: &VerificationMethod{}}
		err := json.Unmarshal(b, &tmp)
		if err != nil {
			return fmt.Errorf("could not parse verificationRelation method: %w", err)
		}
		*v = (VerificationRelationship)(tmp)
	case '"':
		err := json.Unmarshal(b, &v.Reference)
		if err != nil {
			return fmt.Errorf("could not parse verificationRelation key relation DId:%w", err)
		}
	default:
		return errors.New("verificationRelation is invalid")
	}
	return nil
}

// MatchesBlockchainAddress returns true if the verification relationship is a blockchain account and matches the given address.
func (v *VerificationRelationship) MatchesBlockchainAddress(addr string) bool {
	if v.VerificationMethod == nil || !v.VerificationMethod.IsBlockchainAccount() {
		return false
	}
	return strings.EqualFold(v.VerificationMethod.BlockchainAccountId, addr)
}

func (d *DidDocument) ResolveRelationships(vms []VerificationRelationship) *ResolvedDidDocument {
	resolved := &ResolvedDidDocument{
		Id:                   d.Id,
		Context:              d.Context,
		Controller:           d.Controller,
		AlsoKnownAs:          d.AlsoKnownAs,
		VerificationMethod:   d.VerificationMethod,
		Authentication:       []*VerificationRelationship{},
		AssertionMethod:      []*VerificationRelationship{},
		CapabilityDelegation: []*VerificationRelationship{},
		CapabilityInvocation: []*VerificationRelationship{},
		KeyAgreement:         []*VerificationRelationship{},
		Service:              []*Service{},
		Metadata:             d.Metadata,
	}
	return resolved.AddVerificationRelationship(vms)
}

func (d *DidDocument) ResolveMethods(vms []VerificationMethod) *ResolvedDidDocument {
	resolved := &ResolvedDidDocument{
		Id:                   d.Id,
		Context:              d.Context,
		Controller:           d.Controller,
		AlsoKnownAs:          d.AlsoKnownAs,
		VerificationMethod:   d.VerificationMethod,
		Authentication:       []*VerificationRelationship{},
		AssertionMethod:      []*VerificationRelationship{},
		CapabilityDelegation: []*VerificationRelationship{},
		CapabilityInvocation: []*VerificationRelationship{},
		KeyAgreement:         []*VerificationRelationship{},
		Service:              []*Service{},
		Metadata:             d.Metadata,
	}

	// Iterate through VerificationMethod and create a VerificationRelationship for each
	vrs := []VerificationRelationship{}
	for _, vm := range vms {
		vr := VerificationRelationship{
			Reference:          vm.Id,
			VerificationMethod: &vm,
		}
		vrs = append(vrs, vr)
	}
	return resolved.AddVerificationRelationship(vrs)
}

func (r *ResolvedDidDocument) AddVerificationRelationship(vms []VerificationRelationship) *ResolvedDidDocument {
	for _, vm := range vms {
		switch {
		case r.containsRelationship(r.Authentication, vm.Reference):
			r.Authentication = append(r.Authentication, &vm)
		case r.containsRelationship(r.AssertionMethod, vm.Reference):
			r.AssertionMethod = append(r.AssertionMethod, &vm)
		case r.containsRelationship(r.CapabilityDelegation, vm.Reference):
			r.CapabilityDelegation = append(r.CapabilityDelegation, &vm)
		case r.containsRelationship(r.CapabilityInvocation, vm.Reference):
			r.CapabilityInvocation = append(r.CapabilityInvocation, &vm)
		case r.containsRelationship(r.KeyAgreement, vm.Reference):
			r.KeyAgreement = append(r.KeyAgreement, &vm)
		}
	}
	return r
}

func (r *ResolvedDidDocument) containsRelationship(relations []*VerificationRelationship, ref string) bool {
	for _, relation := range relations {
		if relation.Reference == ref {
			return true
		}
	}
	return false
}

// Method returns the DID method of the document
func (d *ResolvedDidDocument) DIDMethod() string {
	return strings.Split(d.Id, ":")[1]
}

// Identifier returns the DID identifier of the document
func (d *ResolvedDidDocument) DIDIdentifier() string {
	return strings.Split(d.Id, ":")[2]
}

// Fragment returns the DID fragment of the document
func (d *ResolvedDidDocument) DIDFragment() string {
	return strings.Split(d.Id, "#")[1]
}
