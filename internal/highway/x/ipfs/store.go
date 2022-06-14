package ipfs

import (
	"context"
	"github.com/ipld/go-ipld-prime/storage"

	"github.com/ipfs/go-datastore"

	ipfslite "github.com/hsanjuan/ipfs-lite"
)

var _ storage.WritableStorage = &Store{}

// Store is a datastore.Batching implementation that stores data in memory.
type Store struct {
	dataStore datastore.Batching
}

// NewStore returns a new Store.
func NewStore() *Store {
	return &Store{
		dataStore: ipfslite.NewInMemoryDatastore(),
	}
}

// Batching returns the Store's datastore.Batching implementation.
func (ms *Store) Batching() datastore.Batching {
	return ms.dataStore
}

// Has checks if the key exists in the store.
func (ms *Store) Has(ctx context.Context, key string) (bool, error) {
	return ms.dataStore.Has(ctx, datastore.NewKey(key))
}

// Get retrieves the value stored in the Store under the given key.
func (ms *Store) Get(ctx context.Context, key string) ([]byte, error) {
	return ms.dataStore.Get(ctx, datastore.NewKey(key))
}

// Put stores the given value, keyed by the given string, into the Store.
func (ms *Store) Put(ctx context.Context, key string, content []byte) error {
	return ms.dataStore.Put(ctx, datastore.NewKey(key), content)
}
