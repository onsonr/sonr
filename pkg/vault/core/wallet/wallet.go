package wallet

import (
	"errors"

	"github.com/sonr-hq/sonr/pkg/vault/core/account"
	v1 "github.com/sonr-hq/sonr/pkg/vault/types/v1"
)

// `Wallet` is an interface that has a method `WalletConfig` that returns a `*v1.WalletConfig` and a
// method `CreateAccount` that takes a `string`, a `string`, and a `string` and returns an `error`.
// @property WalletConfig - This is the configuration of the wallet.
// @property {error} CreateAccount - Creates a new account
// @property GetAccount - Returns the account with the given name.
// @property PrimaryAccount - The primary account is the account that is used to sign transactions.
// @property ListAccounts - Returns a list of all accounts in the wallet
type Wallet interface {
	// The wallet configuration
	WalletConfig() *v1.WalletConfig

	// Creates a new account
	CreateAccount(name string, addrPrefix string, networkName string) error

	// Gets an account by name
	GetAccount(name string) (account.WalletAccount, error)

	// Gets Primary account
	PrimaryAccount() (account.WalletAccount, error)

	// Gets all accounts
	ListAccounts() ([]account.WalletAccount, error)
}

// `walletImpl` is a struct that has a single field, `walletConfig`, which is a pointer to a
// `v1.WalletConfig` struct.
// @property walletConfig - The wallet configuration
type walletImpl struct {
	// The wallet configuration
	walletConfig *v1.WalletConfig
}

// `NewWalletFromConfig` takes a `WalletConfig` and returns a `Wallet` and an error
func NewWalletFromConfig(walletConf *v1.WalletConfig) (Wallet, error) {
	return &walletImpl{
		walletConfig: walletConf,
	}, nil
}

// Returning the wallet configuration.
func (w *walletImpl) WalletConfig() *v1.WalletConfig {
	return w.walletConfig
}

// Creating a new account.
func (w *walletImpl) CreateAccount(name string, addrPrefix string, networkName string) error {
	// The default shards that are added to the MPC wallet
	rootAcc, err := w.GetAccount("Primary")
	if err != nil {
		return err
	}
	acc, err := rootAcc.Bip32Derive(name, uint32(len(w.walletConfig.Accounts)), addrPrefix, networkName)
	if err != nil {
		return err
	}
	w.walletConfig.Accounts[name] = acc.AccountConfig()
	return nil
}

// Returning the account.WalletAccount and error.
func (w *walletImpl) GetAccount(name string) (account.WalletAccount, error) {
	accConf, ok := w.walletConfig.Accounts[name]
	if !ok {
		return nil, errors.New("Account not found")
	}
	return account.NewAccountFromConfig(accConf)
}

// Returning a list of accounts.
func (w *walletImpl) ListAccounts() ([]account.WalletAccount, error) {
	accs := make([]account.WalletAccount, 0, len(w.walletConfig.Accounts))
	for _, accConf := range w.walletConfig.Accounts {
		acc, err := account.NewAccountFromConfig(accConf)
		if err != nil {
			return nil, err
		}
		accs = append(accs, acc)
	}
	return accs, nil
}

// Returning the primary account.
func (w *walletImpl) PrimaryAccount() (account.WalletAccount, error) {
	return w.GetAccount("Primary")
}

// `NewWallet` creates a new wallet with a default root account
func NewWallet() (Wallet, error) {
	// The default shards that are added to the MPC wallet
	rootAcc, err := account.NewAccount("Primary", "snr", "Sonr")
	if err != nil {
		return nil, err
	}
	conf := v1.NewWalletConfigFromRootAccount(rootAcc.AccountConfig())
	return &walletImpl{
		walletConfig: conf,
	}, nil
}
