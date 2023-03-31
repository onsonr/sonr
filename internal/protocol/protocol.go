package protocol

import (
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gofiber/fiber/v2"
	"github.com/sonr-io/sonr/internal/local"
	"github.com/sonr-io/sonr/internal/protocol/transport/rest"
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
				fmt.Sprintf("%s:%s", currPublicHostIP(), snrctx.HighwayPort()),
			)
		}
	}
}
func currPublicHostIP() string {
	if ip := os.Getenv("PUBLC_HOST_IP"); ip != "" {
		return ip
	}
	return "localhost"
}
