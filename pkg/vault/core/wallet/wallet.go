package wallet

import (
	"errors"

	"github.com/sonr-hq/sonr/pkg/vault/core/account"
	v1 "github.com/sonr-hq/sonr/pkg/vault/types/v1"
)

type Wallet interface {
	// The wallet configuration
	WalletConfig() *v1.WalletConfig

	// Creates a new account
	CreateAccount(name string, addrPrefix string, networkName string) error

	// Gets an account by name
	GetAccount(name string) (account.WalletAccount, error)

	// Gets all accounts
	ListAccounts() ([]account.WalletAccount, error)
}

type walletImpl struct {
	// The wallet configuration
	walletConfig *v1.WalletConfig
}

func NewWalletFromConfig(walletConf *v1.WalletConfig) (Wallet, error) {
	return &walletImpl{
		walletConfig: walletConf,
	}, nil
}

func (w *walletImpl) WalletConfig() *v1.WalletConfig {
	return w.walletConfig
}

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

func (w *walletImpl) GetAccount(name string) (account.WalletAccount, error) {
	accConf, ok := w.walletConfig.Accounts[name]
	if !ok {
		return nil, errors.New("Account not found")
	}
	return account.NewAccountFromConfig(accConf)
}

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
