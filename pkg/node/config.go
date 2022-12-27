package node

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/taurusgroup/multi-party-sig/pkg/party"

	files "github.com/ipfs/go-ipfs-files"
	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/kubo/config"
	"github.com/ipfs/kubo/core"
	"github.com/ipfs/kubo/core/coreapi"
	klibp2p "github.com/ipfs/kubo/core/node/libp2p"
	"github.com/ipfs/kubo/plugin/loader"
	"github.com/ipfs/kubo/repo/fsrepo"

	cv1 "github.com/sonr-hq/sonr/pkg/common"
)

//
// Miscellanenous
//
var loadPluginsOnce sync.Once

// TopicMessageHandler is a function that handles a message received on a topic
type TopicMessageHandler func(topic string, msg icore.PubSubMessage) error

//
// Default configuration
//
var (
	// defaultBootstrapMultiaddrs is the default list of bootstrap nodes
	defaultBootstrapMultiaddrs = []string{
		// IPFS Bootstrapper nodes.
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
		// These are the bootstrap nodes for the IPFS network.
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

	// defaultCallback is the default callback for the motor
	defaultCallback = common.DefaultCallback()

	// defaultRendezvousString is the default rendezvous string for the motor
	defaultRendezvousString = "sonr"
)

//
// Options for the node
//

// NodeOption is a function that configures a Node
type NodeOption func(*Node) error

// AddBootstrappers adds additional nodes to start initial connections with
func AddBootstrappers(bootstrappers []string) NodeOption {
	return func(c *Node) error {
		c.bootstrappers = append(c.bootstrappers, bootstrappers...)
		return nil
	}
}

// SetPeerIds sets the peer ids for the node
func SetPeerIds(peerIds ...peer.ID) NodeOption {
	return func(c *Node) error {
		if len(peerIds) > 0 {
			c.mpcPeerIds = peerIds
		}
		return nil
	}
}

// WithNodeCallback sets the callback for the motor
func WithNodeCallback(callback common.NodeCallback) NodeOption {
	return func(c *Node) error {
		c.callback = callback
		return nil
	}
}

// WithPartyId sets the party id for the node. This is to be replaced by the User defined label for the device
func WithPartyId(partyId string) NodeOption {
	return func(c *Node) error {
		c.partyId = party.ID(partyId)
		return nil
	}
}

// WithPeerType sets the type of peer
func WithPeerType(peerType cv1.NodeInfo_Type) NodeOption {
	return func(c *Node) error {
		c.peerType = peerType
		return nil
	}
}

// WithWalletShare sets the wallet share for the node
func WithWalletShare(walletShare common.WalletShare) NodeOption {
	return func(c *Node) error {
		c.walletShare = walletShare
		return nil
	}
}

//
// Node Setup and Initialization
//

// defaultNode creates a new node with default options
func defaultNode(ctx context.Context) *Node {
	return &Node{
		ctx:                ctx,
		bootstrappers:      defaultBootstrapMultiaddrs,
		callback:           defaultCallback,
		peerType:           cv1.NodeInfo_MOTOR,
		rendezvous:         defaultRendezvousString,
		topicEventHandlers: make(map[string]TopicMessageHandler),
		partyId:            party.ID("current"),
	}
}

// It's creating a new node and returning the coreAPI and the node itself.
func (c *Node) Apply(opts ...NodeOption) error {
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return err
		}
	}
	// Spawn a local peer using a temporary path, for testing purposes
	var onceErr error
	loadPluginsOnce.Do(func() {
		onceErr = setupPlugins("")
	})
	if onceErr != nil {
		return onceErr
	}

	// Create a Temporary Repo
	repoPath, err := createTempRepo()
	if err != nil {
		return fmt.Errorf("error creating temporary repo: %s", err)
	}

	node, err := createNode(c.ctx, repoPath)
	if err != nil {
		return err
	}

	api, err := coreapi.NewCoreAPI(node)
	if err != nil {
		return err
	}

	// Set the node and repoPath
	c.node = node
	c.repoPath = repoPath
	c.CoreAPI = api
	return nil
}

// It loads plugins from the `externalPluginsPath` directory and injects them into the application
func setupPlugins(externalPluginsPath string) error {
	// Load any external plugins if available on externalPluginsPath
	plugins, err := loader.NewPluginLoader(filepath.Join(externalPluginsPath, "plugins"))
	if err != nil {
		return fmt.Errorf("error loading plugins: %s", err)
	}

	// Load preloaded and external plugins
	if err := plugins.Initialize(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	if err := plugins.Inject(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	return nil
}

// It creates a temporary directory, initializes a new IPFS repo in that directory, and returns the
// path to the repo
func createTempRepo() (string, error) {
	repoPath, err := os.MkdirTemp("", "ipfs-repo")
	if err != nil {
		return "", fmt.Errorf("failed to get temp dir: %s", err)
	}

	// Create a config with default options and a 2048 bit key
	cfg, err := config.Init(io.Discard, 2048)
	if err != nil {
		return "", err
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
	err = fsrepo.Init(repoPath, cfg)
	if err != nil {
		return "", fmt.Errorf("failed to init ephemeral node: %s", err)
	}

	return repoPath, nil
}

// Creates an IPFS node and returns its coreAPI
func createNode(ctx context.Context, repoPath string) (*core.IpfsNode, error) {
	// Open the repo
	repo, err := fsrepo.Open(repoPath)
	if err != nil {
		return nil, err
	}

	err = repo.SetConfigKey("Pubsub.Enabled", true)
	if err != nil {
		return nil, err
	}
	err = repo.SetConfigKey("Pubsub.Router", "gossipsub")
	if err != nil {
		return nil, err
	}

	// Construct the node
	nodeOptions := &core.BuildCfg{
		Online:  true,
		Routing: klibp2p.DHTServerOption, // This option sets the node to be a full DHT node (both fetching and storing DHT Records)
		// Routing: libp2p.DHTClientOption, // This option sets the node to be a client DHT node (only fetching records)
		Repo: repo,
		ExtraOpts: map[string]bool{
			"pubsub": true,
			"ipnsps": true,
		},
	}

	// Create the node
	return core.NewNode(ctx, nodeOptions)
}

// It takes a path to a file or directory, and returns a UnixFS node
// It takes a path to a file or directory, and returns a UnixFS node
func getUnixfsNode(path string) (files.Node, error) {
	st, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	f, err := files.NewSerialFile(path, false, st)
	if err != nil {
		return nil, err
	}

	return f, nil
}
