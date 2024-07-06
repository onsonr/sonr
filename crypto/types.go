package crypto

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

// PublicKey defines the interface for all public keys
type PublicKey interface {
	Address() cryptotypes.Address
	Bytes() []byte
	String() string
	VerifySignature(msg []byte, sig []byte) bool
	Equals(other cryptotypes.PubKey) bool
	Type() string
}
