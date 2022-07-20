package store

import (
	"context"
	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
	"github.com/sonr-io/sonr/internal/highway/x/store/driver"
	"os"
)

var _ datastore.Batching = (*Store)(nil)

type Store struct {
	fs  *driver.FileSystem
	mem *driver.Memory
}

func (s Store) Get(ctx context.Context, key datastore.Key) (value []byte, err error) {
	//TODO implement me
	panic("implement me")
}

func (s Store) Has(ctx context.Context, key datastore.Key) (exists bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (s Store) GetSize(ctx context.Context, key datastore.Key) (size int, err error) {
	//TODO implement me
	panic("implement me")
}

func (s Store) Query(ctx context.Context, q query.Query) (query.Results, error) {
	//TODO implement me
	panic("implement me")
}

func (s Store) Put(ctx context.Context, key datastore.Key, value []byte) error {
	//TODO implement me
	panic("implement me")
}

func (s Store) Delete(ctx context.Context, key datastore.Key) error {
	//TODO implement me
	panic("implement me")
}

func (s Store) Sync(ctx context.Context, prefix datastore.Key) error {
	//TODO implement me
	panic("implement me")
}

func (s Store) Close() error {
	//TODO implement me
	panic("implement me")
}

func (s Store) Batch(ctx context.Context) (datastore.Batch, error) {
	//TODO implement me
	panic("implement me")
}

func NewStore() *Store {
	memStore := driver.NewMemoryDriver()
	fsStore := driver.NewFileSystemDriver(os.TempDir())

	return &Store{
		fs:  fsStore,
		mem: memStore,
	}
}
