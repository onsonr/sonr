package vault

import (
	"fmt"

	ipfs_path "github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/client/rpc"
	"github.com/ipfs/kubo/core/coreiface/options"

	"github.com/sonrhq/sonr/pkg/shared"
)

// PublishFile publishes a file
func (c *vault) PublishFile() error {
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
func (c *vault) GetFile() error {
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
