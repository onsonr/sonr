package types

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"math/big"

	secp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/sonrhq/kryptology/pkg/accumulator"
	"github.com/sonrhq/kryptology/pkg/core/curves"
	pv1 "github.com/sonrhq/kryptology/pkg/core/protocol"
	"golang.org/x/crypto/hkdf"
)

type ZKAccumulator = accumulator.Accumulator
type Secp256k1PublicKey = *secp256k1.PubKey
type PubKey = cryptotypes.PubKey
type ZKEphemeralKey []byte
type KeyshareRole string

const (
	// KeyshareRolePublic is the default role for the alice dkg
	KeyshareRolePublic KeyshareRole = "alice"

	// KeyshareRoleUser is the role for an encrypted keyshare for a user
	KeyshareRoleUser KeyshareRole = "bob"
)

func (ksr KeyshareRole) isAlice() bool {
	return ksr == KeyshareRolePublic
}

func (ksr KeyshareRole) isBob() bool {
	return ksr == KeyshareRoleUser
}

func SerializeECDSASecp256k1Signature(sig *curves.EcdsaSignature) ([]byte, error) {
	rBytes := sig.R.Bytes()
	sBytes := sig.S.Bytes()

	sigBytes := make([]byte, 66) // V (1 byte) + R (32 bytes) + S (32 bytes)
	sigBytes[0] = byte(sig.V)
	copy(sigBytes[33-len(rBytes):33], rBytes)
	copy(sigBytes[66-len(sBytes):66], sBytes)
	return sigBytes, nil
}

func DeserializeECDSASecp256k1Signature(sigBytes []byte) (*curves.EcdsaSignature, error) {
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

// For DKG bob starts first. For refresh and sign, Alice starts first.
func RunIteratedProtocol(firstParty pv1.Iterator, secondParty pv1.Iterator) (error, error) {
	var (
		message *pv1.Message
		aErr    error
		bErr    error
	)
	for aErr != pv1.ErrProtocolFinished || bErr != pv1.ErrProtocolFinished {
		// Crank each protocol forward one iteration
		message, bErr = firstParty.Next(message)
		if bErr != nil && bErr != pv1.ErrProtocolFinished {
			return nil, bErr
		}

		message, aErr = secondParty.Next(message)
		if aErr != nil && aErr != pv1.ErrProtocolFinished {
			return aErr, nil
		}
	}
	return aErr, bErr
}

func GetSecretFromSecp256k1(publicKey PubKey) (*accumulator.SecretKey, error) {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	var seed [32]byte
	hkdfExtractor := hkdf.New(sha256.New, publicKey.Bytes()[:], nil, nil)

	// Use the HKDF extractor to derive a 32-byte seed
	if _, err := io.ReadFull(hkdfExtractor, seed[:]); err != nil {
		return nil, err
	}

	key, err := new(accumulator.SecretKey).New(curve, seed[:])
	if err != nil {
		return nil, err
	}
	return key, nil
}

func GetEphemeralFromSecp256k1(publicKey PubKey) (ZKEphemeralKey, error) {
	secretKey, err := GetSecretFromSecp256k1(publicKey)
	if err != nil {
		return nil, err
	}
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	pk, err := secretKey.GetPublicKey(curve)
	if err != nil {
		return nil, err
	}
	pkbz, err := pk.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return ZKEphemeralKey(pkbz), nil
}

func GetEncryptionKeyFromSecp256k1(publicKey PubKey) ([]byte, error) {
	ephemeral, err := GetEphemeralFromSecp256k1(publicKey)
	if err != nil {
		return nil, err
	}
	// Use the HKDF extractor to derive a 32-byte seed
	hkdf := hkdf.New(sha256.New, ephemeral, nil, nil)
	aesKey := make([]byte, 32) // 32 bytes for AES-256
	if _, err := io.ReadFull(hkdf, aesKey); err != nil {
		return nil, fmt.Errorf("error generating AES key: %v", err)
	}
	return aesKey, nil
}

func StringToZkElement(str string) accumulator.Element {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	return curve.Scalar.Hash([]byte(str))
}

func StringListToZkElements(strs ...string) []accumulator.Element {
	elements := make([]accumulator.Element, len(strs))
	for i, str := range strs {
		elements[i] = StringToZkElement(str)
	}
	return elements
}
