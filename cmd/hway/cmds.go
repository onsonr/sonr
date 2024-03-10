package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"

	"github.com/sonrhq/sonr/cmd/hway/config"
	"github.com/sonrhq/sonr/pkg/middleware/common"
	"github.com/sonrhq/sonr/pkg/routes"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Sonr Highway",
	Run: func(cmd *cobra.Command, args []string) {
		pterm.Info.Println(fmt.Sprintf("Version: %s", Version))
	},
}

var rootCmd = &cobra.Command{
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
