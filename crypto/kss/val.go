package kss

import (
	"github.com/di-dao/core/crypto"
	"github.com/di-dao/core/crypto/core/curves"
	"github.com/di-dao/core/crypto/core/protocol"
	"github.com/di-dao/core/crypto/tecdsa/dklsv1"
	"github.com/di-dao/core/x/did/types"
	"golang.org/x/crypto/sha3"
)

// RefreshFuncVal is the type for the validator refresh function
type RefreshFuncVal = *dklsv1.AliceRefresh

// SignFuncVal is the type for the validator sign function
type SignFuncVal = *dklsv1.AliceSign

// Val is the interface for the validator keyshare
type Val interface {
	GetSignFunc(msg []byte) (SignFuncVal, error)
	GetRefreshFunc() (RefreshFuncVal, error)
	PublicKey() crypto.PublicKey
}

// validatorKeyshare is the protocol result for the validator keyshare
type validatorKeyshare struct {
	valKSS *protocol.Message
}

// createValidatorKeyshare creates a new ValidatorKeyshare
func createValidatorKeyshare(valKSS *protocol.Message) Val {
	return &validatorKeyshare{
		valKSS: valKSS,
	}
}

// GetSignFunc returns the sign function for the validator keyshare
func (v *validatorKeyshare) GetSignFunc(msg []byte) (SignFuncVal, error) {
	curve := curves.K256()
	aliceSign, err := dklsv1.NewAliceSign(curve, sha3.New256(), msg, v.valKSS, protocol.Version1)
	if err != nil {
		return nil, err
	}
	return aliceSign, nil
}

// GetRefreshFunc returns the refresh function for the validator keyshare
func (v *validatorKeyshare) GetRefreshFunc() (RefreshFuncVal, error) {
	curve := curves.K256()
	aliceRefresh, err := dklsv1.NewAliceRefresh(curve, v.valKSS, protocol.Version1)
	if err != nil {
		return nil, err
	}
	return aliceRefresh, nil
}

// PublicKey is the public key for the keyshare
func (u *validatorKeyshare) PublicKey() crypto.PublicKey {
	aliceOut, err := dklsv1.DecodeAliceDkgResult(u.valKSS)
	if err != nil {
		panic(err)
	}
	pub := &types.PublicKey{
		Key:     aliceOut.PublicKey.ToAffineUncompressed(),
		KeyType: "ecdsa-secp256k1",
	}
	return pub
}
