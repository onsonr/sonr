package wallet

import (
	"encoding/json"

	"github.com/di-dao/sonr/crypto/kss"
)

const filename = "wallet.json"

// Wallet is a struct that contains the information of a wallet to be stored in the vault
type Wallet struct {
	// Accounts of the wallet
	Accounts map[int64][]*Account `json:"accounts"`
}

// New creates a new wallet from kss.Set and coins
func New(kset kss.Set, coinList ...Coin) (*Wallet, error) {
	// Set default coins if none are provided
	if coinList == nil {
		coinList = DefaultCoins
	}

	// Define base wallet
	pubkey := kset.PublicKey()
	wallet := &Wallet{
		Accounts: make(map[int64][]*Account),
	}

	// Create accounts for each coin
	for _, coin := range coinList {
		i := len(wallet.Accounts[coin.GetIndex()])
		account, err := NewAccount(pubkey, coin, i)
		if err != nil {
			return nil, err
		}
		wallet.Accounts[coin.GetIndex()] = append(wallet.Accounts[coin.GetIndex()], account)
	}
	return wallet, nil
}

// Marshal returns the JSON encoding of the Credentials.
func (c *Wallet) Marshal() ([]byte, error) {
	return json.Marshal(c)
}

// Unmarshal parses the JSON-encoded data and stores the result in the Credentials.
func (c *Wallet) Unmarshal(data []byte) error {
	return json.Unmarshal(data, c)
}
