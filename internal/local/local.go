package local

import (
	"os"

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

func setupKeyHandle() {
	if _, err := os.Stat(keysetFile()); os.IsNotExist(err) {
		// If the keyset file doesn't exist, generate a new key handle
		kh, _ = NewKeyHandle()
	} else {
		// If the keyset file exists, load the key handle from the file
		kh, _ = ReadKeyHandle()
	}
}
