package shares

import (
	"fmt"
	"math/big"

	"golang.org/x/crypto/sha3"

	"github.com/sonrhq/sonr/crypto/core/curves"
	"github.com/sonrhq/sonr/crypto/core/protocol"
	"github.com/sonrhq/sonr/crypto/tecdsa/dklsv1"
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

func signV1(msg []byte, aliceDkgResultMessage *protocol.Message, bobDkgResultMessage *protocol.Message) (*curves.EcdsaSignature, error) {
	// New DklsSign
	curve := K_DEFAULT_MPC_CURVE
	aliceSign, err := dklsv1.NewAliceSign(curve, sha3.New256(), msg, aliceDkgResultMessage, protocol.Version1)
	if err != nil {
		return nil, err
	}
	bobSign, err := dklsv1.NewBobSign(curve, sha3.New256(), msg, bobDkgResultMessage, protocol.Version1)
	if err != nil {
		return nil, err
	}

	err = runIteratedProtocol(aliceSign, bobSign)
	if err != nil {
		return nil, err
	}
	// Output
	resultMessage, err := bobSign.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	return dklsv1.DecodeSignature(resultMessage)
}

func refreshV1(aliceDkgResultMessage, bobDkgResultMessage *protocol.Message) (aliceRefreshResultMessage, bobRefreshResultMessage *protocol.Message) {
	curve := K_DEFAULT_MPC_CURVE
	aliceRefresh, err := dklsv1.NewAliceRefresh(curve, aliceDkgResultMessage, protocol.Version1)
	if err != nil {
		return nil, nil
	}
	bobRefresh, err := dklsv1.NewBobRefresh(curve, bobDkgResultMessage, protocol.Version1)
	if err != nil {
		return nil, nil
	}
	err = runIteratedProtocol(aliceRefresh, bobRefresh)
	if err != nil {
		return nil, nil
	}
	aliceRefreshResultMessage, err = aliceRefresh.Result(protocol.Version1)
	if err != nil {
		return nil, nil
	}
	bobRefreshResultMessage, err = bobRefresh.Result(protocol.Version1)
	if err != nil {
		return nil, nil
	}
	return aliceRefreshResultMessage, bobRefreshResultMessage
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

// buildEcPoint builds an elliptic curve point from a compressed byte slice
func buildEcPoint(bz []byte) (*curves.EcPoint, error) {
	crv := K_DEFAULT_MPC_CURVE
	x := new(big.Int).SetBytes(bz[1:33])
	y := new(big.Int).SetBytes(bz[33:])
	ecCurve, err := crv.ToEllipticCurve()
	if err != nil {
		return nil, fmt.Errorf("error converting curve: %v", err)
	}
	return &curves.EcPoint{X: x, Y: y, Curve: ecCurve}, nil
}
