package protocol

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gofiber/fiber/v2"
	"github.com/sonrhq/core/internal/local"
	rest "github.com/sonrhq/core/internal/protocol/transport"
	"github.com/sonrhq/core/pkg/node"
)

func RegisterHighway(ctx client.Context) {
	app := rest.NewHttpTransport(ctx)
	node.StartLocalIPFS()
	go serveFiber(app.App)
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
