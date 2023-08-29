package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func init() {
	setupDefaults()
	viper.SetEnvPrefix("sonr")
	viper.AutomaticEnv()
	confPath := os.Getenv("SONR_LAUNCH_CONFIG")
	if confPath != "" {
		viper.SetConfigFile(confPath)
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())
}
