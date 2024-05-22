package mpc

import (
	"errors"
	"fmt"

	"github.com/di-dao/sonr/crypto/core/curves"
	"github.com/di-dao/sonr/crypto/core/protocol"
	"github.com/di-dao/sonr/crypto/kss"
	"github.com/di-dao/sonr/crypto/tecdsa/dklsv1"
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
