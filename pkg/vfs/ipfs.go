package vfs

import (
	"context"
	"os"
	"path/filepath"

	"github.com/onsonr/hway/internal/env"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/core/coreiface/options"
)

// Constant for the name of the folder where the vaults are stored
const kVaultsFolderName = ".sonr-vaults"

// enclaveRoot is the folder where the vaults are stored
var enclaveRoot Folder

// Package initializes the VaultsFolder
func init() {
	// Initialize VaultsFolder
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	// Create the folder if it does not exist
	enclaveRoot = NewFolder(filepath.Join(homeDir, kVaultsFolderName))
	if !enclaveRoot.Exists() {
		if err := enclaveRoot.Create(); err != nil {
			panic(err)
		}
	}
}

// NewVaultFolder creates a new folder under the VaultsFolder directory
func NewVaultFolder(name string) (Folder, error) {
	vaultFolder := enclaveRoot.Join(name)
	err := vaultFolder.Create()
	if err != nil {
		return "", err
	}
	return vaultFolder, nil
}

// SaveToIPFS saves the Folder to IPFS and returns the IPFS path
func SyncFolderToIPFS(ctx context.Context, f Folder) (path.Path, error) {
	node, err := f.Node()
	if err != nil {
		return nil, err
	}
	c, err := env.GetIPFSClient()
	if err != nil {
		return nil, err
	}

	path, err := c.Unixfs().Add(ctx, node)
	if err != nil {
		return nil, err
	}

	return path, nil
}

// PublishToIPNS publishes the Folder to IPNS
func PublishToIPNS(ctx context.Context, ipfsPath path.Path, f Folder) error {
	c, err := env.GetIPFSClient()
	if err != nil {
		return err
	}
	_, err = c.Name().Publish(ctx, ipfsPath, options.Name.Key(f.Name()))
	if err != nil {
		return err
	}
	return nil
}

// LoadFromIPFS loads a Folder from IPFS
func LoadFromIPFS(ctx context.Context, path string) (Folder, error) {
	c, err := env.GetIPFSClient()
	if err != nil {
		return "", err
	}
	cid, err := ParsePath(path)
	if err != nil {
		return "", err
	}
	node, err := c.Unixfs().Get(ctx, cid)
	if err != nil {
		return "", err
	}
	return LoadNodeInFolder(path, node)
}
