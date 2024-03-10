package main

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/version"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/pterm/pterm"
	"github.com/sonrhq/sonr/cmd/hway/config"
	"github.com/sonrhq/sonr/pkg/middleware/common"
	"github.com/sonrhq/sonr/pkg/routes"
	"github.com/spf13/cobra"
)

var (
	// Commit is set by the compiler via -ldflags.
	Commit = "unset"

	// Version is set by the compiler via -ldflags.
	Version = "unset"
)

// init sets the version flags.
func init() {
	version.Name = "Sonr Highway"
	version.AppName = "hway"
	version.Version = Version
	version.Commit = Commit
}

// main is the entry point for the application.
func main() {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Sonr Highway",
		Run: func(cmd *cobra.Command, args []string) {
			pterm.Info.Println(fmt.Sprintf("Version: %s", Version))
		},
	}

	rootCmd := &cobra.Command{
		Use:                        "hway",
		Short:                      "Serves the Sonr Highway",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		Run: func(c *cobra.Command, _ []string) {
			// Create the config
			cnfg := config.NewHway()
			cnfg.ReadFlags(c)

			// Create the router
			e := echo.New()
			e.Logger.SetLevel(log.INFO)
			common.UseDefaults(e)

			// Register the routes
			routes.RegisterCosmosAPI(e)
			routes.RegisterSonrAPI(e)
			routes.RegisterUI(e)

			// Serve the router
			cnfg.PrintBanner()
			e.Logger.Fatal(e.Start(cnfg.ListenAddress()))
		},
	}

	rootCmd.Flags().String("hway-host", "0.0.0.0", "host")
	rootCmd.Flags().Int("hway-port", 8000, "port")
	rootCmd.Flags().String("hway-psql", "postgresql://sonr:sonr@localhost:5432/sonr?sslmode=disable", "psql connection string")
	rootCmd.Flags().String("val-host", "localhost", "validator host")
	rootCmd.Flags().Int("val-rpc", 26657, "validator rpc port")
	rootCmd.Flags().Int("val-grpc", 9090, "validator grpc port")
	rootCmd.AddCommand(versionCmd)
	err := rootCmd.Execute()
	if err != nil {
		rootCmd.PrintErr(err)
	}
}
