package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"

	"github.com/sonrhq/sonr/pkg/config"
	"github.com/sonrhq/sonr/pkg/middleware"
	"github.com/sonrhq/sonr/pkg/routes"
)

// ServeCommand returns the serve command
func ServeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "hway",
		Short:                      "Serves the Sonr Highway",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		Run:                        serveAction,
	}
	return cmd
}

func serveAction(_ *cobra.Command, _ []string) {
	// 1. Read config from file
	// 2. Check reachable to enabled services
	// 3. Start Gateway router as system service
	// Set the default options
	cnfg := config.NewHighwayOptions()

	// Create the router
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	middleware.UseDefaults(e)
	routes.RegisterCosmosAPI(e)
	routes.RegisterSonrAPI(e)
	routes.RegisterHTMXPages(e)

	// Serve the router
	cnfg.PrintBanner()
	e.Logger.Fatal(e.Start(cnfg.ListenAddress()))
}
