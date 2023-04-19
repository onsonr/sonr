package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TLDRecordKeyPrefix is the prefix to retrieve all TLDRecord
	TLDRecordKeyPrefix = "TLDRecord/value/"
)

// TLDRecordKey returns the store key to retrieve a TLDRecord from the index fields
func TLDRecordKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
