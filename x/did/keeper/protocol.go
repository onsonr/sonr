package keeper

import (
	"errors"
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
		return nil, nil, errors.Join(fmt.Errorf("Error Starting Keyshare MPC Protocol"), err)
	}
	aliceRes, err := alice.Result(protocol.Version1)
	if err != nil {
		return nil, nil, errors.Join(fmt.Errorf("Error Getting Validator Result"), err)
	}
	bobRes, err := bob.Result(protocol.Version1)
	if err != nil {
		return nil, nil, errors.Join(fmt.Errorf("Error Getting User Result"), err)
	}
	usrKs, err := createUserKeyshare(bobRes)
	if err != nil {
		return nil, nil, errors.Join(fmt.Errorf("Error Creating User Keyshare"), err)
	}
	valKs, err := createValidatorKeyshare(aliceRes)
	if err != nil {
		return nil, nil, errors.Join(fmt.Errorf("Error Creating Validator Keyshare"), err)
	}
	return valKs, usrKs, nil
}

// signKSS signs a message with the SignFuncs
func signKSS(valSign ValidatorSignFunc, usrSign UserSignFunc) ([]byte, error) {
	err := startKsProtocol(valSign, usrSign)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("Error Starting Keyshare MPC Protocol"), err)
	}
	// Output
	resultMessage, err := usrSign.Result(protocol.Version1)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("Error Getting User Sign Result"), err)
	}
	sig, err := dklsv1.DecodeSignature(resultMessage)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("Error Decoding Signature"), err)
	}
	return ecdsa.SerializeSecp256k1Signature(sig)
}

// refreshKSS refreshes both keyshares
func refreshKSS(valRefresh ValidatorRefreshFunc, usrRefresh UserRefreshFunc) (*ValidatorKeyshare, *UserKeyshare, error) {
	err := startKsProtocol(valRefresh, usrRefresh)
	if err != nil {
		return nil, nil, errors.Join(fmt.Errorf("Error Starting Keyshare MPC Protocol"), err)
	}
	newAlice, err := valRefresh.Result(protocol.Version1)
	if err != nil {
		return nil, nil, errors.Join(fmt.Errorf("Error Getting Validator Result"), err)
	}
	newBob, err := usrRefresh.Result(protocol.Version1)
	if err != nil {
		return nil, nil, errors.Join(fmt.Errorf("Error Getting User Result"), err)
	}
	usrKs, err := createUserKeyshare(newAlice)
	if err != nil {
		return nil, nil, errors.Join(fmt.Errorf("Error Creating User Keyshare"), err)
	}
	valKs, err := createValidatorKeyshare(newBob)
	if err != nil {
		return nil, nil, errors.Join(fmt.Errorf("Error Creating Validator Keyshare"), err)
	}
	return valKs, usrKs, nil
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
