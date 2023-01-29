// Utility functions for DID Capability Delegation - https://w3c.github.io/did-core/#capability-delegation
// I.e. Verification Material for IPFS Node which stores MPC Configurations
package types

// AddCapabilityDelegation adds a VerificationMethod as CapabilityDelegation
// If the controller is not set, it will be set to the document's ID
func (d *DidDocument) AddCapabilityDelegation(v *VerificationMethod) {
	if v.Controller == "" {
		v.Controller = d.Id
	}
	d.VerificationMethod = append(d.VerificationMethod, v)
	d.CapabilityDelegation = append(d.CapabilityDelegation, &VerificationRelationship{VerificationMethod: v, Reference: d.Id})
}

// AddCapabilityInvocation adds a VerificationMethod as CapabilityInvocation
// If the controller is not set, it will be set to the document's ID
func (d *DidDocument) AddCapabilityInvocation(v *VerificationMethod) {
	if v.Controller == "" {
		v.Controller = d.Id
	}
	d.VerificationMethod = append(d.VerificationMethod, v)
	d.CapabilityInvocation = append(d.CapabilityDelegation, &VerificationRelationship{VerificationMethod: v, Reference: d.Id})
}
