package mpc

import (
	"os"

	"github.com/sonr-hq/sonr/pkg/common"
	// "github.com/sonr-hq/sonr/pkg/node"
)

// SaveToPath saves the wallet to the given path.
func SaveToPath(w common.WalletShare, path string) error {
	bz, err := w.Marshal()
	if err != nil {
		return err
	}
	err = os.WriteFile(path, bz, 0644)
	if err != nil {
		return err
	}
	return nil
}

// LoadFromPath loads a wallet from the given path.
func LoadFromPath(path string) (common.WalletShare, error) {
	w := EmptyWalletShare()
	bz, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = w.Unmarshal(bz)
	if err != nil {
		return nil, err
	}
	return w, nil
}

