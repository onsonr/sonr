package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

func init() {
	viper.SetEnvPrefix("sonr")
	viper.AutomaticEnv()
	viper.SetConfigName("sonr")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("/.sonr")
	viper.AddConfigPath("$HOME/.sonr")
	viper.AddConfigPath("/etc/sonr")
	viper.AddConfigPath("$HOME/etc/sonr")
	viper.AddConfigPath("/etc")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("No config file found - using environment variables only")
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		fmt.Println("No config file found - using environment variables only")
	}
}

// Get returns the configuration values of the application.
func Get() Config {
	return c
}

// ChainID returns the chain ID from the environment variable SONR_CHAIN_ID. (default: sonr-localnet-1)
func ChainID() string {
	return viper.GetString("SONR_CHAIN_ID")
}
// Environment returns the environment from the environment variable SONR_ENVIRONMENT. (default: development)
func Environment() string {
	return viper.GetString("SONR_ENVIRONMENT")
}

// IsProduction returns true if the environment is production
func IsProduction() bool {
	return viper.GetString("SONR_ENVIRONMENT") == "production"
}

// JWTSigningKey returns the JWT signing key
func JWTSigningKey() []byte {
	return []byte(viper.GetString("SONR_LAUNCH_HIGHWAY_SIGNING_KEY"))
}

// HighwayHostAddress returns the host and port of the Highway API
func HighwayHostAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("SONR_LAUNCH_HIGHWAY_API_HOST"), viper.GetInt("SONR_LAUNCH_HIGHWAY_API_PORT"))
}

// HighwayRequestTimeout returns the timeout for Highway API requests
func HighwayRequestTimeout() time.Duration {
	return time.Duration(viper.GetInt("SONR_LAUNCH_HIGHWAY_API_TIMEOUT")) * time.Second
}

// IceFireKVHost returns the host and port of the IceFire KV store
func IceFireKVHost() string {
	return fmt.Sprintf("%s:%d", viper.GetString("SONR_LAUNCH_HIGHWAY_DB_ICEFIREKV_HOST"), viper.GetInt("SONR_LAUNCH_HIGHWAY_DB_ICEFIREKV_PORT"))
}

// NodeAPIHostAddress returns the host and port of the Node API
func NodeAPIHostAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("SONR_LAUNCH_NODE_API_HOST"), viper.GetInt("SONR_LAUNCH_NODE_API_PORT"))
}
// NodeGrpcHostAddress returns the host and port of the Node P2P
func NodeGrpcHostAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("SONR_LAUNCH_NODE_GRPC_HOST"), viper.GetInt("SONR_LAUNCH_NODE_GRPC_PORT"))
}

// NodeP2PHostAddress returns the host and port of the Node P2P
func NodeP2PHostAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("SONR_LAUNCH_NODE_P2P_HOST"), viper.GetInt("SONR_LAUNCH_NODE_P2P_PORT"))
}

// NodeRPCHostAddress returns the host and port of the Node RPC
func NodeRPCHostAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("SONR_LAUNCH_NODE_RPC_HOST"), viper.GetInt("SONR_LAUNCH_NODE_RPC_PORT"))
}

// ValidatorAddress returns the validator address from the environment variable SONR_VALIDATOR_ADDRESS.
func ValidatorAddress() string {
	return viper.GetString("SONR_VALIDATOR_ADDRESS")
}
