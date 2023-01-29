package controller

import (
	"context"
	"errors"
	"strings"

	"github.com/sonrhq/core/pkg/common"
	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/pkg/crypto/wallet"
	"github.com/sonrhq/core/pkg/crypto/wallet/accounts"
	"github.com/sonrhq/core/pkg/crypto/wallet/stores"
	"github.com/sonrhq/core/x/identity/types"
)

// rootWalletAccountName is the name of the root account
const rootWalletAccountName = "Primary"

// `DIDController` is a type that is both a `wallet.Wallet` and a `store.WalletStore`.
// @property GetChallengeResponse - This method is used to get the challenge response from the DID
// controller.
// @property RegisterAuthenticationCredential - This is the method that will be called when the user
// clicks on the "Register" button.
// @property GetAssertionOptions - This method is used to get the options for the assertion.
// @property AuthorizeCredential - This is the method that will be called when the user clicks the
// "Login" button on the login page.
type DIDController interface {
	// Address
	Address() string

	// DID
	ID() string

	// DID Document
	Document() *crypto.DidDocument

	// This method is used to get the challenge response from the DID controller.
	// GetChallengeOptions(aka string) (*v1.ChallengeResponse, error)

	// This is the method that will be called when the user clicks on the "Register" button.
	// RegisterAuthenticationCredential(credentialCreationData string) (*v1.RegisterResponse, error)

	// This method is used to get the options for the assertion.
	// GetAssertionOptions(aka string) (*v1.AssertResponse, error)

	// This is the method that will be called when the user clicks the "Login" button on the login page.
	// AuthorizeCredential(credentialRequestData string) (*v1.LoginResponse, error)

	// Creates a new account
	CreateAccount(name string, coinType common.CoinType) error

	// Gets an account by name
	GetAccount(name string) (wallet.Account, error)

	// Gets all accounts
	ListAccounts() ([]wallet.Account, error)
}

// `DIDControllerImpl` is a type that implements the `DIDController` interface.
// @property  - `wallet.Wallet`: This is the interface that the DID controller implements.
// @property  - `store.WalletStore`: This is the interface that the DID controller implements.
type DIDControllerImpl struct {
	store wallet.Store

	ctx context.Context
	aka string

	accounts       map[string]*wallet.AccountConfig
	didDocument    *crypto.DidDocument
	primaryAccount wallet.Account
}

// `New` creates a new DID controller instance
func New(ctx context.Context, account wallet.Account) (DIDController, error) {
	st, err := stores.New(account.Config())
	if err != nil {
		return nil, err
	}
	docc := &DIDControllerImpl{
		ctx:            ctx,
		primaryAccount: account,
		accounts:       make(map[string]*wallet.AccountConfig),
		store:          st,
	}
	// Create the DID document.
	doc, err := types.NewDocument(account.PubKey())
	if err != nil {
		return nil, err
	}
	docc.didDocument = doc

	return docc, nil
}

// Address returns the address of the DID controller.
func (d *DIDControllerImpl) Address() string {
	addr, err := d.primaryAccount.Config().Address()
	if err != nil {
		return ""
	}
	return addr
}

// ID returns the DID of the DID controller.
func (d *DIDControllerImpl) ID() string {
	return d.primaryAccount.DID()
}

// Document returns the DID document of the DID controller.
func (d *DIDControllerImpl) Document() *crypto.DidDocument {
	return d.didDocument
}

// // This method is used to get the challenge response from the DID controller.
// func (d *DIDControllerImpl) GetChallengeOptions(aka string) (*v1.ChallengeResponse, error) {
// 	return nil, nil
// }

// // This is the method that will be called when the user clicks on the "Register" button.
// func (d *DIDControllerImpl) RegisterAuthenticationCredential(credentialCreationData string) (*v1.RegisterResponse, error) {
// 	return nil, nil
// }

// // This method is used to get the options for the assertion.
// func (d *DIDControllerImpl) GetAssertionOptions(aka string) (*v1.AssertResponse, error) {
// 	return nil, nil
// }

// // This is the method that will be called when the user clicks the "Login" button on the login page.
// func (d *DIDControllerImpl) AuthorizeCredential(credentialRequestData string) (*v1.LoginResponse, error) {
// 	return nil, nil
// }

// Creating a new account.
func (w *DIDControllerImpl) CreateAccount(name string, coinType common.CoinType) error {
	acc, err := w.primaryAccount.Bip32Derive(name, coinType)
	if err != nil {
		return err
	}
	w.accounts[name] = acc.Config()

	addr, err := acc.PubKey().Bech32(acc.Config().CoinType().AddrPrefix())
	if err != nil {
		return err
	}
	err = w.didDocument.SetAssertion(acc.PubKey(), types.WithBlockchainAccount(addr), types.WithController(w.didDocument.Id), types.WithIDFragmentSuffix(acc.Config().Name))
	if err != nil {
		return err
	}
	return nil
}

// Returning the account.WalletAccount and error.
func (w *DIDControllerImpl) GetAccount(name string) (wallet.Account, error) {
	accConf, ok := w.accounts[strings.ToLower(name)]
	if !ok {
		return nil, errors.New("Account not found")
	}
	return accounts.Load(accConf)
}

// Returning a list of accounts.
func (w *DIDControllerImpl) ListAccounts() ([]wallet.Account, error) {
	accs := make([]wallet.Account, 0, len(w.accounts))
	for _, accConf := range w.accounts {
		acc, err := accounts.Load(accConf)
		if err != nil {
			return nil, err
		}
		accs = append(accs, acc)
	}
	return accs, nil
}
