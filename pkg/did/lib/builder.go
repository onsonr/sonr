package did

import (
	"strings"
)

// SONR_PREFIX is the prefix for the SONR DID
const SONR_PREFIX = "did:sonr:"

// DidUrl is a formatted and validated DID URL string.
type DidUrl string

// String returns the string representation of a DidUrl
func (d DidUrl) String() string {
	return string(d)
}

// config is the did string configuration
type config struct {
	fragment   string
	network    string
	paths      []string
	query      string
	identifier string
}

// defaultConfig returns the default configuration
func defaultConfig(identifier string) *config {
	return &config{
		fragment:   "",
		network:    "",
		paths:      []string{},
		query:      "",
		identifier: identifier,
	}
}

// Option is a function that can be used to modify the DidUrl.
type Option func(*config)

// Build creates a new DidUrl from the given options and returns it.
func Build(identifier string, opts ...Option) (DidUrl, error) {
	// Config options
	d := defaultConfig(identifier)
	for _, opt := range opts {
		opt(d)
	}

	// Create base DID
	didStr := SONR_PREFIX + d.network + d.identifier
	for _, v := range d.paths {
		didStr += "/" + v
	}

	// Add query
	if d.query != "" {
		didStr += d.query
	}

	// Add fragment
	if d.fragment != "" {
		didStr += d.fragment
	}

	// Check if the DID is valid
	if !IsValidDid(didStr) {
		return "", ErrParseInvalid
	}

	// Return the DID
	return DidUrl(didStr), nil
}

// WithFragment adds a fragment to a DID
func WithFragment(fragment string) Option {
	return func(d *config) {
		fragment := strings.SplitAfter(fragment, "#")
		d.fragment = ToFragment(fragment[1])
	}
}

// WithNetwork adds a network to a DID
func WithNetwork(network string) Option {
	return func(d *config) {
		// Check if the network is valid
		if ok := IsFragment(network); ok {
			// Check if the network is mainnet
			if network == "mainnet:" {
				network = ":"
			}

			// Check if the network has a trailing colon
			if ContainsString(network, ":") {
				d.network = network
			} else {
				d.network = network + ":"
			}
		} else {
			d.network = "testnet:"
		}
	}
}

// WithPathSegments adds a paths to a DID
func WithPathSegments(p ...string) Option {
	return func(d *config) {
		d.paths = p
	}
}

// WithQuery adds a query to a DID
func WithQuery(query string) Option {
	return func(d *config) {
		query := strings.SplitAfter(query, "?")
		d.query = ToQuery(query[1])
	}
}
