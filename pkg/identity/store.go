package identity

import (
	"fmt"
	"time"

	"github.com/sonr-io/core/pkg/common"
	"google.golang.org/protobuf/proto"
)

// RecentsHistory is a list of recent Peers.
type RecentsHistory map[string]*common.ProfileList

// Bucket Constant Keys
var (
	PROFILE_KEY = []byte("profile")
)

func recentsKey() []byte {
	// Create Key from Time in RFC3339 format
	t := time.Now().Round(time.Hour)
	keyStr := t.Format(time.RFC3339)
	key := []byte(keyStr)
	return []byte(fmt.Sprintf("%s_recents", key))
}

// addToProfileList adds a Peer to the ProfileList.
func (p *IdentityProtocol) addToProfileList(profile *common.Profile) (*common.ProfileList, error) {
	// Get profile list buffer
	oldBuf, err := p.store.Get(recentsKey())
	if err != nil {
		logger.Errorf("%s - Failed to Get old Recents from store")
		return nil, err
	}

	// Unmarshal profile list
	profileList := &common.ProfileList{}
	err = proto.Unmarshal(oldBuf, profileList)
	if err != nil {
		logger.Errorf("%s - Failed to Unmarshal ProfileList")
		return nil, err
	}

	// Add profile to list
	profileList.Add(profile)
	return profileList, nil
}
