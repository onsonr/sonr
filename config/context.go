package config

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetEnvPrefix("sonr")
	viper.AutomaticEnv()
	viper.SetConfigName("sonr")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("$HOME/.sonr")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath("/etc/sonr")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	viper.SetDefault("highway.icefirekv.host", "localhost")
	viper.SetDefault("highway.icefirekv.port", 6001)
}
