package internal

import (
	"errors"

	"github.com/sonrhq/core/pkg/crypto/wallet"
)

type empty struct {}

func (e *empty) GetAccount(name string) (wallet.Account, error) {
	return nil, errors.New("not implemented")
}

func (e *empty) PutAccount(acc wallet.Account) error {
	return errors.New("not implemented")
}


