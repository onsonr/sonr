package main

import (
	"github.com/sonr-io/sonr/cmd/highway-cli/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// PEM_CERT_FILE is the path to the certificate file.
	PEM_CERT_FILE = "cert.pem"

	// PEM_KEY_FILE is the file containing the private key.
	PEM_KEY_FILE = "key.pem"
)

// load environment variables
func loadEnv() error {
	viper.AddConfigPath(".")

	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	return viper.ReadInConfig()
}

func main() {
	cobra.CheckErr(loadEnv())
	cobra.CheckErr(cmd.Execute())
}
