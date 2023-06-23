package types

import (
	"encoding/binary"
	"fmt"
)

var _ binary.ByteOrder

const (
	// ModuleName defines the module name
	ModuleName = "identity"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_identity"

	ClaimableWalletKey      = "ClaimableWallet/value/"
	ClaimableWalletCountKey = "ClaimableWallet/count/"
)

const (
	// ServiceKeyPrefix is the prefix to retrieve all DomainRecord
	AuthenticationKeyPrefix       = "Relationship/authentication/value/"
	AssertionKeyPrefix            = "Relationship/assertion/value/"
	CapabilityDelegationKeyPrefix = "Relationship/capability-delegation/value/"
	CapabilityInvocationKeyPrefix = "Relationship/capability-invocation/value/"
	KeyAgreementKeyPrefix         = "Relationship/key-agreement/value/"
	IdentityKeyPrefix             = "Identification/did-doc/value/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// IdentificationKey returns the store key to retrieve a DidDocument from the index fields
func IdentificationKey(
	did string,
) []byte {
	var key []byte

	didBytes := []byte(did)
	key = append(key, didBytes...)
	key = append(key, []byte("/")...)

	return key
}



// RelationshipKeyPrefix takes the Relationship Name and returns a prefix to retrieve all Relationship
func RelationshipKeyPrefix(typeName string) string {
	return fmt.Sprintf("Relationship/%s/value/", typeName)
}
