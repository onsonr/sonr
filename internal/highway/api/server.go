// @title Highway API
// @version v0.23.0
// @description Manage your Sonr Powered services and blockchain registered types with the Highway API.
// @contact.name Sonr Inc.
// @contact.url https://sonr.io
// @contact.email team@sonr.io
// @license.name OpenGLv3
// @host localhost:8080
// @BasePath /v1

package api

import (
	"context"
	"errors"
	"github.com/sonr-io/sonr/internal/highway/x/store"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kataras/golog"
	"github.com/sonr-io/sonr/internal/highway/x/ipfs"
	metrics "github.com/sonr-io/sonr/internal/highway/x/prometheus"
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/config"
	hn "github.com/sonr-io/sonr/pkg/host"
	"github.com/sonr-io/sonr/pkg/jwt"
	ctv1 "github.com/sonr-io/sonr/x/channel/types"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Error Definitions
var (
	logger                 = golog.Default.Child("node/highway")
	ErrEmptyQueue          = errors.New("No items in Transfer Queue.")
	ErrInvalidQuery        = errors.New("No SName or PeerID provided.")
	ErrMissingParam        = errors.New("Paramater is missing.")
	ErrProtocolsNotSet     = errors.New("Node Protocol has not been initialized.")
	ErrMethodUnimplemented = errors.New("Method is not implemented.")
)

// HighwayServer is the RPC Service for the Custodian Node.
type HighwayServer struct {
	// Config
	ctx    context.Context
	Config *config.Config

	// Clients
	Host     hn.SonrHost
	Cosmos   *client.Cosmos
	JWTToken *jwt.JWT

	// Http Properties
	Router     *gin.Engine
	HTTPServer *http.Server

	// Protocols
	channels     map[string]ctv1.Channel
	ipfsProtocol *ipfs.Protocol
	store        *store.Store
	// matrixProtocol *matrix.MatrixProtocol

	//Prometheus
	Telemetry *metrics.HighwayTelemetry
}

// setupBaseStub creates the base Highway Server.
func CreateStub(ctx context.Context, c *config.Config) (*HighwayServer, error) {
	node, err := hn.NewDefaultHost(ctx, c)
	if err != nil {
		return nil, err
	}

	// Create a new Cosmos Client for Sonr Blockchain
	cosmos, err := client.NewCosmos(ctx, c)
	if err != nil {
		return nil, err
	}

	ds := store.New("tmp")

	ipfs, err := ipfs.New(ctx, ds, node)
	if err != nil {
		return nil, err
	}

	tokenClient := jwt.New(ctx, node)
	metrics, err := metrics.New(ctx, node)

	// TODO: Enabling Matrix Protocol breaks build for Darwin
	// https://github.com/sonr-io/sonr/issues/330

	// Create the RPC Service
	stub := &HighwayServer{
		Cosmos:       cosmos,
		Host:         node,
		ctx:          ctx,
		Router:       gin.Default(),
		Config:       c,
		ipfsProtocol: ipfs,
		store:        ds,
		JWTToken:     tokenClient,
		// matrixProtocol: matrix,
		Telemetry: metrics,
	}
	return stub, nil
}

func (s *HighwayServer) ConfigureRoutes() {
	// Register Cosmos HTTP Routes - Registry
	s.Router.POST("/v1/registry/alias/buy", s.BuyAlias)
	s.Router.POST("/v1/registry/alias/sell", s.SellAlias)
	s.Router.POST("/v1/registry/alias/transfer", s.TransferAlias)

	// Register Cosmos HTTP Routes - Bucket
	s.Router.POST("/v1/bucket/create", s.CreateBucket)
	s.Router.POST("/v1/bucket/update", s.UpdateBucket)
	s.Router.POST("/v1/bucket/deactivate", s.DeactivateBucket)

	// Register Cosmos HTTP Routes - Channel
	s.Router.POST("/v1/channel/create", s.CreateChannel)
	s.Router.POST("/v1/channel/update", s.UpdateChannel)
	s.Router.POST("/v1/channel/deactivate", s.DeactivateChannel)

	// Register Blob HTTP Routes
	s.Router.POST("/v1/blob/upload", s.UploadBlob)
	s.Router.GET("/v1/blob/download/:cid", s.DownloadBlob)
	s.Router.POST("/v1/blob/remove/:cid", s.UnpinBlob)

	// WebAuthn Endpoints
	s.Router.POST("/v1/registry/whois/create", s.CreateWhoIs)
	s.Router.POST("/v1/registry/whois/update", s.UpdateWhoIs)
	s.Router.POST("/v1/registry/whois/deactivate", s.DeactivateWhoIs)

	// Setup Swagger UI
	s.Router.GET("v1/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	s.Router.GET("/metrics", gin.WrapH(s.Telemetry.GetMetricsHandler()))
}

func (s *HighwayServer) ConfigureMiddleware() {
	if s.Router == nil {
		logger.Warn("Cannot configure middleware, router is not yet created")
		return
	}

	s.Router.Use(gin.Logger())

	// Registering middleware for authorization header parsing and creation of a `Token` object
	// Currently disabled and will not invoke on requests
	s.AddMiddlewareDefinition(HighwayMiddleware{
		definition: func(ctx *gin.Context) {
			token := ctx.GetHeader("Authorization")
			if token == "" {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
					Message: "Authorization token not found",
				})

				return
			}
			error := s.JWTToken.BuildJWTParseMiddleware(token)()

			if error != nil {
				logger.Errorf("Error while processing authorization header: %s", error.Error())
				ctx.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
					Message: error.Error(),
				})
				return
			}

			ctx.Next()
		},
		disabled: true,
	})

	// register custom middleware defined within package
	// see Middleware.go for definitions
	s.RegisterMiddleWare()
}

// Serve starts the RPC Service.
func (s *HighwayServer) Serve() {
	// Print the Server Address's
	logger.Infof("Serving HTTP Server on %s", s.Config.HighwayHTTPEndpoint)

	// Start Highway HTTP server on a separate goroutine
	go func() {
		// Start HTTP server (and proxy calls to gRPC server endpoint)
		if err := s.HTTPServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("%s - Failed to start HTTP server", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	<-quit
	logger.Warn("Shutting down HTTP server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.HTTPServer.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown: ", err)
	}
	logger.Info("Goodbye!")
}
