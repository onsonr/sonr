package keeper

import (
	"github.com/di-dao/core/crypto/core/protocol"
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
	ipfs.IPFSClient
	k Keeper
}

// NewController creates a new controller instance.
func (v vaultStore) NewController() (controller, error) {
	valKs, usrKs, err := generateKSS()
	if err != nil {
		return controller{}, err
	}
	controller := controller{
		valKS: valKs,
		usrKS: usrKs,
	}

	return controller, nil
}
