// @title Highway API
// @version v0.23.0
// @description Manage your Sonr Powered services and blockchain registered types with the Highway API.
// @contact.name Sonr Inc.
// @contact.url https://sonr.io
// @contact.email team@sonr.io
// @license.name OpenGLv3
// @host localhost:8080
// @BasePath /v1

package core

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kataras/golog"
	"github.com/sonr-io/sonr/internal/highway/x/ipfs"
	"github.com/sonr-io/sonr/internal/highway/x/jwt"
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/config"
	hn "github.com/sonr-io/sonr/pkg/host"
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
	Webauthn *client.WebAuthn
	JwtToken *jwt.JWT

	// Http Properties
	Router     *gin.Engine
	HTTPServer *http.Server

	// Protocols
	channels     map[string]ctv1.Channel
	ipfsProtocol *ipfs.IPFSProtocol
	// matrixProtocol *matrix.MatrixProtocol
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

	// Create a new WebAuthn Client for Sonr Blockchain
	webauthn, err := client.NewWebauthn(ctx, c)
	if err != nil {
		return nil, err
	}

	// Create the IPFS Protocol
	ipfs, err := ipfs.New(ctx, node)
	if err != nil {
		return nil, err
	}

	tokenClient := jwt.New(ctx, node)
	// TODO: Enabling Matrix Protocol breaks build for Darwin
	// Create the Matrix Protocol
	// matrix, err := matrix.New(ctx, node)
	// if err != nil {
	// 	return nil, err
	// }

	// Create the RPC Service
	stub := &HighwayServer{
		Cosmos:       cosmos,
		Host:         node,
		ctx:          ctx,
		Router:       gin.Default(),
		Config:       c,
		Webauthn:     webauthn,
		ipfsProtocol: ipfs,
		JwtToken:     &tokenClient,
		// matrixProtocol: matrix,
	}
	return stub, nil
}

func (s *HighwayServer) ConfigureRoutes() {
	// Register Cosmos HTTP Routes - Registry
	s.Router.POST("/v1/registry/create/whois", s.CreateWhoIs)
	s.Router.POST("/v1/registry/update/whois", s.UpdateWhoIs)
	s.Router.POST("/v1/registry/deactivate/whois", s.DeactivateWhoIs)
	s.Router.POST("/v1/registry/buy/alias/name", s.BuyNameAlias)
	s.Router.POST("/v1/registry/buy/alias/app", s.BuyAppAlias)
	s.Router.POST("/v1/registry/transfer/alias/name", s.TransferNameAlias)
	s.Router.POST("/v1/registry/transfer/alias/app", s.TransferAppAlias)

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

	// Register IPFS HTTP Routes
	s.Router.POST("/v1/ipfs/upload", s.UploadBlob)
	s.Router.GET("/v1/ipfs/download/:cid", s.DownloadBlob)
	s.Router.POST("/v1/ipfs/remove/:cid", s.RemoveBlob)

	// Register WebAuthn HTTP Routes
	s.Router.GET("/v1/auth/register/start/:username", s.StartRegisterName)
	s.Router.POST("/v1/auth/register/finish/:username", s.FinishRegisterName)
	s.Router.GET("/v1/auth/access/start/:username", s.StartAccessName)
	s.Router.POST("/v1/auth/access/finish/:username", s.FinishAccessName)

	// Setup Swagger UI
	s.Router.GET("v1/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (s *HighwayServer) ConfigureMiddleware() {
	s.Router.Use(func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != "" {
			error := s.JwtToken.BuildJWTParseMiddleware(token)()

			if error != nil {
				logger.Errorf("Error while processing authorization header", error)
				return
			}
			ctx.Next()
		}
	})
}

// Serve starts the RPC Service.
func (s *HighwayServer) Serve() {
	// Print the Server Address's
	logger.Infof("Serving HTTP Server on %s", s.Config.HighwayHTTPEndpoint)

	// Start HTTP server on a separate goroutine
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
