package store

import (
	"github.com/ipfs/go-datastore"
	"github.com/ipld/go-ipld-prime/storage"
)

// Store is a read-write storage with datastore.Batching support.
type Store interface {
	storage.ReadableStorage
	storage.WritableStorage
	Batching() datastore.Batching
}
