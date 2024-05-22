package kss

import (
	"github.com/di-dao/core/crypto"
	"github.com/di-dao/core/crypto/core/protocol"
)

// KssI is the interface for the keyshare set
type Set interface {
	Val() Val
	Usr() User
	PublicKey() crypto.PublicKey
}

// keyshares is the set of keyshares for the protocol
type keyshares struct {
	val Val
	usr User
}

// Usr returns the user keyshare
func (ks *keyshares) Usr() User {
	return ks.usr
}

// Val returns the validator keyshare
func (ks *keyshares) Val() Val {
	return ks.val
}

// PublicKey returns the public key for the keyshare set
func (ks *keyshares) PublicKey() crypto.PublicKey {
	return ks.val.PublicKey()
}

// NewKeyshareSet creates a new KeyshareSet
func NewKeyshareSet(aliceResult *protocol.Message, bobResult *protocol.Message) Set {
	return &keyshares{
		val: createValidatorKeyshare(aliceResult),
		usr: createUserKeyshare(bobResult),
	}
}
