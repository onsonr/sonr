package kss

import "github.com/di-dao/sonr/crypto"

// KssI is the interface for the keyshare set
type EncryptedSet interface {
	Decrypt(key []byte) (Set, error)
	PublicKey() crypto.PublicKey
}
