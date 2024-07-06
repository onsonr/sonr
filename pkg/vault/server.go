package vault

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/onsonr/hway/pkg/vault/middleware"
	"github.com/onsonr/hway/pkg/vault/routes"
	"github.com/labstack/echo/v4"
)

func Serve(ctx context.Context) {
	// Configure the server
	e := echo.New()

	// Use Middlewares
	e.Use(middleware.SessionCookies)
	e.Use(middleware.CacheStores)

	// Setup routes
	routes.RegisterPages(e)
	routes.RegisterAPI(e)

	// Run the server
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
