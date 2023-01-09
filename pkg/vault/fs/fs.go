package fs

import (
	"context"

	files "github.com/ipfs/go-ipfs-files"
	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/sonr-hq/sonr/pkg/common"
)

// Directories created for every user
var k_DEFAULT_DIRS = []string{
	"_auth",
	"mailbox",
	"public",
}

// VaultFS provides an interface for arbitrary Sonr Network Nodes to have IPFS configuration
// for the users secure storage.
type VaultFS interface {
	// CID
	CID() string

	// Address
	Address() string

	// Sync synchronizes the local filesystem with the IPFS network
	Sync() error

	// Upload a file to the public directory
	Upload(data []byte, name string) error

	// Download a file from the public directory
	Download(name string) ([]byte, error)

	// ListMessages returns a list of all messages in the mailbox
	ListMessages() ([][]byte, error)

	// SendMessage sends a message to the given address
	SendMessage(to []byte, message []byte) error

	// SignData signs the given data with the private key
	LoadShares() ([]*common.WalletShareConfig, error)

	// StoreShare stores the given share in the public directory
	StoreShare(share []byte, partyId string) error
}

// `New` creates a new VaultFS instance, initializes it, and returns it
func New(ctx context.Context, address string, ipfs icore.CoreAPI, opts ...Option) (VaultFS, error) {
	// Create a temporary directory to store the file
	config, err := defaultConfig(ctx, address, ipfs)
	if err != nil {
		return nil, err
	}
	err = config.Apply(opts...)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// Returning the IPFS path of the vault
func (c *Config) CID() string {
	return c.ipfsPath.String()
}

// Returning the address of the vault
func (c *Config) Address() string {
	return c.address
}

// Sync synchronizes the local filesystem with the IPFS network.
func (c *Config) Sync() error {
	// Fetch Basepath file info
	fileNode, err := c.ipfs.Unixfs().Get(c.ctx, c.ipfsPath)
	if err != nil {
		return err
	}

	// Copy the file to the temporary directory
	err = files.WriteTo(fileNode, c.localPath)
	if err != nil {
		return err
	}
	c.rootNode = fileNode
	return nil
}
