package highway

import (
	"context"
	"net/http"

	core "github.com/sonr-io/sonr/internal/highway/http"

	"github.com/sonr-io/sonr/pkg/config"
)

// NewHighwayServer creates a new Highway service stub for the node.
func NewHighway(ctx context.Context, opts ...config.Option) (*core.HighwayServer, error) {
	// Create Config
	c := config.DefaultConfig(config.Role_HIGHWAY)
	for _, opt := range opts {
		opt(c)
	}

	// Create the Highway Server
	s, err := core.CreateStub(ctx, c)
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
