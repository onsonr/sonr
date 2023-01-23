package types

import (
	"encoding/binary"
	"fmt"
)

var _ binary.ByteOrder

const (
	// DomainRecordKeyPrefix is the prefix to retrieve all DomainRecord
	DomainRecordKeyPrefix = "DomainRecord/value/"
)

// DomainRecordKey returns the store key to retrieve a DomainRecord from the index fields
func DomainRecordKey(
	domain string,
	tldIndex string,
) []byte {
	var key []byte

	indexBytes := []byte(fmt.Sprintf("%s.%s", domain, tldIndex))
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
