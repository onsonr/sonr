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
	// ValidatorHost is the host of the validator
	ValidatorHost string `mapstructure:"validator_address" json:"validator_address"`

	// ValidatorRPC is the port of the validator for rpc
	ValidatorRPC int `mapstructure:"validator_rpc_port" json:"validator_rpc_port"`

	// ValidatorWS is the port of the validator for websocket
	ValidatorWS int `mapstructure:"validator_ws_port" json:"validator_ws_port"`

	// ValidatorGRPC is the port of the validator for grpc
	ValidatorGRPC int `mapstructure:"validator_grpc_port" json:"validator_grpc_port"`

	// GatewayPort is the port that the gateway listens on
	GatewayPort int `json:"gateway_port"`

	// Host is the host that the gateway listens on
	Host string `json:"host"`

	// PostgresConnection is the connection string for the postgres database
	PostgresConnection string `json:"postgres_connection"`

	// RedisConnection is the connection string for the redis database
	RedisConnection string `json:"redis_connection"`

	// SmtpHost is the host of the smtp server
	SmtpHost string `json:"smtp_host"`

	// SmtpPort is the port of the smtp server
	SmtpPort int `json:"smtp_port"`

	// SmtpUser is the user of the smtp server
	SmtpUser string `json:"smtp_user"`

	// SmtpPassword is the password of the smtp server
	SmtpPassword string `json:"smtp_password"`
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
