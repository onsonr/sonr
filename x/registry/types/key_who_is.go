package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// WhoIsKeyPrefix is the prefix to retrieve all WhoIs
	WhoIsKeyPrefix = "WhoIs/value/"
)

// WhoIsKey returns the store key to retrieve a WhoIs from the did field
func WhoIsKey(did string) []byte {
	var key []byte

	didBytes := []byte(did)
	key = append(key, didBytes...)
	key = append(key, []byte("/")...)

	return key
}
