package highway

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sonr-io/sonr/internal/highway/api"
	"github.com/sonr-io/sonr/pkg/config"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewHighwayServer creates a new Highway service stub for the node.
func NewHighway(ctx context.Context, opts ...config.Option) (*api.HighwayServer, error) {
	// Create Config
	c := config.DefaultConfig(config.Role_HIGHWAY)
	for _, opt := range opts {
		if opt != nil {
			opt(c)
		}
	}

	// Create the Highway Server
	s, err := api.CreateStub(ctx, c)
	if err != nil {
		return nil, err
	}

	// Register Cosmos HTTP Routes - Registry
	// s.Router.POST("/v1/registry/alias/buy", s.BuyAlias)
	// s.Router.POST("/v1/registry/alias/sell", s.SellAlias)
	// s.Router.POST("/v1/registry/alias/transfer", s.TransferAlias)

	// // Register Cosmos HTTP Routes - Bucket
	// s.Router.POST("/v1/bucket/create", s.CreateBucket)
	// s.Router.POST("/v1/bucket/update", s.UpdateBucket)
	// s.Router.POST("/v1/bucket/deactivate", s.DeactivateBucket)

	// // Register Cosmos HTTP Routes - Channel
	// s.Router.POST("/v1/channel/create", s.CreateChannel)
	// s.Router.POST("/v1/channel/update", s.UpdateChannel)
	// s.Router.POST("/v1/channel/deactivate", s.DeactivateChannel)

	// // Register Blob HTTP Routes
	// s.Router.POST("/v1/blob/upload", s.UploadBlob)
	// s.Router.GET("/v1/blob/download/:cid", s.DownloadBlob)
	// s.Router.POST("/v1/blob/remove/:cid", s.UnpinBlob)

	// // WebAuthn Endpoints
	// s.Router.POST("/v1/registry/whois/create", s.CreateWhoIs)
	// s.Router.POST("/v1/registry/whois/update", s.UpdateWhoIs)
	// s.Router.POST("/v1/registry/whois/deactivate", s.DeactivateWhoIs)

	// Setup Swagger UI
	s.Router.GET("v1/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	s.Router.GET("/metrics", gin.WrapH(s.Telemetry.GetMetricsHandler()))

	// Setup HTTP Server
	s.HTTPServer = &http.Server{
		Addr:    s.Config.HighwayHTTPEndpoint,
		Handler: s.Router,
	}
	return s, nil
}
