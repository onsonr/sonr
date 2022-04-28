package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// WhichIsKeyPrefix is the prefix to retrieve all WhichIs
	WhichIsKeyPrefix = "WhichIs/value/"
)

// WhichIsKey returns the store key to retrieve a WhichIs from the did field
func WhichIsKey(
	did string,
) []byte {
	var key []byte

	didBytes := []byte(did)
	key = append(key, didBytes...)
	key = append(key, []byte("/")...)

	return key
}
