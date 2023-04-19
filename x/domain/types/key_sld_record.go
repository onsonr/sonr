package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// SLDRecordKeyPrefix is the prefix to retrieve all SLDRecord
	SLDRecordKeyPrefix = "SLDRecord/value/"
)

// SLDRecordKey returns the store key to retrieve a SLDRecord from the index fields
func SLDRecordKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
