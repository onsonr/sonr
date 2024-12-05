package ipfs

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"path"

	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/client/rpc"
)

type client struct {
	api *rpc.HttpApi
}

func NewClient() (*client, error) {
	api, err := rpc.NewLocalApi()
	if err != nil {
		return nil, err
	}
	return &client{api: api}, nil
}

func (c *client) Add(data []byte) (string, error) {
	file := files.NewBytesFile(data)
	cidFile, err := c.api.Unixfs().Add(context.Background(), file)
	if err != nil {
		return "", err
	}
	return cidFile.String(), nil
}

func (c *client) Get(cid string) ([]byte, error) {
	node, err := c.api.Unixfs().Get(context.Background(), path.New(cid))
	if err != nil {
		return nil, err
	}

	file, ok := node.(files.File)
	if !ok {
		return nil, fmt.Errorf("unexpected node type: %T", node)
	}

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c *client) IsPublished(ipns string) (bool, error) {
	_, err := c.api.Name().Resolve(context.Background(), ipns)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (c *client) Exists(cid string) (bool, error) {
	_, err := c.api.Block().Stat(context.Background(), path.New(cid))
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (c *client) Pin(cid string) error {
	return c.api.Pin().Add(context.Background(), path.New(cid))
}

func (c *client) Unpin(cid string) error {
	return c.api.Pin().Rm(context.Background(), path.New(cid))
}

func (c *client) Publish(cid string, name string) (string, error) {
	result, err := c.api.Name().Publish(context.Background(), path.New(cid))
	if err != nil {
		return "", err
	}
	return result.Name(), nil
}

func (c *client) LsFolder(pathStr string) ([]string, error) {
	node, err := c.api.Unixfs().Ls(context.Background(), path.New(pathStr))
	if err != nil {
		return nil, err
	}

	var files []string
	for entry := range node {
		files = append(files, entry.Name)
	}
	return files, nil
}
