package vault

import (
	"context"

	"github.com/ipfs/kubo/client/rpc"
	iface "github.com/ipfs/kubo/core/coreiface"
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
	SonrPubKey  []byte
	PeerID      peer.ID
}

func New(ctx context.Context) (*Vault, error) {
	c := getIpfsClient()
	kc, err := keychain.New(ctx)
	if err != nil {
		return nil, err
	}
	key, err := c.Key().Generate(ctx, kc.Address)
	if err != nil {
		return nil, err
	}

	return &Vault{
		Key:         key,
		SonrAddress: kc.Address,
		SonrPubKey:  kc.PublicKey,
		PeerID:      key.ID(),
	}, nil
}
