package store

import (
	"errors"
	"time"

	"github.com/sonr-io/core/internal/common"
	bolt "go.etcd.io/bbolt"
	"google.golang.org/protobuf/proto"
)

type RecentsHistory map[string]*common.ProfileList

// GetProfile returns the profile for the user from diskDB
func (s *Store) GetRecents() (RecentsHistory, error) {
	// Create empty map
	recents := make(RecentsHistory)
	key := []byte("TO-DO")

	// Iterate over all profiles
	err := s.db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket(RECENTS_BUCKET)

		// Check if bucket exists
		if b == nil {
			return errors.New("Bucket does not exist")
		}

		// Get profile list buffer
		buf := b.Get(key)
		if buf == nil {
			return nil
		}

		// Unmarshal profile list
		profileList := common.ProfileList{}
		err := proto.Unmarshal(buf, &profileList)
		if err != nil {
			return err
		}

		// Add to map
		recents[string(key)] = &profileList
		return nil
	})
	return recents, err
}

// AddRecent stores the profile for recents in desk and returns list of recent profiles
func (s *Store) AddRecent(profile *common.Profile) error {
	key := []byte("TO-DO")
	t := time.Now().Round(time.Hour)
	t.Format(time.RFC3339)
	// Put in Bucket
	return s.db.Update(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b, err := tx.CreateBucketIfNotExists(RECENTS_BUCKET)
		if err != nil {
			return err
		}

		// Get profile list buffer
		oldBuf := b.Get(key)
		if oldBuf == nil {
			return nil
		}

		// Unmarshal profile list
		profileList := common.ProfileList{}
		err = proto.Unmarshal(oldBuf, &profileList)
		if err != nil {
			return err
		}

		// Add profile to list
		profileList.Add(profile)

		// Marshal profile
		buf, err := proto.Marshal(&profileList)
		if err != nil {
			return err
		}

		// Put profile
		err = b.Put(key, buf)
		if err != nil {
			return err
		}
		return nil
	})
}
