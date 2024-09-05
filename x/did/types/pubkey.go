package types

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

// NewEthPublicKey returns a new ethereum public key
func NewPublicKey(data []byte, keyInfo *KeyInfo) (*PubKey, error) {
	return &PubKey{
		Role: keyInfo.Role,
	}, nil
}

// Address returns the address of the public key
func (k *PubKey) Address() cryptotypes.Address {
	return nil
}

// Bytes returns the raw bytes of the public key
func (k *PubKey) Bytes() []byte {
	return k.GetRaw()
}

// VerifySignature verifies a signature over the given message
func (k *PubKey) VerifySignature(msg []byte, sig []byte) bool {
	return false
}

// Equals returns true if two public keys are equal
func (k *PubKey) Equals(k2 cryptotypes.PubKey) bool {
	if k == nil && k2 == nil {
		return true
	}
	return false
}

// Type returns the type of the public key
func (k *PubKey) Type() string {
	return ""
}
