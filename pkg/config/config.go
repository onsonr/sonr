package config

import (
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
	HighwayEnabled bool   `mapstructure:"highway_enabled" json:"highway_enabled"`
	IPFSConnection string `mapstructure:"ipfs_connection" json:"ipfs_connection"`
	IPFSGateway    string `mapstructure:"ipfs_gateway" json:"ipfs_gateway"`
}

// HighwayConfig represents the highway configuration
type HighwayConfig struct {
	IPFSGateway      string `mapstructure:"ipfs_gateway" json:"ipfs_gateway"`
	MatrixConnection string `mapstructure:"matrix_connection" json:"matrix_connection"`
	NodeGRPCAddress  string `mapstructure:"node_grpc_address" json:"node_grpc_address"`
	ValidatorAddress string `mapstructure:"validator_address" json:"validator_address"`
}

// MatrixConfig represents the nitro configuration
type MatrixConfig struct {
	Server                        string `mapstructure:"matrix_connection" json:"matrix_connection"`
	EventsServiceRegistrationPath string `mapstructure:"events_service_registration_path" json:"events_service_registration_path"`
	ChatServiceRegistrationPath   string `mapstructure:"chat_service_registration_path" json:"chat_service_registration_path"`
}

// Config is the configuration for the application
type Config struct {
	Name     string        `mapstructure:"name" json:"name"`
	Version  string        `mapstructure:"version" json:"version"`
	NodeHome string        `mapstructure:"default_node_home" json:"default_node_home"`
	Sonrd    SonrdConfig   `mapstructure:"sonrd" json:"sonrd"`
	Highway  HighwayConfig `mapstructure:"highway" json:"highway"`
	Matrix   MatrixConfig  `mapstructure:"nitro" json:"nitro"`
}

// LoadConfig loads the configuration from the file
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
