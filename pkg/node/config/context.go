package config

import (
	"context"
	"os"
	"path/filepath"
)

const (
	// CURRENT_CHAIN_ID is the current chain ID.
	CURRENT_CHAIN_ID = "sonrdevnet-1"
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
	HomeDir       string
	NodeRESTUri   string
	NodeGRPCUri   string
	NodeFaucetUri string
	Rendevouz     string
	BsMultiaddrs  []string
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
		NodeRESTUri:   "http://api.sonr.network",
		NodeGRPCUri:   "grpc.sonr.network",
		NodeFaucetUri: "http://faucet.sonr.network",
		Rendevouz:     defaultRendezvousString,
		BsMultiaddrs:  defaultBootstrapMultiaddrs,
	}
	return &ctx, nil
}

// checkPathExists checks if a path exists
func checkPathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

// checkPathNotExists checks if a path does not exist
func checkPathNotExists(path string) bool {
	return !checkPathExists(path)
}
