package ipfs

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/client/rpc"
)

type client struct {
	api *rpc.HttpApi
}

func NewClient() (Client, error) {
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
	p, err := path.NewPath(cid)
	if err != nil {
		return nil, err
	}
	node, err := c.api.Unixfs().Get(context.Background(), p)
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
	p, err := path.NewPath(cid)
	if err != nil {
		return false, err
	}
	_, err = c.api.Block().Stat(context.Background(), p)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (c *client) Pin(cid string) error {
	p, err := path.NewPath(cid)
	if err != nil {
		return err
	}
	return c.api.Pin().Add(context.Background(), p)
}

func (c *client) Unpin(cid string) error {
	p, err := path.NewPath(cid)
	if err != nil {
		return err
	}
	return c.api.Pin().Rm(context.Background(), p)
}

func (c *client) Publish(cid string, name string) (string, error) {
	p, err := path.NewPath(cid)
	if err != nil {
		return "", err
	}
	result, err := c.api.Name().Publish(context.Background(), p)
	if err != nil {
		return "", err
	}
	return result.String(), nil
}

func (c *client) Ls(cid string) ([]string, error) {
	p, err := path.NewPath(cid)
	if err != nil {
		return nil, err
	}
	node, err := c.api.Unixfs().Ls(context.Background(), p)
	if err != nil {
		return nil, err
	}

	var files []string
	for entry := range node {
		files = append(files, entry.Name)
	}
	return files, nil
}
