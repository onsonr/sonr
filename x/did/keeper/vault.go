package keeper

import (
	"github.com/di-dao/core/crypto/core/protocol"
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
	valKs, usrKs, err := generateKSS()
	if err != nil {
		return nil, err
	}
	return CreateController(usrKs, valKs)
}

// formatUserKeyshareDID formats the user keyshare DID
func setUserKeyshareDID(pub *types.PublicKey) (*types.PublicKey, error) {
	addr, err := types.GetIDXAddress(pub)
	if err != nil {
		return nil, err
	}
	pub.Did = addr.DID("ipns")
	return pub, nil
}

// formatValidatorKeyshareDID formats the validator keyshare DID
func setValidatorKeyshareDID(pub *types.PublicKey) (*types.PublicKey, error) {
	addr, err := types.GetIDXAddress(pub)
	if err != nil {
		return nil, err
	}
	pub.Did = addr.DID("vksnr")
	return pub, nil
}
