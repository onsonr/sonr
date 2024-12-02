package mpc

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"

	cometcrypto "github.com/cometbft/cometbft/crypto"
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
	Address() cometcrypto.Address
	Bytes() []byte
	DID() string
	VerifySignature(msg []byte, sig []byte) bool
	Equals(cometcrypto.PubKey) bool
	Type() string
}

type rootPublicKey struct {
	data []byte
	kind string
}

func (k rootPublicKey) Address() cometcrypto.Address {
	return cometcrypto.AddressHash(k.data)
}

func (k rootPublicKey) Bytes() []byte {
	return k.data
}

func (k rootPublicKey) DID() string {
	return fmt.Sprintf("did:sonr:%s", k.Address())
}

func (k rootPublicKey) VerifySignature(msg []byte, sig []byte) bool {
	ok, err := VerifySignature(k.data, msg, sig)
	if err != nil {
		return false
	}
	return ok
}

func (k rootPublicKey) Equals(other cometcrypto.PubKey) bool {
	return bytes.Equal(k.data, other.Bytes())
}

func (k rootPublicKey) Type() string {
	return k.kind
}

func createPublicKey(pk []byte, kind string) PublicKey {
	return rootPublicKey{
		data: pk,
		kind: kind,
	}
}
