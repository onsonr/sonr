package config

import (
	"time"

	dscl "github.com/libp2p/go-libp2p-core/discovery"
)

// WithAddress sets the host address for the Node Stub Client Host
func WithAddress(host string) Option {
	return func(o *Config) {
		o.Libp2pHost = host
	}
}

// WithCosmosAccountName sets the cosmos account name to use. defaults to "alice"
func WithCosmosAccountName(host string) Option {
	return func(o *Config) {
		o.CosmosAccountName = host
	}
}

// WithConnOptions sets the connection manager options. Defaults are (lowWater: 15, highWater: 40, gracePeriod: 5m)
func WithConnOptions(low int, hi int, grace time.Duration) Option {
	return func(o *Config) {
		o.Libp2pLowWater = low
		o.Libp2pHighWater = hi
		o.Libp2pGracePeriod = grace
	}
}

// WithInterval sets the interval for the host. Default is 5 seconds.
func WithInterval(interval time.Duration) Option {
	return func(o *Config) {
		o.Libp2pInterval = interval
	}
}

// WithTTL sets the ttl for the host. Default is 2 minutes.
func WithTTL(ttl time.Duration) Option {
	return func(o *Config) {
		o.Libp2pTTL = dscl.TTL(ttl)
	}
}

// WithPort sets the port for the Node Stub Client
func WithPort(port int) Option {
	return func(o *Config) {
		o.Libp2pPort = port
	}
}

// WithWebAuthn sets the webauthn server Properties
func WithWebAuthn(displayName string, rpId string, rpOrigin string, isDebug bool) Option {
	return func(o *Config) {
		o.WebAuthNRPDisplayName = displayName
		o.WebAuthNRPID = rpId
		o.WebAuthNRPOrigin = rpOrigin
		o.WebAuthNDebug = isDebug
	}
}

// DisableMDNS sets the non-priority of MDNS Discovery
func DisableMDNS() Option {
	return func(o *Config) {
		o.Libp2pMdnsDisabled = true
	}
}
