package fs

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/interface-go-ipfs-core/options"
)

// Upload uploads a file to the vault
func (c *Config) Upload(body []byte, name string) error {
	// Create path for file to be stored and write file
	path := filepath.Join(c.localPath, "public", name)
	err := os.WriteFile(path, body, 0644)
	if err != nil {
		return err
	}

	// Get file stats
	st, err := os.Stat(path)
	if err != nil {
		return err
	}

	// Create IPFS file
	f, err := files.NewSerialFile(path, false, st)
	if err != nil {
		return err
	}

	// Add file to IPFS
	cid, err := c.ipfs.Unixfs().Add(c.ctx, f, options.Unixfs.Pin(true))
	if err != nil {
		return err
	}
	fmt.Println(cid)
	return errors.New("Method not implemented")
}

// A method that is part of the vaultFsImpl struct. It takes a string and returns a byte array and an
// error.
func (c *Config) Download(name string) ([]byte, error) {
	// List all files in the _auth directory
	files, err := os.ReadDir(filepath.Join(c.localPath, "public"))
	if err != nil {
		return nil, err
	}

	// Find the file with the given name
	for _, file := range files {
		if file.Name() == name {
			// Read the file
			body, err := os.ReadFile(filepath.Join(c.localPath, "public", file.Name()))
			if err != nil {
				return nil, err
			}
			return body, nil
		}
	}
	return nil, errors.New("File not found")
}
