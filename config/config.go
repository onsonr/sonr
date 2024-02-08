package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// Config Structure
// {
// 	"sonrd": {
// 		"chain_id": "sonr-testnet-1"
// 	},
// 	"highway": {
// 		"node_grpc_address": "localhost:9090",
// 		"chain_id": "sonr-testnet-1",
// 		"validator_address": "cosmosvaloper1xv3j7yylfwz4e9v3u8h60x5shy7y58m3y2q5qk"
// 	},
// 	"nitro": {
// 		"matrix_host": "http://localhost:3000",
// 		"shared_secret": "secret"
// 	}
// }

// SonrdConfig represents the sonrd configuration
type SonrdConfig struct {
	ChainID        string `mapstructure:"chain_id" json:"chain_id"`
	HighwayEnabled bool   `mapstructure:"highway_enabled" json:"highway_enabled"`
}

// HighwayConfig represents the highway configuration
type HighwayConfig struct {
	NodeGRPCAddress  string `mapstructure:"node_grpc_address" json:"node_grpc_address"`
	ChainID          string `mapstructure:"chain_id" json:"chain_id"`
	ValidatorAddress string `mapstructure:"validator_address" json:"validator_address"`
}

// NitroConfig represents the nitro configuration
type NitroConfig struct {
	MatrixHost   string `mapstructure:"matrix_host" json:"matrix_host"`
	SharedSecret string `mapstructure:"shared_secret" json:"shared_secret"`
}

// Config is the configuration for the application
type Config struct {
	Name            string        `mapstructure:"name" json:"name"`
	Version         string        `mapstructure:"version" json:"version"`
	Description     string        `mapstructure:"description" json:"description"`
	DefaultNodeHome string        `mapstructure:"default_node_home" json:"default_node_home"`
	Sonrd           SonrdConfig   `mapstructure:"sonrd" json:"sonrd"`
	Highway         HighwayConfig `mapstructure:"highway" json:"highway"`
	Nitro           NitroConfig   `mapstructure:"nitro" json:"nitro"`
}

// LoadConfig loads the configuration from the file
func LoadConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := v.Unmarshal(&config, func(dc *mapstructure.DecoderConfig) {
		dc.TagName = "mapstructure"
	}); err != nil {
		return nil, err
	}

	return &config, nil
}
