package schemas

import (
	"context"
	"sync"

	"github.com/ipld/go-ipld-prime/storage"
	"github.com/sonr-io/sonr/pkg/client"
)

type ReadableStore interface {
	storage.ReadableStorage
}

// Store implementation to abstract store operations
type ReadStoreImpl struct {
	mu     sync.Mutex
	cache  map[string][]byte
	Client *client.Client
}

func (rs *ReadStoreImpl) GetCache() map[string][]byte {
	if rs.cache == nil {
		rs.cache = make(map[string][]byte)
	}
	return rs.cache
}

func (rs *ReadStoreImpl) Has(ctx context.Context, key string) (bool, error) {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	if rs.cache == nil {
		rs.cache = make(map[string][]byte)
		return false, nil
	}

	if rs.cache[key] != nil {
		return true, nil
	}

	return false, nil
}

func (rs *ReadStoreImpl) Get(ctx context.Context, key string) ([]byte, error) {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	if rs.cache == nil {
		rs.cache = make(map[string][]byte)
	}

	if rs.cache[key] != nil {
		return rs.cache[key], nil
	}

	wi, err := rs.Client.QueryWhatIsByDid(key)

	if err != nil {
		return nil, err
	}

	buf, err := wi.Marshal()

	if err != nil {
		return nil, err
	}

	rs.cache[key] = buf

	return buf, nil
}
