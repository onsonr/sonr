package fs

import (
	"context"
	"crypto/rand"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/client"
	"golang.org/x/crypto/nacl/box"
)

// `VaultConfig` is a struct that contains the local path to the vault, the IPFS node to use, the IPFS path
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
type VaultConfig struct {
	cctx client.Context
	// The Local Directory
	localRootDir Folder
	// The Auth Directory
	authDir Folder
	// The Mailbox Directory
	mailboxDir Folder
	// The Public Directory
	publicDir Folder
	address   string
	// Context
	ctx                   context.Context
	encryptionPubKeyPath  string
	encryptionPrivKeyPath string
}

// Option is a function that configures a `Config` object.
type Option func(*VaultConfig) error

// Apply applies the given options to the `Config` object.
func (c *VaultConfig) Apply(opts ...Option) error {
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return err
		}
	}
	err := setupLocalDirs(c)
	if err != nil {
		return err
	}
	return nil
}

//
// Helper Functions
//

// defaultConfig returns a `Config` object with default values.
func defaultConfig(addr string) (*VaultConfig, error) {
	// Create a temporary directory to store the file
	outputBasePath, err := os.MkdirTemp("", addr)
	if err != nil {
		return nil, err
	}
	c := &VaultConfig{
		address:      addr,
		localRootDir: Folder(filepath.Join(outputBasePath, addr)),
	}
	err = setupLocalDirs(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// It takes a path to a file or directory, and returns a UnixFS node
func setupLocalDirs(c *VaultConfig) error {
	// Configure default paths
	var err error
	c.authDir, err = c.localRootDir.CreateFolder("auth")
	if err != nil {
		return err
	}
	c.mailboxDir, err = c.localRootDir.CreateFolder("mailbox")
	if err != nil {
		return err
	}
	c.publicDir, err = c.localRootDir.CreateFolder("public")
	if err != nil {
		return err
	}
	return nil
}

// WithEncryptionKeyPath sets the encryption private key for the node from a file
func WithClientContext(cctx client.Context, generate bool) Option {
	return func(c *VaultConfig) error {
		c.cctx = cctx
		if hasKeys(cctx) {
			c.encryptionPrivKeyPath = kEncPrivKeyPath(cctx)
			c.encryptionPubKeyPath = kEncPubKeyPath(cctx)
		}
		if generate {
			err := generateBoxKeys(cctx)
			if err != nil {
				return err
			}
			c.encryptionPrivKeyPath = kEncPrivKeyPath(cctx)
			c.encryptionPubKeyPath = kEncPubKeyPath(cctx)
		}
		return nil
	}
}

func kEncPrivKeyPath(cctx client.Context) string {
	return filepath.Join(cctx.HomeDir, ".sonr", "highway", "encryption_key")
}

func kEncPubKeyPath(cctx client.Context) string {
	return filepath.Join(cctx.HomeDir, ".sonr", "highway", "encryption_key.pub")
}

func hasEncryptionKey(cctx client.Context) bool {
	_, err := os.Stat(kEncPrivKeyPath(cctx))
	return err == nil
}

func hasEncryptionPubKey(cctx client.Context) bool {
	_, err := os.Stat(kEncPubKeyPath(cctx))
	return err == nil
}

func hasKeys(cctx client.Context) bool {
	return hasEncryptionKey(cctx) && hasEncryptionPubKey(cctx)
}

func generateBoxKeys(cctx client.Context) error {
	pub, priv, err := box.GenerateKey(rand.Reader)
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Dir(kEncPrivKeyPath(cctx)), 0755)
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Dir(kEncPubKeyPath(cctx)), 0755)
	if err != nil {
		return err
	}
	err = os.WriteFile(kEncPrivKeyPath(cctx), priv[:], 0600)
	if err != nil {
		return err
	}
	err = os.WriteFile(kEncPubKeyPath(cctx), pub[:], 0600)
	if err != nil {
		return err
	}
	return nil
}

func loadBoxKeys(cctx client.Context) (*[32]byte, *[32]byte, error) {
	if !hasKeys(cctx) {
		return nil, nil, fmt.Errorf("no keys found")
	}
	priv, err := os.ReadFile(kEncPrivKeyPath(cctx))
	if err != nil {
		return nil, nil, err
	}
	pub, err := os.ReadFile(kEncPubKeyPath(cctx))
	if err != nil {
		return nil, nil, err
	}
	var privKey [32]byte
	var pubKey [32]byte
	copy(privKey[:], priv)
	copy(pubKey[:], pub)
	return &privKey, &pubKey, nil
}
