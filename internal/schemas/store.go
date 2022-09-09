package schemas

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/ipld/go-ipld-prime/storage"
	"github.com/sonr-io/sonr/pkg/client"
)

type ReadableStore interface {
	storage.ReadableStorage
}

// Store implementation to abstract store operations
type readStoreImpl struct {
	mu     sync.Mutex
	cache  map[string][]byte
	client *client.Client
}

func (rs *readStoreImpl) Has(ctx context.Context, key string) (bool, error) {
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

func (rs *readStoreImpl) Get(ctx context.Context, key string) ([]byte, error) {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	if rs.cache == nil {
		rs.cache = make(map[string][]byte)
	}

	if rs.cache[key] != nil {
		return rs.cache[key], nil
	}

	time_stamp := fmt.Sprintf("%d", time.Now().Unix())

	out_path := filepath.Join(os.TempDir(), key+time_stamp+".txt")
	defer os.Remove(out_path)

	wi, err := rs.client.QueryWhatIsByDid(key)

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
