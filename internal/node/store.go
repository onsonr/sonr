package node

import (
	"bytes"
	"context"
	"time"

	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/host"

	"github.com/sonr-io/core/tools/state"
	bolt "go.etcd.io/bbolt"
	"google.golang.org/protobuf/proto"
)

// RecentsHistory is a list of recent Peers.
type RecentsHistory map[string]*common.ProfileList

// Buckets in Database
var (
	RECENTS_BUCKET = []byte("recents")
	USER_BUCKET    = []byte("user")
)

// Bucket Constant Keys
var (
	PROFILE_KEY = []byte("profile")
)

// openStore creates a new Store instance for Node
func (n *Node) openStore(ctx context.Context, h *host.SNRHost, em *state.Emitter) error {
	path, err := device.NewDatabasePath("sonr-bolt.db")
	if err != nil {
		logger.Child("Store").Error("Failed to get DB Path", err)
		return err
	}

	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open(path, 0600, &bolt.Options{})
	if err != nil {
		logger.Child("Store").Error("Failed to open Database", err)
		return err
	}
	n.store = db

	// Create Profile Bucket
	err = n.createBucket(USER_BUCKET, PROFILE_KEY, n.options.profileBuf)
	if err != nil {
		return err
	}

	// Create Recents Bucket
	err = n.createBucket(RECENTS_BUCKET, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

// createBucket creates a new bucket in the store.
func (n *Node) createBucket(name []byte, key []byte, val []byte) error {
	// Check if Store is open
	if n.store == nil {
		logger.Child("Store").Error("Failed to Create Bucket", ErrStoreNotCreated)
		return ErrStoreNotCreated
	}
	if err := n.store.Update(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b, err := tx.CreateBucketIfNotExists(name)
		if err != nil {
			logger.Child("Store").Error("Failed to Create Bucket", ErrStoreNotCreated)
			return err
		}

		// Set Initial Value
		retVal := b.Get(key)
		if retVal != nil {
			return nil
		}

		// Check if Value is not nil
		if val != nil && key != nil {
			// Put in Bucket
			err = b.Put(key, val)
			if err != nil {
				return err
			}
		} else {
			logger.Warn("No initial key/value provided skipping first Put")
		}
		return nil
	}); err != nil {
		logger.Child("Store").Error("Failed to Create bucket", err)
		return err
	}
	return nil
}

// AddRecent stores the profile for recents in desk and returns list of recent profiles
func (n *Node) AddRecent(profile *common.Profile) error {
	// Check if Store is open
	if n.store == nil {
		logger.Child("Store").Error("Failed to Add Recent", ErrStoreNotCreated)
		return ErrStoreNotCreated
	}

	// Check if profile is nil
	if profile == nil {
		return ErrProfileNotProvided
	}

	// Create Key from Time in RFC3339 format
	t := time.Now().Round(time.Hour)
	keyStr := t.Format(time.RFC3339)
	key := []byte(keyStr)

	// Put in Bucket
	if err := n.store.Update(func(tx *bolt.Tx) error {
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
	}); err != nil {
		logger.Child("Store").Error("Failed to ADD Recent to Store", err)
		return err
	}
	return nil
}

// GetProfile returns the profile for the user from diskDB
func (n *Node) GetRecents() (RecentsHistory, error) {
	// Check if Store is open
	if n.store == nil {
		logger.Child("Store").Error("Failed to Get Recents", ErrStoreNotCreated)
		return nil, ErrStoreNotCreated
	}

	// Create empty map
	recents := make(RecentsHistory)
	now := time.Now().Round(time.Hour)
	start := now.Add(-time.Hour * 24 * 7)

	// Set Time Range for Recent Profiles
	nowStr := now.Format(time.RFC3339)
	startStr := start.Format(time.RFC3339)

	// Iterate over all profiles
	err := n.store.View(func(tx *bolt.Tx) error {
		// Get bucket
		b := tx.Bucket(RECENTS_BUCKET)
		if b == nil {
			return ErrRecentsNotCreated
		}

		// Assume our events bucket exists and has RFC3339 encoded time keys.
		c := b.Cursor()

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
	if err != nil {
		logger.Child("Store").Error("Failed to GET Recents in Store", err)
		return nil, err
	}
	return recents, nil
}

// Profile returns the profile for the user from diskDB
func (n *Node) Profile() (*common.Profile, error) {
	// Check if Store is open
	if n.store == nil {
		logger.Error("Failed to Get Profile", ErrStoreNotCreated)
		return nil, ErrStoreNotCreated
	}

	var profile common.Profile
	err := n.store.View(func(tx *bolt.Tx) error {
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
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

// SetProfile stores the profile for the user in diskDB
func (n *Node) SetProfile(profile *common.Profile) error {
	// Check if Store is open
	if n.store == nil {
		logger.Child("Store").Error("Failed to Set Profile", ErrStoreNotCreated)
		return ErrStoreNotCreated
	}

	// Check if profile is nil
	if profile == nil {
		return ErrProfileNotProvided
	}

	// Put in Bucket
	if err := n.store.Update(func(tx *bolt.Tx) error {
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
	}); err != nil {
		logger.Child("Store").Error("Failed to SET Profile in Store", err)
		return err
	}
	return nil
}
