package ante

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gogo/protobuf/grpc"
)

type SonrConfigurator interface {
	module.Configurator

	GatewayServer() grpc.Server
}

type sonrConfigurator struct {
	gatewayServer grpc.Server
	module.Configurator
}

func NewSonrConfigurator(cdc codec.Codec, msgServer grpc.Server, queryServer grpc.Server, gatewayServer grpc.Server) SonrConfigurator {
	c := module.NewConfigurator(cdc, msgServer, queryServer)
	return &sonrConfigurator{
		Configurator:  c,
		gatewayServer: gatewayServer,
	}
}

func (c *sonrConfigurator) GatewayServer() grpc.Server {
	return c.gatewayServer
}
