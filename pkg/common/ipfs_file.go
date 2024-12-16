package common

import (
	"context"

	"github.com/ipfs/boxo/files"
)

type file struct {
	files.File
	name string
}

func (f *file) Name() string {
	return f.name
}

func NewFile(name string, data []byte) File {
	return &file{File: files.NewBytesFile(data), name: name}
}

func (c *client) AddFile(file File) (string, error) {
	cidFile, err := c.api.Unixfs().Add(context.Background(), file)
	if err != nil {
		return "", err
	}
	return cidFile.String(), nil
}
