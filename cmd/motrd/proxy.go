package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"

	"github.com/onsonr/sonr/pkg/nebula"
	"github.com/onsonr/sonr/pkg/nebula/pages"
)

func NewProxyCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "proxy",
		Short: "Starts the DWN proxy server for the local IPFS node",
		Run: func(cmd *cobra.Command, args []string) {
			// Echo instance
			e := echo.New()
			e.Logger.SetLevel(log.INFO)

			// Configure the server
			if err := nebula.UseAssets(e); err != nil {
				e.Logger.Fatal(err)
			}

			e.GET("/", pages.Home)
			e.GET("/allocate", pages.Profile)

			if err := e.Start(":1323"); err != http.ErrServerClosed {
				log.Fatal(err)
			}
		},
	}
}
