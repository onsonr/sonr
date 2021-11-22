package wallet

import (
	"errors"

	"github.com/kataras/golog"
)

// Error Definitions
var (
	logger             = golog.Default.Child("core/wallet")
	ErrInvalidKeyType  = errors.New("Invalid KeyPair Type provided")
	ErrKeychainUnready = errors.New("Keychain has not been loaded")
	ErrNoPrivateKey    = errors.New("No private key in KeyPair")
	ErrNoPublicKey     = errors.New("No public key in KeyPair")
)

// Option is a function that modifies the wallet options.
type Option func(*options)

// WithPassphrase sets the passphrase of the wallet.
func WithPassphrase(passphrase string) Option {
	return func(o *options) {
		o.passphrase = passphrase
	}
}

// WithSName sets the name of the wallet.
func WithSName(sname string) Option {
	return func(o *options) {
		o.sname = sname
	}
}

// options is a collection of options for the wallet.
type options struct {
	passphrase string
	sname      string
}

// defaultOptions returns the default wallet options.
func defaultOptions() *options {
	return &options{
		passphrase: "wagmi",
		sname:      "test",
	}
}
