package kss

import (
	"encoding/json"

	"github.com/di-dao/sonr/crypto"
	"github.com/di-dao/sonr/crypto/core/protocol"
)

// KssI is the interface for the keyshare set
type Set interface {
	BytesUsr() []byte
	BytesVal() []byte
	PublicKey() crypto.PublicKey
	Usr() User
	Val() Val
}

// keyshares is the set of keyshares for the protocol
type keyshares struct {
	val Val
	usr User

	valBz []byte
	usrBz []byte
}

// Usr returns the user keyshare
func (ks *keyshares) Usr() User {
	return ks.usr
}

// Val returns the validator keyshare
func (ks *keyshares) Val() Val {
	return ks.val
}

// BytesUsr returns the user keyshare as bytes
func (ks *keyshares) BytesUsr() []byte {
	return ks.usrBz
}

// BytesVal returns the validator keyshare as bytes
func (ks *keyshares) BytesVal() []byte {
	return ks.valBz
}

// PublicKey returns the public key for the keyshare set
func (ks *keyshares) PublicKey() crypto.PublicKey {
	return ks.val.PublicKey()
}

// NewKeyshareSet creates a new KeyshareSet
func NewKeyshareSet(aliceResult *protocol.Message, bobResult *protocol.Message) (Set, error) {
	valBz, err := json.Marshal(aliceResult)
	if err != nil {
		return nil, err
	}
	usrBz, err := json.Marshal(bobResult)
	if err != nil {
		return nil, err
	}
	return &keyshares{
		val:   createValidatorKeyshare(aliceResult),
		usr:   createUserKeyshare(bobResult),
		valBz: valBz,
		usrBz: usrBz,
	}, nil
}
