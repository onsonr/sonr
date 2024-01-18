package shares

import (
	"fmt"
	"math/big"

	"github.com/sonrhq/sonr/crypto/core/curves"
	"github.com/sonrhq/sonr/crypto/core/protocol"
)

// For DKG bob starts first. For refresh and sign, Alice starts first.
func runIteratedProtocol(firstParty protocol.Iterator, secondParty protocol.Iterator) error {
	var (
		message *protocol.Message
		aErr    error
		bErr    error
	)

	for aErr != protocol.ErrProtocolFinished || bErr != protocol.ErrProtocolFinished {
		// Crank each protocol forward one iteration
		message, bErr = firstParty.Next(message)
		if bErr != nil && bErr != protocol.ErrProtocolFinished {
			return bErr
		}

		message, aErr = secondParty.Next(message)
		if aErr != nil && aErr != protocol.ErrProtocolFinished {
			return aErr
		}
	}
	return checkProtocolErrors(aErr, bErr)
}

func checkProtocolErrors(aErr, bErr error) error {
	if aErr == protocol.ErrProtocolFinished && bErr == protocol.ErrProtocolFinished {
		return nil
	}
	if aErr != nil && bErr != nil {
		return fmt.Errorf("both parties failed: %v, %v", aErr, bErr)
	}
	if aErr != nil {
		return fmt.Errorf("alice failed: %v", aErr)
	}
	if bErr != nil {
		return fmt.Errorf("bob failed: %v", bErr)
	}
	return nil
}

// BuildEcPoint builds an elliptic curve point from a compressed byte slice
func BuildEcPoint(crv *curves.Curve, bz []byte) (*curves.EcPoint, error) {
	x := new(big.Int).SetBytes(bz[1:33])
	y := new(big.Int).SetBytes(bz[33:])
	ecCurve, err := crv.ToEllipticCurve()
	if err != nil {
		return nil, fmt.Errorf("error converting curve: %v", err)
	}
	return &curves.EcPoint{X: x, Y: y, Curve: ecCurve}, nil
}

// // PublicPoint returns the public point of the party
// func (p *share) PublicPoint() (*curves.EcPoint, error) {
// 	// Decode the result message
// 	bobRes, err := dklsv1.DecodeBobDkgResult(nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return buildEcPoint(K_DEFAULT_MPC_CURVE, bobRes.PublicKey.ToAffineCompressed())
// }

// // PubKeyHex returns the public key of the party in hex format
// func (p *share) PubKeyHex() (string, error) {
// 	pp, err := p.PublicPoint()
// 	if err != nil {
// 		return "", err
// 	}
// 	ppbz, err := pp.MarshalBinary()
// 	if err != nil {
// 		return "", err
// 	}
// 	return hex.EncodeToString(ppbz), nil
// }

// // Verify verifies the signature of the message
// func (p *share) Verify(msg []byte, sigBz []byte) (bool, error) {
// 	sig, err := ecdsa.DeserializeSecp256k1Signature(sigBz)
// 	if err != nil {
// 		return false, fmt.Errorf("error deserializing signature: %v", err)
// 	}
// 	hash := sha3.New256()
// 	_, err = hash.Write(msg)
// 	if err != nil {
// 		return false, fmt.Errorf("error hashing message: %v", err)
// 	}
// 	digest := hash.Sum(nil)
// 	publicKey, err := p.PublicPoint()
// 	if err != nil {
// 		return false, fmt.Errorf("error getting public key: %v", err)
// 	}
// 	return curves.VerifyEcdsa(publicKey, digest[:], sig), nil
// }
