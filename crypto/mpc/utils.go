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
