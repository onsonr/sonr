package protocol

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gofiber/fiber/v2"
	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/internal/protocol/transport/rest"
)
func RegisterHighway(ctx client.Context) {
	app := rest.NewHttpTransport(ctx)
	go serveFiber(app.App)
}

func serveFiber(app *fiber.App) {
	snrctx := local.NewContext()
	if snrctx.HasTlsCert() {
		app.ListenTLS(
			fmt.Sprintf(":%s", snrctx.HighwayPort()),
			snrctx.TlsCertPath,
			snrctx.TlsKeyPath,
		)
	} else {
		if snrctx.IsDev() {

			app.Listen(
				fmt.Sprintf(":%s", snrctx.HighwayPort()),
			)
		} else {
			app.Listen(
				fmt.Sprintf("%s:%s", snrctx.GrpcEndpoint(), snrctx.HighwayPort()),
			)
		}
	}
}
