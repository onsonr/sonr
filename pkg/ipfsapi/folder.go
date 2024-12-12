package ipfsapi

import (
	"context"

	"github.com/ipfs/boxo/files"
)

type Folder = files.Directory

func NewFolder(fs ...File) Folder {
	return files.NewMapDirectory(convertFilesToMap(fs))
}

func (c *client) AddFolder(folder Folder) (string, error) {
	cidFile, err := c.api.Unixfs().Add(context.Background(), folder)
	if err != nil {
		return "", err
	}
	return cidFile.String(), nil
}
