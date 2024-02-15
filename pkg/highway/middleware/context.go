package middleware

import (
	bankv1beta1 "cosmossdk.io/api/cosmos/bank/v1beta1"
	govv1 "cosmossdk.io/api/cosmos/gov/v1"
	stakingv1beta1 "cosmossdk.io/api/cosmos/staking/v1beta1"
	cmtcservice "github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/gorilla/sessions"
	"github.com/ipfs/kubo/client/rpc"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	identityv1 "github.com/sonrhq/sonr/api/sonr/identity/v1"
	servicev1 "github.com/sonrhq/sonr/api/sonr/service/v1"
)

// GRPCConn is a gRPC client connection.
type GRPCConn = *grpc.ClientConn

// ContextMiddleware defines a middleware with context helpers like secure, HTTPOnly flags.
type ContextMiddleware struct {
	echo.Context
	Secure   bool
	HTTPOnly bool
}

// Context returns an http.Handler middleware that sets default middleware options.
func Context(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &ContextMiddleware{
			Context:  c,
			Secure:   true,
			HTTPOnly: true,
		}
		return next(cc)
	}
}

// UseDefaults adds chi provided middleware libraries to the router.
func UseDefaults(e *echo.Echo) {
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.Decompress())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(HTMX)
}

// IPFSClient creates an IPFS HTTP client.
func IPFSClient(e echo.Context) *rpc.HttpApi {
	// The `IPFSClient()` function is a method of the `context` struct that returns an instance of the `rpc.HttpApi` type.
	ipfsC, err := rpc.NewLocalApi()
	if err != nil {
		return nil
	}
	return ipfsC
}

// BankClient returns a new bank client.
func BankClient(e echo.Context) bankv1beta1.QueryClient {
	if cc := GrpcClientConn(e); cc != nil {
		return bankv1beta1.NewQueryClient(cc)
	}
	return nil
}

// CometClient returns a new comet client.
func CometClient(e echo.Context) cmtcservice.ServiceClient {
	if cc := GrpcClientConn(e); cc != nil {
		return cmtcservice.NewServiceClient(cc)
	}
	return nil
}

// GovClient creates a new gov client.
func GovClient(e echo.Context) govv1.QueryClient {
	if cc := GrpcClientConn(e); cc != nil {
		return govv1.NewQueryClient(cc)
	}
	return nil
}

// IdentityClient creates a new identity client.
func IdentityClient(e echo.Context) identityv1.QueryClient {
	if cc := GrpcClientConn(e); cc != nil {
		return identityv1.NewQueryClient(cc)
	}
	return nil
}

// ServiceClient creates a new service client.
func ServiceClient(e echo.Context) servicev1.QueryClient {
	if cc := GrpcClientConn(e); cc != nil {
		return servicev1.NewQueryClient(cc)
	}
	return nil
}

// StakingClient creates a new staking client.
func StakingClient(e echo.Context) stakingv1beta1.QueryClient {
	if cc := GrpcClientConn(e); cc != nil {
		return stakingv1beta1.NewQueryClient(cc)
	}
	return nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                      Helper GRPC Client Wrapper Functions                      ||
// ! ||--------------------------------------------------------------------------------||

// GrpcClientConn creates a gRPC client connection.
func GrpcClientConn(e echo.Context) *grpc.ClientConn {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		"sonrd:9090", // your gRPC server address.
		grpc.WithTransportCredentials(insecure.NewCredentials()), // The Cosmos SDK doesn't support any transport security mechanism.
		// This instantiates a general gRPC codec which handles proto bytes. We pass in a nil interface registry
		// if the request/response types contain interface instead of 'nil' you should pass the application specific codec.
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
	)
	if err != nil {
		return nil
	}
	return grpcConn
}

// ipfsGateway returns the address of the user from the session cookie.
func ipfsGateway(e echo.Context) (nodeGrpcAddress string) {
	cookie, err := e.Cookie("ipfsGateway")
	if err != nil {
		return
	}
	return cookie.Value
}

// matrixConnection returns the address of the user from the session cookie.
func matrixConnection(e echo.Context) (nodeGrpcAddress string) {
	cookie, err := e.Cookie("matrixConnection")
	if err != nil {
		return
	}
	return cookie.Value
}

// nodeGrpcAddress returns the address of the user from the session cookie.
func nodeGrpcAddress(e echo.Context) (nodeGrpcAddress string) {
	cookie, err := e.Cookie("nodeGrpcAddress")
	if err != nil {
		return
	}
	return cookie.Value
}
