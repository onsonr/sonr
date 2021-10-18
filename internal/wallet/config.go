package wallet

import (
	"github.com/kataras/golog"
)

// Error Definitions
var (
	logger = golog.Child("internal/wallet")
)

// NodeOption is a function that modifies the node options.
type WalletOption func(*walletOptions)

// walletOptions is a collection of options for the node.
type walletOptions struct {
	directory string
}

// defaultNodeOptions returns the default node options.
func defaultNodeOptions() *walletOptions {
	// path, _ := device.NewSupportPath(".wallet", device.CreateDirIfNotExist())
	return &walletOptions{
		// directory: path,
	}
}
