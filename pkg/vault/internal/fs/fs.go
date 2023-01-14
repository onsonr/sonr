package fs

import (
	"fmt"

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
	// Service returns the DID Service configuration of the Vault
	Service(cid string) *types.Service

	// Address
	Address() string

	// Sync synchronizes the local filesystem with the IPFS network
	Sync() error

	// StoreShare stores the given share in the public directory
	StoreShare(share []byte, partyId string, index int) error
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

// Returning the IPFS path of the vault
func (c *Config) Service(cid string) *types.Service {
	return &types.Service{
		ID:              fmt.Sprintf("did:snr:%s#vault", c.address),
		Type:            types.ServiceType_ServiceType_ENCRYPTED_DATA_VAULT,
		ServiceEndpoint: cid,
	}
}

// Returning the address of the vault
func (c *Config) Address() string {
	return c.address
}

// Sync synchronizes the local filesystem with the IPFS network.
func (c *Config) Sync() error { // ipfsApi coreapi.CoreAPI
	return nil
}
