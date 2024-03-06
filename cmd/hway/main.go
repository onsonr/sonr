package main

import (
	"fmt"
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"

	"github.com/sonrhq/sonr/app"
	"github.com/sonrhq/sonr/app/params"
	"github.com/sonrhq/sonr/cmd/sonrd/cmds"
	"github.com/sonrhq/sonr/config"
	"github.com/sonrhq/sonr/pkg/middleware/common"
	"github.com/sonrhq/sonr/pkg/routes"
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
	version.BuildTags = "netgo"
}

// main is the entry point for the application.
func main() {
	params.SetAddressPrefixes()
	rootCmd := cmds.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "SONR", app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}

// HwayCommand returns the serve command for the Sonr Highway
func HwayCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "hway",
		Short:                      "Serves the Sonr Highway",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		Run:                        serveHighway,
	}
	cmd.Flags().String("hway-host", "0.0.0.0", "host")
	cmd.Flags().Int("hway-port", 8000, "port")
	cmd.Flags().String("hway-psql", "postgresql://sonr:sonr@localhost:5432/sonr?sslmode=disable", "psql connection string")
	cmd.Flags().String("val-host", "sonrd", "validator host")
	cmd.Flags().Int("val-rpc", 26657, "validator rpc port")
	cmd.Flags().Int("val-grpc", 9090, "validator grpc port")
	return cmd
}

func serveHighway(c *cobra.Command, _ []string) {
	// Create the config
	cnfg := config.NewHway()
	cnfg.ReadFlags(c)

	// Create the router
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	common.UseDefaults(e)
	err := common.UseCache(e)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Register the routes
	routes.RegisterCosmosAPI(e)
	routes.RegisterSonrAPI(e)
	routes.RegisterUI(e)

	// Serve the router
	cnfg.PrintBanner()
	e.Logger.Fatal(e.Start(cnfg.ListenAddress()))
}
