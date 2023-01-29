package internal

import (
	"fmt"
	"sync"
	"time"

	"github.com/sonrhq/core/pkg/common"
	"github.com/sonrhq/core/pkg/crypto/wallet"
	"github.com/sonrhq/core/pkg/crypto/wallet/accounts/internal/mpc"
	"github.com/sonrhq/core/pkg/crypto/wallet/accounts/internal/network"
	"github.com/sonrhq/core/pkg/crypto/wallet/accounts/internal/token"
	v1 "github.com/sonrhq/core/x/identity/types/vault/v1"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
	"github.com/ucan-wg/go-ucan"
)

// BaseAccountFromConfig returns a new base account
func BaseAccountFromConfig(conf *v1.AccountConfig, rootCnf *cmp.Config) wallet.Account {
	return &baseAccountImpl{
		accountConfig: conf,
	}
}

// The baseAccountImpl type is a struct that has a single field, accountConfig, which is a pointer to
// a v1.AccountConfig.
// @property accountConfig - The account configuration
type baseAccountImpl struct {
	// The account configuration
	accountConfig *v1.AccountConfig
}

// Deriving a new account from a BIP32 path.
func (w *baseAccountImpl) Bip32Derive(name string, coinType common.CoinType) (wallet.Account, error) {
	newCnf, err := w.rootCmpConf().DeriveBIP32(uint32(coinType.Index()))
	if err != nil {
		return nil, err
	}
	deri, err := v1.NewDerivedAccountConfig(name, coinType, newCnf)
	if err != nil {
		return nil, err
	}
	return BaseAccountFromConfig(deri, newCnf), nil
}

// CoinType returns the account coin type
func (w *baseAccountImpl) CoinType() common.CoinType {
	return w.accountConfig.CoinType()
}

// Config returns the account configuration
func (w *baseAccountImpl) Config() *v1.AccountConfig {
	return w.accountConfig
}

// DID returns the account DID
func (w *baseAccountImpl) DID() string {
	return w.accountConfig.DID()
}

// Info returns the account information
func (w *baseAccountImpl) Info() map[string]string {
	addr, _ := w.accountConfig.Address()

	return map[string]string{
		"name":    w.accountConfig.Name,
		"network": w.accountConfig.CoinType().Name(),
		"address": addr,
	}
}

// Marshal returns the local config protobuf bytes
func (w *baseAccountImpl) Marshal() ([]byte, error) {
	return w.accountConfig.Marshal()
}

// NewOriginToken returns a new origin token for the account.
func (w *baseAccountImpl) NewOriginToken(audienceDID string, att ucan.Attenuations, fct []ucan.Fact, notBefore, expires time.Time) (string, error) {
	return token.NewUnsignedUCAN(w.accountConfig, audienceDID, nil, att, fct, notBefore, expires)
}

// NewAttenuatedToken returns a new attenuated token for the account.
func (w *baseAccountImpl) NewAttenuatedToken(parent *ucan.Token, audienceDID string, att ucan.Attenuations, fct []ucan.Fact, notBefore, expires time.Time) (string, error) {
	if !parent.Attenuations.Contains(att) {
		return "", fmt.Errorf("scope of ucan attenuations must be less than it's parent")
	}
	return token.NewUnsignedUCAN(w.accountConfig, audienceDID, append(parent.Proofs, ucan.Proof(parent.Raw)), att, fct, notBefore, expires)
}

// PubKey returns secp256k1 public key
func (w *baseAccountImpl) PubKey() common.SNRPubKey {
	pbKey, _ := w.accountConfig.PubKey()
	return pbKey
}

// Signing a transaction.
func (w *baseAccountImpl) Sign(msg []byte) ([]byte, error) {
	net := network.NewOfflineNetwork(w.accountConfig.PartyIDs())
	configs, err := v1.DeserializeConfigList(w.accountConfig.Shares)
	if err != nil {
		return nil, err
	}

	doneChan := make(chan []byte, 1)
	var wg sync.WaitGroup
	for _, c := range configs {
		wg.Add(1)
		go func(conf *cmp.Config) {
			pl := pool.NewPool(0)
			defer pl.TearDown()
			sig, err := mpc.CmpSign(conf, msg, net.Ls(), net, &wg, pl)
			if err != nil {
				return
			}
			if conf.ID == "current" {
				doneChan <- sig
			}
		}(c)
	}
	wg.Wait()
	return <-doneChan, nil
}

// Type returns the account type
func (w *baseAccountImpl) Type() string {
	return fmt.Sprintf("%s/ecdsa-secp256k1", w.accountConfig.CoinType().Name())
}

// Unmarshal returns the local config protobuf bytes
func (w *baseAccountImpl) Unmarshal(bz []byte) error {
	conf := v1.AccountConfig{}
	if err := conf.Unmarshal(bz); err != nil {
		return err
	}
	w.accountConfig = &conf
	return nil
}

// Verifying a signature.
func (w *baseAccountImpl) Verify(bz []byte, sig []byte) (bool, error) {
	pubKey, err := w.accountConfig.PubKey()
	if err != nil {
		return false, err
	}
	return pubKey.VerifySignature(bz, sig), nil
}

func (w *baseAccountImpl) rootCmpConf() *cmp.Config {
	cmps, err := v1.DeserializeConfigList(w.accountConfig.Shares)
	if err != nil {
		panic(err)
	}
	for _, c := range cmps {
		if c.ID == "current" {
			return c
		}
	}
	panic("no current config")
}
