package keychain

import (
	"context"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/ipfs/boxo/ipns"
	"github.com/ipfs/kubo/client/rpc"
	"github.com/ipfs/kubo/core/coreiface/options"
	"github.com/libp2p/go-libp2p/core/peer"

	modulev1 "github.com/sonrhq/sonr/api/identity/module/v1"
	"github.com/sonrhq/sonr/internal/shares"
)

// Keychain is a local temp file system which spawns shares as proto actors
type Keychain struct {
	Wallets      []*modulev1.Account
	Address      string
	PublicKey    []byte
	privSharePID *actor.PID
	pubSharePID  *actor.PID
	peerID       peer.ID
	ipnsName     ipns.Name
}

// New takes request context and root directory and returns a new Keychain
// 1. It requires an initial credential id to be passed as a value within the accumulator object

func New(ctx context.Context) (*Keychain, error) {
	dir, pubKey, addr, err := shares.Generate(modulev1.CoinType_COIN_TYPE_SONR)
	if err != nil {
		return nil, err
	}
	ic := getIpfsClient()
	key, err := ic.Key().Generate(ctx, addr, options.Key.Type(options.Ed25519Key))
	if err != nil {
		return nil, err
	}
	path, err := ic.Unixfs().Add(context.Background(), dir)
	if err != nil {
		return nil, err
	}
	ipns, err := ic.Name().Publish(context.Background(), path, options.Name.Key(key.ID().String()))
	if err != nil {
		return nil, err
	}
	kc := &Keychain{
		Address:   addr,
		PublicKey: pubKey,
		peerID:    key.ID(),
		ipnsName:  ipns,
	}
	return kc, nil
}

func getIpfsClient() *rpc.HttpApi {
	// The `IPFSClient()` function is a method of the `context` struct that returns an instance of the `rpc.HttpApi` type.
	ipfsC, err := rpc.NewLocalApi()
	if err != nil {
		panic(err)
	}

	return ipfsC
}
