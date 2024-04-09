package vault

import (
	"context"
	"fmt"

	"github.com/ipfs/kubo/client/rpc"
	"github.com/ipfs/kubo/core/coreiface/options"

	"github.com/didao-org/sonr/x/identity/types"
	"github.com/didao-org/sonr/x/identity/wallet"
)

// GenerateKey generates a new key
func (c *vault) GenerateKey() error {
	address := c.Param("address")
	ipfsC, err := rpc.NewLocalApi()
	if err != nil {
		return ErrFailedIPFSClient
	}
	key, err := ipfsC.Key().Generate(c.Request().Context(), address, options.Key.Type(options.Ed25519Key))
	if err != nil {
		return fmt.Errorf("failed to generate key: %w", err)
	}
	return c.JSON(200, key)
}

// GenerateIdentity generates a new fully scoped Sonr identity
func (c *vault) GenerateIdentity() (*types.IdentityController, error) {
	ipfsC, err := rpc.NewLocalApi()
	if err != nil {
		return nil, ErrFailedIPFSClient
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
	cnt := &types.IdentityController{
		Address:   kc.Address,
		PeerId:    key.ID().String(),
		PublicKey: kc.PublicKey,
		Ipns:      name.String(),
	}
	return cnt, nil
}
