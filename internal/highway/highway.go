package highway

import (
	"context"
	"net/http"

	"github.com/sonr-io/sonr/internal/highway/api"
	"github.com/sonr-io/sonr/pkg/config"
)

// NewHighwayServer creates a new Highway service stub for the node.
func NewHighway(ctx context.Context, opts ...config.Option) (*api.HighwayServer, error) {
	// Create Config
	c := config.DefaultConfig(config.Role_HIGHWAY, "")
	for _, opt := range opts {
		if opt != nil {
			opt(c)
		}
	}

	// Create the Highway Server
	s, err := api.CreateStub(ctx, c)
	s.ConfigureMiddleware()
	s.ConfigureRoutes()

	if err != nil {
		return nil, err
	}

	// Setup HTTP Server
	s.HTTPServer = &http.Server{
		Addr:    s.Config.HighwayHTTPEndpoint,
		Handler: s.Router,
	}

	return s, nil
}
