package ipfs

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	orbitdb "berty.tech/go-orbit-db"
	"berty.tech/go-orbit-db/iface"
	files "github.com/ipfs/go-ipfs-files"
	icore "github.com/ipfs/interface-go-ipfs-core"
	config "github.com/ipfs/kubo/config"
	"github.com/ipfs/kubo/core"
	"github.com/ipfs/kubo/core/coreapi"
	"github.com/ipfs/kubo/core/node/libp2p"
	"github.com/ipfs/kubo/plugin/loader"
	"github.com/ipfs/kubo/repo/fsrepo"
	"github.com/sonrhq/core/internal/local"
	nodeconfig "github.com/sonrhq/core/pkg/node/config"
)

// Initialize creates a new local IPFS node
func Initialize(c *nodeconfig.Config) (nodeconfig.IPFSNode, error) {
	ipfsDoneCh := make(chan nodeconfig.IPFSNode)
	ipfsErrCh := make(chan error)
	n := defaultNode(c)
	snrctx := local.NewContext()
	go func() {
		// Apply the options
		err := n.initialize()
		if err != nil {
			ipfsErrCh <- err
			return
		}
		// Connect to the bootstrap nodes
		err = n.Connect(snrctx.BsMultiaddrs...)
		if err != nil {
			ipfsErrCh <- err
			return
		}
		db, err := orbitdb.NewOrbitDB(n.ctx, n.CoreAPI(), &orbitdb.NewOrbitDBOptions{})
		if err != nil {
			ipfsErrCh <- err
			return
		}
		n.orbitDb = db
		ipfsDoneCh <- n
	}()

	select {
	case err := <-ipfsErrCh:
		return nil, err
	case <-ipfsDoneCh:
		return n, nil
	}
}

// Miscellanenous
var loadPluginsOnce sync.Once

// TopicMessageHandler is a function that handles a message received on a topic
type TopicMessageHandler func(topic string, msg icore.PubSubMessage) error

//
// Node Setup and Initialization
//

// defaultNode creates a new node with default options
func defaultNode(cnfg *nodeconfig.Config) *localIpfs {
	return &localIpfs{
		ctx:   context.Background(),
		config: cnfg,
	}
}

// It's creating a new node and returning the coreAPI and the node itself.
func (c *localIpfs) initialize() error {
	snrctx := local.NewContext()
	c.repoPath = snrctx.IPFSRepoPath

	// Spawn a local peer using a temporary path, for testing purposes
	var onceErr error
	loadPluginsOnce.Do(func() {
		onceErr = setupPlugins("")
	})
	if onceErr != nil {
		return onceErr
	}

	// Create a config with default options and a 2048 bit key
	cfg, err := config.Init(io.Discard, 2048)
	if err != nil {
		return err
	}

	// https://github.com/ipfs/kubo/blob/master/docs/experimental-features.md#ipfs-filestore
	cfg.Experimental.FilestoreEnabled = true
	// https://github.com/ipfs/kubo/blob/master/docs/experimental-features.md#ipfs-p2p
	cfg.Experimental.Libp2pStreamMounting = true

	// Create the repo with the config
	err = fsrepo.Init(c.repoPath, cfg)
	if err != nil {
		return err
	}

	node, err := createNode(c.ctx, c.repoPath)
	if err != nil {
		return err
	}

	api, err := coreapi.NewCoreAPI(node)
	if err != nil {
		return err
	}

	// Set the node and repoPath
	c.node = node
	c.api = api
	return nil
}

// It loads plugins from the `externalPluginsPath` directory and injects them into the application
func setupPlugins( externalPluginsPath string) error {
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
		Routing: libp2p.DHTOption,
		Repo:    repo,
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

///
/// Helper functions
///

// fetchDocsAddress fetches the address of the document store for a given username
func fetchDocsAddress(orb iface.OrbitDB, username string) (string, error) {
	addr, err := orb.DetermineAddress(context.Background(), username, nodeconfig.DB_DOCUMENT_STORE.String(), nil)
	if err != nil {
		return "", err
	}
	return addr.String(), nil
}

// fetchEventLogAddress fetches the address of the event log for a given username
func fetchEventLogAddress(orb iface.OrbitDB, username string) (string, error) {
	addr, err := orb.DetermineAddress(context.Background(), username, nodeconfig.DB_EVENT_LOG_STORE.String(), nil)
	if err != nil {
		return "", err
	}
	return addr.String(), nil
}

// fetchKeyValueAddress fetches the address of the key value store for a given username
func fetchKeyValueAddress(orb iface.OrbitDB, username string) (string, error) {
	addr, err := orb.DetermineAddress(context.Background(), username, nodeconfig.DB_KEY_VALUE_STORE.String(), nil)
	if err != nil {
		return "", err
	}
	return addr.String(), nil
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
