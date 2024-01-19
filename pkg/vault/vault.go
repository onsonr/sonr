package vault

import (
	"context"

	"github.com/ipfs/boxo/ipns"
	"github.com/ipfs/kubo/client/rpc"
	iface "github.com/ipfs/kubo/core/coreiface"
	"github.com/ipfs/kubo/core/coreiface/options"
	"github.com/libp2p/go-libp2p/core/peer"

	"github.com/sonrhq/sonr/internal/keychain"
)

func getIpfsClient() *rpc.HttpApi {
	// The `IPFSClient()` function is a method of the `context` struct that returns an instance of the `rpc.HttpApi` type.
	ipfsC, err := rpc.NewLocalApi()
	if err != nil {
		panic(err)
	}

	return ipfsC
}

type Vault struct {
	localPath   string
	Key         iface.Key
	SonrAddress string
	PeerID      peer.ID
	IPNS        ipns.Name
}

func New(ctx context.Context) (*Vault, error) {
	c := getIpfsClient()
	kc, err := keychain.New(ctx)
	if err != nil {
		return nil, err
	}
	key, err := c.Key().Generate(context.Background(), kc.Address, options.Key.Type(options.Ed25519Key))
	if err != nil {
		return nil, err
	}
	path, err := c.Unixfs().Add(context.Background(), kc.Directory)
	if err != nil {
		return nil, err
	}
	name, err := c.Name().Publish(context.Background(), path, options.Name.Key(key.ID().String()))
	if err != nil {
		return nil, err
	}
	return &Vault{
		Key:         key,
		SonrAddress: kc.Address,
		PeerID:      key.ID(),
		IPNS:        name,
	}, nil
}
