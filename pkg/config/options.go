package config

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	dscl "github.com/libp2p/go-libp2p-core/discovery"
)

// Option configures your client.
type Option func(*Config)

// WithHighwayAPISettings sets the host address for the Node Stub Client Host
func WithHighwayAPISettings(network string, grpcHost string, grpcPort int, httpPort int) Option {
	return func(o *Config) {
		o.HighwayGRPCNetwork = network
		o.HighwayGRPCEndpoint = fmt.Sprintf("%s:%d", grpcHost, grpcPort)
		o.HighwayHTTPEndpoint = fmt.Sprintf(":%d", httpPort)
	}
}

func WithJWTTokenOptions(secret string, signingMethod jwt.SigningMethod, exp int64) Option {
	return func(o *Config) {
		o.JWTSecret = secret
		o.JWTSigningMethod = signingMethod
		o.JWTExpiration = exp
	}
}

// WithLibp2pConnOptions sets the connection manager options. Defaults are (lowWater: 15, highWater: 40, gracePeriod: 5m)
func WithLibp2pConnOptions(low int, hi int, grace time.Duration) Option {
	return func(o *Config) {
		o.Libp2pLowWater = low
		o.Libp2pHighWater = hi
		o.Libp2pGracePeriod = grace
	}
}

// WithLibp2pRendevouz sets the interval and timeout for the DHT rendezvous strategy
func WithLibp2pRendevouz(point string, ttl time.Duration, interval time.Duration) Option {
	return func(o *Config) {
		o.Libp2pInterval = interval
		o.Libp2pTTL = dscl.TTL(ttl)
		o.Libp2pRendezvous = point
	}
}

// WithLibp2pMDNS sets the non-priority of MDNS Discovery
func WithLibp2pMDNS(isActive bool) Option {
	return func(o *Config) {
		o.Libp2pMdnsDisabled = isActive
	}
}

// WithAccountAddress sets the account address
func WithAccountAddress(addr string) Option {
	return func(o *Config) {
		o.AccountAddress = addr
	}
}

// WithDeviceID sets the device ID
func WithDeviceID(id string) Option {
	return func(o *Config) {
		// Set Home Directory
		if id != "" {
			o.DeviceID = id
		}
	}
}

// WithHomePath sets the Home Directory
func WithHomePath(p string) Option {
	return func(o *Config) {
		// Set Home Directory
		if p != "" {
			o.HomeDirPath = p
		}
	}
}

// WithTempPath sets the Temporary Directory
func WithTempPath(p string) Option {
	return func(o *Config) {
		// Set Home Directory
		if p != "" {
			o.TempDirPath = p
		}
	}
}

// WithSupportPath sets the Support Directory
func WithSupportPath(p string) Option {
	return func(o *Config) {
		// Set Home Directory
		if p != "" {
			o.SupportDirPath = p
		}
	}
}
