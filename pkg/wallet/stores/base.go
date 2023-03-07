package stores

import (
	"github.com/sonrhq/core/pkg/common"
	"github.com/sonrhq/core/pkg/wallet"
	"github.com/sonrhq/core/pkg/wallet/stores/internal"
)

// NewWalletStore returns a new WalletStore
func New(acc wallet.Account, opts ...Option) (wallet.Store, error) {
	cfg := &storeConfig{
		acc: acc,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg.Apply()
}

// storeConfig is the configuration for the store
type storeConfig struct {
	acc      wallet.Account
	path     string
	ipfsNode common.IPFSNode
	isFile   bool
	isIPFS   bool
	pwd      []byte
}

// Apply applies the configuration to the store
func (cfg *storeConfig) Apply() (wallet.Store, error) {
	if cfg.isFile {
		return internal.NewFileStore(cfg.path, cfg.pwd, cfg.acc.Config())
	}
	if cfg.isIPFS {
		return internal.NewIPFSStore(cfg.ipfsNode, cfg.acc.Config())
	}
	return internal.NewMemoryStore(cfg.acc.Config())
}

// Option is a function that configures the store
type Option func(*storeConfig)

// SetFileStore sets the store to use a file store. Password must be 32 bytes and already hashed
func SetFileStore(path string, password []byte) Option {
	return func(cfg *storeConfig) {
		cfg.path = path
		cfg.isFile = true
		cfg.isIPFS = false
		cfg.pwd = password
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
