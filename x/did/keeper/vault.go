package keeper

import (
	"github.com/di-dao/core/crypto/core/curves"
	"github.com/di-dao/core/crypto/core/protocol"
	"github.com/di-dao/core/crypto/tecdsa/dklsv1"
	"github.com/di-dao/core/pkg/ipfs"
	"github.com/di-dao/core/x/did/types"
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
	k    Keeper
}

// NewController creates a new controller instance.
func (v vaultStore) NewController() (Controller, error) {
	kss, err := GenerateKSS()
	if err != nil {
		return nil, err
	}
	return CreateController(kss)
}

// GenerateKSS generates both keyshares
func GenerateKSS() (*types.KeyshareSet, error) {
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
	return types.NewKeyshareSet(aliceRes, bobRes), nil
}
