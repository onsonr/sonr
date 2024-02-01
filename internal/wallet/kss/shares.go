package kss

import (
	"fmt"

	"golang.org/x/crypto/sha3"

	"github.com/sonrhq/sonr/pkg/crypto/core/curves"
	"github.com/sonrhq/sonr/pkg/crypto/core/protocol"
	"github.com/sonrhq/sonr/pkg/crypto/signatures/ecdsa"
	"github.com/sonrhq/sonr/pkg/crypto/tecdsa/dklsv1"
)

// K_DEFAULT_MPC_CURVE is the default curve for the controller.
var K_DEFAULT_MPC_CURVE = curves.P256()

func Generate() ([]byte, *protocol.Message, *protocol.Message, error) {
	bob := dklsv1.NewBobDkg(K_DEFAULT_MPC_CURVE, protocol.Version1)
	alice := dklsv1.NewAliceDkg(K_DEFAULT_MPC_CURVE, protocol.Version1)
	err := runIteratedProtocol(bob, alice)
	if err != nil {
		return nil, nil, nil, err
	}
	aliceRes, err := alice.Result(protocol.Version1)
	if err != nil {
		return nil, nil, nil, err
	}
	bobRes, err := bob.Result(protocol.Version1)
	if err != nil {
		return nil, nil, nil, err
	}
	aliceOut, err := dklsv1.DecodeAliceDkgResult(aliceRes)
	if err != nil {
		return nil, nil, nil, err
	}
	pub := aliceOut.PublicKey.ToAffineUncompressed()
	return pub, aliceRes, bobRes, nil
}

// VerifySignature verifies a signature with a PublicKey and a message
func VerifySignature(pubKey []byte, msg []byte, sigBz []byte) (bool, error) {
	pp, err := buildEcPoint(pubKey)
	if err != nil {
		return false, err
	}
	sig, err := ecdsa.DeserializeSecp256k1Signature(sigBz)
	if err != nil {
		return false, fmt.Errorf("error deserializing signature: %v", err)
	}
	hash := sha3.New256()
	_, err = hash.Write(msg)
	if err != nil {
		return false, fmt.Errorf("error hashing message: %v", err)
	}
	digest := hash.Sum(nil)
	return curves.VerifyEcdsa(pp, digest[:], sig), nil
}
