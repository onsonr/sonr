package identity

import (
	"fmt"
	"time"

	"github.com/sonr-io/core/pkg/common"
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
