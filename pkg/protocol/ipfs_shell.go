package protocol

import (
	"bytes"
	"context"

	"github.com/ipfs/go-datastore"
	shell "github.com/ipfs/go-ipfs-api"
)

// IPFSShell implements a protocol.IPFS featuring:
// 	- IPFS initialization and error fallback
//	- In-memory cache mechanism
type IPFSShell struct {
	*shell.Shell
	cache datastore.Datastore
}

func NewIPFSShell(url string, cacheStore datastore.Datastore) *IPFSShell {
	return &IPFSShell{Shell: shell.NewShell(url), cache: cacheStore}
}

func (i *IPFSShell) GetData(ctx context.Context, cid string) ([]byte, error) {
	panic("implement me")
}

func (i *IPFSShell) PutData(ctx context.Context, data []byte) (string, error) {
	cidStr, err := i.Add(bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}

	err = i.cache.Put(ctx, datastore.NewKey(cidStr), data)
	if err != nil {
		return "", err
	}

	return cidStr, nil
}

func (i *IPFSShell) PinFile(ctx context.Context, cidstr string) error {
	panic("implement me")
}

func (i *IPFSShell) RemoveFile(ctx context.Context, cidstr string) error {
	panic("implement me")
}
