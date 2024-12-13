package mpc

import (
	"context"
	"encoding/hex"
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
	"golang.org/x/crypto/sha3"
)

func addEnclaveIPFS(enclave KeyEnclave, ipc *rpc.HttpApi) (KeyEnclave, error) {
	jsonEnclave, err := json.Marshal(enclave)
	if err != nil {
		return nil, err
	}
	// Save enclave to IPFS
	cid, err := ipc.Unixfs().Add(context.Background(), files.NewBytesFile(jsonEnclave))
	if err != nil {
		return nil, err
	}
	enclave[kVaultCIDKey] = []byte(cid.String())
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

func getKeyShareArrayPoint(kss []KeyShare) (Point, error) {
	for _, ks := range kss {
		if ks.Role() == RoleUser {
			msg, err := ks.Message()
			if err != nil {
				return nil, err
			}
			return getBobPubPoint(msg)
		}
		if ks.Role() == RoleValidator {
			msg, err := ks.Message()
			if err != nil {
				return nil, err
			}
			return getAlicePubPoint(msg)
		}
	}
	return nil, fmt.Errorf("invalid share role")
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

func getRefreshFunc(ks KeyShare) (RefreshFunc, error) {
	curve := curves.K256()
	msg, err := ks.Message()
	if err != nil {
		return nil, err
	}
	switch ks.Role() {
	case RoleUser:
		return dklsv1.NewBobRefresh(curve, msg, protocol.Version1)
	case RoleValidator:
		return dklsv1.NewAliceRefresh(curve, msg, protocol.Version1)
	default:
		return nil, fmt.Errorf("invalid share role")
	}
}

func getSignFunc(ks KeyShare, msgBz []byte) (SignFunc, error) {
	curve := curves.K256()
	msg, err := ks.Message()
	if err != nil {
		return nil, err
	}
	switch ks.Role() {
	case RoleUser:
		return dklsv1.NewBobSign(curve, sha3.New256(), msgBz, msg, protocol.Version1)
	case RoleValidator:
		return dklsv1.NewAliceSign(curve, sha3.New256(), msgBz, msg, protocol.Version1)
	default:
		return nil, fmt.Errorf("invalid share role")
	}
}

// SerializeSecp256k1Signature serializes an ECDSA signature into a byte slice
func serializeSignature(sig *curves.EcdsaSignature) ([]byte, error) {
	rBytes := sig.R.Bytes()
	sBytes := sig.S.Bytes()

	sigBytes := make([]byte, 66) // V (1 byte) + R (32 bytes) + S (32 bytes)
	sigBytes[0] = byte(sig.V)
	copy(sigBytes[33-len(rBytes):33], rBytes)
	copy(sigBytes[66-len(sBytes):66], sBytes)
	return sigBytes, nil
}

// DeserializeSecp256k1Signature deserializes an ECDSA signature from a byte slice
func deserializeSignature(sigBytes []byte) (*curves.EcdsaSignature, error) {
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

func marshalPointJSON(point curves.Point) ([]byte, error) {
	m := make(map[string]string, 2)
	m["type"] = point.CurveName()
	m["value"] = hex.EncodeToString(point.ToAffineCompressed())
	return json.Marshal(m)
}

func unmarshalPointJSON(input []byte) (curves.Point, error) {
	var m map[string]string

	err := json.Unmarshal(input, &m)
	if err != nil {
		return nil, err
	}
	curve := curves.GetCurveByName(m["type"])
	if curve == nil {
		return nil, fmt.Errorf("invalid type")
	}
	p, err := hex.DecodeString(m["value"])
	if err != nil {
		return nil, err
	}
	P, err := curve.Point.FromAffineCompressed(p)
	if err != nil {
		return nil, err
	}
	return P, nil
}
