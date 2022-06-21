package store

import (
	"context"
	"errors"
	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
	"github.com/ipld/go-ipld-prime/storage/fsstore"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	_ datastore.PersistentDatastore = (*Store)(nil)
	_ datastore.Batching            = (*Store)(nil)
)

var ObjectKeySuffix = ".dsobject"

// Store is a datastore.Batching implementation that stores data in memory.
type Store struct {
	path      string
	dataStore *fsstore.Store
}

// New returns a new Store.
func New(path string) *Store {
	store := &fsstore.Store{}
	err := store.InitDefaults(path)
	if err != nil {
		panic(err)
	}

	return &Store{
		path:      path,
		dataStore: store,
	}
}

func (s *Store) Get(ctx context.Context, key datastore.Key) (value []byte, err error) {
	return s.dataStore.Get(ctx, key.String())
}

func (s *Store) Has(ctx context.Context, key datastore.Key) (exists bool, err error) {
	return s.dataStore.Has(ctx, key.String())
}

func (s *Store) GetSize(ctx context.Context, key datastore.Key) (size int, err error) {
	return 0, errors.New("TODO")
}

func (s *Store) Query(ctx context.Context, q query.Query) (query.Results, error) {
	log.Printf("querying for %+v", q)

	results := make(chan query.Result)

	walkFn := func(path string, info os.FileInfo, _ error) error {
		// remove ds path prefix
		relPath, err := filepath.Rel(s.path, path)
		if err == nil {
			path = filepath.ToSlash(relPath)
		}

		if !info.IsDir() {
			path = strings.TrimSuffix(path, ObjectKeySuffix)
			var result query.Result
			key := datastore.NewKey(path)
			result.Entry.Key = key.String()
			if !q.KeysOnly {
				result.Entry.Value, result.Error = s.Get(ctx, key)
			}
			results <- result
		}
		return nil
	}

	defer func() {
		filepath.Walk(s.path, walkFn)
		close(results)
	}()

	r := query.ResultsWithChan(q, results)
	r = query.NaiveQueryApply(q, r)
	return r, nil
}

func (s *Store) Put(ctx context.Context, key datastore.Key, value []byte) error {
	return s.dataStore.Put(ctx, key.String(), value)

}

// KeyFilename returns the filename associated with `key`
func (s *Store) KeyFilename(key datastore.Key) string {
	return filepath.Join(s.path, key.String(), ObjectKeySuffix)
}

func (s *Store) Delete(ctx context.Context, key datastore.Key) error {
	fn := s.KeyFilename(key)
	if !isFile(fn) {
		return nil
	}

	err := os.Remove(fn)
	if os.IsNotExist(err) {
		err = nil
	}
	return err
}

// isDir returns whether given path is a directory
func isDir(path string) bool {
	finfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	return finfo.IsDir()
}

// isFile returns whether given path is a file
func isFile(path string) bool {
	finfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	return !finfo.IsDir()
}

func (s *Store) Sync(ctx context.Context, prefix datastore.Key) error {
	return errors.New("TODO")
}

func (s *Store) Close() error {
	return errors.New("TODO")
}

func (s *Store) DiskUsage(ctx context.Context) (uint64, error) {
	return 0, errors.New("TODO")
}

func (s *Store) Batch(ctx context.Context) (datastore.Batch, error) {
	return nil, errors.New("TODO")
}
