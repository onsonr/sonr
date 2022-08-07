package protocol

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/ipfs/go-datastore"
	shell "github.com/ipfs/go-ipfs-api"
)

var _ IPFS = (*IPFSShell)(nil)

// IPFSShell implements a protocol.IPFS featuring:
// 	- IPFS initialization and error fallback
//	- In-memory cache mechanism
type IPFSShell struct {
	shell *shell.Shell
	cache datastore.Datastore
}

func NewIPFSShell(url string, cacheStore datastore.Datastore) *IPFSShell {
	return &IPFSShell{shell: shell.NewShell(url), cache: cacheStore}
}

func (i *IPFSShell) DagGet(ctx context.Context, ref string, out interface{}) error {
	key := datastore.NewKey(ref)
	cached, err := i.cache.Get(ctx, key)

	if err != nil {
		return err
	}

	if len(cached) > 0 {
		out = cached
	}

	return i.shell.DagGet(ref, out)
}

func (i *IPFSShell) DagPut(ctx context.Context, data interface{}, inputCodec, storeCodec string) (string, error) {
	panic("TODO")
}

func (i *IPFSShell) GetData(ctx context.Context, cid string) ([]byte, error) {
	key := datastore.NewKey(cid)
	data, err := i.cache.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	if data != nil {
		return data, nil
	}

	outPath := path.Join(os.TempDir(), "data", fmt.Sprintf("%s-%d", cid, time.Now().Unix()))
	if err = i.shell.Get(cid, outPath); err != nil {
		return nil, err
	}

	buf, err := os.ReadFile(outPath)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(buf, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func (i *IPFSShell) PutData(ctx context.Context, data []byte) (string, error) {
	cidStr, err := i.shell.Add(bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}

	err = i.cache.Put(ctx, datastore.NewKey(cidStr), data)
	if err != nil {
		return "", err
	}

	return cidStr, nil
}

func (i *IPFSShell) PinFile(ctx context.Context, cidStr string) error {
	return i.shell.Pin(cidStr)
}

func (i *IPFSShell) RemoveFile(ctx context.Context, cidstr string) error {
	return errors.New("not implemented yet")
}
