package types

import (
	"github.com/sonr-io/sonr/pkg/crypto"
	"lukechampine.com/blake3"
)

const (
	// ModuleName defines the module name
	ModuleName = "domain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_domain"
)

const EmailMethod = "email"

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func EmailIndex(email string) string {
	hash := blake3.Sum256([]byte(email))
	val := crypto.Base64Encode(hash[:])
	return val
}
