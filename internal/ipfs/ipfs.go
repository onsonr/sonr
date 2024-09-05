package ipfs

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/client/rpc"
)

var (
	initialized bool
	localMount  bool
	ipfsClient  *rpc.HttpApi
	ipfsFS      FileSystem
)

// init initializes the IPFS client and checks for local mounts
func init() {
	var err error
	ipfsClient, err = rpc.NewLocalApi()
	if err != nil {
		initialized = false
		localMount = false
		return
	}

	initialized = true
	ipfsFS = &IPFSFileSystem{client: ipfsClient}

	// Check if /ipfs and /ipns are mounted using os package
	_, errIPFS := os.Stat("/ipfs")
	_, errIPNS := os.Stat("/ipns")
	localMount = !os.IsNotExist(errIPFS) && !os.IsNotExist(errIPNS)
}

// GetFileSystem returns the IPFS FileSystem implementation
func GetFileSystem() FileSystem {
	return ipfsFS
}

// AddFile adds a single file to IPFS and returns its CID
func AddFile(ctx context.Context, filePath string) (string, error) {
	if !initialized {
		return "", ErrNotInitialized
	}

	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	fileNode := files.NewReaderFile(file)

	cidFile, err := ipfsClient.Unixfs().Add(ctx, fileNode)
	if err != nil {
		return "", fmt.Errorf("failed to add file to IPFS: %w", err)
	}

	return cidFile.String(), nil
}

// AddFolder adds a folder and its contents to IPFS and returns the CID of the folder
func AddFolder(ctx context.Context, folderPath string) (string, error) {
	if !initialized {
		return "", ErrNotInitialized
	}

	stat, err := os.Stat(folderPath)
	if err != nil {
		return "", fmt.Errorf("failed to get folder info: %w", err)
	}

	if !stat.IsDir() {
		return "", fmt.Errorf("provided path is not a directory")
	}

	folderNode, err := files.NewSerialFile(folderPath, false, stat)
	if err != nil {
		return "", fmt.Errorf("failed to create folder node: %w", err)
	}

	cidFolder, err := ipfsClient.Unixfs().Add(ctx, folderNode)
	if err != nil {
		return "", fmt.Errorf("failed to add folder to IPFS: %w", err)
	}

	return cidFolder.String(), nil
}

func GetCID(ctx context.Context, cid string) ([]byte, error) {
	if !initialized {
		return nil, ErrNotInitialized
	}

	if localMount {
		// Try to read from local filesystem first
		data, err := os.ReadFile(filepath.Join("/ipfs", cid))
		if err == nil {
			return data, nil
		}
		// If local read fails, fall back to IPFS client
	}

	// Use IPFS client to fetch the data
	p, err := path.NewPath("/ipfs/" + cid)
	if err != nil {
		return nil, err
	}
	n, err := ipfsClient.Unixfs().Get(ctx, p)
	if err != nil {
		return nil, err
	}

	return readNodeData(n)
}

func GetIPNS(ctx context.Context, name string) ([]byte, error) {
	if !initialized {
		return nil, ErrNotInitialized
	}

	if localMount {
		// Try to read from local filesystem first
		data, err := os.ReadFile(filepath.Join("/ipns", name))
		if err == nil {
			return data, nil
		}
		// If local read fails, fall back to IPFS client
	}

	// Use IPFS client to fetch the data
	p, err := path.NewPath("/ipns/" + name)
	if err != nil {
		return nil, err
	}
	n, err := ipfsClient.Unixfs().Get(ctx, p)
	if err != nil {
		return nil, err
	}

	return readNodeData(n)
}

func PinCID(ctx context.Context, cid string, name string) error {
	if !initialized {
		return ErrNotInitialized
	}

	p, err := path.NewPath(cid)
	if err != nil {
		return ErrNotInitialized
	}
	err = ipfsClient.Pin().Add(ctx, p)
	if err != nil {
		return ErrInternal
	}
	return nil
}

func readNodeData(n files.Node) ([]byte, error) {
	switch n := n.(type) {
	case files.File:
		return io.ReadAll(n)
	default:
		return nil, fmt.Errorf("unsupported node type: %T", n)
	}
}
