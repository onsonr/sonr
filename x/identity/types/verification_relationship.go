package types

import (
	"strings"
)

// NewAuthenticationRelationship creates a new authentication relationship from the given verification method.
func NewAuthenticationRelationship(method *VerificationMethod) VerificationRelationship {
	owner, _ := GetOwnerFromDID(method.Id)
	return VerificationRelationship{VerificationMethod: method, Reference: method.Id, Owner: owner, Type: AuthenticationRelationshipName}
}

// NewAssertionRelationship creates a new assertion relationship from the given verification method.
func NewAssertionRelationship(method *VerificationMethod) VerificationRelationship {
	owner, _ := GetOwnerFromDID(method.Id)
	return VerificationRelationship{VerificationMethod: method, Reference: method.Id, Owner: owner, Type: AssertionRelationshipName}
}

// NewKeyAgreementRelationship creates a new key agreement relationship from the given verification method.
func NewKeyAgreementRelationship(method *VerificationMethod) VerificationRelationship {
	owner, _ := GetOwnerFromDID(method.Id)
	return VerificationRelationship{VerificationMethod: method, Reference: method.Id, Owner: owner, Type: KeyAgreementRelationshipName}
}

// NewCapabilityDelegationRelationship creates a new capability delegation relationship from the given verification method.
func NewCapabilityDelegationRelationship(method *VerificationMethod) VerificationRelationship {
	owner, _ := GetOwnerFromDID(method.Id)
	return VerificationRelationship{VerificationMethod: method, Reference: method.Id, Owner: owner, Type: CapabilityDelegationRelationshipName}
}

// NewCapabilityInvocationRelationship creates a new capability invocation relationship from the given verification method.
func NewCapabilityInvocationRelationship(method *VerificationMethod) VerificationRelationship {
	owner, _ := GetOwnerFromDID(method.Id)
	return VerificationRelationship{VerificationMethod: method, Reference: method.Id, Owner: owner, Type: CapabilityInvocationRelationshipName}
}

// GetOwnerFromDID returns the owner of the given DID.
func GetOwnerFromDID(did string) (string, bool) {
	ptrs := strings.Split(did, ":")
	if len(ptrs) != 3 {
		return "", false
	}
	if ptrs[1] != "sonr" {
		return "", false
	}
	return ptrs[2], true
}
