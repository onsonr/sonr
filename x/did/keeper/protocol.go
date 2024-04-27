package keeper

import (
	"github.com/di-dao/core/crypto/core/curves"
	"github.com/di-dao/core/crypto/core/protocol"
	"github.com/di-dao/core/crypto/tecdsa/dklsv1"
)

// GenerateKSS generates both keyshares
func GenerateKSS() (*ValidatorKeyshare, *UserKeyshare, error) {
	defaultCurve := curves.P256()
	bob := dklsv1.NewBobDkg(defaultCurve, protocol.Version1)
	alice := dklsv1.NewAliceDkg(defaultCurve, protocol.Version1)
	err := StartKsProtocol(bob, alice)
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
	valKs := createValidatorKeyshare(aliceRes)
	usrKs := createUserKeyshare(bobRes)
	return valKs, usrKs, nil
}
