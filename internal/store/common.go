package store

import "github.com/sonr-io/core/internal/common"

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
