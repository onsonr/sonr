package vault

import (
	"context"
	"fmt"

	"github.com/ipfs/kubo/client/rpc"
	"github.com/ipfs/kubo/core/coreiface/options"

	modulev1 "github.com/sonrhq/sonr/api/sonr/identity/module/v1"
	"github.com/sonrhq/sonr/internal/wallet"
	"github.com/sonrhq/sonr/pkg/middleware/shared"
)

// GenerateKey generates a new key
func (c *vault) GenerateKey() error {
	address := c.Param("address")
	ipfsC, err := rpc.NewLocalApi()
	if err != nil {
		return shared.ErrFailedIPFSClient
	}
	key, err := ipfsC.Key().Generate(c.Request().Context(), address, options.Key.Type(options.Ed25519Key))
	if err != nil {
		return fmt.Errorf("failed to generate key: %w", err)
	}
	return c.JSON(200, key)
}

// GenerateIdentity generates a new fully scoped Sonr identity
func (c *vault) GenerateIdentity() (*modulev1.Controller, error) {
	ipfsC, err := rpc.NewLocalApi()
	if err != nil {
		return nil, shared.ErrFailedIPFSClient
	}
	dir, kc, err := wallet.New(c.Request().Context())
	if err != nil {
		return nil, err
	}
	key, err := ipfsC.Key().Generate(context.Background(), kc.Address, options.Key.Type(options.Ed25519Key))
	if err != nil {
		return nil, fmt.Errorf("failed to generate key: %w", err)
	}
	keyIDAssociatedBytes, err := key.ID().MarshalBinary()
	if err != nil {
		return nil, err
	}
	encDir, err := kc.Encrypt(dir, keyIDAssociatedBytes)
	if err != nil {
		return nil, err
	}
	path, err := ipfsC.Unixfs().Add(context.Background(), encDir)
	if err != nil {
		return nil, err
	}
	name, err := ipfsC.Name().Publish(context.Background(), path, options.Name.Key(key.ID().String()))
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
