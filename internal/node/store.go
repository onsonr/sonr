package node

import (
	"context"
	"time"

	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"

	"git.mills.io/prologic/bitcask"
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
func (n *Node) openStore(ctx context.Context, opts *nodeOptions) error {
	path, err := device.NewDatabasePath("sonr_bitcask")
	if err != nil {
		logger.Error("Failed to get DB Path", err)
		return err
	}

	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bitcask.Open(path)
	if err != nil {
		logger.Error("Failed to open Database", err)
		return err
	}
	n.store = db

	// Create Profile Bucket
	err = n.SetProfile(opts.profile)
	if err != nil {
		logger.Error("Failed to Set Profile", err)
		return err
	}
	return nil
}

// AddRecent stores the profile for recents in desk and returns list of recent profiles
func (n *Node) AddRecent(profile *common.Profile) error {
	// Check if Store is open
	if n.store == nil {
		logger.Error("Failed to Add Recent", ErrProtocolsNotSet)
		return ErrProtocolsNotSet
	}

	// Check if profile is nil
	if profile == nil {
		logger.Error("Failed to Add Recent", ErrProfileNotProvided)
		return ErrProfileNotProvided
	}

	// Create Key from Time in RFC3339 format
	t := time.Now().Round(time.Hour)
	keyStr := t.Format(time.RFC3339)
	key := []byte(keyStr)

	// Put in Bucket
	if n.store.Has(key) {
		// Get profile list buffer
		oldBuf, err := n.store.Get(key)
		if err != nil {
			logger.Error("Failed to Get old Recents from store")
			return err
		}

		// Unmarshal profile list
		profileList := common.ProfileList{}
		err = proto.Unmarshal(oldBuf, &profileList)
		if err != nil {
			logger.Error("Failed to Unmarshal ProfileList")
			return err
		}

		// Add profile to list
		profileList.Add(profile)

		// Marshal profile
		val, err := proto.Marshal(&profileList)
		if err != nil {
			logger.Error("Failed to Marshal ProfileList")
			return err
		}

		err = n.store.Put(key, val)
		if err != nil {
			return err
		}
	}
	profileList := common.ProfileList{}
	profileList.Add(profile)
	val, err := proto.Marshal(&profileList)
	if err != nil {
		return err
	}
	err = n.store.Put(key, val)
	if err != nil {
		return err
	}
	return nil
}

// Profile returns the profile for the user from diskDB
func (n *Node) Profile() (*common.Profile, error) {
	// Check if Store is open
	if n.store == nil {
		logger.Error("Failed to Get Profile", ErrProtocolsNotSet)
		return common.NewDefaultProfile(), ErrProtocolsNotSet
	}
	if n.store.Has(PROFILE_KEY) {
		pbuf, err := n.store.Get(PROFILE_KEY)
		if err != nil {
			return common.NewDefaultProfile(), err
		}

		profile := common.Profile{}
		err = proto.Unmarshal(pbuf, &profile)
		if err != nil {
			return nil, err
		}
		return &profile, nil
	}
	return common.NewDefaultProfile(), ErrProfileNotCreated
}

// SetProfile stores the profile for the user in diskDB
func (n *Node) SetProfile(profile *common.Profile) error {
	// Check if Store is open
	if n.store == nil {
		logger.Error("Failed to Set Profile", ErrProtocolsNotSet)
		return ErrProtocolsNotSet
	}

	// Check if profile is nil
	if profile == nil {
		return ErrProfileNotProvided
	}
	pbuf, err := profile.Buffer()
	if err != nil {
		return err
	}
	err = n.store.Put(PROFILE_KEY, pbuf)
	if err != nil {
		return err
	}
	return nil
}
