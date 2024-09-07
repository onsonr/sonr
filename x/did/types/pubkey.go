package types

import (
	fmt "fmt"
	"math/big"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/onsonr/crypto/core/curves"
	"github.com/onsonr/crypto/signatures/ecdsa"
	"golang.org/x/crypto/sha3"
)

// NewEthPublicKey returns a new ethereum public key
func NewPublicKey(data []byte, keyInfo *KeyInfo) (*PubKey, error) {
	encKey, err := keyInfo.Encoding.EncodeRaw(data)
	if err != nil {
		return nil, err
	}

	return &PubKey{
		Raw:       encKey,
		Role:      keyInfo.Role,
		Encoding:  keyInfo.Encoding,
		Algorithm: keyInfo.Algorithm,
		Curve:     keyInfo.Curve,
		KeyType:   keyInfo.Type,
	}, nil
}

// Address returns the address of the public key
func (k *PubKey) Address() cryptotypes.Address {
	return nil
}

// Bytes returns the raw bytes of the public key
func (k *PubKey) Bytes() []byte {
	bz, _ := k.GetEncoding().DecodeRaw(k.GetRaw())
	return bz
}

// Clone returns a copy of the public key
func (k *PubKey) Clone() cryptotypes.PubKey {
	return &PubKey{
		Raw:       k.GetRaw(),
		Role:      k.GetRole(),
		Encoding:  k.GetEncoding(),
		Algorithm: k.GetAlgorithm(),
		Curve:     k.GetCurve(),
		KeyType:   k.GetKeyType(),
	}
}

// VerifySignature verifies a signature over the given message
func (k *PubKey) VerifySignature(msg []byte, sig []byte) bool {
	pp, err := buildEcPoint(k.Bytes())
	if err != nil {
		return false
	}
	sigEd, err := ecdsa.DeserializeSecp256k1Signature(sig)
	if err != nil {
		return false
	}
	hash := sha3.New256()
	_, err = hash.Write(msg)
	if err != nil {
		return false
	}
	digest := hash.Sum(nil)
	return curves.VerifyEcdsa(pp, digest[:], sigEd)
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

// VerifySignature verifies the signature of a message
func VerifySignature(key []byte, msg []byte, sig []byte) bool {
	pp, err := buildEcPoint(key)
	if err != nil {
		return false
	}
	sigEd, err := ecdsa.DeserializeSecp256k1Signature(sig)
	if err != nil {
		return false
	}
	hash := sha3.New256()
	_, err = hash.Write(msg)
	if err != nil {
		return false
	}
	digest := hash.Sum(nil)
	return curves.VerifyEcdsa(pp, digest[:], sigEd)
}

// BuildEcPoint builds an elliptic curve point from a compressed byte slice
func buildEcPoint(pubKey []byte) (*curves.EcPoint, error) {
	crv := curves.K256()
	x := new(big.Int).SetBytes(pubKey[1:33])
	y := new(big.Int).SetBytes(pubKey[33:])
	ecCurve, err := crv.ToEllipticCurve()
	if err != nil {
		return nil, fmt.Errorf("error converting curve: %v", err)
	}
	return &curves.EcPoint{X: x, Y: y, Curve: ecCurve}, nil
}
