package types

import (
	"encoding/binary"
)

var _ binary.ByteOrder

const (
	// ServiceKeyPrefix is the prefix to retrieve all DomainRecord
	ServiceKeyPrefix = "Service/value/"
)

// ServiceKey returns the store key to retrieve a DomainRecord from the index fields
func ServiceKey(
	did string,
) []byte {
	var key []byte

	didBytes := []byte(did)
	key = append(key, didBytes...)
	key = append(key, []byte("/")...)
	return key
}
