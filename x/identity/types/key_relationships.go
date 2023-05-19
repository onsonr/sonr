package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ServiceKeyPrefix is the prefix to retrieve all DomainRecord
	AuthenticationKeyPrefix       = "Relationship/authentication/value/"
	AssertionKeyPrefix            = "Relationship/assertion/value/"
	CapabilityDelegationKeyPrefix = "Relationship/capability-delegation/value/"
	CapabilityInvocationKeyPrefix = "Relationship/capability-invocation/value/"
	KeyAgreementKeyPrefix         = "Relationship/key-agreement/value/"
)

// ServiceKey returns the store key to retrieve a DomainRecord from the index fields
func RelationshipKey(
	did string,
) []byte {
	var key []byte

	didBytes := []byte(did)
	key = append(key, didBytes...)
	key = append(key, []byte("/")...)
	return key
}
