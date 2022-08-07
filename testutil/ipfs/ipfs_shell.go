package ipfs

import (
	"context"
	"crypto/sha256"
	"errors"

	"github.com/ipfs/go-datastore"
)

type IpfsShellMock struct {
	cache datastore.Datastore
}

func NewMockShell(cacheStore datastore.Datastore) *IpfsShellMock {
	return &IpfsShellMock{cache: cacheStore}
}

func (i *IpfsShellMock) DagGet(ctx context.Context, ref string, out interface{}) error {
	return errors.New("not yet implemented")
}

func (i *IpfsShellMock) DagPut(ctx context.Context, data interface{}, inputCodec, storeCodec string) (string, error) {
	return "", errors.New("not yet implemented")
}

func (i *IpfsShellMock) GetData(ctx context.Context, cid string) ([]byte, error) {
	key := datastore.NewKey(cid)
	data, err := i.cache.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (i *IpfsShellMock) PutData(ctx context.Context, data []byte) (string, error) {
	cidStr := string(sha256.New().Sum(nil))

	err := i.cache.Put(ctx, datastore.NewKey(cidStr), data)
	if err != nil {
		return "", err
	}

	return cidStr, nil
}

func (i *IpfsShellMock) PinFile(ctx context.Context, cidStr string) error {
	return errors.New("not implemented yet")
}

func (i *IpfsShellMock) RemoveFile(ctx context.Context, cidstr string) error {
	return errors.New("not implemented yet")
}
