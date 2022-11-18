//go:build !wasm
// +build !wasm

package main

import (
	"log"
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/getsentry/sentry-go"
	"github.com/kataras/golog"
	"github.com/sonr-io/sonr/app"
	 "github.com/sonr-io/sonr/cmd/sonrd/cmd"
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
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SentryDSN"),
		TracesSampleRate: 0.7,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	sentry.CaptureMessage("Sentry exception catching enabled, proceeding with binary execution.")
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
