package vault

import (
	"context"

	"github.com/ipfs/kubo/core/coreiface/options"

	modulev1 "github.com/sonrhq/sonr/api/sonr/identity/module/v1"
	snrctx "github.com/sonrhq/sonr/internal/context"
	"github.com/sonrhq/sonr/internal/wallet"
)

// Create takes request context and root directory and returns a new Root Identity Controller
func Create(ctx context.Context) (*modulev1.Controller, error) {
	c := snrctx.GetIpfsClient()
	dir, kc, err := wallet.New(ctx)
	if err != nil {
		return nil, err
	}
	key, err := c.Key().Generate(context.Background(), kc.Address, options.Key.Type(options.Ed25519Key))
	if err != nil {
		return nil, err
	}
	keyIDAssociatedBytes, err := key.ID().MarshalBinary()
	if err != nil {
		return nil, err
	}
	encDir, err := kc.Encrypt(dir, keyIDAssociatedBytes)
	if err != nil {
		return nil, err
	}
	path, err := c.Unixfs().Add(context.Background(), encDir)
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
