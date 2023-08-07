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

var (
	// ChainID returns the chain ID from the environment variable SONR_CHAIN_ID. (default: sonr-localnet-1)
	ChainID = c.ChainID

	// HighwayHostAddress returns the host and port of the Highway API
	HighwayHostAddress = fmt.Sprintf("%s:%d", c.Launch.Highway.API.Host, c.Launch.Highway.API.Port)

	// HighwayRequestTimeout returns the timeout for Highway API requests
	HighwayRequestTimeout = time.Duration(c.Launch.Highway.API.Timeout) * time.Second

	// IceFireHost returns the host and port of the IceFire KV store
	IceFireHost = fmt.Sprintf("%s:%d", c.Launch.Highway.DB.IcefireKV.Host, c.Launch.Highway.DB.IcefireKV.Port)

	// JWTSigningKey returns the JWT signing key
	JWTSigningKey = []byte(c.Launch.Highway.JWTSigningKey)

	// NodeAPIHostAddress returns the host and port of the Node API
	NodeAPIHostAddress = fmt.Sprintf("%s:%d", c.Launch.Node.API.Host, c.Launch.Node.API.Port)

	// NodeGrpcHostAddress returns the host and port of the Node gRPC
	NodeGrpcHostAddress = fmt.Sprintf("%s:%d", c.Launch.Node.GRPC.Host, c.Launch.Node.GRPC.Port)

	// NodeP2PHostAddress returns the host and port of the Node P2P
	NodeP2PHostAddress = fmt.Sprintf("%s:%d", c.Launch.Node.P2P.Host, c.Launch.Node.P2P.Port)

	// NodeRPCHostAddress returns the host and port of the Node RPC
	NodeRPCHostAddress = fmt.Sprintf("%s:%d", c.Launch.Node.RPC.Host, c.Launch.Node.RPC.Port)

	// Environment returns the environment from the environment variable SONR_ENVIRONMENT. (default: development)
	Environment = c.Environment

	// IsProduction returns true if the environment is production
	IsProduction = c.Environment == "production"

	// ValidatorAddress returns the validator address from the environment variable SONR_VALIDATOR_ADDRESS.
	ValidatorAddress = c.Launch.Address
)
