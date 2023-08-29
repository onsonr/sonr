package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DIDDocumentKeyPrefix is the prefix to retrieve all DIDDocument
	DIDDocumentKeyPrefix = "DIDDocument/value/"
)

// DIDDocumentKey returns the store key to retrieve a DIDDocument from the index fields
func DIDDocumentKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

func ControllerAccountKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

func EscrowAccountKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

