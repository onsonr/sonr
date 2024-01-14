package keychain

import (
	"os"

	"github.com/asynkron/protoactor-go/actor"

	modulev1 "github.com/sonrhq/sonr/api/identity/module/v1"
	"github.com/sonrhq/sonr/x/identity/internal/shares"
	"github.com/sonrhq/sonr/x/identity/internal/vfs"
)

// Keychain is a local temp file system which spawns shares as proto actors
type Keychain struct {
	RootDir      string
	Wallets      *vfs.Wallets
	Credentials  *vfs.Credentials
	PrivSharePID *actor.PID
	PubSharePID  *actor.PID
}

// New takes request context and root directory and returns a new Keychain
// 1. It requires an initial credential id to be passed as a value within the accumulator object

func New() (*Keychain, error) {
	rootDir, err := os.MkdirTemp("", "sonr-keychain")
	if err != nil {
		return nil, err
	}
	privId, pubId, err := shares.Generate(rootDir, modulev1.CoinType_COIN_TYPE_SONR)
	if err != nil {
		return nil, err
	}
	kc := &Keychain{
		RootDir:      rootDir,
		PrivSharePID: privId,
		PubSharePID:  pubId,
	}
	return kc, nil
}

// Burn removes the root directory of the keychain
func (kc *Keychain) Burn() error {
	return os.RemoveAll(kc.RootDir)
}
