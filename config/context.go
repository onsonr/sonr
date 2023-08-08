package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	setupDefaults()
	viper.SetEnvPrefix("sonr")
	viper.AutomaticEnv()

	configPath := viper.GetString("launch.config")
	if configPath != "" {
		viper.SetConfigFile(configPath)
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())
	PrintEnv()
}

// PrintEnv prints the environment variables
func PrintEnv() {
	opts := viper.GetViper().AllSettings()
	for k, v := range opts {
		fmt.Printf("%s: %v\n", k, v)
	}
}
