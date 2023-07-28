package mpc

import (
	"fmt"
	"sync"

	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/internal/mpc/base"
	models "github.com/sonrhq/core/internal/mpc/base"
	v0algo "github.com/sonrhq/core/internal/mpc/protocol/cmp"
	v1algo "github.com/sonrhq/core/internal/mpc/protocol/dkls"
	v0types "github.com/sonrhq/core/internal/mpc/types/v0"
	v1types "github.com/sonrhq/core/internal/mpc/types/v1"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

type AccountV1 = base.AccountV1
type KeyshareV0 = v0types.Keyshare
type KeyshareV1 = v1types.Keyshare
type KeyShareCollection = v0types.KeyShareCollection
type KeyshareSet = v1types.KeyshareSet
type EncKeyshareSet = v1types.EncKeyshareSet
type ZKSet = v1types.ZKSet

// The function GenerateV0 generates a collection of key shares with a given prefix, count, and coin type.
func GenerateV0(prefix string, count int, coinType crypto.CoinType) (KeyShareCollection, error) {
	getPartyID := func(i int) crypto.PartyID {
		return crypto.PartyID(fmt.Sprintf("%s-%d", prefix, i))
	}
	peers := make([]crypto.PartyID, count)
	for i := 0; i < count; i++ {
		peers[i] = getPartyID(i + 1)
	}
	cmps, err := KeygenV0(getPartyID(1), v0types.WithPartyIDs(peers...))
	if err != nil {
		return nil, err
	}
	return v0types.NewKSSFromCoin(coinType, cmps...), nil
}

// The function KeygenV0 generates cryptographic keys based on the provided options.
func KeygenV0(current crypto.PartyID, option ...v0types.KeygenOption) ([]*cmp.Config, error) {
	opts := v0types.DefaultKeygenOpts(current)
	opts.Apply(option...)
	net := opts.GetNetwork()
	var mtx sync.Mutex
	var wg sync.WaitGroup
	confs := make([]*cmp.Config, 0)
	for _, id := range net.Ls() {
		wg.Add(1)
		go func(id party.ID) {
			pl := pool.NewPool(0)
			defer pl.TearDown()
			conf, err := v0algo.Keygen(id, net.Ls(), net, opts.Threshold, &wg, pl)
			if err != nil {
				return
			}
			mtx.Lock()
			confs = append(confs, conf)
			mtx.Unlock()
		}(id)
	}
	wg.Wait()
	return confs, nil
}

// GenerateV2 generates a new account with a given ID.
func GenerateV2(name string, ct crypto.CoinType) (*models.AccountV1, KeyshareSet, error) {
	return models.NewAccountV1(name, ct)
}

// The function KeygenV1 generates a keyshare set.
func KeygenV1() (KeyshareSet, error) {
	kss, err := v1algo.DKLSKeygen()
	if err != nil {
		return v1types.EmptyKeyshareSet(), err
	}
	return kss, nil
}

// NewKSS creates a new keyshare set from a list of keyshares.
func NewKSS(pub *KeyshareV1, priv *KeyshareV1) KeyshareSet {
	return v1types.NewKSS(pub, priv)
}

// NewZKSet creates a new zero-knowledge set from a list of zero-knowledge proofs.
func NewZKSet(pubKey *crypto.Secp256k1PubKey, initialIds ...string) (ZKSet, error) {
	return v1types.CreateZkSet(pubKey, initialIds...)
}

// LoadZKSet loads a zero-knowledge set from a list of zero-knowledge proofs.
func LoadZKSet(str string) (ZKSet, error) {
	return v1types.OpenZkSet(str)
}
