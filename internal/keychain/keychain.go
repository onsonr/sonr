package keychain

import (
	"context"
	"fmt"
	"os"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/ipfs/kubo/client/rpc"
	"github.com/libp2p/go-libp2p/core/peer"

	modulev1 "github.com/sonrhq/sonr/api/identity/module/v1"
	"github.com/sonrhq/sonr/internal/shares"
)

// Keychain is a local temp file system which spawns shares as proto actors
type Keychain struct {
	RootDir      string
	Wallets      []*modulev1.Account
	Address      string
	PublicKey    []byte
	privSharePID *actor.PID
	pubSharePID  *actor.PID
	peerID       peer.ID
}

// New takes request context and root directory and returns a new Keychain
// 1. It requires an initial credential id to be passed as a value within the accumulator object

func New(ctx context.Context) (*Keychain, error) {
	rootDir, err := os.MkdirTemp("", "sonr-keychain")
	if err != nil {
		return nil, err
	}
	dir, pubKey, addr, err := shares.Generate(rootDir, modulev1.CoinType_COIN_TYPE_SONR)
	if err != nil {
		return nil, err
	}
	ic := getIpfsClient()
	key, err := ic.Key().Generate(ctx, addr)
	if err != nil {
		return nil, err
	}
	path, err := ic.Unixfs().Add(ctx, dir)
	if err != nil {
		return nil, err
	}
	err = ic.Pin().Add(ctx, path)
	if err != nil {
		return nil, err
	}
	ipns, err := ic.Name().Publish(ctx, path)
	if err != nil {
		return nil, err
	}
	fmt.Println(ipns)
	kc := &Keychain{
		Address:   addr,
		PublicKey: pubKey,
		RootDir:   rootDir,
		peerID:    key.ID(),
	}

	return kc, nil
}

// Burn removes the root directory of the keychain
func (kc *Keychain) Burn() error {
	return os.RemoveAll(kc.RootDir)
}

func getIpfsClient() *rpc.HttpApi {
	// The `IPFSClient()` function is a method of the `context` struct that returns an instance of the `rpc.HttpApi` type.
	ipfsC, err := rpc.NewLocalApi()
	if err != nil {
		panic(err)
	}

	return ipfsC
}
