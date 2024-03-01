package cmds

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"

	"github.com/sonrhq/sonr/config"
	"github.com/sonrhq/sonr/pkg/middleware/common"
	"github.com/sonrhq/sonr/pkg/routes"
)

// HwayCommand returns the serve command for the Sonr Highway
func HwayCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "hway",
		Short:                      "Serves the Sonr Highway",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		Run:                        serveAction,
	}
	cmd.Flags().String("hway-host", "0.0.0.0", "host")
	cmd.Flags().Int("hway-port", 8000, "port")
	cmd.Flags().String("hway-psql", "postgresql://sonr:sonr@localhost:5432/sonr?sslmode=disable", "psql connection string")
	cmd.Flags().String("val-host", "sonrd", "validator host")
	cmd.Flags().Int("val-rpc", 26657, "validator rpc port")
	cmd.Flags().Int("val-grpc", 9090, "validator grpc port")
	return cmd
}

func serveAction(c *cobra.Command, _ []string) {
	cnfg := config.CreateHwayConfig()
	cnfg.ReadFlags(c)

	// Create the router
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	common.UseDefaults(e)

	// Register the routes
	routes.RegisterCosmosAPI(e)
	routes.RegisterSonrAPI(e)
	routes.RegisterStaticAssets(e, cnfg.Assets)
	routes.RegisterHTMXPages(e)
	routes.RegisterHTMXModals(e)

	// Serve the router
	cnfg.PrintBanner()
	e.Logger.Fatal(e.Start(cnfg.ListenAddress()))
}