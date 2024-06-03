package local

import (
	"os"
	"time"

	"github.com/bool64/cache"
	"github.com/tink-crypto/tink-go/v2/keyset"
)

var (
	chainID = "testnet"
	valAddr = "val1"
	nodeDir = ".sonr"

	defaultNodeHome = os.ExpandEnv("$HOME/") + nodeDir

	kh *keyset.Handle
)

// Initialize initializes the local configuration values
func Initialize() {
	setupCache()
	setupKeyHandle()
}

// SetLocalContextSessionID sets the session ID for the local context
func SetLocalValidatorAddress(address string) {
	valAddr = address
}

// SetLocalContextChainID sets the chain ID for the local
func SetLocalChainID(id string) {
	chainID = id
}

// setupCache configures cache and inital settings for proxy.
func setupCache() {
	// Setup cache for session.
	baseSessionCache = cache.NewFailoverOf(func(cfg *cache.FailoverConfigOf[session]) {
		// Using last 30 seconds of 5m TTL for background update.
		cfg.MaxStaleness = 1 * time.Hour
		cfg.BackendConfig.TimeToLive = 2*time.Hour - cfg.MaxStaleness
	})
	authorizedSessionCache = cache.NewFailoverOf(func(cfg *cache.FailoverConfigOf[authorizedSession]) {
		// Using last 30 seconds of 5m TTL for background update.
		cfg.MaxStaleness = 30 * time.Minute
		cfg.BackendConfig.TimeToLive = 1*time.Hour - cfg.MaxStaleness
	})
}

func setupKeyHandle() {
	if _, err := os.Stat(keysetFile()); os.IsNotExist(err) {
		// If the keyset file doesn't exist, generate a new key handle
		kh, _ = NewKeyHandle()
	} else {
		// If the keyset file exists, load the key handle from the file
		kh, _ = ReadKeyHandle()
	}
}
