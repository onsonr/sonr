package types

import (
	"encoding/binary"
	"fmt"
	"strings"
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
	AlsoKnownAsPrefix       = "AlsoKnownAs/value/"
)

const (
	// ServiceKeyPrefix is the prefix to retrieve all DomainRecord
	AuthenticationKeyPrefix       = "Relationship/authentication/value/"
	AssertionKeyPrefix            = "Relationship/assertion/value/"
	CapabilityDelegationKeyPrefix = "Relationship/capability-delegation/value/"
	CapabilityInvocationKeyPrefix = "Relationship/capability-invocation/value/"
	KeyAgreementKeyPrefix         = "Relationship/key-agreement/value/"
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

// IdentificationKeyPrefix takes a did string splits to find the method and returns a prefix to retrieve all DidDocument
func IdentificationKeyPrefix(did string) (string, bool) {
	ptrs := strings.Split(did, ":")
	method := ptrs[1]
	params := DefaultParams()
	if ok := params.IsSupportedDidMethod(method); !ok {
		return "", false
	}
	return fmt.Sprintf("Identification/%s/value/", method), true
}

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

// RelationshipKeyPrefix takes the Relationship Name and returns a prefix to retrieve all Relationship
func RelationshipKeyPrefix(typeName string) string {
	return fmt.Sprintf("Relationship/%s/value/", typeName)
}
