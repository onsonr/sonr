package store

import (
	"context"

	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/state"
	bolt "go.etcd.io/bbolt"
)


type Store struct {
	ctx     context.Context
	db      *bolt.DB
	emitter *state.Emitter
	host    *host.SNRHost
}

// NewStore creates a new Store instance.
func NewStore(ctx context.Context, h *host.SNRHost, em *state.Emitter) (*Store, error) {
	path, err := device.NewDatabasePath("sonr-bolt.db")
	if err != nil {
		return nil, logger.Error("Failed to get DB Path", err)
	}

	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open(path, 0600, &bolt.Options{})
	if err != nil {
		return nil, logger.Error("Failed to open Database", err)
	}

	return &Store{
		ctx:     ctx,
		db:      db,
		emitter: em,
		host:    h,
	}, nil
}

// Close closes the store.
func (s *Store) Close() error {
	return s.db.Close()
}
