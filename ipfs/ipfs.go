package ipfs

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	config "github.com/ipfs/go-ipfs-config"
	files "github.com/ipfs/go-ipfs-files"
	icore "github.com/ipfs/interface-go-ipfs-core"
	iface "github.com/ipfs/interface-go-ipfs-core"
	icorepath "github.com/ipfs/interface-go-ipfs-core/path"
	ma "github.com/multiformats/go-multiaddr"

	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/core/node/libp2p"
	"github.com/ipfs/go-ipfs/plugin/loader" // This package is needed so that all the preloaded plugins are loaded automatically
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	"github.com/libp2p/go-libp2p-core/peer"
	status "github.com/sonr-io/core/errors"
)

var flagExp = flag.Bool("experimental", false, "enable experimental features")

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

func createTempRepo() (string, error) {
	repoPath, err := ioutil.TempDir("", "ipfs-shell")
	if err != nil {
		return "", fmt.Errorf("failed to get temp dir: %s", err)
	}

	// Create a config with default options and a 2048 bit key
	cfg, err := config.Init(ioutil.Discard, 2048)
	if err != nil {
		return "", err
	}

	// When creating the repository, you can define custom settings on the repository, such as enabling experimental
	// features (See experimental-features.md) or customizing the gateway endpoint.
	// To do such things, you should modify the variable `cfg`. For example:
	if *flagExp {
		// https://github.com/ipfs/go-ipfs/blob/master/docs/experimental-features.md#ipfs-filestore
		cfg.Experimental.FilestoreEnabled = true
		// https://github.com/ipfs/go-ipfs/blob/master/docs/experimental-features.md#ipfs-urlstore
		cfg.Experimental.UrlstoreEnabled = true
		// https://github.com/ipfs/go-ipfs/blob/master/docs/experimental-features.md#ipfs-p2p
		cfg.Experimental.Libp2pStreamMounting = true
		// https://github.com/ipfs/go-ipfs/blob/master/docs/experimental-features.md#p2p-http-proxy
		cfg.Experimental.P2pHttpProxy = true
		// https://github.com/ipfs/go-ipfs/blob/master/docs/experimental-features.md#strategic-providing
		cfg.Experimental.StrategicProviding = true
	}

	// Create the repo with the config
	err = fsrepo.Init(repoPath, cfg)
	if err != nil {
		return "", fmt.Errorf("failed to init ephemeral node: %s", err)
	}

	return repoPath, nil
}

/// ------ Spawning the node

// Creates an IPFS node and returns its coreAPI
func createNode(ctx context.Context, repoPath string) (icore.CoreAPI, error) {
	// Open the repo
	repo, err := fsrepo.Open(repoPath)
	if err != nil {
		return nil, err
	}

	// Construct the node

	nodeOptions := &core.BuildCfg{
		Online:  true,
		Routing: libp2p.DHTOption, // This option sets the node to be a full DHT node (both fetching and storing DHT Records)
		// Routing: libp2p.DHTClientOption, // This option sets the node to be a client DHT node (only fetching records)
		Repo: repo,
	}

	node, err := core.NewNode(ctx, nodeOptions)
	if err != nil {
		return nil, err
	}

	// Attach the Core API to the constructed node
	return coreapi.NewCoreAPI(node)
}

// Spawns a node on the default repo location, if the repo exists
func SpawnDefault(ctx context.Context) (icore.CoreAPI, error) {
	defaultPath, err := config.PathRoot()
	if err != nil {
		// shouldn't be possible
		return nil, err
	}

	if err := setupPlugins(defaultPath); err != nil {
		return nil, err

	}

	return createNode(ctx, defaultPath)
}

// Spawns a node to be used just for this run (i.e. creates a tmp repo)
func SpawnEphemeral(ctx context.Context) (icore.CoreAPI, error) {
	if err := setupPlugins(""); err != nil {
		return nil, err
	}

	// Create a Temporary Repo
	repoPath, err := createTempRepo()
	if err != nil {
		return nil, fmt.Errorf("failed to create temp repo: %s", err)
	}

	// Spawning an ephemeral IPFS node
	return createNode(ctx, repoPath)
}

