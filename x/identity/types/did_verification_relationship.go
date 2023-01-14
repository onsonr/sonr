package types

import (
	"encoding/json"
	"errors"
	fmt "fmt"
	"strings"

	"github.com/sonr-hq/sonr/x/identity/types/internal/marshal"
)

// Count returns the number of VerificationRelationships in the slice
func (vmr *VerificationRelationships) Count() int {
	return len(vmr.GetData())
}

// Filter returns a new VerificationRelationships slice with all VerificationRelationships that match the filter
func (vmr *VerificationRelationships) Filter(filter func(*VerificationRelationship) bool) *VerificationRelationships {
	vrs := make([]*VerificationRelationship, 0)
	for _, r := range vmr.GetData() {
		if filter(r) {
			vrs = append(vrs, r)
		}
	}
	return &VerificationRelationships{Data: vrs}
}

// FindByID returns the first VerificationRelationship that matches with the id.
// For comparison both the ID of the embedded VerificationMethod and reference is used.
func (vmr *VerificationRelationships) FindByID(id string) *VerificationMethod {
	for _, r := range vmr.GetData() {
		if r.VerificationMethod != nil {
			if r.VerificationMethod.ID == id {
				return r.VerificationMethod
			}
		}
	}
	return nil
}

// FindByFragment returns the first VerificationRelationship that contans the fragment.
// For comparison both the ID of the embedded VerificationMethod and reference is used.
func (vmr *VerificationRelationships) FindByFragment(frag string) []*VerificationMethod {
	vms := make([]*VerificationMethod, 0)
	for _, r := range vmr.GetData() {
		if r.VerificationMethod != nil {
			if strings.Contains(r.VerificationMethod.ID, frag) {
				vms = append(vms, r.VerificationMethod)
			}
		}
	}
	return vms
}

// Remove removes a VerificationRelationship from the slice.
// If a VerificationRelationship was removed with the given DID, it will be returned
func (vmr VerificationRelationships) Remove(id string) *VerificationRelationship {
	var (
		filteredVMRels []*VerificationRelationship
		removedRel     *VerificationRelationship
	)
	for _, r := range vmr.GetData() {
		if r.Reference == id {
			filteredVMRels = append(filteredVMRels, r)
		} else {
			removedRel = r
		}
	}
	vmr.Data = filteredVMRels
	return removedRel
}

// Add adds a verificationMethod to a relationship collection.
// When the collection already contains the method it will not be added again.
func (vmr *VerificationRelationships) Add(vm *VerificationMethod) {
	for _, rel := range vmr.GetData() {
		if rel.Reference == vm.ID {
			return
		}
	}
	vmr.Data = append(vmr.GetData(), &VerificationRelationship{vm, vm.ID})
}

func (d *DidDocument) MarshalJSON() ([]byte, error) {
	type alias *DidDocument
	tmp := alias(d)
	if data, err := json.Marshal(tmp); err != nil {
		return nil, err
	} else {
		return marshal.NormalizeDocument(data, marshal.Unplural(contextKey), marshal.Unplural(controllerKey))
	}
}

func (d *DidDocument) UnmarshalJSON(b []byte) error {
	type alias DidDocument
	normalizedDoc, err := marshal.NormalizeDocument(b, pluralContext, marshal.Plural(controllerKey))
	if err != nil {
		return err
	}
	doc := alias{}
	err = json.Unmarshal(normalizedDoc, &doc)
	if err != nil {
		return err
	}
	*d = (DidDocument)(doc)

	const errMsg = "unable to resolve all '%s' references: %w"
	if err = resolveVerificationRelationships(d.Authentication.Data, d.VerificationMethod.Data); err != nil {
		return fmt.Errorf(errMsg, authenticationKey, err)
	}
	if err = resolveVerificationRelationships(d.AssertionMethod.Data, d.VerificationMethod.Data); err != nil {
		return fmt.Errorf(errMsg, assertionMethodKey, err)
	}
	if err = resolveVerificationRelationships(d.KeyAgreement.Data, d.VerificationMethod.Data); err != nil {
		return fmt.Errorf(errMsg, keyAgreementKey, err)
	}
	if err = resolveVerificationRelationships(d.CapabilityInvocation.Data, d.VerificationMethod.Data); err != nil {
		return fmt.Errorf(errMsg, capabilityInvocationKey, err)
	}
	if err = resolveVerificationRelationships(d.CapabilityDelegation.Data, d.VerificationMethod.Data); err != nil {
		return fmt.Errorf(errMsg, capabilityDelegationKey, err)
	}
	return nil
}

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
		if method.ID == reference {
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
			return fmt.Errorf("could not parse verificationRelation key relation DID: %w", err)
		}
	default:
		return errors.New("verificationRelation is invalid")
	}
	return nil
}
