package main

import (
	"fmt"
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/sonrhq/sonr/app"
	"github.com/sonrhq/sonr/app/params"
	"github.com/sonrhq/sonr/cmd/sonrd/cmd"
)

var (
  // Commit is set by the compiler via -ldflags.
  Commit = "unset"

  // Version is set by the compiler via -ldflags.
  Version = "unset"
)


// init sets the version flags.
func init() {
	version.Name = "sonr"
	version.AppName = "sonrd"
	version.Version = Version
	version.Commit = Commit
}

// main is the entry point for the application.
func main() {
	params.SetAddressPrefixes()
	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "SONR", app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
