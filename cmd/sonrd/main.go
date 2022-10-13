//go:build !wasm
// +build !wasm

package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/kataras/golog"
	"github.com/sonr-io/sonr/app"
	"github.com/spf13/viper"
)

const (
	CONFIG_PATH_ROOT = "../../"
)

var (
	logger = golog.Default.Child("sonrd")
)

// load environment variables
func loadEnv() error {
	viper.AddConfigPath(CONFIG_PATH_ROOT)

	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	return viper.ReadInConfig()
}

func main() {
	rootCmd, _ := NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
