package mpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/ipfs/boxo/files"
	"github.com/ipfs/kubo/client/rpc"
	"github.com/onsonr/sonr/crypto/core/curves"
	"github.com/onsonr/sonr/crypto/core/protocol"
	"github.com/onsonr/sonr/crypto/tecdsa/dklsv1"
)

func addEnclaveIPFS(enclave *KeyEnclave, ipc *rpc.HttpApi) (Enclave, error) {
	jsonEnclave, err := json.Marshal(enclave)
	if err != nil {
		return nil, err
	}
	// Save enclave to IPFS
	cid, err := ipc.Unixfs().Add(context.Background(), files.NewBytesFile(jsonEnclave))
	if err != nil {
		return nil, err
	}
	enclave.VaultCID = cid.String()
	return enclave, nil
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

func computeSonrAddr(pp Point) (string, error) {
	pk := pp.ToAffineCompressed()
	sonrAddr, err := bech32.ConvertAndEncode("idx", pk)
	if err != nil {
		return "", err
	}
	return sonrAddr, nil
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

// SerializeSecp256k1Signature serializes an ECDSA signature into a byte slice
func serializeSignature(sig *curves.EcdsaSignature) ([]byte, error) {
	if sig == nil || sig.R == nil || sig.S == nil {
		return nil, errors.New("invalid signature: nil values")
	}

	rBytes := sig.R.Bytes()
	sBytes := sig.S.Bytes()

	// Ensure R and S are exactly 32 bytes
	rPadded := make([]byte, 32)
	sPadded := make([]byte, 32)
	copy(rPadded[32-len(rBytes):], rBytes)
	copy(sPadded[32-len(sBytes):], sBytes)

	// Combine V, R, and S
	sigBytes := make([]byte, 65) // V (1 byte) + R (32 bytes) + S (32 bytes)
	sigBytes[0] = byte(sig.V)
	copy(sigBytes[1:33], rPadded)
	copy(sigBytes[33:], sPadded)
	return sigBytes, nil
}

// DeserializeSecp256k1Signature deserializes an ECDSA signature from a byte slice
func deserializeSignature(sigBytes []byte) (*curves.EcdsaSignature, error) {
	if len(sigBytes) != 65 {
		return nil, fmt.Errorf("malformed signature: expected 65 bytes, got %d", len(sigBytes))
	}

	r := new(big.Int).SetBytes(sigBytes[1:33])
	s := new(big.Int).SetBytes(sigBytes[33:65])
	
	if r.Sign() == 0 || s.Sign() == 0 {
		return nil, errors.New("invalid signature: R or S is zero")
	}

	return &curves.EcdsaSignature{
		V: int(sigBytes[0]),
		R: r,
		S: s,
	}, nil
}
