package kss

import (
	"github.com/di-dao/core/crypto"
	"github.com/di-dao/core/crypto/core/curves"
	"github.com/di-dao/core/crypto/core/protocol"
	"github.com/di-dao/core/crypto/tecdsa/dklsv1"
	"github.com/di-dao/core/x/did/types"
	"golang.org/x/crypto/sha3"
)

// RefreshFuncUser is the type for the user refresh function
type RefreshFuncUser = *dklsv1.BobRefresh

// SignFuncUser is the type for the user sign function
type SignFuncUser = *dklsv1.BobSign

// User is the interface for the user keyshare
type User interface {
	GetSignFunc(msg []byte) (SignFuncUser, error)
	GetRefreshFunc() (RefreshFuncUser, error)
	PublicKey() crypto.PublicKey
}

// userKeyshare is the protocol result for the user keyshare
type userKeyshare struct {
	usrKSS *protocol.Message
}

// createUserKeyshare creates a new UserKeyshare and stores it into IPFS
func createUserKeyshare(usrKSS *protocol.Message) User {
	return &userKeyshare{
		usrKSS: usrKSS,
	}
}

// GetSignFunc returns the sign function for the user keyshare
func (u *userKeyshare) GetSignFunc(msg []byte) (SignFuncUser, error) {
	curve := curves.K256()
	bobSign, err := dklsv1.NewBobSign(curve, sha3.New256(), msg, u.usrKSS, protocol.Version1)
	if err != nil {
		return nil, err
	}
	return bobSign, nil
}

// GetRefreshFunc returns the refresh function for the user keyshare
func (u *userKeyshare) GetRefreshFunc() (RefreshFuncUser, error) {
	curve := curves.K256()
	bobRefresh, err := dklsv1.NewBobRefresh(curve, u.usrKSS, protocol.Version1)
	if err != nil {
		return nil, err
	}
	return bobRefresh, nil
}

// PublicKey is the public key for the keyshare
func (u *userKeyshare) PublicKey() crypto.PublicKey {
	bobOut, err := dklsv1.DecodeBobDkgResult(u.usrKSS)
	if err != nil {
		panic(err)
	}
	pub := &types.PublicKey{
		Key:     bobOut.PublicKey.ToAffineUncompressed(),
		KeyType: "ecdsa-secp256k1",
	}
	return pub
}
