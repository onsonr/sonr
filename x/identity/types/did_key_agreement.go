// Utility functions for DID Key Agreement - https://w3c.github.io/did-core/#key-agreement
// I.e. Verification Material for Key Exchange. Used for X25119 of Validator Node Key
package types

// KeyAgreementCount returns the number of Assertion Methods
func (vm *DidDocument) KeyAgreementCount() int {
	return len(vm.KeyAgreement)
}

// AddKeyAgreement adds a VerificationMethod as KeyAgreement
// If the controller is not set, it will be set to the document's ID
func (d *DidDocument) AddKeyAgreement(v *VerificationMethod) {
	if v.Controller == "" {
		v.Controller = d.Id
	}
	d.VerificationMethod = append(d.VerificationMethod, v)
	d.KeyAgreement = append(d.KeyAgreement, &VerificationRelationship{
		Reference:          v.Id,
		VerificationMethod: v,
	})
}

func (d *DidDocument) GetKeyAgreements() []*VerificationRelationship {
	return d.KeyAgreement
}
