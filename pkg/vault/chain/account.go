package chain

import (
	"github.com/di-dao/core/crypto"
	"github.com/di-dao/core/crypto/bip32"
)

// Account is a struct that contains the information of a wallet account
type Account struct {
	// Address of the wallet
	Address string `json:"address"`

	// PublicKey of the wallet
	PublicKey []byte `json:"publicKey"`

	// Index of the wallet
	Index int `json:"index"`

	// CoinType of the account
	CoinType int64 `json:"coinType"`
}

// NewAccount creates a new account from a public key, coin, and index
func NewAccount(pubkey crypto.PublicKey, coin Coin, index int) (*Account, error) {
	expbz := pubkey.Bytes()
	pubBz, err := bip32.ComputePublicKey(expbz, coin.GetPath(), index)
	if err != nil {
		return nil, err
	}
	addr, err := coin.FormatAddress(pubBz)
	if err != nil {
		return nil, err
	}
	return &Account{
		PublicKey: pubBz,
		Index:     index,
		Address:   addr,
		CoinType:  coin.GetIndex(),
	}, nil
}
