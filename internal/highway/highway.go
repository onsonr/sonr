package highway

import (
	"context"
	"log"
	"net/http"

	"github.com/sonr-io/sonr/internal/highway/x/core"
	_ "github.com/sonr-io/sonr/internal/highway/x/core"
	"github.com/sonr-io/sonr/pkg/config"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "go.buf.build/grpc/go/sonr-io/highway/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// NewHighwayServer creates a new Highway service stub for the node.
func NewHighway(ctx context.Context, opts ...config.Option) (*core.HighwayServer, error) {
	// Create Config
	c := config.DefaultConfig(config.Role_HIGHWAY)
	for _, opt := range opts {
		opt(c)
	}

	// Create the Highway Server
	stub, err := core.CreateStub(ctx, c)
	if err != nil {
		return nil, err
	}

	// Register RPC Service
	if err := setupAPI(ctx, stub); err != nil {
		return nil, err
	}
	return stub, nil
}

// setupAPI creates the gRPC Service for the RPC Service.
func setupAPI(ctx context.Context, s *core.HighwayServer) error {
	// Register the RPC Service and reflection
	v1.RegisterHighwayServer(s.GRPCServer, s)
	reflection.Register(s.GRPCServer)

	// Set up a connection to the server.
	conn, err := grpc.Dial(s.Config.HighwayGRPCEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return err
	}

	// Create a new gRPC Client
	s.GRPCConn = conn
	s.GRPCClient = v1.NewHighwayClient(conn)

	// Register WebAuthn HTTP Routes
	s.Router.GET("/v1/name/register/start/:username", s.StartRegisterName)
	s.Router.POST("/v1/name/register/finish/:username", s.FinishRegisterName)
	s.Router.GET("/v1/name/access/start/:username", s.StartAccessName)
	s.Router.POST("/v1/name/access/finish/:username", s.FinishAccessName)

	// Register Cosmos HTTP Routes
	s.Router.POST("/v1/bucket/create", s.CreateBucketHTTP)
	s.Router.POST("/v1/bucket/update", s.UpdateBucketHTTP)
	s.Router.POST("/v1/bucket/deactivate", s.DeactivateBucketHTTP)
	s.Router.POST("/v1/channel/create", s.CreateChannelHTTP)
	s.Router.POST("/v1/channel/update", s.UpdateChannelHTTP)
	s.Router.POST("/v1/channel/deactivate", s.DeactivateChannelHTTP)
	s.Router.GET("/v1/channel/listen", s.ListenChannelHTTP)
	s.Router.POST("/v1/object/create", s.CreateObjectHTTP)
	s.Router.POST("/v1/object/update", s.UpdateObjectHTTP)
	s.Router.POST("/v1/object/deactivate", s.DeactivateObjectlHTTP)

	// Register IPFS HTTP Routes
	s.Router.POST("/v1/blob/upload", s.UploadBlobHTTP)
	s.Router.GET("/v1/blob/download/:cid", s.DownloadBlobHTTP)
	s.Router.POST("/v1/blob/remove/:cid", s.RemoveBlobHTTP)

	// Setup Swagger UI
	s.Router.GET("v1/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Setup HTTP Server
	s.HTTPServer = &http.Server{
		Addr:    s.Config.HighwayHTTPEndpoint,
		Handler: s.Router,
	}
	return nil
}
