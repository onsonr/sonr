package fs

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	files "github.com/ipfs/go-ipfs-files"
	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/ipfs/interface-go-ipfs-core/path"
)

// `Config` is a struct that contains the local path to the vault, the IPFS node to use, the IPFS path
// to the vault, the IPFS key to use, the IPFS entry to use, the root node of the vault, the address of
// the vault, and the authentication shares.
// @property {string} localPath - The local path to the vault
// @property ipfs - The IPFS node to use.
// @property ipfsPath - The IPFS path to the vault.
// @property key - The IPFS key to use
// @property entry - The IPNS entry that points to the root node of the vault.
// @property rootNode - The root node of the vault. This is the node that contains all the other nodes
// in the vault.
// @property {string} address - The address of the vault. This is the address that the vault will be
// @property {[]*common.WalletShareConfig} authShares - The authentication shares.
type Config struct {
	// The local path to the vault
	localPath string
	// The IPFS node to use
	ipfs icore.CoreAPI
	// The IPFS path to the vault
	ipfsPath path.Path
	// The IPFS key to use
	key icore.Key
	// The IPFS entry to use
	entry icore.IpnsEntry
	// The root node of the vault
	rootNode files.Node
	// The address of the vault
	address string
	// Context
	ctx context.Context
	// IsExisting
	isExisting bool
	// ResolverURL
	resolverUrl string
}

// Option is a function that configures a `Config` object.
type Option func(*Config) error

// WithIPFSPath sets the IPFS path to the vault.
func WithIPFSPath(ipfsPath string) Option {
	return func(c *Config) error {
		c.ipfsPath = path.New(ipfsPath)
		c.isExisting = true
		return nil
	}
}

// Apply applies the given options to the `Config` object.
func (c *Config) Apply(opts ...Option) error {
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return err
		}
	}
	if !c.isExisting {
		ctx, cancel := context.WithCancel(c.ctx)
		defer cancel()
		// Set Root Node
		rn, err := setupLocalDirs(c)
		if err != nil {
			return err
		}
		c.rootNode = rn

		// Pin default user directory
		cid, err := c.ipfs.Unixfs().Add(ctx, c.rootNode, options.Unixfs.Pin(true))
		if err != nil {
			return err
		}
		c.ipfsPath = cid
		fmt.Printf("Pinned user directory with CID %s", cid)
	} else {
		return c.Sync()
	}
	return nil
}

//
// Helper Functions
//

// defaultConfig returns a `Config` object with default values.
func defaultConfig(ctx context.Context, addr string, ipfs icore.CoreAPI) (*Config, error) {
	// Create a temporary directory to store the file
	outputBasePath, err := os.MkdirTemp("", addr)
	if err != nil {
		return nil, err
	}
	c := &Config{
		ctx:         ctx,
		address:     addr,
		ipfs:        ipfs,
		localPath:   filepath.Join(outputBasePath, addr),
		resolverUrl: "https://ipfs.sonr.network",
	}

	return c, nil
}

// It takes a path to a file or directory, and returns a UnixFS node
func setupLocalDirs(c *Config) (files.Node, error) {
	// Configure default paths
	paths := []string{}
	for _, p := range k_DEFAULT_DIRS {
		paths = append(paths, filepath.Join(c.localPath, p))
	}

	// Recursively create directories
	for _, p := range paths {
		err := os.MkdirAll(p, 0777)
		if err != nil {
			return nil, err
		}
	}

	// Fetch Basepath file info
	st, err := os.Stat(c.localPath)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Created default directories %s", c.localPath)

	// Create new NewSerialFile
	f, err := files.NewSerialFile(c.localPath, false, st)
	if err != nil {
		return nil, err
	}
	return f, nil
}
