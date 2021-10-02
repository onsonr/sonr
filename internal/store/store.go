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

// createBucket creates a new bucket in the store.
func (s *Store) createBucket(key []byte) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		_, err := tx.CreateBucketIfNotExists(USER_BUCKET)
		if err != nil {
			return logger.Error("Failed to create new bucket", err)
		}
		return nil
	})
}

// Close closes the store.
func (s *Store) Close() error {
	return s.db.Close()
}

// checkGetErr checks if an error occurred and if so, handles it.
func (s *Store) checkGetErr(err error) error {
	if err != nil {
		// Check if profile bucket not created
		if err == ErrProfileNotCreated {
			logger.Debug("No Profile Bucket found, Creating new one...")

			// Check if bucket was created
			err = s.createBucket(USER_BUCKET)
			if err != nil {
				return err
			}
			return nil
		}

		// Check if recents bucket not created
		if err == ErrRecentsNotCreated {
			logger.Debug("No Recents Bucket found, Creating new one...")

			// Check if bucket was created
			err = s.createBucket(RECENTS_BUCKET)
			if err != nil {
				return err
			}
			return nil
		}

		// Other error
		return err
	}
	return nil
}
