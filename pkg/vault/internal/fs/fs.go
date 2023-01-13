package fs

import (
	"context"
	"fmt"
	"strings"

	files "github.com/ipfs/go-ipfs-files"
	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/node/config"
	"github.com/sonr-hq/sonr/x/identity/types"
)

// Directories created for every user
var k_DEFAULT_DIRS = []string{
	"_auth",
	"mailbox",
	"public",
}

// VaultFS provides an interface for arbitrary Sonr Network Nodes to have IPFS configuration
// for the users secure storage.
// @property {string} CID - The CID of the vault
// @property Service - The DID Service configuration of the Vault
// @property {string} Address - The IPFS address of the vault
// @property {error} Sync - Synchronizes the local filesystem with the IPFS network
// @property {error} Upload - Uploads a file to the public directory
// @property Download - Downloads a file from the public directory
// @property ListMessages - Returns a list of all messages in the mailbox
// @property {error} SendMessage - This is the function that will be used to send a message to the
// given address.
// @property LoadShares - Loads the shares from the public directory
// @property {error} StoreShare - Stores the share in the public directory
type VaultFS interface {
	// CID
	CID() string

	// Service returns the DID Service configuration of the Vault
	Service() *types.Service

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
	LoadShares(password string) ([]*common.WalletShareConfig, error)

	// StoreShare stores the given share in the public directory
	StoreShare(share []byte, partyId string, password string) error
}

// `New` creates a new VaultFS instance, initializes it, and returns it
func New(ctx context.Context, address string, node config.IPFSNode, opts ...Option) (VaultFS, error) {
	// Create a temporary directory to store the file
	config, err := defaultConfig(ctx, address, node.CoreAPI())
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

// Returning the IPFS path of the vault
func (c *Config) Service() *types.Service {
	return &types.Service{
		ID:              fmt.Sprintf("did:ipfs:%s", c.address),
		Type:            types.ServiceType_ServiceType_ENCRYPTED_DATA_VAULT,
		ServiceEndpoint: fmt.Sprintf("%s/%s", c.resolverUrl, strings.TrimPrefix(c.ipfsPath.String(), "/")),
	}
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
