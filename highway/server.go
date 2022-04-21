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
	"github.com/sonr-io/core/highway/client"
	"github.com/sonr-io/core/highway/config"
	hn "github.com/sonr-io/core/host"
	"github.com/sonr-io/core/host/ipfs"
	v1 "go.buf.build/sonr-io/grpc-gateway/sonr-io/core/highway/v1"
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
	v1.HighwayServer

	// Clients
	node     hn.HostImpl
	cosmos   *client.Cosmos
	webauthn *client.WebAuthn

	// Config
	ctx      context.Context
	config   *config.Config
	listener net.Listener
	gin      *gin.Engine
	httpSrv  *http.Server

	// Grpc Server
	grpc       *grpc.Server
	grpcConn   *grpc.ClientConn
	grpcClient v1.HighwayClient

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
	// Start the gRPC Server
	go func() {
		// Print the gRPC Server Address
		logger.Infof("Serving RPC Server on %s", s.listener.Addr().String())

		// Start gRPC server (and proxy calls to gRPC server endpoint)
		if err := s.grpc.Serve(s.listener); err != nil {
			logger.Errorf("%s - Failed to start HTTP server", err)
		}
	}()

	// Start HTTP server on a separate goroutine
	go func() {
		// Print the gRPC Server Address
		logger.Infof("Serving HTTP Server on %s", s.config.HighwayHTTPEndpoint)

		// Start HTTP server (and proxy calls to gRPC server endpoint)
		if err := s.httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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
	if err := s.httpSrv.Shutdown(ctx); err != nil {
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
	lst, err := node.Listener()
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
		gin:    gin.Default(),
		grpc:   grpc.NewServer(),

		listener: lst,
		webauthn: webauthn,
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
	s.gin.GET("/register/name/start/{username}", s.StartRegisterName)
	s.gin.POST("/register/name/finish/{username}", s.FinishRegisterName)
	s.gin.GET("/access/name/start/{username}", s.StartAccessName)
	s.gin.POST("/access/name/finish/{username}", s.FinishAccessName)

	// Register Cosmos HTTP Routes
	s.gin.POST("/create/bucket", s.CreateBucketHTTP)
	s.gin.POST("/create/channel", s.CreateChannelHTTP)
	s.gin.POST("/create/object", s.CreateObjectHTTP)
	s.gin.POST("/update/bucket", s.UpdateBucketHTTP)
	s.gin.POST("/update/channel", s.UpdateChannelHTTP)
	s.gin.POST("/update/object", s.UpdateObjectHTTP)

	// Setup HTTP Server
	s.httpSrv = &http.Server{
		Addr:    s.config.HighwayHTTPEndpoint,
		Handler: s.gin,
	}
	return nil
}
