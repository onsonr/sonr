package common

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/ipfs/kubo/config"
	"github.com/ipfs/kubo/repo/fsrepo"
	"golang.org/x/crypto/nacl/box"
)

// Default configuration
var (
	// defaultBootstrapMultiaddrs is the default list of bootstrap nodes
	defaultBootstrapMultiaddrs = []string{
		// IPFS Bootstrapper nodes.
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
		// "/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
		// "/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		// "/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",

		// IPFS Cluster Pinning nodes
		// "/ip4/138.201.67.219/tcp/4001/p2p/QmUd6zHcbkbcs7SMxwLs48qZVX3vpcM8errYS7xEczwRMA",
		// "/ip4/138.201.67.219/udp/4001/quic/p2p/QmUd6zHcbkbcs7SMxwLs48qZVX3vpcM8errYS7xEczwRMA",
		// "/ip4/138.201.67.220/tcp/4001/p2p/QmNSYxZAiJHeLdkBg38roksAR9So7Y5eojks1yjEcUtZ7i",
		// "/ip4/138.201.67.220/udp/4001/quic/p2p/QmNSYxZAiJHeLdkBg38roksAR9So7Y5eojks1yjEcUtZ7i",
		// "/ip4/138.201.68.74/tcp/4001/p2p/QmdnXwLrC8p1ueiq2Qya8joNvk3TVVDAut7PrikmZwubtR",
		// "/ip4/138.201.68.74/udp/4001/quic/p2p/QmdnXwLrC8p1ueiq2Qya8joNvk3TVVDAut7PrikmZwubtR",
		// "/ip4/94.130.135.167/tcp/4001/p2p/QmUEMvxS2e7iDrereVYc5SWPauXPyNwxcy9BXZrC1QTcHE",
		// "/ip4/94.130.135.167/udp/4001/quic/p2p/QmUEMvxS2e7iDrereVYc5SWPauXPyNwxcy9BXZrC1QTcHE",

		// You can add more nodes here, for example, another IPFS node you might have running locally, mine was:
		// "/ip4/127.0.0.1/tcp/4010/p2p/QmZp2fhDLxjYue2RiUvLwT9MWdnbDxam32qYFnGmxZDh5L",
		// "/ip4/127.0.0.1/udp/4010/quic/p2p/QmZp2fhDLxjYue2RiUvLwT9MWdnbDxam32qYFnGmxZDh5L",
	}

	// defaultRendezvousString is the default rendezvous string for the motor
	defaultRendezvousString = "sonr"
)

// `Context` is a struct that contains the information needed to run the `go-ipfs` node.
// @property {string} HomeDir - The home directory of the user running the application.
// @property {string} RepoPath - The path to the IPFS repo.
// @property {string} NodeRESTUri - The REST endpoint of the node.
// @property {string} NodeGRPCUri - The GRPC endpoint of the node.
// @property {string} NodeFaucetUri - The URI of the faucet service.
// @property {string} Rendevouz - The rendevouz point for the swarm.
// @property {[]string} BsMultiaddrs - The bootstrap multiaddrs.
// @property encPubKey - The public key of the encryption key pair.
// @property encPrivKey - The private key used to encrypt the data.
type Context struct {
	Ctx           context.Context
	ClientContext client.Context
	HomeDir       string
	RepoPath      string
	NodeRESTUri   string
	NodeGRPCUri   string
	NodeFaucetUri string
	Rendevouz     string
	BsMultiaddrs  []string

	encPubKey  *[32]byte
	encPrivKey *[32]byte
}

// NewContext creates a new context object, initializes the encryption keys, and returns the context object
func NewContext(c context.Context) (*Context, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	ctx := Context{
		Ctx:           c,
		HomeDir:       filepath.Join(userHomeDir, ".sonr"),
		RepoPath:      filepath.Join(userHomeDir, ".sonr", "ipfs"),
		NodeRESTUri:   "http://api.sonr.network",
		NodeGRPCUri:   "grpc.sonr.network",
		NodeFaucetUri: "http://faucet.sonr.network",
		Rendevouz:     defaultRendezvousString,
		BsMultiaddrs:  defaultBootstrapMultiaddrs,
	}
	return ctx.initialize()
}

