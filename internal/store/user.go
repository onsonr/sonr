package store

import (
	"time"

	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/tools/logger"
	bolt "go.etcd.io/bbolt"
	"google.golang.org/protobuf/proto"
)

// GetProfile returns the profile for the user from diskDB
func (s *Store) GetProfile() (*common.Profile, error) {
	var profile common.Profile
	err := s.db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket(USER_BUCKET)

		// Check if bucket exists
		if b == nil {
			return ErrProfileNotCreated
		}

		// Get profile buffer
		buf := b.Get(PROFILE_KEY)
		if buf == nil {
			return nil
		}

		// Unmarshal profile
		err := proto.Unmarshal(buf, &profile)
		if err != nil {
			return err
		}
		return nil
	})
	return &profile, s.checkGetErr(err)
}

// SetProfile stores the profile for the user in diskDB
func (s *Store) SetProfile(profile *common.Profile) error {
	// Check if profile is nil
	if profile == nil {
		return ErrProfileNotProvided
	}

	// Compare current profile with new profile
	isNewProfile := false
	currentProfile, err := s.GetProfile()
	if err != nil {
		if err == ErrProfileNotCreated {
			profile.LastModified = time.Now().Unix()
			isNewProfile = true
		} else {
			return logger.Error("Failed to set Profile", err)
		}
	}

	// Check if given profile has Timestamp
	if !isNewProfile && profile.GetLastModified() == 0 {
		return ErrProfileNoTimestamp
	}

	// Verify timestamp
	if !isNewProfile && profile.LastModified < currentProfile.GetLastModified() {
		return ErrProfileIsOlder
	}

	// Put in Bucket
	return s.db.Update(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b, err := tx.CreateBucketIfNotExists(USER_BUCKET)
		if err != nil {
			return err
		}

		// Marshal profile
		buf, err := proto.Marshal(profile)
		if err != nil {
			return err
		}

		// Put profile
		err = b.Put(PROFILE_KEY, buf)
		if err != nil {
			return err
		}
		return nil
	})
}
