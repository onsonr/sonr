//go:build !wasm
// +build !wasm

package main

import (
	"github.com/kataras/golog"
	"github.com/sonr-io/sonr/cmd/highway-cli/highwaycmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// PEM_CERT_FILE is the path to the certificate file.
	PEM_CERT_FILE = "cert.pem"

	// PEM_KEY_FILE is the file containing the private key.
	PEM_KEY_FILE = "key.pem"

	CONFIG_PATH_LOCAL = "."

	CONFIG_PATH_ROOT = "../../"
)

var (
	logger = golog.Default.Child("highway-cli")
)

// load environment variables
func loadEnv() error {
	viper.AddConfigPath(CONFIG_PATH_LOCAL)
	viper.AddConfigPath(CONFIG_PATH_ROOT)

	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	return viper.ReadInConfig()
}

func main() {
	logger.Info("Starting highway-cli, loading env variables")
	cobra.CheckErr(loadEnv())
	cobra.CheckErr(highwaycmd.Execute())
}
