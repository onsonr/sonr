package keeper

import (
	"fmt"

	"github.com/di-dao/core/crypto/core/protocol"
	"github.com/di-dao/core/crypto/signatures/ecdsa"
	"github.com/di-dao/core/crypto/tecdsa/dklsv1"
)

// generateKSS generates both keyshares
func generateKSS() (*ValidatorKeyshare, *UserKeyshare, error) {
	bob := dklsv1.NewBobDkg(defaultCurve, protocol.Version1)
	alice := dklsv1.NewAliceDkg(defaultCurve, protocol.Version1)
	err := startKsProtocol(bob, alice)
	if err != nil {
		return nil, nil, err
	}
	aliceRes, err := alice.Result(protocol.Version1)
	if err != nil {
		return nil, nil, err
	}
	bobRes, err := bob.Result(protocol.Version1)
	if err != nil {
		return nil, nil, err
	}
	return NewValidatorKeyshare(aliceRes), NewUserKeyshare(bobRes), nil
}

// signKSS signs a message with the SignFuncs
func signKSS(valSign ValidatorSignFunc, usrSign UserSignFunc) ([]byte, error) {
	err := startKsProtocol(valSign, usrSign)
	if err != nil {
		return nil, err
	}
	// Output
	resultMessage, err := usrSign.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	sig, err := dklsv1.DecodeSignature(resultMessage)
	if err != nil {
		return nil, err
	}
	return ecdsa.SerializeSecp256k1Signature(sig)
}

// refreshKSS refreshes both keyshares
func refreshKSS(valRefresh ValidatorRefreshFunc, usrRefresh UserRefreshFunc) (*ValidatorKeyshare, *UserKeyshare, error) {
	err := startKsProtocol(valRefresh, usrRefresh)
	if err != nil {
		return nil, nil, err
	}
	newValKsMsg, err := valRefresh.Result(protocol.Version1)
	if err != nil {
		return nil, nil, err
	}
	newUsrKsMsg, err := usrRefresh.Result(protocol.Version1)
	if err != nil {
		return nil, nil, err
	}
	return NewValidatorKeyshare(newValKsMsg), NewUserKeyshare(newUsrKsMsg), nil
}

// startKsProtocol runs the keyshare protocol between two parties
func startKsProtocol(firstParty protocol.Iterator, secondParty protocol.Iterator) error {
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
	if aErr == protocol.ErrProtocolFinished && bErr == protocol.ErrProtocolFinished {
		return nil
	}
	if aErr != nil && bErr != nil {
		return fmt.Errorf("both parties failed: %v, %v", aErr, bErr)
	}
	if aErr != nil {
		return fmt.Errorf("validator keyshare failed: %v", aErr)
	}
	if bErr != nil {
		return fmt.Errorf("user keyshare failed: %v", bErr)
	}
	return nil
}
