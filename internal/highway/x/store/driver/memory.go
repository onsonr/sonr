package driver

import (
	"context"
	ipfslite "github.com/hsanjuan/ipfs-lite"
	"github.com/ipfs/go-datastore"
	"github.com/ipld/go-ipld-prime/storage"
)

// Memory is a datastore.Batching implementation that stores data in memory.
type Memory struct {
	storage.WritableStorage
	dataStore datastore.Batching
}

// NewMemoryDriver returns a new Memory.
func NewMemoryDriver() *Memory {
	return &Memory{
		dataStore: ipfslite.NewInMemoryDatastore(),
	}
}

// Batching returns the Memory's datastore.Batching implementation.
func (ms *Memory) Batching() datastore.Batching {
	return ms.dataStore
}

// Get retrieves the value stored in the Memory under the given key.
func (ms *Memory) Get(ctx context.Context, key string) ([]byte, error) {
	return ms.dataStore.Get(ctx, datastore.NewKey(key))
}

// Put stores the given value, keyed by the given string, into the Memory.
func (ms *Memory) Put(ctx context.Context, key string, content []byte) error {
	return ms.dataStore.Put(ctx, datastore.NewKey(key), content)
}
