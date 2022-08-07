package schemas

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/ipld/go-ipld-prime/storage"
)

type ReadableStore interface {
	storage.ReadableStorage
}

// Store implementation to abstract store operations
type readStoreImpl struct {
	mu    sync.Mutex
	cache map[string][]byte
	shell *shell.Shell
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

	err := rs.shell.Get(key, out_path)

	if err != nil {
		return nil, err
	}

	buf, err := os.ReadFile(out_path)

	if err != nil {
		return nil, err
	}

	rs.cache[key] = buf

	return buf, nil
}
