package fs

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/sonr-hq/sonr/pkg/common"
)

// Storing the share of the wallet.
func (c *Config) StoreShare(share []byte, partyId string) error {
	// Verify WalletConfigShare
	shareConfig := &common.WalletShareConfig{}
	err := shareConfig.Unmarshal(share)
	if err != nil {
		return err
	}

	// Create path for file to be stored and write file
	path := filepath.Join(c.localPath, "_auth", partyId)
	err = os.WriteFile(path, share, 0644)
	if err != nil {
		return err
	}

	// Add file to IPFS
	cid, err := c.ipfs.Unixfs().Add(c.ctx, c.rootNode, options.Unixfs.Pin(true))
	if err != nil {
		return err
	}
	c.ipfsPath = cid
	return nil
}

func (c *Config) LoadShares() ([]*common.WalletShareConfig, error) {
	shares := []*common.WalletShareConfig{}
	// List all files in the _auth directory
	files, err := os.ReadDir(filepath.Join(c.localPath, "_auth"))
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, errors.New("No shares found")
	}

	// Iterate over all files
	for _, file := range files {
		// Read file
		bz, err := os.ReadFile(filepath.Join(c.localPath, "_auth", file.Name()))
		if err != nil {
			return nil, err
		}
		// Unmarshal share
		share := &common.WalletShareConfig{}
		err = share.Unmarshal(bz)
		if err != nil {
			continue
		}
		// Add share to list
		shares = append(shares, share)
	}
	return shares, nil
}
