package store

import (
	"errors"

	"github.com/sonr-io/core/internal/common"
)

// RecentsHistory is a list of recent Peers.
type RecentsHistory map[string]*common.ProfileList

// Error Definitions
var (
	ErrRecentsNotCreated  = errors.New("Recents has not been created yet.")
	ErrProfileNotCreated  = errors.New("Profile has not been created yet.")
	ErrProfileNotProvided = errors.New("Profile has not been provided to Store.")
	ErrProfileIsOlder     = errors.New("Profile is older than the oldest one on disk.")
	ErrProfileNoTimestamp = errors.New("Profile has no timestamp.")
)

// Buckets in Database
var (
	RECENTS_BUCKET = []byte("recents")
	USER_BUCKET    = []byte("user")
)

// Bucket Constant Keys
var (
	PROFILE_KEY = []byte("profile")
)
