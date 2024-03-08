package config

import (
	"github.com/spf13/viper"
)

// Highway represents the highway configuration
type Highway struct {
	// ValidatorHost is the host of the validator
	ValidatorHost string `mapstructure:"validator_address" json:"validator_address" yaml:"validator_address"`

	// ValidatorRPC is the port of the validator for rpc
	ValidatorRPC int `mapstructure:"validator_rpc_port" json:"validator_rpc_port" yaml:"validator_rpc_port"`

	// ValidatorWS is the port of the validator for websocket
	ValidatorWS int `mapstructure:"validator_ws_port" json:"validator_ws_port" yaml:"validator_ws_port"`

	// ValidatorGRPC is the port of the validator for grpc
	ValidatorGRPC int `mapstructure:"validator_grpc_port" json:"validator_grpc_port" yaml:"validator_grpc_port"`

	// GatewayPort is the port that the gateway listens on
	GatewayPort int `json:"gateway_port" yaml:"gateway_port"`

	// Host is the host that the gateway listens on
	Host string `json:"host" yaml:"host"`

	// PostgresConnection is the connection string for the postgres database
	PostgresConnection string `json:"postgres_connection" yaml:"postgres_connection"`

	// RedisConnection is the connection string for the redis database
	RedisConnection string `json:"redis_connection" yaml:"redis_connection"`

	// SmtpHost is the host of the smtp server
	SmtpHost string `json:"smtp_host" yaml:"smtp_host"`

	// SmtpPort is the port of the smtp server
	SmtpPort int `json:"smtp_port" yaml:"smtp_port"`

	// SmtpUser is the user of the smtp server
	SmtpUser string `json:"smtp_user" yaml:"smtp_user"`

	// SmtpPassword is the password of the smtp server
	SmtpPassword string `json:"smtp_password" yaml:"smtp_password"`
}

// LoadConfig loads the configuration from the file
func LoadConfig() (*Highway, error) {
	viper.SetConfigName("highway")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Highway
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