//

func ConnectToPeers(ctx context.Context, ipfs icore.CoreAPI, peers []string) error {
	var wg sync.WaitGroup
	peerInfos := make(map[peer.ID]*peer.AddrInfo, len(peers))
	for _, addrStr := range peers {
		addr, err := ma.NewMultiaddr(addrStr)
		if err != nil {
			return err
		}
		pii, err := peer.AddrInfoFromP2pAddr(addr)
		if err != nil {
			return err
		}
		pi, ok := peerInfos[pii.ID]
		if !ok {
			pi = &peer.AddrInfo{ID: pii.ID}
			peerInfos[pi.ID] = pi
		}
		pi.Addrs = append(pi.Addrs, pii.Addrs...)
	}

	wg.Add(len(peerInfos))
	for _, peerInfo := range peerInfos {
		go func(peerInfo *peer.AddrInfo) {
			defer wg.Done()
			err := ipfs.Swarm().Connect(ctx, *peerInfo)
			if err != nil {
				log.Printf("failed to connect to %s: %s", peerInfo.ID, err)
			}
		}(peerInfo)
	}
	wg.Wait()
	return nil
}

func getUnixfsFile(path string) (files.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	st, err := file.Stat()
	if err != nil {
		return nil, err
	}

	f, err := files.NewReaderPathFile(path, file, st)
	if err != nil {
		return nil, err
	}

	return f, nil
}

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

type UploadResponse struct {
	Status int
	Cid    string
	Path   string
}

// External Functions
func UploadData(data []byte, node iface.CoreAPI) (UploadResponse, error) {
	var ctx context.Context

	// TODO add CID preprocessor and name file off of that
	// use --only-hash option for preprocessor

	var permissions uint32 = 0644 // or whatever you need
	err := ioutil.WriteFile("./tmp/data.txt", data, fs.FileMode(permissions))
	if err != nil {
		return UploadResponse{Status: status.StatusInternalServerError}, err
	}

	// TODO revisit the file strategy and see if there is another way

	file, err := getUnixfsNode("./tmp/data.txt")
	if err != nil {
		panic(fmt.Errorf("Could not get File: %s", err))
	}

	cidFile, err := node.Unixfs().Add(ctx, file)
	if err != nil {
		panic(fmt.Errorf("Could not add File: %s", err))
	}

	fmt.Printf("Added file to IPFS with CID %s\n", cidFile.String())

	oldLocation := "./tmp/data.txt"
	newLocation := "./ipfs-storage/" + cidFile.String() + "/data.txt"
	err = os.Rename(oldLocation, newLocation)
	if err != nil {
		return UploadResponse{Status: status.StatusInternalServerError}, err
	}

	return UploadResponse{
		Status: status.StatusOK,
		Cid:    cidFile.String(),
	}, nil
}

// TODO discuss this with Josh
type DownloadResponse struct {
	Status    int
	FileData  []byte
	NameSpace string
	Path      string
}

func DownloadData(cid string, node iface.CoreAPI) (DownloadResponse, error) {
	var ctx context.Context
	cidPath := icorepath.New(cid)
	path := "./ipfs/" + cid + "/data.txt"

	// check local file system first
	if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return DownloadResponse{Status: status.StatusInternalServerError}, err
		}
		return DownloadResponse{
			Status:    status.StatusOK,
			FileData:  data,
			Path:      cidPath.String(),
			NameSpace: cidPath.Namespace(),
		}, nil
	}

	rootNodeFile, err := node.Unixfs().Get(ctx, cidPath)
	if err != nil {
		return DownloadResponse{Status: status.StatusInternalServerError}, err
	}
	err = files.WriteTo(rootNodeFile, path)
	if err != nil {
		return DownloadResponse{Status: status.StatusInternalServerError}, err
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return DownloadResponse{Status: status.StatusInternalServerError}, err
	}
	return DownloadResponse{
		Status:    status.StatusOK,
		FileData:  data,
		Path:      cidPath.String(),
		NameSpace: cidPath.Namespace(),
	}, nil
}
