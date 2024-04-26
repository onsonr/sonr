package keeper

import (
	"github.com/di-dao/core/crypto/core/protocol"
	"github.com/di-dao/core/crypto/tecdsa/dklsv1"
	"github.com/di-dao/core/crypto/tecdsa/dklsv1/dkg"

	"golang.org/x/crypto/sha3"
)

// ValidatorKSOutput is the protocol result for the validator keyshare output
type ValidatorKSOutput = *dkg.AliceOutput

// ValidatorSignFunc is the type for the validator sign function
type ValidatorSignFunc = *dklsv1.AliceSign

// ValidatorRefreshFunc is the type for the validator refresh function
type ValidatorRefreshFunc = *dklsv1.AliceRefresh

// ValidatorKeyshare is the protocol result for the validator keyshare
type ValidatorKeyshare struct {
	Keyshare
	valKSS *protocol.Message
}

// newValidatorKeyshare creates a new ValidatorKeyshare
func NewValidatorKeyshare(valKSS *protocol.Message) *ValidatorKeyshare {
	return &ValidatorKeyshare{
		valKSS: valKSS,
	}
}

// DecodeOutput decodes the output from the validator keyshare
func (v *ValidatorKeyshare) DecodeOutput() (ValidatorKSOutput, error) {
	aliceOut, err := dklsv1.DecodeAliceDkgResult(v.valKSS)
	if err != nil {
		return nil, err
	}
	return aliceOut, nil
}

// GetSignFunc returns the sign function for the validator keyshare
func (v *ValidatorKeyshare) GetSignFunc(msg []byte) (ValidatorSignFunc, error) {
	curve := defaultCurve
	aliceSign, err := dklsv1.NewAliceSign(curve, sha3.New256(), msg, v.valKSS, protocol.Version1)
	if err != nil {
		return nil, err
	}
	return aliceSign, nil
}

// GetRefreshFunc returns the refresh function for the validator keyshare
func (v *ValidatorKeyshare) GetRefreshFunc() (ValidatorRefreshFunc, error) {
	curve := defaultCurve
	aliceRefresh, err := dklsv1.NewAliceRefresh(curve, v.valKSS, protocol.Version1)
	if err != nil {
		return nil, err
	}
	return aliceRefresh, nil
}

// PublicKey is the public key for the keyshare
func (v *ValidatorKeyshare) PublicKey() ([]byte, error) {
	aliceOut, err := v.DecodeOutput()
	if err != nil {
		return nil, err
	}
	return aliceOut.PublicKey.ToAffineUncompressed(), nil
}
