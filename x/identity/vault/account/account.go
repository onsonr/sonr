package account

import (
	"errors"
	"fmt"
	"sync"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/shengdoushi/base58"
	"github.com/sonr-hq/sonr/x/identity/types"
	"github.com/sonr-hq/sonr/x/identity/vault/account/internal/mpc"
	"github.com/sonr-hq/sonr/x/identity/vault/account/internal/network"

	v1 "github.com/sonr-hq/sonr/x/identity/types/vault/v1"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

// `WalletAccount` is an interface that defines the methods that a wallet account must implement.
// @property AccountConfig - The account configuration
// @property Bip32Derive - This is a method that derives a new account from a BIP32 path.
// @property GetAssertionMethod - returns the verification method for the account.
// @property {bool} IsPrimary - returns true if the account is the primary account
// @property ListConfigs - This is a list of all the configurations that are needed to sign a
// transaction.
// @property Sign - This is the function that signs a transaction.
// @property Verify - Verifies a signature
type WalletAccount interface {
	// The account configuration
	AccountConfig() *v1.AccountConfig

	// Bip32Derive derives a new account from a BIP32 path
	Bip32Derive(accName string, idx uint32, addrPrefix string, network string) (WalletAccount, error)

	// GetAssertionMethod returns the verification method for the account
	GetAssertionMethod() *types.VerificationMethod

	// IsPrimary returns true if the account is the primary account
	IsPrimary() bool

	// ListConfigs returns the list of all the configurations that are needed to
	// sign a transaction.
	ListConfigs() ([]*cmp.Config, error)

	// PubKey returns secp256k1 public key
	PubKey() (*secp256k1.PubKey, error)

	// Signs a transaction
	Sign(bz []byte) ([]byte, error)

	// Verifies a signature
	Verify(bz []byte, sig []byte) (bool, error)
}

// The walletAccountImpl type is a struct that has a single field, accountConfig, which is a pointer to
// a v1.AccountConfig.
// @property accountConfig - The account configuration
type walletAccountImpl struct {
	// The account configuration
	accountConfig *v1.AccountConfig
}

// It creates a new account with the given name, address prefix, and network name
func NewAccount(accName string, addrPrefix string, networkName string) (WalletAccount, error) {
	// The default shards that are added to the MPC wallet
	parties := party.IDSlice{"vault", "current"}
	net := network.NewOfflineNetwork(parties)
	accConf, err := mpc.Keygen(accName, "current", 1, net, addrPrefix, networkName)
	if err != nil {
		return nil, err
	}
	return &walletAccountImpl{
		accountConfig: accConf,
	}, nil
}

// > This function creates a new wallet account from the given account configuration
func NewAccountFromConfig(accConf *v1.AccountConfig) (WalletAccount, error) {
	return &walletAccountImpl{
		accountConfig: accConf,
	}, nil
}

// It returns the account configuration.
func (w *walletAccountImpl) AccountConfig() *v1.AccountConfig {
	return w.accountConfig
}

// Deriving a new account from a BIP32 path.
func (w *walletAccountImpl) Bip32Derive(accName string, idx uint32, addrPrefix string, network string) (WalletAccount, error) {
	if !w.IsPrimary() {
		return nil, errors.New("cannot derive from non-primary account")
	}
	oldConfs, err := w.ListConfigs()
	if err != nil {
		return nil, err
	}
	shares := make([]*v1.ShareConfig, 0)
	for _, conf := range oldConfs {
		c, err := conf.DeriveBIP32(idx)
		if err != nil {
			return nil, err
		}
		shares = append(shares, v1.NewShareConfig(network, c))
	}
	accConf, err := v1.NewAccountConfigFromShares(accName, 0, addrPrefix, shares)
	if err != nil {
		return nil, err
	}
	return NewAccountFromConfig(accConf)
}

// Returning the verification method for the account.
func (w *walletAccountImpl) GetAssertionMethod() *types.VerificationMethod {
	return &types.VerificationMethod{
		ID:                  fmt.Sprintf("%s#%s", types.ConvertAccAddressToDid(w.accountConfig.Address), w.accountConfig.Name),
		Type:                types.KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019,
		Controller:          types.ConvertAccAddressToDid(w.accountConfig.Address),
		PublicKeyMultibase:  base58.Encode(w.accountConfig.PublicKey, base58.BitcoinAlphabet),
		BlockchainAccountId: w.accountConfig.Address,
		Metadata: []*types.KeyValuePair{
			{
				Key:   "label",
				Value: w.accountConfig.Name,
			},
			{
				Key:   "network",
				Value: w.accountConfig.Shares[0].Network,
			},
			{
				Key:   "algorithm",
				Value: "cmp",
			},
		},
	}
}

// It returns true if the account is the primary account.
func (w *walletAccountImpl) IsPrimary() bool {
	return w.accountConfig.Name == "primary"
}

// Returning the list of all the configurations that are needed to sign a transaction.
func (w *walletAccountImpl) ListConfigs() ([]*cmp.Config, error) {
	confMap := w.accountConfig.GetConfigMap()
	configs := make([]*cmp.Config, 0, len(confMap))
	for _, conf := range confMap {
		configs = append(configs, conf)
	}
	return configs, nil
}

// Returning the secp256k1 public key.
func (w *walletAccountImpl) PubKey() (*secp256k1.PubKey, error) {
	return w.accountConfig.Shares[0].GetPubKeySecp256k1()
}

// Signing a transaction.
func (w *walletAccountImpl) Sign(bz []byte) ([]byte, error) {
	return signWithAccount(w.accountConfig, bz)
}

// Verifying a signature.
func (w *walletAccountImpl) Verify(bz []byte, sig []byte) (bool, error) {
	conf := w.accountConfig.GetConfigMap()
	return mpc.CmpVerify(conf["current"], bz, sig)
}

// signWithAccount signs a message with the given account configuration
func signWithAccount(a *v1.AccountConfig, msg []byte) ([]byte, error) {
	net := network.NewOfflineNetwork(a.PartyIDs())
	configs := a.GetConfigMap()
	doneChan := make(chan []byte, 1)
	var wg sync.WaitGroup
	for _, id := range net.Ls() {
		wg.Add(1)
		go func(id party.ID) {
			pl := pool.NewPool(0)
			defer pl.TearDown()
			sig, err := mpc.CmpSign(configs[id], msg, net.Ls(), net, &wg, pl)
			if err != nil {
				return
			}
			if id == "current" {
				doneChan <- sig
			}
		}(id)
	}
	wg.Wait()
	return <-doneChan, nil
}
