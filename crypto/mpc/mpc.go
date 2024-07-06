package mpc

import (
	"errors"
	"fmt"
	"github.com/onsonr/hway/crypto/core/curves"
	"github.com/onsonr/hway/crypto/core/protocol"
	"github.com/onsonr/hway/crypto/kss"
	"github.com/onsonr/hway/crypto/signatures/ecdsa"
	"github.com/onsonr/hway/crypto/tecdsa/dklsv1"
)

// GenerateKss generates both keyshares
func GenerateKss() (kss.Set, error) {
	defaultCurve := curves.K256()
	bob := dklsv1.NewBobDkg(defaultCurve, protocol.Version1)
	alice := dklsv1.NewAliceDkg(defaultCurve, protocol.Version1)
	err := RunProtocol(bob, alice)
	if err != nil {
		return nil, err
	}
	aliceRes, err := alice.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	bobRes, err := bob.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	return kss.NewKeyshareSet(aliceRes, bobRes)
}

// RunProtocol runs the keyshare protocol between two parties
func RunProtocol(firstParty protocol.Iterator, secondParty protocol.Iterator) error {
	var (
		message *protocol.Message
		aErr    error
		bErr    error
	)

	for aErr != protocol.ErrProtocolFinished || bErr != protocol.ErrProtocolFinished {
		// Crank each protocol forward one iteration
		message, bErr = firstParty.Next(message)
		if bErr != nil && bErr != protocol.ErrProtocolFinished {
			return errors.Join(fmt.Errorf("validator failed to process mpc message"), bErr)
		}

		message, aErr = secondParty.Next(message)
		if aErr != nil && aErr != protocol.ErrProtocolFinished {
			return errors.Join(fmt.Errorf("user failed to process mpc message"), aErr)
		}
	}
	if aErr == protocol.ErrProtocolFinished && bErr == protocol.ErrProtocolFinished {
		return nil
	}
	if aErr != nil {
		return fmt.Errorf("validator keyshare failed: %v", aErr)
	}
	if bErr != nil {
		return fmt.Errorf("user keyshare failed: %v", bErr)
	}
	return nil
}

// RunSignProtocol runs the generic dkls protocol using the kss.SignFuncVal and kss.SignFuncUser
func RunSignProtocol(valSignFunc kss.SignFuncVal, userSignFunc kss.SignFuncUser) ([]byte, error) {
	err := RunProtocol(valSignFunc, userSignFunc)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to get validator sign function"))
	}
	resultMessage, err := userSignFunc.Result(protocol.Version1)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("error Getting User Sign Result"), err)
	}
	sig, err := dklsv1.DecodeSignature(resultMessage)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("error Decoding Signature"), err)
	}
	return ecdsa.SerializeSecp256k1Signature(sig)
}

// RunRefreshProtocol runs the generic dkls protocol using the kss.RefreshFuncVal and kss.RefreshFuncUser
func RunRefreshProtocol(valRefreshFunc kss.RefreshFuncVal, userRefreshFunc kss.RefreshFuncUser) (kss.Set, error) {
  err := RunProtocol(valRefreshFunc, userRefreshFunc)
  if err != nil {
    return nil, err
  }
	valRes, err := valRefreshFunc.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	userRes, err := userRefreshFunc.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	return kss.NewKeyshareSet(valRes, userRes)
}
