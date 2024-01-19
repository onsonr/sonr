package vault

import (
	"context"

	"github.com/ipfs/kubo/client/rpc"
	"github.com/ipfs/kubo/core/coreiface/options"

	modulev1 "github.com/sonrhq/sonr/api/identity/module/v1"
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

func NewController(ctx context.Context) (*modulev1.Controller, error) {
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
	cnt := &modulev1.Controller{
		Address:   kc.Address,
		PeerId:    key.ID().String(),
		PublicKey: kc.PublicKey,
		Ipns:      name.String(),
	}
	return cnt, nil
}
