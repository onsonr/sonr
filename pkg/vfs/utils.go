package vfs

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
)

// Helper function to parse IPFS path
func ParsePath(p string) (path.Path, error) {
	return path.NewPath(p)
}

// FetchVaultPath returns the path to the vault with the given name
func FetchVaultPath(name string) string {
	return enclaveRoot.Join(name).Path()
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
