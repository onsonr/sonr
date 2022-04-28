package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// HowIsKeyPrefix is the prefix to retrieve all HowIs
	HowIsKeyPrefix = "HowIs/value/"
)

// HowIsKey returns the store key to retrieve a HowIs from the did field
func HowIsKey(did string) []byte {
	var key []byte

	didBytes := []byte(did)
	key = append(key, didBytes...)
	key = append(key, []byte("/")...)

	return key
}
