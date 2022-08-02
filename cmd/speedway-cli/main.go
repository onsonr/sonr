package main

import (
	"github.com/kataras/golog"
	"github.com/sonr-io/sonr/cmd/speedway-cli/speedwaycmd"
	"github.com/spf13/cobra"
)

var (
	logger = golog.Default.Child("motor-cli")
)

func main() {
	logger.Info("Starting motor-cli")
	cobra.CheckErr(speedwaycmd.Execute())
}
