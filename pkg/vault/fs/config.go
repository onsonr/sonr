package fs

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	files "github.com/ipfs/go-ipfs-files"
	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/ipfs/interface-go-ipfs-core/path"
)

// Directories created for every user
var k_DEFAULT_DIRS = []string{
	"_auth",
	"mailbox",
	"public",
}

// VaultFS provides an interface for arbitrary Sonr Network Nodes to have IPFS configuration
// for the users secure storage.
type VaultFS interface {
	Add(data []byte, name string) error
	Get(name string) ([]byte, error)
	ListMessages() ([][]byte, error)
	SendMessage(to []byte, message []byte) error
	SignData(data []byte) ([]byte, []byte, error)
	StoreShare(share []byte, partyId string) error
	VerifyData(data []byte, signature []byte) bool
}

func New(ipfs icore.CoreAPI, address string) (VaultFS, error) {
	// Add config/api
	impl := &vaultFsImpl{
		ipfs:    ipfs,
		address: address,
	}

	// Initialize Node: GenKey, MkDir, Publish
	err := impl.init()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("vaultfs failed to initialize %s", err))
	}
	return impl, nil
}

func Load(ipfs icore.CoreAPI, address string) (VaultFS, error) {
	// Add config/api
	impl := &vaultFsImpl{
		ipfs:    ipfs,
		address: address,
	}

	// Initialize Node: GenKey, MkDir, Publish
	fil, err := impl.loadDefaultDirs()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("vaultfs failed to initialize %s", err))
	}
	impl.rootNode = fil
	return impl, nil
}

type vaultFsImpl struct {
	ipfs      icore.CoreAPI
	address   string
	entry     icore.IpnsEntry
	key       icore.Key
	cidHash   path.Resolved
	rootNode  files.Node
	localPath string
}

func (vfs *vaultFsImpl) init() error {
	// Generate Key for Address
	key, err := vfs.ipfs.Key().Generate(context.Background(), vfs.address)
	if err != nil {
		return err
	}
	vfs.key = key
	fmt.Printf("Generated key for %s", vfs.address)

	// Get default file node
	fileNode, err := vfs.setupDefaultDirs()
	if err != nil {
		return err
	}
	vfs.rootNode = fileNode

	// Pin default user directory
	cid, err := vfs.ipfs.Unixfs().Add(context.Background(), fileNode, options.Unixfs.Pin(true))
	if err != nil {
		return err
	}
	vfs.cidHash = cid
	fmt.Printf("Pinned user directory with CID %s", cid)

	// Publish Name with Key
	entry, err := vfs.ipfs.Name().Publish(context.Background(), vfs.cidHash, options.Name.Key(vfs.address))
	if err != nil {
		return err
	}
	vfs.entry = entry
	fmt.Printf("Published user directory with name %s and value %s", entry.Name(), entry.Value())
	return nil
}

func (vfs *vaultFsImpl) loadDefaultDirs() (files.Node, error) {
	// Fetch Basepath file info
	fileNode, err := vfs.ipfs.Unixfs().Get(context.Background(), vfs.cidHash)
	if err != nil {
		return nil, err
	}
	// Create a temporary directory to store the file
	outputBasePath, err := os.MkdirTemp("", vfs.address)
	if err != nil {
		return nil, err
	}
	outputPath := filepath.Join(outputBasePath, vfs.address)
	vfs.localPath = outputPath
	// Copy the file to the temporary directory
	err = files.WriteTo(fileNode, outputPath)
	if err != nil {
		return nil, err
	}
	return fileNode, nil
}

// It takes a path to a file or directory, and returns a UnixFS node
func (vfs *vaultFsImpl) setupDefaultDirs() (files.Node, error) {
	// Create root temp directory
	rootPath, err := os.MkdirTemp("", vfs.address)
	if err != nil {
		return nil, err
	}

	// Configure default paths
	paths := []string{}
	for _, p := range k_DEFAULT_DIRS {
		paths = append(paths, filepath.Join(rootPath, vfs.address, p))
	}

	// Recursively create directories
	for _, p := range paths {
		err := os.MkdirAll(p, 0777)
		if err != nil {
			return nil, err
		}
	}

	// Fetch Basepath file info
	basePath := filepath.Join(rootPath, vfs.address)
	vfs.localPath = basePath
	st, err := os.Stat(basePath)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Created default directories %s", basePath)

	// Create new NewSerialFile
	f, err := files.NewSerialFile(basePath, false, st)
	if err != nil {
		return nil, err
	}
	return f, nil
}
