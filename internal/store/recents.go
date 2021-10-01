package store

import (
	"bytes"
	"time"

	"github.com/sonr-io/core/internal/common"
	bolt "go.etcd.io/bbolt"
	"google.golang.org/protobuf/proto"
)



// GetProfile returns the profile for the user from diskDB
func (s *Store) GetRecents() (RecentsHistory, error) {
	// Create empty map
	recents := make(RecentsHistory)
	now := time.Now().Round(time.Hour)
	start := now.Add(-time.Hour * 24 * 7)

	// Set Time Range for Recent Profiles
	nowStr := now.Format(time.RFC3339)
	startStr := start.Format(time.RFC3339)

	// Iterate over all profiles
	err := s.db.View(func(tx *bolt.Tx) error {
		// Assume our events bucket exists and has RFC3339 encoded time keys.
		c := tx.Bucket(RECENTS_BUCKET).Cursor()

		// Our time range spans the 90's decade.
		min := []byte(startStr)
		max := []byte(nowStr)

		// Iterate over the 90's.
		for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
			// Unmarshal profile list
			profileList := common.ProfileList{}
			err := proto.Unmarshal(v, &profileList)
			if err != nil {
				return err
			}

			// Add to map
			recents[string(k)] = &profileList
		}
		return nil
	})
	return recents, err
}

// AddRecent stores the profile for recents in desk and returns list of recent profiles
func (s *Store) AddRecent(profile *common.Profile) error {
	// Create Key from Time in RFC3339 format
	t := time.Now().Round(time.Hour)
	keyStr := t.Format(time.RFC3339)
	key := []byte(keyStr)

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
