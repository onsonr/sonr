package host

import "github.com/spf13/viper"

const (
	CONFIG_PATH_ROOT = "../../"
)

// load environment variables
func loadEnv() error {
	viper.AddConfigPath(CONFIG_PATH_ROOT)

	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	return viper.ReadInConfig()
}
