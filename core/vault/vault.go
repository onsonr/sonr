package vault

import (
	"context"
	"fmt"
	"io/ioutil"
	"sync"

	crypto "github.com/libp2p/go-libp2p/core/crypto"
	peer "github.com/libp2p/go-libp2p/core/peer"
	"github.com/sonr-hq/sonr/core/vault/x/mpc"
	"github.com/sonr-hq/sonr/internal/node"
	"github.com/sonr-hq/sonr/pkg/wallet"
	"github.com/sonr-io/multi-party-sig/pkg/party"
	"github.com/sonr-io/multi-party-sig/pkg/pool"
	tm_crypto "github.com/tendermint/tendermint/crypto"
	tm_json "github.com/tendermint/tendermint/libs/json"
)

var (
	v *Vault
)

func loadPrivKeyFromJsonPath(path string) (crypto.PrivKey, error) {
	// Load the key from the given path.
	key, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// Create new private key interface
	var vnPk tm_crypto.PrivKey

	// Unmarshal the key into the interface.
	err = tm_json.Unmarshal(key, &vnPk)
	if err != nil {
		return nil, err
	}
	priv, err := crypto.UnmarshalPrivateKey(vnPk.Bytes())
	if err != nil {
		return nil, err
	}
	return priv, nil
}

type Vault struct {
	nodeKeyPath string
	nodePrivKey crypto.PrivKey
	p2pNode     *node.Node
	ctx         context.Context
}

func InitVault(path string) error {
	ctx := context.Background()
	key, err := loadPrivKeyFromJsonPath(path)
	if err != nil {
		return err
	}
	n, err := node.New(ctx)
	v = &Vault{
		nodeKeyPath: path,
		nodePrivKey: key,
		p2pNode:     n,
		ctx:         ctx,
	}
	return nil
}

func (v *Vault) GenerateWallet(id peer.ID, threshold int) (wallet.WalletShare, error) {
	var wg sync.WaitGroup
	pl := pool.NewPool(0)
	defer pl.TearDown()
	ids := party.IDSlice{
		party.ID(v.p2pNode.ID()),
		party.ID(id),
	}
	name := fmt.Sprintf("/sonr/v0.2.0/mpc/keygen/%s-%s", ids[0], ids[1])
	th, err := v.p2pNode.Subscribe(name)
	if err != nil {
		panic(err)
	}
	conf, err := mpc.CmpKeygen(party.ID(id), ids, th, threshold, &wg, pl)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
