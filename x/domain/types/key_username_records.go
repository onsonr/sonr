package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UsernameRecordsKeyPrefix is the prefix to retrieve all UsernameRecord
	UsernameRecordsKeyPrefix = "UsernameRecord/value/"
)

// UsernameRecordsKey returns the store key to retrieve a UsernameRecord from the index fields
func UsernameRecordsKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
