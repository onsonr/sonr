package middleware

import (
	"context"
	"fmt"

	ipfs_path "github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/client/rpc"
	"github.com/ipfs/kubo/core/coreiface/options"
	"github.com/labstack/echo/v4"

	modulev1 "github.com/sonrhq/sonr/api/sonr/identity/module/v1"
	"github.com/sonrhq/sonr/internal/wallet"
	"github.com/sonrhq/sonr/pkg/shared"
)

var Vault = &vault{}

type vault struct {
	echo.Context
}

// GenerateKey generates a new key
func (v *vault) GenerateKey(c echo.Context) error {
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
func (v *vault) GenerateIdentity(c echo.Context) error {
	ipfsC, err := rpc.NewLocalApi()
	if err != nil {
		return shared.ErrFailedIPFSClient
	}
	dir, kc, err := wallet.New(c.Request().Context())
	if err != nil {
		return err
	}
	key, err := ipfsC.Key().Generate(context.Background(), kc.Address, options.Key.Type(options.Ed25519Key))
	if err != nil {
		return fmt.Errorf("failed to generate key: %w", err)
	}
	keyIDAssociatedBytes, err := key.ID().MarshalBinary()
	if err != nil {
		return err
	}
	encDir, err := kc.Encrypt(dir, keyIDAssociatedBytes)
	if err != nil {
		return err
	}
	path, err := ipfsC.Unixfs().Add(context.Background(), encDir)
	if err != nil {
		return err
	}
	name, err := ipfsC.Name().Publish(context.Background(), path, options.Name.Key(key.ID().String()))
	if err != nil {
		return err
	}
	cnt := &modulev1.Controller{
		Address:   kc.Address,
		PeerId:    key.ID().String(),
		PublicKey: kc.PublicKey,
		Ipns:      name.String(),
	}
	return c.JSON(200, cnt)
}

// PublishFile publishes a file
func (v *vault) PublishFile(c echo.Context) error {
	keyID := c.Param("keyID")
	ipfsPath := c.Param("cid")
	path, err := ipfs_path.NewPath(ipfsPath)
	if err != nil {
		return fmt.Errorf("failed to create path: %w", err)
	}
	ipfsC, err := rpc.NewLocalApi()
	if err != nil {
		return shared.ErrFailedIPFSClient
	}
	name, err := ipfsC.Name().Publish(c.Request().Context(), path, options.Name.Key(keyID))
	if err != nil {
		return fmt.Errorf("failed to publish file: %w", err)
	}
	return c.JSON(200, name)
}

// GetFile gets a file
func (v *vault) GetFile(c echo.Context) error {
	path := c.Param("cid")
	ipfsC, err := rpc.NewLocalApi()
	if err != nil {
		return shared.ErrFailedIPFSClient
	}
	cid, err := ipfs_path.NewPath(path)
	if err != nil {
		return fmt.Errorf("failed to create path: %w", err)
	}
	file, err := ipfsC.Unixfs().Get(c.Request().Context(), cid)
	if err != nil {
		return fmt.Errorf("failed to get file: %w", err)
	}
	return c.JSON(200, file)
}
