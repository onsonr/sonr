package types

import (
	"golang.org/x/crypto/sha3"

	"github.com/di-dao/core/crypto/core/curves"
	"github.com/di-dao/core/crypto/core/protocol"
	"github.com/di-dao/core/crypto/tecdsa/dklsv1"
)

// UserSignFunc is the type for the user sign function
type UserSignFunc = *dklsv1.BobSign

// UserRefreshFunc is the type for the user refresh function
type UserRefreshFunc = *dklsv1.BobRefresh

// ValidatorSignFunc is the type for the validator sign function
type ValidatorSignFunc = *dklsv1.AliceSign

// ValidatorRefreshFunc is the type for the validator refresh function
type ValidatorRefreshFunc = *dklsv1.AliceRefresh

// KeyshareSet is the set of keyshares for the protocol
type KeyshareSet struct {
	Val ValidatorKeyshare
	Usr UserKeyshare
}

// PublicKey returns the public key for the keyshare set
func (ks *KeyshareSet) PublicKey() *PublicKey {
	return ks.Val.PublicKey()
}

// UserKeyshare is the interface for the user keyshare
type UserKeyshare interface {
	GetSignFunc(msg []byte) (UserSignFunc, error)
	GetRefreshFunc() (UserRefreshFunc, error)
	PublicKey() *PublicKey
}

// ValidatorKeyshare is the interface for the validator keyshare
type ValidatorKeyshare interface {
	GetSignFunc(msg []byte) (ValidatorSignFunc, error)
	GetRefreshFunc() (ValidatorRefreshFunc, error)
	PublicKey() *PublicKey
}

// NewKeyshareSet creates a new KeyshareSet
func NewKeyshareSet(aliceResult *protocol.Message, bobResult *protocol.Message) *KeyshareSet {
	return &KeyshareSet{
		Val: createValidatorKeyshare(aliceResult),
		Usr: createUserKeyshare(bobResult),
	}
}

// GetSignFunc returns the sign function for the user keyshare
func (u *userKeyshare) GetSignFunc(msg []byte) (UserSignFunc, error) {
	curve := curves.P256()
	bobSign, err := dklsv1.NewBobSign(curve, sha3.New256(), msg, u.usrKSS, protocol.Version1)
	if err != nil {
		return nil, err
	}
	return bobSign, nil
}

// GetRefreshFunc returns the refresh function for the user keyshare
func (u *userKeyshare) GetRefreshFunc() (UserRefreshFunc, error) {
	curve := curves.P256()
	bobRefresh, err := dklsv1.NewBobRefresh(curve, u.usrKSS, protocol.Version1)
	if err != nil {
		return nil, err
	}
	return bobRefresh, nil
}

// PublicKey is the public key for the keyshare
func (u *userKeyshare) PublicKey() *PublicKey {
	bobOut, err := dklsv1.DecodeBobDkgResult(u.usrKSS)
	if err != nil {
		panic(err)
	}
	pub := &PublicKey{
		Key:     bobOut.PublicKey.ToAffineUncompressed(),
		KeyType: "ecdsa-secp256k1",
	}
	return pub
}

// GetSignFunc returns the sign function for the validator keyshare
func (v *validatorKeyshare) GetSignFunc(msg []byte) (ValidatorSignFunc, error) {
	curve := curves.P256()
	aliceSign, err := dklsv1.NewAliceSign(curve, sha3.New256(), msg, v.valKSS, protocol.Version1)
	if err != nil {
		return nil, err
	}
	return aliceSign, nil
}

// GetRefreshFunc returns the refresh function for the validator keyshare
func (v *validatorKeyshare) GetRefreshFunc() (ValidatorRefreshFunc, error) {
	curve := curves.P256()
	aliceRefresh, err := dklsv1.NewAliceRefresh(curve, v.valKSS, protocol.Version1)
	if err != nil {
		return nil, err
	}
	return aliceRefresh, nil
}

// PublicKey is the public key for the keyshare
func (u *validatorKeyshare) PublicKey() *PublicKey {
	aliceOut, err := dklsv1.DecodeAliceDkgResult(u.valKSS)
	if err != nil {
		panic(err)
	}
	pub := &PublicKey{
		Key:     aliceOut.PublicKey.ToAffineUncompressed(),
		KeyType: "ecdsa-secp256k1",
	}
	return pub
}

// userKeyshare is the protocol result for the user keyshare
type userKeyshare struct {
	usrKSS *protocol.Message
}

// validatorKeyshare is the protocol result for the validator keyshare
type validatorKeyshare struct {
	valKSS *protocol.Message
}

// createUserKeyshare creates a new UserKeyshare and stores it into IPFS
func createUserKeyshare(usrKSS *protocol.Message) UserKeyshare {
	return &userKeyshare{
		usrKSS: usrKSS,
	}
}

// createValidatorKeyshare creates a new ValidatorKeyshare
func createValidatorKeyshare(valKSS *protocol.Message) ValidatorKeyshare {
	return &validatorKeyshare{
		valKSS: valKSS,
	}
}
