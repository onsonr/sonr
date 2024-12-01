package mpc

import (
	"crypto/ecdsa"

	cosmoscrypto "github.com/cometbft/cometbft/crypto"
)

type PublicKeyType string

const (
	PublicKeyTypeRaw      PublicKeyType = "secp256k1"
	PublicKeyTypeCosmos   PublicKeyType = "cosmos"
	PublicKeyTypeBitcoin  PublicKeyType = "bitcoin"
	PublicKeyTypeEthereum PublicKeyType = "ethereum"
	PublicKeyTypeSonr     PublicKeyType = "sonr"
)

type ECDSAPublicKey *ecdsa.PublicKey

type PublicKey interface {
	// Bytes returns the byte representation of the public key
	Bytes() []byte

	// Equals checks if two public keys are equal
	Equals(other PublicKey) bool

	// ToCosmosPubKey converts the public key to a Cosmos compatible public key
	ToCosmosPubKey() cosmoscrypto.PubKey

	// DID returns the DID of the public key
	DID() string

	// ToIPFSPubKey converts the public key to an IPFS compatible public key
	// ToIPFSPubKey() ipfs.PubKey // TODO: Implement this
}
