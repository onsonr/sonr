package fs

import (
	"fmt"
	"strings"

	files "github.com/ipfs/go-ipfs-files"
	"github.com/sonrhq/core/pkg/common"
	"github.com/sonrhq/core/x/identity/types"
)

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
	// Address
	Address() string

	// Import imports the filesystem from IPFS with the DIDDocument
	Import(node common.IPFSNode, doc *types.DidDocument) error

	// Export exports the filesystem to IPFS and returns the CID
	Export(node common.IPFSNode) (*types.Service, error)

	// Sync synchronizes the local filesystem with the IPFS network

}

// `New` creates a new VaultFS instance, initializes it, and returns it
func New(address string, opts ...Option) (VaultFS, error) {
	// Create a temporary directory to store the file
	config, err := defaultConfig(address)
	if err != nil {
		return nil, err
	}
	err = config.Apply(opts...)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// Returning the address of the vault
func (c *VaultConfig) Address() string {
	return c.address
}

// Sync synchronizes the local filesystem with the IPFS network.
func (c *VaultConfig) Sync() error { // ipfsApi coreapi.CoreAPI
	return nil
}

// Export exports the filesystem to IPFS and returns the CID within a DID Service
func (c *VaultConfig) Export(node common.IPFSNode) (*types.Service, error) {
	if !c.localRootDir.Exists("") {
		return nil, fmt.Errorf("vault is not initialized")
	}
	cid, err := node.AddPath(c.localRootDir.Path())
	if err != nil {
		return nil, err
	}
	return &types.Service{
		Id:              fmt.Sprintf("did:snr:%s#vault", strings.TrimPrefix(c.address, "snr")),
		Type:            types.ServiceType_ServiceType_ENCRYPTED_DATA_VAULT,
		ServiceEndpoint: cid,
	}, nil
}

// Import imports the filesystem from IPFS with the DIDDocument
func (c *VaultConfig) Import(node common.IPFSNode, doc *types.DidDocument) error {
	service := doc.GetVaultService()
	if service == nil {
		return fmt.Errorf("no vault service found")
	}
	fileMap, err := node.GetPath(service.ServiceEndpoint)
	if err != nil {
		return err
	}
	for path, file := range fileMap {
		err := files.WriteTo(file, c.localRootDir.JoinPath(path))
		if err != nil {
			return err
		}
	}
	return nil
}
