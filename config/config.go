package config

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/spf13/viper"
)

//go:embed config.yaml
var gatewayConfigYaml []byte

var Config SonrConfig

// Define structs to match your configuration structure
type SonrConfig struct {
    Gateway    Gateway    `mapstructure:"gateway"`
    Services   Services   `mapstructure:"services"`
    Blockchain Blockchain `mapstructure:"blockchain"`
}

type Gateway struct {
    Port int    `mapstructure:"port"`
    Host string `mapstructure:"host"`
}

type Services struct {
    Sonrd  Sonrd  `mapstructure:"sonrd"`
    Matrix Matrix `mapstructure:"matrix"`
}

type Sonrd struct {
    Grpc Grpc `mapstructure:"grpc"`
}

type Grpc struct {
    Port int    `mapstructure:"port"`
    Host string `mapstructure:"host"`
}

type Matrix struct {
    Server string `mapstructure:"server"`
}

type Blockchain struct {
    ChainID   string `mapstructure:"chain_id"`
    MinGasFees string `mapstructure:"min_gas_fees"`
}

func init() {
    var conf SonrConfig
 	// Using viper.ReadConfig to read the embedded config
    if err := viper.ReadConfig(bytes.NewBuffer(gatewayConfigYaml)); err != nil {
        log.Fatalf("Error reading config: %s", err)
    }

    // Unmarshal the configuration into the Config struct
    if err := viper.Unmarshal(&conf); err != nil {
        log.Fatalf("Unable to decode into struct: %s", err)
    }

    // Now you can use the conf object which is filled with the data from your config
    log.Printf("Config: %+v\n", conf)
	Config = conf
}
