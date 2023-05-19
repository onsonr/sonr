package gateway

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gofiber/fiber/v2"
	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/x/identity/client/gateway/config"
	"github.com/sonrhq/core/x/identity/client/gateway/routes"
)

func RegisterGateway(ctx client.Context) {
	conf := config.NewConfig(ctx)
	routes.SetupRoutes(conf)
	go serveFiber(conf.App)
}

func serveFiber(app *fiber.App) {
	if local.Context().HasTlsCert() {
		app.ListenTLS(
			local.Context().FiberListenAddress(),
			local.Context().TlsCertPath,
			local.Context().TlsKeyPath,
		)
	} else {
		app.Listen(local.Context().FiberListenAddress())
	}
}
