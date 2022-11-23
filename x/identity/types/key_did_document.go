package types

import (
	"encoding/binary"
	"strings"
)

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

// DidDocumentKeyToMethod returns the method of a DidDocument from the bytes of the did.
func DidDocumentKeyToMethod(
	didKey []byte,
) string {
	var method string
	did := string(didKey)
	method = strings.Split(did, ":")[1]
	return method
}
