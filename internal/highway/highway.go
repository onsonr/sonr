package highway

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
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
	s.Engine.POST("/v1/registry/alias/buy", s.BuyAlias)
	s.Engine.POST("/v1/registry/alias/sell", s.SellAlias)
	s.Engine.POST("/v1/registry/alias/transfer", s.TransferAlias)

	// Register Cosmos HTTP Routes - Object
	s.Engine.POST("/v1/object/create", s.CreateObject)
	s.Engine.POST("/v1/object/update", s.UpdateObject)
	s.Engine.POST("/v1/object/deactivate", s.DeactivateObject)

	// Register Cosmos HTTP Routes - Bucket
	s.Engine.POST("/v1/bucket/create", s.CreateBucket)
	s.Engine.POST("/v1/bucket/update", s.UpdateBucket)
	s.Engine.POST("/v1/bucket/deactivate", s.DeactivateBucket)

	// Register Cosmos HTTP Routes - Channel
	s.Engine.POST("/v1/channel/create", s.CreateChannel)
	s.Engine.POST("/v1/channel/update", s.UpdateChannel)
	s.Engine.POST("/v1/channel/deactivate", s.DeactivateChannel)

	// Register Blob HTTP Routes
	s.Engine.POST("/v1/blob/upload", s.UploadBlob)
	s.Engine.GET("/v1/blob/download/:cid", s.DownloadBlob)
	s.Engine.POST("/v1/blob/remove/:cid", s.RemoveBlob)

	// WebAuthn Endpoints
	s.Engine.POST("/v1/registry/whois/create", s.CreateWhoIs)
	s.Engine.POST("/v1/registry/whois/update", s.UpdateWhoIs)
	s.Engine.POST("/v1/registry/whois/deactivate", s.DeactivateWhoIs)

	// Setup Swagger UI
	s.Engine.GET("v1/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	s.Engine.GET("/metrics", gin.WrapH(s.Telemetry.GetMetricsHandler()))

	// Setup the Webauthn HTTP Server
	s.Router.HandleFunc("/register/begin/{username}", s.StartRegisterName).Methods("GET")
	s.Router.HandleFunc("/register/finish/{username}", s.FinishRegisterName).Methods("POST")
	s.Router.HandleFunc("/login/begin/{username}", s.StartAccessName).Methods("GET")
	s.Router.HandleFunc("/login/finish/{username}", s.FinishAccessName).Methods("POST")
	s.Router.PathPrefix("/").Handler(http.FileServer(http.Dir(c.WebAuthNFSDir)))

	// Setup HTTP Server
	s.HTTPServer = &http.Server{
		Addr:    s.Config.HighwayHTTPEndpoint,
		Handler: s.Engine,
	}
	return s, nil
}
