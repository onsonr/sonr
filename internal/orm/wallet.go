package orm

import (
	"github.com/onsonr/hway/crypto"
	"github.com/onsonr/hway/crypto/bip32"
	"github.com/onsonr/hway/pkg/coins"
	"gorm.io/gorm"
)

// Wallet is a struct that contains the information of a wallet account
type Wallet struct {
	gorm.Model
	Address    string `json:"address"`
	Controller string `json:"controller"`
	Name       string `json:"name"`
	ChainID    string `json:"chainId"`
	Network    string `json:"network"`
	Label      string `json:"label"`
	DID        string `json:"did"`
	PublicKey  []byte `json:"publicKey"`
	Index      int    `json:"index"`
	CoinType   int64  `json:"coinType"`
}

// NewWallet creates a new account from a public key, coin, and index
func NewWallet(pubkey crypto.PublicKey, coin coins.Coin, index int) (*Wallet, error) {
	expbz := pubkey.Bytes()
	pubBz, err := bip32.ComputePublicKey(expbz, coin.GetPath(), index)
	if err != nil {
		return nil, err
	}
	addr, err := coin.FormatAddress(pubBz)
	if err != nil {
		return nil, err
	}
	return &Wallet{
		PublicKey: pubBz,
		Index:     index,
		Address:   addr,
		CoinType:  coin.GetIndex(),
	}, nil
}
