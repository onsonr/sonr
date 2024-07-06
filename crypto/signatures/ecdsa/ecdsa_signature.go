package ecdsa

import (
	"errors"
	"math/big"

	"github.com/onsonr/hway/crypto/core/curves"
)

// SerializeSecp256k1Signature serializes an ECDSA signature into a byte slice
func SerializeSecp256k1Signature(sig *curves.EcdsaSignature) ([]byte, error) {
	rBytes := sig.R.Bytes()
	sBytes := sig.S.Bytes()

	sigBytes := make([]byte, 66) // V (1 byte) + R (32 bytes) + S (32 bytes)
	sigBytes[0] = byte(sig.V)
	copy(sigBytes[33-len(rBytes):33], rBytes)
	copy(sigBytes[66-len(sBytes):66], sBytes)
	return sigBytes, nil
}

// DeserializeSecp256k1Signature deserializes an ECDSA signature from a byte slice
func DeserializeSecp256k1Signature(sigBytes []byte) (*curves.EcdsaSignature, error) {
	if len(sigBytes) != 66 {
		return nil, errors.New("malformed signature: not the correct size")
	}
	sig := &curves.EcdsaSignature{
		V: int(sigBytes[0]),
		R: new(big.Int).SetBytes(sigBytes[1:33]),
		S: new(big.Int).SetBytes(sigBytes[33:66]),
	}
	return sig, nil
}
