package highway

import (
	client "github.com/cosmos/cosmos-sdk/client"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sonrhq/core/internal/highway/handler"
	"github.com/sonrhq/core/internal/highway/router"
	"google.golang.org/grpc"
)

var hway *Instance

// Instance is the local process instance of the Highway Service Server.
type Instance struct {
	grpcServer *grpc.Server
}

func init() {
	hway = &Instance{
		grpcServer: grpc.NewServer(),
	}
}

// RegisterHighwayGateway registers the Highway Service Server.
func RegisterHighwayGateway(cctx client.Context, mux *runtime.ServeMux) {
	bAPIURL := getBaseAPIURL(cctx)
	handler.RegisterHandlers(cctx, hway.grpcServer)
	router.RegisterRouter(mux, bAPIURL)
}

// utility function to get the base API URL
func getBaseAPIURL(cctx client.Context) string {
	b := "localhost:1317"
	tAPIAddr := cctx.GRPCClient.Target()
	if tAPIAddr != "" {
		b = tAPIAddr
	}
	vAPIAddr := cctx.Viper.GetString("api.address")
	if vAPIAddr != "" && vAPIAddr != "localhost:1317" && vAPIAddr != b {
		b = vAPIAddr
	}
	return b
}
