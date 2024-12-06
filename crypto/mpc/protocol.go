package mpc

import (
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/onsonr/sonr/crypto/core/curves"
	"github.com/onsonr/sonr/crypto/core/protocol"
	"github.com/onsonr/sonr/crypto/tecdsa/dklsv1"
)

// NewKeyshareSource generates a new MPC keyshare
func NewKeyset() (Keyset, error) {
	curve := curves.K256()
	valKs := dklsv1.NewAliceDkg(curve, protocol.Version1)
	userKs := dklsv1.NewBobDkg(curve, protocol.Version1)
	aErr, bErr := RunProtocol(userKs, valKs)
	if err := checkIteratedErrors(aErr, bErr); err != nil {
		return nil, err
	}
	valRes, err := valKs.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	valShare, err := NewValKeyshare(valRes)
	if err != nil {
		return nil, err
	}
	userRes, err := userKs.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	userShare, err := NewUserKeyshare(userRes)
	if err != nil {
		return nil, err
	}
	addr, err := computeSonrAddr(valShare.CompressedPublicKey())
	if err != nil {
		return nil, err
	}
	return keyset{val: valShare, user: userShare, addr: addr}, nil
}

// ExecuteSigning runs the MPC signing protocol
func ExecuteSigning(signFuncVal SignFunc, signFuncUser SignFunc) (Signature, error) {
	aErr, bErr := RunProtocol(signFuncVal, signFuncUser)
	if err := checkIteratedErrors(aErr, bErr); err != nil {
		return nil, err
	}
	out, err := signFuncUser.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	return dklsv1.DecodeSignature(out)
}

// ExecuteRefresh runs the MPC refresh protocol
func ExecuteRefresh(refreshFuncVal RefreshFunc, refreshFuncUser RefreshFunc) (Keyset, error) {
	aErr, bErr := RunProtocol(refreshFuncVal, refreshFuncUser)
	if err := checkIteratedErrors(aErr, bErr); err != nil {
		return nil, err
	}
	valRefreshResult, err := refreshFuncVal.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	valShare, err := NewValKeyshare(valRefreshResult)
	if err != nil {
		return nil, err
	}
	userRefreshResult, err := refreshFuncUser.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	userShare, err := NewUserKeyshare(userRefreshResult)
	if err != nil {
		return nil, err
	}
	addr, err := computeSonrAddr(valShare.CompressedPublicKey())
	if err != nil {
		return nil, err
	}
	return keyset{val: valShare, user: userShare, addr: addr}, nil
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
