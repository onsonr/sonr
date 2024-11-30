package gateway

import (
	"net/http"

	"github.com/ipfs/kubo/client/rpc"
	"github.com/labstack/echo/v4"
)

type IPFSClient = *rpc.HttpApi

var ipfs IPFSClient

func RegisterRoutes(e *echo.Echo, client IPFSClient) {
	ipfs = client

	// Custom error handler for gateway
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			// Log the error if needed
			c.Logger().Errorf("Gateway error: %v", he.Message)
		}
		// Redirect to main site
		c.Redirect(http.StatusFound, "http://localhost:3000")
	}
	gw := New(client)
	e.POST("/_spawn", SpawnVault)
	e.Any("/*", gw.Handler())
}
