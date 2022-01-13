package did

import (
	"strings"
)

type Network int

const (
	Method = "did:sonr"
)

type Option func(*Did)

// WithFragment adds a fragment to a DID
func WithFragment(fragment string) Option {
	return func(d *Did) {
		fragment := strings.SplitAfter(fragment, "#")
		d.Fragment = fragment[1]
	}
}

// WithNetwork adds a network to a DID
func WithNetwork(network string) Option {
	return func(d *Did) {
		if ok := IsValidNetworkPrefix(network); ok {
			if network == "mainnet" {
				network = ""
			}
			d.Network = network
		} else {
			d.Network = "testnet"
		}
	}
}

// WithPath adds a path to a DID
func WithPath(p string) Option {
	return func(d *Did) {
		rawpaths := strings.SplitAfter(p, "/")
		paths := strings.Split(strings.Join(rawpaths[1:], ""), "/")
		d.Paths = paths
	}
}

// WithQuery adds a query to a DID
func WithQuery(query string) Option {
	return func(d *Did) {
		query := strings.SplitAfter(query, "?")
		d.Query = query[1]
	}
}

// AddPath adds a path to the DID struct and returns the new DID
func (d *Did) AddPath(path string) *Did {
	// Check if the path contains a leading slash
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}

	// Split and join the path
	paths := strings.Split(path, "/")
	d.Paths = append(d.Paths, paths...)
	return d
}

// AddFragment adds a fragment to the DID struct and returns the new DID
func (d *Did) AddFragment(fragment string) *Did {
	// Check if the fragment contains a leading hash
	if !strings.HasPrefix(fragment, "#") {
		fragment = "#" + fragment
	}

	d.Fragment = fragment
	return d
}

// AddQuery adds a query to the DID struct and returns the new DID
func (d *Did) AddQuery(query string) *Did {
	// Check if the query contains a leading question mark
	if !strings.HasPrefix(query, "?") {
		query = "?" + query
	}

	d.Query = query
	return d
}
