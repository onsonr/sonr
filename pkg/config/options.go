package config

import (
	"fmt"
	"time"

	dscl "github.com/libp2p/go-libp2p-core/discovery"
)

// Option configures your client.
type Option func(*Config)

func WithCosmosAccountSettings(accountName string, addrPrefix string, homeDir string, keyringServiceName string, keyringBackend string) Option {
	return func(c *Config) {
		c.CosmosAccountName = accountName
		c.CosmosAddressPrefix = addrPrefix
		c.CosmosHomePath = homeDir
		c.CosmosKeyringServiceName = keyringServiceName
		c.CosmosKeyringBackend = keyringBackend
	}
}

// WithCosmosNodeAddress sets the node address of your chain. When this option is not provided
// `http://localhost:26657` is used as default.
func WithCosmosNodeAddress(addr string) Option {
	return func(c *Config) {
		c.CosmosNodeAddress = addr
	}
}

func WithUseFaucet(faucetAddress, denom string, minAmount uint64) Option {
	return func(c *Config) {
		c.CosmosUseFaucet = true
		c.CosmosFaucetAddress = faucetAddress
		if denom != "" {
			c.CosmosFaucetDenom = denom
		}
		if minAmount != 0 {
			c.CosmosFaucetMinAmount = minAmount
		}
	}
}

// WithHighwayAPISettings sets the host address for the Node Stub Client Host
func WithHighwayAPISettings(network string, grpcHost string, grpcPort int, httpPort int) Option {
	return func(o *Config) {
		o.HighwayGRPCNetwork = network
		o.HighwayGRPCEndpoint = fmt.Sprintf("%s:%d", grpcHost, grpcPort)
		o.HighwayHTTPEndpoint = fmt.Sprintf(":%d", httpPort)
	}
}

type MotorOption func(o *Config)

// WithLibp2pConnOptions sets the connection manager options. Defaults are (lowWater: 15, highWater: 40, gracePeriod: 5m)
func WithLibp2pConnOptions(low int, hi int, grace time.Duration) MotorOption {
	return func(o *Config) {
		o.Libp2pLowWater = low
		o.Libp2pHighWater = hi
		o.Libp2pGracePeriod = grace
	}
}

// WithLibp2pRendevouz sets the interval and timeout for the DHT rendezvous strategy
func WithLibp2pRendevouz(point string, ttl time.Duration, interval time.Duration) MotorOption {
	return func(o *Config) {
		o.Libp2pInterval = interval
		o.Libp2pTTL = dscl.TTL(ttl)
		o.Libp2pRendezvous = point
	}
}

// WithLibp2pMDNS sets the non-priority of MDNS Discovery
func WithLibp2pMDNS(isActive bool) MotorOption {
	return func(o *Config) {
		o.Libp2pMdnsDisabled = isActive
	}
}

// WithDeviceID sets the device ID
func WithDeviceID(id string) MotorOption {
	return func(o *Config) {
		// Set Home Directory
		if id != "" {
			o.DeviceID = id
		}
	}
}

// WithDeviceID sets the device ID
func WithHostName(n string) MotorOption {
	return func(o *Config) {
		// Set Home Directory
		if n != "" {
			o.HostName = n
		}
	}
}

// WithHomePath sets the Home Directory
func WithHomePath(p string) MotorOption {
	return func(o *Config) {
		// Set Home Directory
		if p != "" {
			o.HomeDirPath = p
		}
	}
}

// WithTempPath sets the Temporary Directory
func WithTempPath(p string) MotorOption {
	return func(o *Config) {
		// Set Home Directory
		if p != "" {
			o.TempDirPath = p
		}
	}
}

// WithSupportPath sets the Support Directory
func WithSupportPath(p string) MotorOption {
	return func(o *Config) {
		// Set Home Directory
		if p != "" {
			o.SupportDirPath = p
		}
	}
}
