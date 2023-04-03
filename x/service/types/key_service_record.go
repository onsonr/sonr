package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ServiceRecordKeyPrefix is the prefix to retrieve all ServiceRecord
	ServiceRecordKeyPrefix = "ServiceRecord/value/"
)

// ServiceRecordKey returns the store key to retrieve a ServiceRecord from the index fields
func ServiceRecordKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
