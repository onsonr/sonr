package fs

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/core/coreiface/options"
)

// Constant for the name of the folder where the vaults are stored
const kVaultsFolderName = ".sonr-vaults"

// VaultsFolder is the folder where the vaults are stored
var VaultsFolder Folder

// Folder represents a folder in the filesystem
type Folder string

// Package initializes the VaultsFolder
func init() {
	// Initialize VaultsFolder
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	// Create the folder if it does not exist
	VaultsFolder = NewFolder(filepath.Join(homeDir, kVaultsFolderName))
	if !VaultsFolder.Exists() {
		if err := VaultsFolder.Create(); err != nil {
			panic(err)
		}
	}
}

// NewFolder creates a new Folder instance
func NewFolder(path string) Folder {
	return Folder(path)
}

// NewVaultFolder creates a new folder under the VaultsFolder directory
func NewVaultFolder(name string) (Folder, error) {
	vaultFolder := VaultsFolder.Join(name)
	err := vaultFolder.Create()
	if err != nil {
		return "", err
	}
	return vaultFolder, nil
}

// Name returns the name of the folder
func (f Folder) Name() string {
	return filepath.Base(string(f))
}

// Path returns the path of the folder
func (f Folder) Path() string {
	return string(f)
}

// Create creates the folder if it doesn't exist
func (f Folder) Create() error {
	return os.MkdirAll(string(f), os.ModePerm)
}

// Exists checks if the folder exists
func (f Folder) Exists() bool {
	info, err := os.Stat(string(f))
	return err == nil && info.IsDir()
}

// Ls lists the contents of the folder
func (f Folder) Ls() ([]fs.DirEntry, error) {
	return os.ReadDir(string(f))
}

// WriteFile writes data to a file in the folder
func (f Folder) WriteFile(name string, data []byte, perm os.FileMode) error {
	return os.WriteFile(filepath.Join(string(f), name), data, perm)
}

// ReadFile reads the contents of a file in the folder
func (f Folder) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(filepath.Join(string(f), name))
}

// CopyFile copies a file from src to dst within the folder
func (f Folder) CopyFile(src, dst string) error {
	srcPath := filepath.Join(string(f), src)
	dstPath := filepath.Join(string(f), dst)

	sourceFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

// DeleteFile deletes a file from the folder
func (f Folder) DeleteFile(name string) error {
	return os.Remove(filepath.Join(string(f), name))
}

// Remove removes the folder and its contents
func (f Folder) Remove() error {
	return os.RemoveAll(string(f))
}

// Join joins the folder path with the given elements
func (f Folder) Join(elem ...string) Folder {
	return Folder(filepath.Join(append([]string{string(f)}, elem...)...))
}

// IsDir checks if the folder is a directory
func (f Folder) IsDir() bool {
	info, err := os.Stat(string(f))
	return err == nil && info.IsDir()
}

// Touch creates an empty file if it doesn't exist, or updates its access and modification times if it does
func (f Folder) Touch(name string) (File, error) {
	var err error
	filePath := filepath.Join(string(f), name)
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return File(filePath), err
		}
		if err := file.Close(); err != nil {
			return File(filePath), err
		}
		return File(filePath), nil
	}
	return File(filePath), err
}

// Node returns a files.Node representation of the folder
func (f Folder) Node() (files.Node, error) {
	entries, err := f.Ls()
	if err != nil {
		return nil, err
	}
	dirEntries := make([]files.DirEntry, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			subFolder := f.Join(entry.Name())
			subNode, err := subFolder.Node()
			if err != nil {
				return nil, err
			}
			dirEntries = append(dirEntries, files.FileEntry(entry.Name(), subNode))
		} else {
			file := File(filepath.Join(string(f), entry.Name()))
			fileNode, err := file.Node()
			if err != nil {
				return nil, err
			}
			dirEntries = append(dirEntries, files.FileEntry(entry.Name(), fileNode))
		}
	}
	return files.NewSliceDirectory(dirEntries), nil
}

// SaveToIPFS saves the Folder to IPFS and returns the IPFS path
func (f Folder) SaveToIPFS(ctx context.Context) (path.Path, error) {
	node, err := f.Node()
	if err != nil {
		return nil, err
	}
	c, err := getIPFSClient()
	if err != nil {
		return nil, err
	}
	return c.Unixfs().Add(ctx, node)
}

// PublishToIPNS publishes the Folder to IPNS
func (f Folder) PublishToIPNS(ctx context.Context, ipfsPath path.Path) error {
	c, err := getIPFSClient()
	if err != nil {
		return err
	}
	_, err = c.Name().Publish(ctx, ipfsPath, options.Name.Key(f.Name()))
	return err
}

// LoadFromIPFS loads a Folder from IPFS
func LoadFromIPFS(ctx context.Context, path string) (Folder, error) {
	c, err := getIPFSClient()
	if err != nil {
		return "", err
	}
	cid, err := ParsePath(path)
	if err != nil {
		return "", err
	}
	node, err := c.Unixfs().Get(ctx, cid)
	if err != nil {
		return "", err
	}
	return LoadNodeInFolder(path, node)
}

// LoadNodeInFolder loads an IPFS node into a Folder
func LoadNodeInFolder(path string, node files.Node) (Folder, error) {
	folder := NewFolder(path)
	if err := folder.Create(); err != nil {
		return "", err
	}

	switch n := node.(type) {
	case files.File:
		f, err := os.Create(folder.Path())
		if err != nil {
			return "", err
		}
		defer f.Close()
		if _, err := io.Copy(f, n); err != nil {
			return "", err
		}
	case files.Directory:
		entries := n.Entries()
		for entries.Next() {
			name := entries.Name()
			childNode := entries.Node()
			childPath := filepath.Join(folder.Path(), name)
			if _, err := LoadNodeInFolder(childPath, childNode); err != nil {
				return "", err
			}
		}
		if entries.Err() != nil {
			return "", entries.Err()
		}
	default:
		return "", fmt.Errorf("unsupported node type: %T", n)
	}

	return folder, nil
}
