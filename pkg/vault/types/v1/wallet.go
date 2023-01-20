package v1

import (
	"errors"
	"strings"
)

func NewWalletConfigFromRootAccount(account *AccountConfig) *WalletConfig {
	return &WalletConfig{
		Address:   account.Address,
		PublicKey: account.PublicKey,
		Algorithm: "cmp",
		Accounts: map[string]*AccountConfig{
			strings.ToLower(account.Name): account,
		},
	}
}

func (w *WalletConfig) GetAccountByName(name string) (*AccountConfig, error) {
	for _, acc := range w.Accounts {
		if acc.Name == strings.ToLower(name) {
			return acc, nil
		}
	}
	return nil, errors.New("account not found")
}

func (w *WalletConfig) GetAccountByIndex(index int) (*AccountConfig, error) {
	if index >= len(w.Accounts) {
		return nil, errors.New("account not found")
	}
	for _, acc := range w.Accounts {
		if acc.Index == uint32(index) {
			return acc, nil
		}
	}
	return nil, errors.New("account not found")
}
