package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// ChainID returns the chain ID from the configuration. (default: sonr-localnet-1)
func ChainID() string {
	viper.SetDefault("launch.chain-id", "sonr-localnet-1")
	return viper.GetString("launch.chain-id")
}

// Environment returns the environment from the configuration. (default: development)
func Environment() string {
	viper.SetDefault("launch.environment", "development")
	return viper.GetString("launch.environment")
}

// Moniker returns the moniker from the configuration. (default: alice)
func Moniker() string {
	viper.SetDefault("launch.moniker", "alice")
	return viper.GetString("launch.moniker")
}

// IsProduction returns true if the environment is production
func IsProduction() bool {
	return viper.GetString("launch.environment") == "production"
}

// JWTSigningKey returns the JWT signing key
func JWTSigningKey() []byte {
	viper.SetDefault("highway.jwt.key", "@re33lyb@dsecret")
	return []byte(viper.GetString("highway.jwt.key"))
}

// HighwayHostAddress returns the host and port of the Highway API
func HighwayHostAddress() string {
	viper.SetDefault("highway.api.host", "0.0.0.0")
	return fmt.Sprintf("%s:%d", viper.GetString("highway.api.host"), viper.GetInt("highway.api.port"))
}

// HighwayRequestTimeout returns the timeout for Highway API requests
func HighwayRequestTimeout() time.Duration {
	viper.SetDefault("highway.api.timeout", 15)
	return time.Duration(viper.GetInt("highway.api.timeout")) * time.Second
}

// IceFireKVHost returns the host and port of the IceFire KV store
func IceFireKVHost() string {
	viper.SetDefault("highway.icefirekv.host", "localhost")
	return fmt.Sprintf("%s:%d", viper.GetString("highway.icefirekv.host"), viper.GetInt("highway.icefirekv.port"))
}

// IceFireSQLHost returns the host and port of the IceFire KV store
func IceFireSQLHost() string {
	viper.SetDefault("highway.icefiresql.host", "localhost")
	viper.SetDefault("highway.icefiresql.port", 23306)
	return fmt.Sprintf("%s:%d", viper.GetString("highway.icefiresql.host"), viper.GetInt("highway.icefiresql.port"))
}

// NodeAPIHostAddress returns the host and port of the Node API
func NodeAPIHostAddress() string {
	viper.SetDefault("node.api.host", "0.0.0.0")
	viper.SetDefault("node.api.port", 1317)
	return fmt.Sprintf("%s:%d", viper.GetString("node.api.host"), viper.GetInt("node.api.port"))
}

// NodeGrpcHostAddress returns the host and port of the Node P2P
func NodeGrpcHostAddress() string {
	viper.SetDefault("node.grpc.host", "0.0.0.0")
	viper.SetDefault("node.grpc.port", 9090)
	return fmt.Sprintf("%s:%d", viper.GetString("node.grpc.host"), viper.GetInt("node.grpc.port"))
}

// NodeP2PHostAddress returns the host and port of the Node P2P
func NodeP2PHostAddress() string {
	viper.SetDefault("node.p2p.host", "validator")
	viper.SetDefault("node.p2p.port", 26656)
	return fmt.Sprintf("%s:%d", viper.GetString("node.p2p.host"), viper.GetInt("node.p2p.port"))
}

// NodeRPCHostAddress returns the host and port of the Node RPC
func NodeRPCHostAddress() string {
	viper.SetDefault("node.p2p.host", "0.0.0.0")
	viper.SetDefault("node.p2p.port", 26657)
	return fmt.Sprintf("%s:%d", viper.GetString("node.rpc.host"), viper.GetInt("node.rpc.port"))
}

// ValidatorAddress returns the validator address from the configuration.
func ValidatorAddress() string {
	viper.SetDefault("launch.val_address", "")
	return viper.GetString("launch.val_address")
}
