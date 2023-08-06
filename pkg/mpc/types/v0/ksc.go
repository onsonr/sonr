package types

import (
	"fmt"
	"sync"

	"github.com/sonrhq/core/pkg/crypto"
	algo "github.com/sonrhq/core/pkg/mpc/protocol/cmp"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

type KeyShareCollection []*Keyshare

// NewKSS creates a new KeyShareCollection
func NewKSS(kss ...*Keyshare) KeyShareCollection {
	return KeyShareCollection(kss)
}

// NewKSSFromConfig creates a new KeyShareCollection from a config
func NewKSSFromCoin(coinType crypto.CoinType, configs ...*cmp.Config) KeyShareCollection {
	kslist := make([]*Keyshare, len(configs))
	for i, c := range configs {
		kslist[i] = CreateKeyshare(c, coinType)
	}
	return NewKSS(kslist...)
}

// Address returns the address of the account based on the coin type
func (a KeyShareCollection) Address() string {
	return a.Index(0).Address
}

// CoinType returns the coin type of the account
func (a KeyShareCollection) CoinType() crypto.CoinType {
	return a.Index(0).ParseCoinType()
}

// DID returns the DID of the account
func (wa KeyShareCollection) DID() string {
	did, _ := wa.CoinType().FormatDID(wa.PubKey())
	return did
}

// GetAccountData returns the proto representation of the account
func (wa KeyShareCollection) GetAccountData() *crypto.AccountData {
	dat, err := crypto.NewDefaultAccountData(wa.CoinType(), wa.PubKey())
	if err != nil {
		return nil
	}
	return dat
}

// Index returns the keyshare at the index
func (a KeyShareCollection) Index(i int) *Keyshare {
	return a[i]
}

// IsValid returns true if the keyshare collection is valid. A valid keyshare collection has at least 2 keyshares
func (a KeyShareCollection) IsValid() bool {
	return len(a) >= 2
}

// PubKey returns secp256k1 public key
func (kss KeyShareCollection) PubKey() *crypto.PubKey {
	return kss[0].PubKey()
}

// PubKey returns secp256k1 public key
func (kss KeyShareCollection) PubKeyType() string {
	return kss[0].PubKey().KeyType
}

// Signs a message using the account
func (kss KeyShareCollection) Sign(bz []byte) ([]byte, error) {
	var configs []*cmp.Config
	for _, ks := range kss {
		configs = append(configs, ks.ParseConfig())
	}
	peers := make([]crypto.PartyID, len(configs))
	for i, c := range configs {
		peers[i] = c.ID
	}
	net := algo.NewOfflineNetwork(peers...)
	doneChan := make(chan *crypto.MPCECDSASignature, 1)
	var wg sync.WaitGroup
	for _, c := range configs {
		wg.Add(1)
		go func(conf *cmp.Config) {
			pl := crypto.NewMPCPool(0)
			defer pl.TearDown()
			sig, err := algo.Sign(conf, bz, net.Ls(), net, &wg, pl)
			if err != nil {
				return
			}
			doneChan <- sig
		}(c)
	}
	wg.Wait()
	sig := <-doneChan
	return SerializeECDSASecp256k1Signature(sig)
}

// Type returns the type of the account
func (wa KeyShareCollection) Type() string {
	return fmt.Sprintf("%s/ecdsa-secp256k1", wa.CoinType().Name())
}

// Verify verifies a signature for a message using the topmost keyshare
func (kss KeyShareCollection) Verify(msg []byte, sig []byte) bool {
	return kss.Index(0).Verify(msg, sig)
}
