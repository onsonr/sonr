package mpc

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"math/big"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/onsonr/sonr/internal/crypto/core/curves"
	"github.com/onsonr/sonr/internal/crypto/core/protocol"
	"github.com/onsonr/sonr/internal/crypto/tecdsa/dklsv1"
	"golang.org/x/crypto/sha3"
)

func checkIteratedErrors(aErr, bErr error) error {
	if aErr == protocol.ErrProtocolFinished && bErr == protocol.ErrProtocolFinished {
		return nil
	}
	if aErr != protocol.ErrProtocolFinished {
		return aErr
	}
	if bErr != protocol.ErrProtocolFinished {
		return bErr
	}
	return nil
}

func computeSonrAddr(pp Point) (string, error) {
	pk := pp.ToAffineCompressed()
	sonrAddr, err := bech32.ConvertAndEncode("idx", pk)
	if err != nil {
		return "", err
	}
	return sonrAddr, nil
}

func hashKey(key []byte) []byte {
	hash := sha3.New256()
	hash.Write(key)
	return hash.Sum(nil)[:32] // Use first 32 bytes of hash
}

func decryptKeyshare(msg []byte, key []byte, nonce []byte) ([]byte, error) {
	hashedKey := hashKey(key)
	block, err := aes.NewCipher(hashedKey)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	plaintext, err := aesgcm.Open(nil, nonce, msg, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func encryptKeyshare(msg Message, key []byte, nonce []byte) ([]byte, error) {
	hashedKey := hashKey(key)
	msgBytes, err := protocol.EncodeMessage(msg)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(hashedKey)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	ciphertext := aesgcm.Seal(nil, nonce, []byte(msgBytes), nil)
	return ciphertext, nil
}

func getAliceOut(msg *protocol.Message) (AliceOut, error) {
	return dklsv1.DecodeAliceDkgResult(msg)
}

func getAlicePubPoint(msg *protocol.Message) (Point, error) {
	out, err := dklsv1.DecodeAliceDkgResult(msg)
	if err != nil {
		return nil, err
	}
	return out.PublicKey, nil
}

func getBobOut(msg *protocol.Message) (BobOut, error) {
	return dklsv1.DecodeBobDkgResult(msg)
}

func getBobPubPoint(msg *protocol.Message) (Point, error) {
	out, err := dklsv1.DecodeBobDkgResult(msg)
	if err != nil {
		return nil, err
	}
	return out.PublicKey, nil
}

// getEcdsaPoint builds an elliptic curve point from a compressed byte slice
func getEcdsaPoint(pubKey []byte) (*curves.EcPoint, error) {
	crv := curves.K256()
	x := new(big.Int).SetBytes(pubKey[1:33])
	y := new(big.Int).SetBytes(pubKey[33:])
	ecCurve, err := crv.ToEllipticCurve()
	if err != nil {
		return nil, fmt.Errorf("error converting curve: %v", err)
	}
	return &curves.EcPoint{X: x, Y: y, Curve: ecCurve}, nil
}

func serializeSignature(sig *curves.EcdsaSignature) ([]byte, error) {
	if sig == nil {
		return nil, errors.New("nil signature")
	}

	rBytes := sig.R.Bytes()
	sBytes := sig.S.Bytes()

	// Ensure both components are 32 bytes
	rPadded := make([]byte, 32)
	sPadded := make([]byte, 32)
	copy(rPadded[32-len(rBytes):], rBytes)
	copy(sPadded[32-len(sBytes):], sBytes)

	// Concatenate R and S
	result := make([]byte, 64)
	copy(result[0:32], rPadded)
	copy(result[32:64], sPadded)

	return result, nil
}

func deserializeSignature(sigBytes []byte) (*curves.EcdsaSignature, error) {
	if len(sigBytes) != 64 {
		return nil, fmt.Errorf("invalid signature length: expected 64 bytes, got %d", len(sigBytes))
	}

	r := new(big.Int).SetBytes(sigBytes[:32])
	s := new(big.Int).SetBytes(sigBytes[32:])

	return &curves.EcdsaSignature{
		R: r,
		S: s,
	}, nil
}

func userSignFunc(k *keyEnclave, bz []byte) (SignFunc, error) {
	curve := curves.K256()
	return dklsv1.NewBobSign(curve, sha3.New256(), bz, k.UserShare, protocol.Version1)
}

func userRefreshFunc(k *keyEnclave) (RefreshFunc, error) {
	curve := curves.K256()
	return dklsv1.NewBobRefresh(curve, k.UserShare, protocol.Version1)
}

func valSignFunc(k *keyEnclave, bz []byte) (SignFunc, error) {
	curve := curves.K256()
	return dklsv1.NewAliceSign(curve, sha3.New256(), bz, k.ValShare, protocol.Version1)
}

func valRefreshFunc(k *keyEnclave) (RefreshFunc, error) {
	curve := curves.K256()
	return dklsv1.NewAliceRefresh(curve, k.ValShare, protocol.Version1)
}
