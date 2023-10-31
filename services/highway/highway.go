package highway

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sonrhq/sonr/services/highway/handler"
)

var hway *Instance

// Instance is the local process instance of the Highway Service Server.
type Instance struct {
	ctx 	  context.Context
	addr 	string
}

// RegisterHighwayGateway registers the Highway Service Server.
func RegisterHighwayGateway(c *Config) {
hway = &Instance{
		ctx: context.Background(),
		addr: fmt.Sprintf("%s:%s", c.Host, c.Port),
	}
	mux := runtime.NewServeMux()
	handler.RegisterHandlers(hway.ctx, mux)
	go hway.Serve(mux)
}

// utility function to get the base API URL
func (i *Instance) Serve(mux *runtime.ServeMux) {
	s := &http.Server{
		Addr:    i.addr,
		Handler: mux,
	}

	fmt.Printf("Starting Highway Gateway on %s\n", i.addr)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		panic(err)
	}
}

type Config struct {
	Host string
	Port string
}

func NewConfig(host string, port string) *Config {
	return &Config{
		Host: host,
		Port: port,
	}
}

func DefaultConfig() *Config {
	return &Config{
		Host: "localhost",
		Port: "8080",
	}
}
