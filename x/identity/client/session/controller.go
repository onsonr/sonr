package session

import (
	"fmt"

	"github.com/sonrhq/core/x/vault"
)


type SessionController struct {
	accounts []vault.Account
}

func NewSessionController(accounts ...vault.Account) *SessionController {
	return &SessionController{
		accounts: accounts,
	}
}

// The `GetAccounts()` function is a method of the `SessionController` struct that returns a slice of `vault.Account` objects. It allows access to the `accounts` field of the `SessionController` struct from outside the struct.
func (sc *SessionController) GetAccounts() []vault.Account {
	return sc.accounts
}

// This function takes in a DID (decentralized identifier) string and a byte slice of data, and returns a byte slice of the signature and an error. It searches through the accounts in the SessionController to find the account with the matching DID, and then calls the Sign method on
// that account to sign the data. If the account is not found, it returns an error.
func (sc *SessionController) SignWithDID(did string, data []byte) ([]byte, error) {
	for _, account := range sc.accounts {
		if account.Did() == did {
			return account.Sign(data)
		}
	}
	return nil, fmt.Errorf("account with did %s not found", did)
}

// The `VerifyWithDID` function is a method of the `SessionController` struct that takes in a DID (decentralized identifier) string, a byte slice of data, and a byte slice of signature. It searches through the accounts in the `SessionController` to find the account with the matching
// DID, and then calls the `Verify` method on that account to verify the signature against the data. If the account is not found, it returns an error. The function returns a boolean value indicating whether the signature is valid and an error if there is any.
func (sc *SessionController) VerifyWithDID(did string, data []byte, signature []byte) (bool, error) {
	for _, account := range sc.accounts {
		if account.Did() == did {
			return account.Verify(data, signature)
		}
	}
	return false, fmt.Errorf("account with did %s not found", did)
}
