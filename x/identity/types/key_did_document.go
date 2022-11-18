package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DidDocumentKeyPrefix is the prefix to retrieve all DidDocument
	DidDocumentKeyPrefix = "DidDocument/value/"
)

// DidDocumentKey returns the store key to retrieve a DidDocument from the index fields
func DidDocumentKey(
	did string,
) []byte {
	var key []byte

	didBytes := []byte(did)
	key = append(key, didBytes...)
	key = append(key, []byte("/")...)

	return key
}
