package models

import (
	"github.com/di-dao/sonr/crypto"
	"github.com/di-dao/sonr/crypto/bip32"
	"github.com/di-dao/sonr/x/did/types"
	"gorm.io/gorm"
)

// Coin represents a cryptocurrency
type Coin interface {
	// FormatAddress formats a public key into an address
	FormatAddress(pubKey []byte) (string, error)

	// GetIndex returns the coin type index
	GetIndex() int64

	// GetPath returns the coin component path
	GetPath() uint32

	// GetSymbol returns the coin symbol
	GetSymbol() string

	// GetName returns the coin name
	GetName() string
}

// DefaultCoins is a list of default coins used in the vault
var DefaultCoins = []Coin{
	types.CoinBTC,
	types.CoinETH,
	types.CoinSNR,
}

// CoinBTCType is the coin type for BTC
const CoinBTCType = int64(0)

// CoinETHType is the coin type for ETH
const CoinETHType = int64(60)

// CoinSNRType is the coin type for SNR
const CoinSNRType = int64(703)

// Wallet is a struct that contains the information of a wallet account
type Wallet struct {
	gorm.Model
	Address    string `json:"address"`
	Controller string `json:"controller"`
	PublicKey  []byte `json:"publicKey"`
	Index      int    `json:"index"`
	CoinType   int64  `json:"coinType"`
}

// NewWallet creates a new account from a public key, coin, and index
func NewWallet(pubkey crypto.PublicKey, coin Coin, index int) (*Wallet, error) {
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
