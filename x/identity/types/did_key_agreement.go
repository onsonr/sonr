// Utility functions for DID Key Agreement - https://w3c.github.io/did-core/#key-agreement
// I.e. Verification Material for Key Exchange. Used for X25119 of Validator Node Key
package types

// AddKeyAgreement adds a VerificationMethod as KeyAgreement
// If the controller is not set, it will be set to the document's ID
func (d *DidDocument) AddKeyAgreement(v *VerificationMethod) {
	if v.Controller == "" {
		v.Controller = d.ID
	}
	d.VerificationMethod.Add(v)
	d.KeyAgreement.Add(v)
}

func (d *DidDocument) GetKeyAgreements() VerificationRelationships {
	return *d.KeyAgreement
}