// SetClientContext sets the client context
func (ctx *Context) WrapClientContext(c client.Context) *Context {
	ctx.ClientContext = c
	return ctx
}

// It creates a temporary directory, initializes a new IPFS repo in that directory, and returns the
// path to the repo
func (ctx *Context) initialize() (*Context, error) {
	if !hasKeys(ctx) {
		err := generateBoxKeys(ctx)
		if err != nil {
			return ctx, err
		}
	}
	pk, pb, err := loadBoxKeys(ctx)
	if err != nil {
		return ctx, err
	}
	ctx.encPrivKey = pk
	ctx.encPubKey = pb

	// Create the home directory if it doesn't exist
	_, err = os.Stat(ctx.RepoPath)
	if err != nil {
		if os.IsNotExist(err) {
			// Create a config with default options and a 2048 bit key
			cfg, err := config.Init(io.Discard, 2048)
			if err != nil {
				return ctx, err
			}
			// https://github.com/ipfs/kubo/blob/master/docs/experimental-features.md#ipfs-filestore
			cfg.Experimental.FilestoreEnabled = true
			// https://github.com/ipfs/kubo/blob/master/docs/experimental-features.md#ipfs-urlstore
			cfg.Experimental.UrlstoreEnabled = true
			// https://github.com/ipfs/kubo/blob/master/docs/experimental-features.md#ipfs-p2p
			cfg.Experimental.Libp2pStreamMounting = true
			// https://github.com/ipfs/kubo/blob/master/docs/experimental-features.md#p2p-http-proxy
			cfg.Experimental.P2pHttpProxy = true

			// Create the repo with the config
			err = fsrepo.Init(ctx.RepoPath, cfg)
			if err != nil {
				return ctx, fmt.Errorf("failed to init ephemeral node: %s", err)
			}
		}
	}
	return ctx, nil
}

// Write encrypts a message using the box algorithm
// This encrypts msg and appends the result to the nonce.
func (b *Context) EncryptMessage(msg []byte, peerPk []byte) []byte {
	return box.Seal(nil, msg, b.findNonce(peerPk), b.encPubKey, b.encPrivKey)
}

// The recipient can decrypt the message using their private key and the
// sender's public key. When you decrypt, you must use the same nonce you
// used to encrypt the message. One way to achieve this is to store the
// nonce alongside the encrypted message. Above, we stored the nonce in the
// first 24 bytes of the encrypted text.
func (b *Context) DecryptMessage(encMsg []byte, peerPk []byte) ([]byte, bool) {
	return box.Open(nil, encMsg, b.findNonce(peerPk), b.encPubKey, b.encPrivKey)
}

// Nonce returns the nonce for the box
func (b *Context) findNonce(peerPk []byte) *[24]byte {
	var nonce [24]byte
	copy(nonce[:], peerPk[:24])
	return &nonce
}

// It returns the path to the file that contains the encryption key for the private key
func kEncPrivKeyPath(cctx *Context) string {
	return filepath.Join(cctx.HomeDir, "config", "highway", "encryption_key")
}

// `kEncPubKeyPath` returns the path to the public encryption key
func kEncPubKeyPath(cctx *Context) string {
	return filepath.Join(cctx.HomeDir, "config", "highway", "encryption_key.pub")
}

// If the file `/.lotus/encryption-key` exists, then return true, otherwise return false
func hasEncryptionKey(cctx *Context) bool {
	_, err := os.Stat(kEncPrivKeyPath(cctx))
	return err == nil
}

// It checks if the file `/.lotus/encryption-pubkey` exists
func hasEncryptionPubKey(cctx *Context) bool {
	_, err := os.Stat(kEncPubKeyPath(cctx))
	return err == nil
}

// If the context has both an encryption key and an encryption public key, then return true
func hasKeys(cctx *Context) bool {
	return hasEncryptionKey(cctx) && hasEncryptionPubKey(cctx)
}

// It generates a new encryption key pair, and writes the private key to the file
// `/.lotus/key.priv` and the public key to the file `/.lotus/key.pub`
func generateBoxKeys(cctx *Context) error {
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

// It reads the private and public keys from the file system and returns them
func loadBoxKeys(cctx *Context) (*[32]byte, *[32]byte, error) {
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
