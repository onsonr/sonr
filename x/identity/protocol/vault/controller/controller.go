package controller

import (
	"context"
	"errors"
	"fmt"

	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/pkg/crypto/wallet"
	"github.com/sonrhq/core/pkg/crypto/wallet/accounts"
	"github.com/sonrhq/core/pkg/crypto/wallet/stores"
	"github.com/sonrhq/core/x/identity/types"
)

// `DIDControllerImpl` is a type that implements the `DIDController` interface.
// @property  - `wallet.Wallet`: This is the interface that the DID controller implements.
// @property  - `store.WalletStore`: This is the interface that the DID controller implements.
type DIDControllerImpl struct {
	store wallet.Store

	ctx context.Context
	aka string

	didDocument    *types.DidDocument
	primaryAccount wallet.Account
	authentication *types.VerificationMethod
}

// `New` creates a new DID controller instance
func New(account wallet.Account, opts ...stores.Option) (DIDController, error) {
	if account == nil {
		return nil, errors.New("account is nil")
	}
	// Create the wallet store.
	st, err := stores.New(account, opts...)
	if err != nil {
		return nil, err
	}
	docc := &DIDControllerImpl{
		ctx:            context.Background(),
		primaryAccount: account,
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
	return d.primaryAccount.Address()
}

// ID returns the DID of the DID controller.
func (d *DIDControllerImpl) ID() string {
	return d.primaryAccount.DID()
}

// Document returns the DID document of the DID controller.
func (d *DIDControllerImpl) Document() *types.DidDocument {
	return d.didDocument
}

// Creating a new account.
func (w *DIDControllerImpl) CreateAccount(name string, coinType crypto.CoinType) (*types.VerificationMethod, error) {
	acc, err := w.primaryAccount.Bip32Derive(name, coinType)
	if err != nil {
		return nil, err
	}
	// Set account in list
	err = w.store.PutAccount(acc, name)
	if err != nil {
		return nil, err
	}

	vm, err := w.didDocument.SetAssertion(acc.PubKey(), types.WithBlockchainAccount(acc.Address()),
		types.WithController(w.didDocument.Id),
		types.WithIDFragmentSuffix(acc.Config().Name),
	)
	if err != nil {
		return nil, err
	}
	vm.SetMetadataValue(kDIDMetadataKeyAccName, acc.Name())
	vm.SetMetadataValue(kDIDMetadataKeyCoin, acc.CoinType().Name())
	w.didDocument.UpdateAssertion(vm)
	return vm, nil
}

// Returning the account.WalletAccount and error.
func (w *DIDControllerImpl) GetAccount(name string) (wallet.Account, error) {
	accConf, err := w.store.GetAccount(name)
	if err != nil {
		return nil, err
	}
	return accConf, nil
}

// Get Sonr account
func (w *DIDControllerImpl) GetSonrAccount() (wallet.CosmosAccount, error) {
	return accounts.GetCosmosAccount(w.primaryAccount, w.primaryAccount, nil), nil
}

// ListAccounts returns the list of accounts.
func (w *DIDControllerImpl) ListAccounts() ([]wallet.Account, error) {
	vms := w.didDocument.ListBlockchainAccounts()
	if len(vms) == 0 {
		return nil, fmt.Errorf("no accounts found")
	}
	accounts := make([]wallet.Account, 0, len(vms))
	for _, vm := range vms {
		name, ok := vm.GetMetadataValue(kDIDMetadataKeyAccName)
		if !ok {
			return nil, fmt.Errorf("account name not found in metadata")
		}
		acc, err := w.store.GetAccount(name)
		if err != nil {
			return nil, fmt.Errorf("failed to get account %s: %w", name, err)
		}
		accounts = append(accounts, acc)
	}
	return accounts, nil
}

// Sign signs the data with the given account.
func (w *DIDControllerImpl) Sign(data []byte) ([]byte, error) {
	return w.primaryAccount.Sign(data)
}

// Verify verifies the signature with the given account.
func (w *DIDControllerImpl) Verify(data, sig []byte) (bool, error) {
	return w.primaryAccount.Verify(data, sig)
}
