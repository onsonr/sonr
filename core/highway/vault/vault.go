package vault

import (
	"context"
	"fmt"
	"io/ioutil"

	crypto "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/sonr-hq/sonr/core/highway/vault/internal"
	v1 "github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/node"
	"github.com/sonr-hq/sonr/pkg/wallet"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	mpc "github.com/taurusgroup/multi-party-sig/pkg/protocol"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
	tm_crypto "github.com/tendermint/tendermint/crypto"
	tm_json "github.com/tendermint/tendermint/libs/json"
)

var (
	v *Vault
)

// > Load the key from the given path, unmarshal the key into the interface, and return the private key
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

// > This function creates a new wallet for two parties
func NewTwoPartyWallet() {

}

// `NewMultiWallet()` creates a new MultiWallet
func NewMultiWallet(motor *node.Node) (wallet.WalletShare, error) {
	fmt.Println("Creating new multi wallet")
	pl := pool.NewPool(0)
	err := motor.Connect(v.P2P.MultiAddr())
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to motor")

	peerIds := []peer.ID{motor.ID(), v.P2P.ID()}
	mtrSession, err := internal.NewMultiSession(motor, peerIds)
	if err != nil {
		return nil, err
	}
	vaultSession, err := internal.NewMultiSession(v.P2P, peerIds)
	if err != nil {
		return nil, err
	}
	fmt.Println("Created sessions")

	vaultHand, err := mpc.NewMultiHandler(cmp.Keygen(curve.Secp256k1{}, vaultSession.SelfID(), vaultSession.PartyIds(), 1, pl), nil)
	if err != nil {
		fmt.Println("Error creating handler")
		return nil, err
	}
	 vaultSession.RunProtocol(vaultHand)

	mtrHand, err := mpc.NewMultiHandler(cmp.Keygen(curve.Secp256k1{}, mtrSession.SelfID(), mtrSession.PartyIds(), 1, pl), nil)
	if err != nil {
		fmt.Println("Error creating handler")
		return nil, err
	}
	 mtrSession.RunProtocol(mtrHand)

	w, err := mtrHand.Result()
	if err != nil {
		return nil, err
	}
	return wallet.NewWalletImpl(w), nil
}

// RunSignProtocol runs the sign protocol.
func RunSignProtocol() {

}

// RunRefreshProtocol() is a function that does nothing.
func RunRefreshProtocol() {
}
