package controller

import (
	"errors"
	"fmt"

	"github.com/di-dao/core/crypto/core/curves"
	"github.com/di-dao/core/crypto/core/protocol"
	"github.com/di-dao/core/crypto/tecdsa/dklsv1"
	"github.com/di-dao/core/pkg/ipfs"
)

// RefreshFunc is the type for the refresh function
type RefreshFunc = protocol.Iterator

// SignFunc is the type for the sign function
type SignFunc = protocol.Iterator

// Keyshare is the interface for the keyshare
type Keyshare interface {
	DecodeOutput() (interface{}, error)
	GetSignFunc(msg []byte) (SignFunc, error)
	GetRefreshFunc() (RefreshFunc, error)
	PublicKey() ([]byte, error)
}

// vaultStore is the interface for interacting with Keyshares in the IPFS network.
type vaultStore struct {
	ipfs ipfs.IPFSClient
}

// NewController creates a new controller instance.
func (v vaultStore) NewController() (Controller, error) {
	kss, err := GenKSS()
	if err != nil {
		return nil, err
	}
	return Create(kss)
}

// GenKSS generates both keyshares
func GenKSS() (KeyshareSet, error) {
	defaultCurve := curves.P256()
	bob := dklsv1.NewBobDkg(defaultCurve, protocol.Version1)
	alice := dklsv1.NewAliceDkg(defaultCurve, protocol.Version1)
	err := runMpcProtocol(bob, alice)
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
	return NewKeyshareSet(aliceRes, bobRes), nil
}

// runMpcProtocol runs the keyshare protocol between two parties
func runMpcProtocol(firstParty protocol.Iterator, secondParty protocol.Iterator) error {
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
