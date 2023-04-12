package config

import (
	"github.com/taurusgroup/multi-party-sig/pkg/party"
)

// StoreType is the type of a store
type StoreType string

const (
	// DB_EVENT_LOG_STORE is a store that stores events
	DB_EVENT_LOG_STORE StoreType = "eventlog"

	// DB_KEY_VALUE_STORE is a store that stores key-value pairs
	DB_KEY_VALUE_STORE StoreType = "keyvalue"

	// DB_DOCUMENT_STORE is a store that stores documents
	DB_DOCUMENT_STORE StoreType = "docstore"
)

// A method of the StoreType type.
func (st StoreType) String() string {
	return string(st)
}

// Config is the configuration for the node
type Config struct {
	// Callback is the callback for the motor
	Callback NodeCallback

	// GroupIDs is the list of peer ids for the node
	GroupIDs []party.ID

	// SelfPartyID is the party id for the node
	SelfPartyID party.ID
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		SelfPartyID: party.ID("current"),
		Callback:    DefaultCallback(),
	}
}

// Apply applies the options to the configuration
func (c *Config) Apply(opts ...Option) error {
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return err
		}
	}
	return nil
}

// Option is a function that configures a Node
type Option func(*Config) error

// WithGroupIds sets the peer ids for the node
func WithGroupIds(partyIds ...party.ID) Option {
	return func(c *Config) error {
		if len(partyIds) > 0 {
			c.GroupIDs = partyIds
		}
		return nil
	}
}

// WithNodeCallback sets the callback for the motor
func WithNodeCallback(callback NodeCallback) Option {
	return func(c *Config) error {
		c.Callback = callback
		return nil
	}
}

// WithPartyId sets the party id for the node. This is to be replaced by the User defined label for the device
func WithPartyId(partyId string) Option {
	return func(c *Config) error {
		c.SelfPartyID = party.ID(partyId)
		return nil
	}
}

type node struct {
	host   PeerNode
	ipfs   IPFSNode
	config *Config
}

func (n *node) Host() PeerNode {
	return n.host
}

func (n *node) IPFS() IPFSNode {
	return n.ipfs
}
