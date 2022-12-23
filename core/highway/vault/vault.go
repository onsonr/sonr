package vault

import (
	"context"
	"io/ioutil"

	crypto "github.com/libp2p/go-libp2p/core/crypto"
	v1 "github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/node"
	tm_crypto "github.com/tendermint/tendermint/crypto"
	tm_json "github.com/tendermint/tendermint/libs/json"
)

var (
	v *Vault
)

// A Vault is a node that has a private key and a P2P node.
// @property {string} nodeKeyPath - The path to the node's private key.
// @property nodePrivKey - The private key of the vault node.
// @property P2P - This is the node that will be used to connect to the network.
// @property ctx - The context.Context is a way to cancel the node.
type Vault struct {
	nodeKeyPath string
	nodePrivKey crypto.PrivKey
	P2P         *node.Node
	ctx         context.Context
}

// > Initialize() creates a new node and stores it in the Vault
func Initialize() error {
	ctx := context.Background()
	// key, err := loadPrivKeyFromJsonPath(path)
	// if err != nil {
	// 	return err
	// }3d
	n, err := node.New(ctx, node.WithPeerType(v1.Peer_HIGHWAY))
	if err != nil {
		return err
	}

	v = &Vault{
		P2P: n,
		ctx: ctx,
	}
	return nil
}

//
// Helper functions
//

// load the key from the given path, unmarshal the key into the interface, and return the private key
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
