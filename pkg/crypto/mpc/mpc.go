package mpc

import (
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/onsonr/sonr/pkg/crypto/core/curves"
	"github.com/onsonr/sonr/pkg/crypto/core/protocol"
	"github.com/onsonr/sonr/pkg/crypto/tecdsa/dklsv1"
)

// GenerateKeyshares generates a new MPC keyshare
func GenerateKeyshares() ([]Share, error) {
	curve := curves.K256()
	valKs := dklsv1.NewAliceDkg(curve, protocol.Version1)
	userKs := dklsv1.NewBobDkg(curve, protocol.Version1)
	aErr, bErr := RunProtocol(valKs, userKs)
	if aErr != protocol.ErrProtocolFinished {
		return nil, aErr
	}
	if bErr != protocol.ErrProtocolFinished {
		return nil, bErr
	}
	valRes, err := valKs.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	userRes, err := userKs.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	return NewKeyshareArray(valRes, userRes)
}

// RunSignProtocol runs the MPC signing protocol
func RunSignProtocol(signFuncVal SignFunc, signFuncUser SignFunc) (Signature, error) {
	aErr, bErr := RunProtocol(signFuncVal, signFuncUser)
	if aErr != protocol.ErrProtocolFinished {
		return nil, aErr
	}
	if bErr != protocol.ErrProtocolFinished {
		return nil, bErr
	}
	out, err := signFuncUser.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	return dklsv1.DecodeSignature(out)
}

// RunRefreshProtocol runs the MPC refresh protocol
func RunRefreshProtocol(refreshFuncVal RefreshFunc, refreshFuncUser RefreshFunc) ([]Share, error) {
	aErr, bErr := RunProtocol(refreshFuncVal, refreshFuncUser)
	if aErr != protocol.ErrProtocolFinished {
		return nil, aErr
	}
	if bErr != protocol.ErrProtocolFinished {
		return nil, bErr
	}
	valRefreshResult, err := refreshFuncVal.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	userRefreshResult, err := refreshFuncUser.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	return NewKeyshareArray(valRefreshResult, userRefreshResult)
}

// SerializeSecp256k1Signature serializes an ECDSA signature into a byte slice
func SerializeSignature(sig Signature) ([]byte, error) {
	rBytes := sig.R.Bytes()
	sBytes := sig.S.Bytes()

	sigBytes := make([]byte, 66) // V (1 byte) + R (32 bytes) + S (32 bytes)
	sigBytes[0] = byte(sig.V)
	copy(sigBytes[33-len(rBytes):33], rBytes)
	copy(sigBytes[66-len(sBytes):66], sBytes)
	return sigBytes, nil
}

// DeserializeSecp256k1Signature deserializes an ECDSA signature from a byte slice
func DeserializeSignature(sigBytes []byte) (Signature, error) {
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

// VerifyMPCSignature verifies an MPC signature
func VerifyMPCSignature(sig Signature, msg []byte, publicKey *ecdsa.PublicKey) bool {
	return ecdsa.Verify(publicKey, msg, sig.R, sig.S)
}
