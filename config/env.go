package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// setupDefaults sets the default values for the configuration parameters.
func setupDefaults() {
	viper.SetDefault("launch.config", "")
	viper.SetDefault("launch.chain-id", "sonr-localnet-1")
	viper.SetDefault("launch.environment", "development")
	viper.SetDefault("launch.moniker", "alice")
	viper.SetDefault("launch.val_address", "")

	viper.SetDefault("highway.jwt.key", "@re33lyb@dsecret")
	viper.SetDefault("highway.api.host", "localhost")
	viper.SetDefault("highway.api.timeout", 15)

	viper.SetDefault("highway.icefirekv.host", "localhost")
	viper.SetDefault("highway.icefirekv.port", 6001)

	viper.SetDefault("highway.icefiresql.host", "localhost")
	viper.SetDefault("highway.icefiresql.port", 23306)

	viper.SetDefault("node.api.host", "localhost")
	viper.SetDefault("node.api.port", 1317)
	viper.SetDefault("node.grpc.host", "localhost")
	viper.SetDefault("node.grpc.port", 9090)
	viper.SetDefault("node.p2p.host", "localhost")
	viper.SetDefault("node.p2p.port", 26657)
	viper.SetDefault("node.p2p.host", "validator")
	viper.SetDefault("node.p2p.port", 26656)
}

// ChainID returns the chain ID from the configuration. (default: sonr-localnet-1)
func ChainID() string {
	return viper.GetString("launch.chain-id")
}

// Environment returns the environment from the configuration. (default: development)
func Environment() string {
	return viper.GetString("launch.environment")
}

// Moniker returns the moniker from the configuration. (default: alice)
func Moniker() string {
	return viper.GetString("launch.moniker")
}

// IsProduction returns true if the environment is production
func IsProduction() bool {
	return viper.GetString("launch.environment") == "production"
}

// JWTSigningKey returns the JWT signing key
func JWTSigningKey() []byte {
	return []byte(viper.GetString("highway.jwt.key"))
}

// HighwayHostAddress returns the host and port of the Highway API
func HighwayHostAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("highway.api.host"), viper.GetInt("highway.api.port"))
}

// HighwayRequestTimeout returns the timeout for Highway API requests
func HighwayRequestTimeout() time.Duration {
	return time.Duration(viper.GetInt("highway.api.timeout")) * time.Second
}

// IceFireKVHost returns the host and port of the IceFire KV store
func IceFireKVHost() string {
	return fmt.Sprintf("%s:%d", viper.GetString("highway.icefirekv.host"), viper.GetInt("highway.icefirekv.port"))
}

// IceFireSQLHost returns the host and port of the IceFire KV store
func IceFireSQLHost() string {
	return fmt.Sprintf("%s:%d", viper.GetString("highway.icefiresql.host"), viper.GetInt("highway.icefiresql.port"))
}

// NodeAPIHostAddress returns the host and port of the Node API
func NodeAPIHostAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("node.api.host"), viper.GetInt("node.api.port"))
}

// NodeGrpcHostAddress returns the host and port of the Node P2P
func NodeGrpcHostAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("node.grpc.host"), viper.GetInt("node.grpc.port"))
}

// NodeP2PHostAddress returns the host and port of the Node P2P
func NodeP2PHostAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("node.p2p.host"), viper.GetInt("node.p2p.port"))
}

// NodeRPCHostAddress returns the host and port of the Node RPC
func NodeRPCHostAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("node.rpc.host"), viper.GetInt("node.rpc.port"))
}

// ValidatorAddress returns the validator address from the configuration.
func ValidatorAddress() string {
	return viper.GetString("launch.val_address")
}
