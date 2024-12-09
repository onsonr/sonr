package mpc

import (
	"crypto/ecdsa"
	genericecdsa "crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/onsonr/sonr/crypto/core/curves"
	"github.com/onsonr/sonr/crypto/core/protocol"
	"github.com/onsonr/sonr/crypto/tecdsa/dklsv1"
	"golang.org/x/crypto/sha3"
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

// For DKG bob starts first. For refresh and sign, Alice starts first.
func RunProtocol(firstParty protocol.Iterator, secondParty protocol.Iterator) (error, error) {
	var (
		message *protocol.Message
		aErr    error
		bErr    error
	)

	for aErr != protocol.ErrProtocolFinished || bErr != protocol.ErrProtocolFinished {
		// Crank each protocol forward one iteration
		message, bErr = firstParty.Next(message)
		if bErr != nil && bErr != protocol.ErrProtocolFinished {
			return nil, bErr
		}

		message, aErr = secondParty.Next(message)
		if aErr != nil && aErr != protocol.ErrProtocolFinished {
			return aErr, nil
		}
	}
	return aErr, bErr
}

// ComputeEcPoint builds an elliptic curve point from a compressed byte slice
func ComputeEcPoint(pubKey []byte) (*curves.EcPoint, error) {
	crv := curves.K256()
	x := new(big.Int).SetBytes(pubKey[1:33])
	y := new(big.Int).SetBytes(pubKey[33:])
	ecCurve, err := crv.ToEllipticCurve()
	if err != nil {
		return nil, fmt.Errorf("error converting curve: %v", err)
	}
	return &curves.EcPoint{X: x, Y: y, Curve: ecCurve}, nil
}

func ComputeEcdsaPublicKey(pubKey []byte) (*genericecdsa.PublicKey, error) {
	pk, err := ComputeEcPoint(pubKey)
	if err != nil {
		return nil, err
	}
	return &genericecdsa.PublicKey{
		Curve: pk.Curve,
		X:     pk.X,
		Y:     pk.Y,
	}, nil
}

// VerifySignature verifies the signature of a message
func VerifySignature(pk []byte, msg []byte, sig []byte) (bool, error) {
	pp, err := ComputeEcPoint(pk)
	if err != nil {
		return false, err
	}
	sigEd, err := DeserializeSignature(sig)
	if err != nil {
		return false, err
	}
	hash := sha3.New256()
	_, err = hash.Write(msg)
	if err != nil {
		return false, err
	}
	digest := hash.Sum(nil)
	return curves.VerifyEcdsa(pp, digest[:], sigEd), nil
}

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
