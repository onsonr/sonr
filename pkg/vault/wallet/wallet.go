package wallet

import (
	"encoding/json"

	"github.com/di-dao/core/crypto/kss"
	"github.com/di-dao/core/pkg/auth"
	"github.com/di-dao/core/pkg/vault/chain"
)

// Wallet is a struct that contains the information of a wallet to be stored in the vault
type Wallet struct {
	// Accounts of the wallet
	Accounts map[int64][]*chain.Account `json:"accounts"`

	// Credentials of the wallet
	Credentials map[string][]*auth.Credential `json:"data"`
}

// New creates a new wallet from kss.Set and coins
func New(kset kss.Set, coinList ...chain.Coin) (*Wallet, error) {
	if coinList == nil {
		coinList = chain.DefaultCoins
	}
	pubkey := kset.PublicKey()
	wallet := &Wallet{
		Accounts: make(map[int64][]*chain.Account),
	}
	for _, coin := range coinList {
		i := len(wallet.Accounts[coin.GetIndex()])
		account, err := chain.NewAccount(pubkey, coin, i)
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
