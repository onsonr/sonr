package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

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

			// Start server
			ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
			defer stop()
			// Start server
			go func() {
				if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
					e.Logger.Fatal("shutting down the server")
				}
			}()

			// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
			<-ctx.Done()
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			// Shutdown the server with 10 seconds timeout.
			if err := e.Shutdown(ctx); err != nil {
				e.Logger.Fatal(err)
			}
		},
	}
}
