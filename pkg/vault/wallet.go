package vault

import (
	"github.com/cometbft/cometbft/libs/bytes"
	"github.com/di-dao/core/pkg/kss"
	"github.com/di-dao/core/x/did/types"
)

// WalletData is a struct that contains the information of a wallet to be stored in the vault
type WalletData struct {
	// Address of the wallet
	Address bytes.HexBytes `json:"address"`

	// Keys of the wallet
	Keys []*Key `json:"keys"`

	// Did of the wallet
	Did string `json:"did"`

	// PubKey of the wallet
	PublicKey *types.PublicKey `json:"publicKey"`
}

// NewWallet creates a new wallet from kss.SetI
func NewWallet(keyshares kss.SetI) (*WalletData, error) {
	pubkey := keyshares.PublicKey()
	wallet := &WalletData{
		Address:   pubkey.Address(),
		Keys:      make([]*Key, 0),
		Did:       pubkey.Did,
		PublicKey: pubkey,
	}
	if err := createBtcKey(pubkey, wallet); err != nil {
		return nil, err
	}
	if err := createSonrKey(pubkey, wallet); err != nil {
		return nil, err
	}
	return wallet, nil
}
