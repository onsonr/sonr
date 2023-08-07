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
	viper.SetDefault("highway.icefirekv.host", "localhost")
	viper.SetDefault("highway.icefirekv.port", 6001)
	viper.SetDefault("highway.jwt.key", "@re33lyb@dsecret")
	viper.ReadInConfig()
}
