package common

import (
	fmt "fmt"
	"strings"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

// `SNRPubKey` is a `PubKey` that has a `DID` and a `Multibase`
// @property {string} DID - The DID of the SNR
// @property {string} Multibase - The multibase encoding of the DID.
type SNRPubKey interface {
	cryptotypes.PubKey

	Bech32(pfix string) (string, error)
	DID(opts ...DIDOption) string
	Multibase() string
	Raw() []byte
}

type DIDConfig struct {
	// Method is the DID method name
	Method string `json:"method"`

	// Network is the DID network name
	Network string `json:"network"`

	// Identifier is the DID identifier
	Identifier string `json:"identifier"`

	// Path is the DID path
	Path string `json:"path"`

	// Fragment is the DID fragment
	Fragment string `json:"fragment"`
}

// DefaultDidUriConfig returns a new DID URI config with default values
func DefaultDidUriConfig() *DIDConfig {
	return &DIDConfig{
		Method:     "sonr",
		Network:    "devnet",
		Identifier: "",
		Path:       "",
		Fragment:   "",
	}
}

// Apply applies the options and returns a constructed valid DID URI
func (c *DIDConfig) Apply(opts ...DIDOption) string {
	for _, opt := range opts {
		opt(c)
	}
	if c.Path != "" {
		c.Path = fmt.Sprintf("/%s", strings.TrimLeft(c.Path, "/"))
	}
	if c.Fragment != "" {
		c.Fragment = fmt.Sprintf("#%s", strings.TrimLeft(c.Fragment, "#"))
	}
	return fmt.Sprintf("did:%s:%s%s%s", c.Method, c.Identifier, c.Path, c.Fragment)
}

// DIDOption is a function that configures a DID URI config
type DIDOption func(*DIDConfig)

// WithMethod sets the DID method name
func WithMethod(method string) DIDOption {
	return func(c *DIDConfig) {
		c.Method = method
	}
}

// WithNetwork sets the DID network name
func WithNetwork(network string) DIDOption {
	return func(c *DIDConfig) {
		c.Network = network
	}
}

// WithIdentifier sets the DID identifier
func WithIdentifier(identifier string) DIDOption {
	return func(c *DIDConfig) {
		c.Identifier = identifier
	}
}

// WithPath sets the DID path
func WithPath(path string) DIDOption {
	return func(c *DIDConfig) {
		c.Path = path
	}
}

// WithFragment sets the DID fragment
func WithFragment(fragment string) DIDOption {
	return func(c *DIDConfig) {
		c.Fragment = fragment
	}
}
