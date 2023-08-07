package config

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetEnvPrefix("sonr")
	// viper.SetConfigName("sonr")
	// viper.SetConfigType("yml")

	// viper.AddConfigPath(".")
	// viper.AddConfigPath("../")
	// viper.AddConfigPath("../../")
	// viper.AddConfigPath("/.sonr")
	// viper.AddConfigPath("$HOME/.sonr")
	// viper.AddConfigPath("$HOME")

	viper.AutomaticEnv()
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	panic(err)
	// }
	// err = viper.Unmarshal(&c)
	// if err != nil {
	// 	panic(err)
	// }
}
