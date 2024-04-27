package keeper

import (
	"github.com/di-dao/core/crypto/core/protocol"
	"github.com/di-dao/core/crypto/tecdsa/dklsv1"
	"github.com/di-dao/core/crypto/tecdsa/dklsv1/dkg"
	"github.com/di-dao/core/x/did/types"
	"golang.org/x/crypto/sha3"
)

// UserKSOutput is the protocol result for the user keyshare output
type UserKSOutput = *dkg.BobOutput

// UserSignFunc is the type for the user sign function
type UserSignFunc = *dklsv1.BobSign

// UserRefreshFunc is the type for the user refresh function
type UserRefreshFunc = *dklsv1.BobRefresh

// UserKeyshare is the protocol result for the user keyshare
type UserKeyshare struct {
	Keyshare
	usrKSS *protocol.Message
	pubKey *types.PublicKey
}

// createUserKeyshare creates a new UserKeyshare and stores it into IPFS
func createUserKeyshare(usrKSS *protocol.Message) (*UserKeyshare, error) {
	bobOut, err := dklsv1.DecodeBobDkgResult(usrKSS)
	if err != nil {
		return nil, err
	}
	pub := &types.PublicKey{
		Key:     bobOut.PublicKey.ToAffineUncompressed(),
		KeyType: "ecdsa-secp256k1",
	}
	pub, err = setUserKeyshareDID(pub)
	if err != nil {
		return nil, err
	}
	return &UserKeyshare{
		usrKSS: usrKSS,
		pubKey: pub,
	}, nil
}

// GetSignFunc returns the sign function for the user keyshare
func (u *UserKeyshare) GetSignFunc(msg []byte) (UserSignFunc, error) {
	curve := defaultCurve
	bobSign, err := dklsv1.NewBobSign(curve, sha3.New256(), msg, u.usrKSS, protocol.Version1)
	if err != nil {
		return nil, err
	}
	return bobSign, nil
}

// GetRefreshFunc returns the refresh function for the user keyshare
func (u *UserKeyshare) GetRefreshFunc() (UserRefreshFunc, error) {
	curve := defaultCurve
	bobRefresh, err := dklsv1.NewBobRefresh(curve, u.usrKSS, protocol.Version1)
	if err != nil {
		return nil, err
	}
	return bobRefresh, nil
}

// PublicKey is the public key for the keyshare
func (u *UserKeyshare) PublicKey() *types.PublicKey {
	return u.pubKey
}
