package vault

import (
	"context"
	"io/ioutil"

	crypto "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/sonr-hq/sonr/pkg/node"
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
	P2P         *node.Node
	ctx         context.Context
}

func Initialize() error {
	ctx := context.Background()
	// key, err := loadPrivKeyFromJsonPath(path)
	// if err != nil {
	// 	return err
	// }3d
	n, err := node.New(ctx)
	if err != nil {
		return err
	}

	v = &Vault{
		//nodeKeyPath: path,
		//		nodePrivKey: key,
		P2P: n,
		ctx: ctx,
	}
	return nil
}

func NewTwoPartyWallet() {

}

func NewMultiWallet() {

}

func RunSignProtocol() {

}

func RunRefreshProtocol() {
}
