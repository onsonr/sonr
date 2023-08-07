package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

func init() {
	viper.SetEnvPrefix("sonr")
	viper.SetConfigName("sonr")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("/.sonr")
	viper.AddConfigPath("$HOME/.sonr")
	viper.AutomaticEnv()
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

	// HighwayHostPort returns the host and port of the Highway API
	HighwayHostPort = fmt.Sprintf("%s:%d", c.Launch.Highway.API.Host, c.Launch.Highway.API.Port)

	// HighwayRequestTimeout returns the timeout for Highway API requests
	HighwayRequestTimeout = time.Duration(c.Launch.Highway.API.Timeout) * time.Second

	// IceFireHost returns the host and port of the IceFire KV store
	IceFireHost = fmt.Sprintf("%s:%d", c.Launch.Highway.DB.IcefireKV.Host, c.Launch.Highway.DB.IcefireKV.Port)

	// JWTSigningKey returns the JWT signing key
	JWTSigningKey = []byte(c.Launch.Highway.SigningKey)

	// NodeAPIHost returns the host and port of the Node API
	NodeAPIHost = fmt.Sprintf("%s:%d", c.Launch.Node.API.Host, c.Launch.Node.API.Port)

	// NodeGrpcHost returns the host and port of the Node gRPC
	NodeGrpcHost = fmt.Sprintf("%s:%d", c.Launch.Node.GRPC.Host, c.Launch.Node.GRPC.Port)

	// NodeP2PHost returns the host and port of the Node P2P
	NodeP2PHost = fmt.Sprintf("%s:%d", c.Launch.Node.P2P.Host, c.Launch.Node.P2P.Port)

	// NodeRPCHost returns the host and port of the Node RPC
	NodeRPCHost = fmt.Sprintf("%s:%d", c.Launch.Node.RPC.Host, c.Launch.Node.RPC.Port)

	// Environment returns the environment from the environment variable SONR_ENVIRONMENT. (default: development)
	Environment = c.Environment

	// IsProduction returns true if the environment is production
	IsProduction = c.Environment == "production"

	// ValidatorAddress returns the validator address from the environment variable SONR_VALIDATOR_ADDRESS.
	ValidatorAddress = c.Launch.Address
)
