package ipfs

import (
	"context"

	"github.com/ipfs/go-datastore"

	ipfslite "github.com/hsanjuan/ipfs-lite"
	"github.com/ipld/go-ipld-prime/storage"
)

// MemoryStore is a datastore.Batching implementation that stores data in memory.
type MemoryStore struct {
	storage.WritableStorage
	dataStore datastore.Batching
}

// NewMemoryStore returns a new MemoryStore.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		dataStore: ipfslite.NewInMemoryDatastore(),
	}
}

// Batching returns the MemoryStore's datastore.Batching implementation.
func (ms *MemoryStore) Batching() datastore.Batching {
	return ms.dataStore
}

// Get retrieves the value stored in the MemoryStore under the given key.
func (ms *MemoryStore) Get(ctx context.Context, key string) ([]byte, error) {
	return ms.dataStore.Get(ctx, datastore.NewKey(key))
}

// Put stores the given value, keyed by the given string, into the MemoryStore.
func (ms *MemoryStore) Put(ctx context.Context, key string, content []byte) error {
	return ms.dataStore.Put(ctx, datastore.NewKey(key), content)
}
