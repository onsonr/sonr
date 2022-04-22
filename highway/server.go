// @title Highway API
// @version v0.23.0
// @description Manage your Sonr Powered services and blockchain registered types with the Highway API.
// @contact.name Sonr Inc.
// @contact.url https://sonr.io
// @contact.email team@sonr.io
// @license.name OpenGLv3
// @host localhost:8081
// @BasePath /v1

package highway

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kataras/golog"
	"github.com/sonr-io/core/channel"
	"github.com/sonr-io/core/device"
	docs "github.com/sonr-io/core/docs"
	"github.com/sonr-io/core/highway/client"
	"github.com/sonr-io/core/highway/config"
	hn "github.com/sonr-io/core/host"
	"github.com/sonr-io/core/host/ipfs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "go.buf.build/grpc/go/sonr-io/core/highway/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	v1.HighwayServer
	ctx    context.Context
	config *config.Config

	// Clients
	node     hn.HostImpl
	cosmos   *client.Cosmos
	webauthn *client.WebAuthn

	// Http Properties
	router     *gin.Engine
	httpServer *http.Server

	// Grpc Properties
	grpc         *grpc.Server
	grpcConn     *grpc.ClientConn
	grpcClient   v1.HighwayClient
	grpcListener net.Listener

	// Protocols
	channels     map[string]channel.Channel
	ipfsProtocol *ipfs.IPFSProtocol
}

// NewHighwayServer creates a new Highway service stub for the node.
func NewHighway(ctx context.Context, opts ...config.Option) (*HighwayServer, error) {
	// Create Config
	c := config.DefaultConfig()
	for _, opt := range opts {
		opt(c)
	}

	// Create the Highway Server
	stub, err := setupBaseStub(ctx, c)
	if err != nil {
		return nil, err
	}

	// Setup Protocols
	if err := setupProtocols(ctx, stub); err != nil {
		return nil, err
	}

	// Register RPC Service
	if err := setupAPI(ctx, stub); err != nil {
		return nil, err
	}
	return stub, nil
}

// Serve starts the RPC Service.
func (s *HighwayServer) Serve() {
	// Print the Server Address's
	logger.Infof("Serving RPC Server on %s", s.grpcListener.Addr().String())
	logger.Infof("Serving HTTP Server on %s", s.config.HighwayHTTPEndpoint)

	// Start the gRPC Server
	go func() {
		// Start gRPC server (and proxy calls to gRPC server endpoint)
		if err := s.grpc.Serve(s.grpcListener); err != nil {
			logger.Errorf("%s - Failed to start HTTP server", err)
		}
	}()

	// Start HTTP server on a separate goroutine
	go func() {
		// Start HTTP server (and proxy calls to gRPC server endpoint)
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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
	if err := s.httpServer.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown: ", err)
	}

	// Close the gRPC server
	logger.Warn("Shutting down gRPC server...")
	s.grpc.GracefulStop()
	logger.Info("Goodbye!")
}

// setupBaseStub creates the base Highway Server.
func setupBaseStub(ctx context.Context, c *config.Config) (*HighwayServer, error) {
	node, err := hn.NewHost(ctx, device.Role_HIGHWAY, c)
	if err != nil {
		return nil, err
	}

	// Get the Listener for the Host
	lst, err := net.Listen(c.HighwayGRPCNetwork, c.HighwayGRPCEndpoint)
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

	// Create the RPC Service
	stub := &HighwayServer{
		cosmos: cosmos,
		node:   node,
		ctx:    ctx,
		router: gin.Default(),
		grpc:   grpc.NewServer(),
		config: c,

		grpcListener: lst,
		webauthn:     webauthn,
	}
	return stub, nil
}

// setupProtocols creates the protocols for the node.
func setupProtocols(ctx context.Context, stub *HighwayServer) error {
	// Create the IPFS Protocol
	ipfs, err := ipfs.New(ctx, stub.node)
	if err != nil {
		return err
	}
	stub.ipfsProtocol = ipfs
	return nil
}

// setupAPI creates the gRPC Service for the RPC Service.
func setupAPI(ctx context.Context, s *HighwayServer) error {
	// Register the RPC Service and reflection
	v1.RegisterHighwayServer(s.grpc, s)
	reflection.Register(s.grpc)

	// Set up a connection to the server.
	conn, err := grpc.Dial(s.config.HighwayGRPCEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return err
	}

	// Create a new gRPC Client
	s.grpcConn = conn
	s.grpcClient = v1.NewHighwayClient(conn)

	// Register WebAuthn HTTP Routes
	s.router.GET("/v1/name/register/start/:username", s.StartRegisterName)
	s.router.POST("/v1/name/register/finish/:username", s.FinishRegisterName)
	s.router.GET("/v1/name/access/start/:username", s.StartAccessName)
	s.router.POST("/v1/name/access/finish/:username", s.FinishAccessName)

	// Register Cosmos HTTP Routes
	s.router.POST("/v1/bucket/create", s.CreateBucketHTTP)
	s.router.POST("/v1/bucket/update", s.UpdateBucketHTTP)
	s.router.POST("/v1/bucket/deactivate", s.DeactivateBucketHTTP)
	s.router.POST("/v1/channel/create", s.CreateChannelHTTP)
	s.router.POST("/v1/channel/update", s.UpdateChannelHTTP)
	s.router.POST("/v1/channel/deactivate", s.DeactivateChannelHTTP)
	s.router.GET("/v1/channel/listen", s.ListenChannelHTTP)
	s.router.POST("/v1/object/create", s.CreateObjectHTTP)
	s.router.POST("/v1/object/update", s.UpdateObjectHTTP)
	s.router.POST("/v1/object/deactivate", s.DeactivateObjectlHTTP)

	// Register IPFS HTTP Routes
	s.router.POST("/v1/blob/upload", s.UploadBlobHTTP)
	s.router.GET("/v1/blob/download/:cid", s.DownloadBlobHTTP)
	s.router.POST("/v1/blob/remove/:cid", s.RemoveBlobHTTP)

	// Setup Swagger UI
	docs.SwaggerInfo.BasePath = "/v1"
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Setup HTTP Server
	s.httpServer = &http.Server{
		Addr:    s.config.HighwayHTTPEndpoint,
		Handler: s.router,
	}
	return nil
}
