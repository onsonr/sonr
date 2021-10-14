package node

import (
	"context"
	"fmt"
	"time"

	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"

	"git.mills.io/prologic/bitcask"
	"google.golang.org/protobuf/proto"
)

// RecentsHistory is a list of recent Peers.
type RecentsHistory map[string]*common.ProfileList

// Bucket Constant Keys
var (
	PROFILE_KEY = []byte("profile")
)

func historyKey() []byte {
	// Create Key from Time in RFC3339 format
	t := time.Now().Round(time.Hour)
	keyStr := t.Format(time.RFC3339)
	key := []byte(keyStr)
	return []byte(fmt.Sprintf("%s_history", key))
}

func recentsKey() []byte {
	// Create Key from Time in RFC3339 format
	t := time.Now().Round(time.Hour)
	keyStr := t.Format(time.RFC3339)
	key := []byte(keyStr)
	return []byte(fmt.Sprintf("%s_recents", key))
}

// initStore creates a new Store instance for Node
func (n *Node) initStore(ctx context.Context, opts *nodeOptions) error {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := n.openStore()
	if err != nil {
		logger.Error("Failed to Open Bitcask Store", err)
		return err
	}
	defer db.Close()

	// Create Profile Bucket
	err = n.SetProfile(opts.profile)
	if err != nil {
		logger.Error("Failed to Set Profile", err)
		return err
	}
	return nil
}

// openStore opens the store file
func (n *Node) openStore() (*bitcask.Bitcask, error) {
	path, err := device.NewDatabasePath("sonr_bitcask")
	if err != nil {
		logger.Error("Failed to get DB Path", err)
		return nil, err
	}
	return bitcask.Open(path, bitcask.WithAutoRecovery(true))
}

// AddHistory adds payload to the store file history
func (n *Node) AddHistory(payload *common.Payload) error {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := n.openStore()
	if err != nil {
		logger.Error("Failed to Open Bitcask Store", err)
		return err
	}
	defer db.Close()

	// Check if profile is nil
	if payload == nil {
		logger.Error("Failed to Add Payload", ErrMissingParam)
		return ErrMissingParam
	}

	// Put in Bucket
	if db.Has(historyKey()) {
		// Get profile list buffer
		oldBuf, err := db.Get(historyKey())
		if err != nil {
			logger.Error("Failed to Get History from store")
			return err
		}

		// Unmarshal profile list
		payloadList := common.PayloadList{}
		err = proto.Unmarshal(oldBuf, &payloadList)
		if err != nil {
			logger.Error("Failed to Unmarshal PayloadList")
			return err
		}

		// Add profile to list
		payloadList.Add(payload)

		// Marshal profile
		val, err := proto.Marshal(&payloadList)
		if err != nil {
			logger.Error("Failed to Marshal PayloadList")
			return err
		}

		err = db.Put(historyKey(), val)
		if err != nil {
			return err
		}
		return nil
	}

	payloadList := common.PayloadList{
		Key: string(recentsKey()),
	}
	payloadList.Add(payload)
	val, err := proto.Marshal(&payloadList)
	if err != nil {
		return err
	}
	err = db.Put(historyKey(), val)
	if err != nil {
		return err
	}
	return nil
}

// AddRecent stores the profile for recents in desk and returns list of recent profiles
func (n *Node) AddRecent(profile *common.Profile) error {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := n.openStore()
	if err != nil {
		logger.Error("Failed to Open Bitcask Store", err)
		return err
	}
	defer db.Close()

	// Check if profile is nil
	if profile == nil {
		logger.Error("Failed to Add Recent Profile", ErrMissingParam)
		return ErrMissingParam
	}

	// Put in Bucket
	if db.Has(recentsKey()) {
		// Get profile list buffer
		oldBuf, err := db.Get(recentsKey())
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

		err = db.Put(recentsKey(), val)
		if err != nil {
			return err
		}
		return nil
	}
	profileList := common.ProfileList{
		Key: string(recentsKey()),
	}
	profileList.Add(profile)
	val, err := proto.Marshal(&profileList)
	if err != nil {
		return err
	}
	err = db.Put(recentsKey(), val)
	if err != nil {
		return err
	}
	return nil
}

// GetHistory returns the history of profiles
func (n *Node) GetHistory() (*common.PayloadList, error) {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := n.openStore()
	if err != nil {
		logger.Error("Failed to Open Bitcask Store", err)
		return nil, err
	}
	defer db.Close()

	// Check for Key
	if db.Has(historyKey()) {
		rbuf, err := db.Get(historyKey())
		if err != nil {
			logger.Error("Failed to Get History from store")
			return nil, err
		}

		// Unmarshal profile list
		payloadList := common.PayloadList{}
		err = proto.Unmarshal(rbuf, &payloadList)
		if err != nil {
			logger.Error("Failed to Unmarshal PayloadList")
			return nil, err
		}
		return &payloadList, nil
	}
	return &common.PayloadList{}, nil
}

// GetRecents returns the list of recent profiles
func (n *Node) GetRecents() (*common.ProfileList, error) {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := n.openStore()
	if err != nil {
		logger.Error("Failed to Open Bitcask Store", err)
		return nil, err
	}
	defer db.Close()

	// Check for Key
	if db.Has(recentsKey()) {
		rbuf, err := db.Get(recentsKey())
		if err != nil {
			logger.Error("Failed to Get Recents from store")
			return nil, err
		}

		// Unmarshal profile list
		profileList := common.ProfileList{}
		err = proto.Unmarshal(rbuf, &profileList)
		if err != nil {
			logger.Error("Failed to Unmarshal ProfileList")
			return nil, err
		}
		return &profileList, nil
	}
	return &common.ProfileList{}, nil
}

// GetProfile returns the profile for the user from diskDB
func (n *Node) GetProfile() (*common.Profile, error) {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := n.openStore()
	if err != nil {
		logger.Error("Failed to Open Bitcask Store", err)
		return nil, err
	}
	defer db.Close()

	if db.Has(PROFILE_KEY) {
		pbuf, err := db.Get(PROFILE_KEY)
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
	return common.NewDefaultProfile(), nil
}

// SetProfile stores the profile for the user in diskDB
func (n *Node) SetProfile(profile *common.Profile) error {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := n.openStore()
	if err != nil {
		logger.Error("Failed to Open Bitcask Store", err)
		return err
	}
	defer db.Close()

	// Check if profile is nil
	if profile == nil {
		return ErrMissingParam
	}
	pbuf, err := profile.Buffer()
	if err != nil {
		return err
	}
	err = db.Put(PROFILE_KEY, pbuf)
	if err != nil {
		return err
	}
	return nil
}
