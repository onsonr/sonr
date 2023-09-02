package config

import "github.com/spf13/viper"

// EnvChainID returns the chain ID from the configuration. (default: sonr-localnet-1)
func EnvChainID() string {
	return viper.GetString("launch.chain-id")
}
