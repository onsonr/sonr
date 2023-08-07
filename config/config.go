package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// ChainID returns the chain ID from the configuration. (default: sonr-localnet-1)
func ChainID() string {
	return viper.GetString("chain-id")
}

// Environment returns the environment from the configuration. (default: development)
func Environment() string {
	return viper.GetString("environment")
}

// IsProduction returns true if the environment is production
func IsProduction() bool {
	return viper.GetString("environment") == "production"
}

// JWTSigningKey returns the JWT signing key
func JWTSigningKey() []byte {
	return []byte(viper.GetString("launch.highway.signing-key"))
}

// HighwayHostAddress returns the host and port of the Highway API
func HighwayHostAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("launch.highway.api.host"), viper.GetInt("launch.highway.api.port"))
}

// HighwayRequestTimeout returns the timeout for Highway API requests
func HighwayRequestTimeout() time.Duration {
	return time.Duration(viper.GetInt("launch.highway.api.timeout")) * time.Second
}

// IceFireKVHost returns the host and port of the IceFire KV store
func IceFireKVHost() string {
	return fmt.Sprintf("%s:%d", viper.GetString("launch.highway.db.icefirekv.host"), viper.GetInt("launch.highway.db.icefirekv.port"))
}

// NodeAPIHostAddress returns the host and port of the Node API
func NodeAPIHostAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("launch.node.api.host"), viper.GetInt("launch.node.api.port"))
}

// NodeGrpcHostAddress returns the host and port of the Node P2P
func NodeGrpcHostAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("launch.node.grpc.host"), viper.GetInt("launch.node.grpc.port"))
}

// NodeP2PHostAddress returns the host and port of the Node P2P
func NodeP2PHostAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("launch.node.p2p.host"), viper.GetInt("launch.node.p2p.port"))
}

// NodeRPCHostAddress returns the host and port of the Node RPC
func NodeRPCHostAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("launch.node.rpc.host"), viper.GetInt("launch.node.rpc.port"))
}

// ValidatorAddress returns the validator address from the configuration.
func ValidatorAddress() string {
	return viper.GetString("launch.address")
}
