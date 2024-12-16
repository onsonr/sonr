package common

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/client/rpc"
	"github.com/ipfs/kubo/core/coreiface/options"
)

// IPFS represents a wrapper interface abstracting the localhost api
type IPFS interface {
	Add(data []byte) (string, error)
	AddFile(file File) (string, error)
	AddFolder(folder Folder) (string, error)
	Exists(cid string) (bool, error)
	Get(cid string) ([]byte, error)
	IsPinned(ipns string) (bool, error)
	Ls(cid string) ([]string, error)
	Pin(cid string, name string) error
	Unpin(cid string) error
}

type File interface {
	files.File
	Name() string
}

func NewFileMap(vs []File) map[string]files.Node {
	m := make(map[string]files.Node)
	for _, f := range vs {
		m[f.Name()] = f
	}
	return m
}

type client struct {
	api *rpc.HttpApi
}

func NewIPFS() (IPFS, error) {
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

func (c *client) IsPinned(ipns string) (bool, error) {
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

func (c *client) Pin(cid string, name string) error {
	p, err := path.NewPath(cid)
	if err != nil {
		return err
	}
	return c.api.Pin().Add(context.Background(), p, options.Pin.Name(name))
}

func (c *client) Unpin(cid string) error {
	p, err := path.NewPath(cid)
	if err != nil {
		return err
	}
	return c.api.Pin().Rm(context.Background(), p)
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
