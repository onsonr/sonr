package config

import (
	"fmt"

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
}

func setupDefaults() {
	viper.SetDefault("node.api.host", "localhost:1317")
	viper.SetDefault("node.grpc.host", "localhost:9090")
	viper.SetDefault("icefire.kv.host", "localhost:6379")
}
