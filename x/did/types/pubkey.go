package types

import (
	"bytes"
	"encoding/hex"
	fmt "fmt"
	"math/big"

	cmtcrypto "github.com/cometbft/cometbft/crypto"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"golang.org/x/crypto/sha3"

	"github.com/di-dao/core/crypto/core/curves"
	"github.com/di-dao/core/crypto/signatures/ecdsa"
)

func (k *PublicKey) Address() cryptotypes.Address {
	return cmtcrypto.AddressHash(k.Key)
}

func (k *PublicKey) Bytes() []byte {
	return k.Key
}

func (k *PublicKey) String() string {
	return hex.EncodeToString(k.Key)
}

func (k *PublicKey) VerifySignature(msg []byte, sig []byte) bool {
	pp, err := BuildEcPoint(k.Key)
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

// Equals returns true if the public keys are equal
func (k *PublicKey) Equals(other cryptotypes.PubKey) bool {
	return bytes.Equal(k.Key, other.Bytes())
}

// Type returns the public key type
func (k *PublicKey) Type() string {
	return k.KeyType
}

// BuildEcPoint builds an elliptic curve point from a compressed byte slice
func BuildEcPoint(pubKey []byte) (*curves.EcPoint, error) {
	crv := curves.P256()
	x := new(big.Int).SetBytes(pubKey[1:33])
	y := new(big.Int).SetBytes(pubKey[33:])
	ecCurve, err := crv.ToEllipticCurve()
	if err != nil {
		return nil, fmt.Errorf("error converting curve: %v", err)
	}
	return &curves.EcPoint{X: x, Y: y, Curve: ecCurve}, nil
}
