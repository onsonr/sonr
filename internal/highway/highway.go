package highway

import (
	"context"
	"net/http"

	core "github.com/sonr-io/sonr/internal/highway/http"
	"github.com/sonr-io/sonr/pkg/config"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// Register Cosmos HTTP Routes - Registry
	s.Router.POST("/v1/registry/buy/alias", s.BuyAlias)
	s.Router.POST("/v1/registry/sell/alias", s.SellAlias)
	s.Router.POST("/v1/registry/transfer/alias", s.TransferAlias)

	// Register Cosmos HTTP Routes - Object
	s.Router.POST("/v1/object/create", s.CreateObject)
	s.Router.POST("/v1/object/update", s.UpdateObject)
	s.Router.POST("/v1/object/deactivate", s.DeactivateObject)

	// Register Cosmos HTTP Routes - Bucket
	s.Router.POST("/v1/bucket/create", s.CreateBucket)
	s.Router.POST("/v1/bucket/update", s.UpdateBucket)
	s.Router.POST("/v1/bucket/deactivate", s.DeactivateBucket)

	// Register Cosmos HTTP Routes - Channel
	s.Router.POST("/v1/channel/create", s.CreateChannel)
	s.Router.POST("/v1/channel/update", s.UpdateChannel)
	s.Router.POST("/v1/channel/deactivate", s.DeactivateChannel)

	// Register WebAuthn HTTP Routes
	s.Router.GET("/v1/auth/register/start/:username", s.StartRegisterName)
	s.Router.POST("/v1/auth/register/finish/:username", s.FinishRegisterName)
	s.Router.GET("/v1/auth/access/start/:username", s.StartAccessName)
	s.Router.POST("/v1/auth/access/finish/:username", s.FinishAccessName)

	// Register Blob HTTP Routes
	s.Router.POST("/v1/blob/upload", s.UploadBlob)
	s.Router.GET("/v1/blob/download/:cid", s.DownloadBlob)
	s.Router.POST("/v1/blob/remove/:cid", s.RemoveBlob)

	// WebAuthn Endpoints
	s.Router.POST("/v1/registry/create/whois", s.CreateWhoIs)
	s.Router.POST("/v1/registry/update/whois", s.UpdateWhoIs)
	s.Router.POST("/v1/registry/deactivate/whois", s.DeactivateWhoIs)

	// Setup Swagger UI
	s.Router.GET("v1/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Setup HTTP Server
	s.HTTPServer = &http.Server{
		Addr:    s.Config.HighwayHTTPEndpoint,
		Handler: s.Router,
	}
	return s, nil
}
