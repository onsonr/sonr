package stores

import (
	"github.com/sonrhq/core/pkg/common"
	"github.com/sonrhq/core/pkg/crypto/wallet"
	 "github.com/sonrhq/core/pkg/crypto/wallet/stores/internal"
	v1 "github.com/sonrhq/core/x/identity/types/vault/v1"
)

// NewWalletStore returns a new WalletStore
func New(cnfg *v1.AccountConfig, opts ...Option) (wallet.Store, error) {
	cfg := &storeConfig{
		accCfg: cnfg,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg.Apply()
}

// storeConfig is the configuration for the store
type storeConfig struct {
	accCfg   *v1.AccountConfig
	path     string
	ipfsNode common.IPFSNode
	isFile   bool
	isIPFS   bool
}

// Apply applies the configuration to the store
func (cfg *storeConfig) Apply() (wallet.Store, error) {
	if cfg.isFile {
		return internal.NewFileStore(cfg.path, cfg.accCfg)
	}
	if cfg.isIPFS {
		return internal.NewIPFSStore(cfg.ipfsNode, cfg.accCfg)
	}
	return internal.NewMemoryStore(cfg.accCfg)
}

// Option is a function that configures the store
type Option func(*storeConfig)

// SetFileStore sets the store to use a file store
func SetFileStore(path string) Option {
	return func(cfg *storeConfig) {
		cfg.path = path
		cfg.isFile = true
		cfg.isIPFS = false
	}
}

// SetIPFSStore sets the store to use an IPFS store
func SetIPFSStore(ipfsNode common.IPFSNode) Option {
	return func(cfg *storeConfig) {
		cfg.ipfsNode = ipfsNode
		cfg.isFile = false
		cfg.isIPFS = true
	}
}
