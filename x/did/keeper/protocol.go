package keeper

import (
	"errors"
	"fmt"

	"github.com/di-dao/core/crypto/core/protocol"
	"github.com/di-dao/core/crypto/tecdsa/dklsv1"
)

// GenerateKSS generates both keyshares
func GenerateKSS() (*ValidatorKeyshare, *UserKeyshare, error) {
	bob := dklsv1.NewBobDkg(defaultCurve, protocol.Version1)
	alice := dklsv1.NewAliceDkg(defaultCurve, protocol.Version1)
	err := StartKsProtocol(bob, alice)
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
