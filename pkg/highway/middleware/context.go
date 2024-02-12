package middleware

import (
	"fmt"
	"net/http"

	bankv1beta1 "cosmossdk.io/api/cosmos/bank/v1beta1"
	govv1 "cosmossdk.io/api/cosmos/gov/v1"
	stakingv1beta1 "cosmossdk.io/api/cosmos/staking/v1beta1"
	cmtcservice "github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/ipfs/kubo/client/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	identityv1 "github.com/sonrhq/sonr/api/sonr/identity/v1"
	servicev1 "github.com/sonrhq/sonr/api/sonr/service/v1"
	"github.com/sonrhq/sonr/pkg/config"
)

// GRPCConn is a gRPC client connection.
type GRPCConn = *grpc.ClientConn

// Context returns an http.Handler middleware that sets default middleware options.
func Context(next http.Handler) http.Handler {
	mw := ContextMiddleware{
		Next:     next,
		Secure:   true,
		HTTPOnly: true,
	}
	return mw
}

// ContextMiddleware defines a middleware with context helpers like secure, HTTPOnly flags.
type ContextMiddleware struct {
	Next     http.Handler
	Secure   bool
	HTTPOnly bool
}

// ServeHTTP calls the next middleware handler.
func (mw ContextMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := nodeGrpcAddress(r)
	if id == "" {
		c, err := config.LoadConfig()
		if err != nil {
			return
		}
		http.SetCookie(w, &http.Cookie{Name: "ipfsGateway", Value: c.Highway.IPFSGateway, Secure: mw.Secure, HttpOnly: mw.HTTPOnly})
		http.SetCookie(w, &http.Cookie{Name: "matrixConnection", Value: c.Highway.MatrixConnection, Secure: mw.Secure, HttpOnly: mw.HTTPOnly})
		http.SetCookie(w, &http.Cookie{Name: "nodeGrpcAddress", Value: c.Highway.NodeGRPCAddress, Secure: mw.Secure, HttpOnly: mw.HTTPOnly})
	}
	mw.Next.ServeHTTP(w, r)
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                            Grpc Client Module Proxy                            ||
// ! ||--------------------------------------------------------------------------------||

// IPFSClient creates an IPFS HTTP client.
func IPFSClient(r *http.Request, w http.ResponseWriter) *rpc.HttpApi {
	// The `IPFSClient()` function is a method of the `context` struct that returns an instance of the `rpc.HttpApi` type.
	ipfsC, err := rpc.NewLocalApi()
	if err != nil {
		InternalServerError(w, err)
		return nil
	}
	return ipfsC
}

// BankClient returns a new bank client.
func BankClient(r *http.Request, w http.ResponseWriter) bankv1beta1.QueryClient {
	if cc := grpcClientConn(r, w); cc != nil {
		return bankv1beta1.NewQueryClient(cc)
	}
	return nil
}

// CometClient returns a new comet client.
func CometClient(r *http.Request, w http.ResponseWriter) cmtcservice.ServiceClient {
	if cc := grpcClientConn(r, w); cc != nil {
		return cmtcservice.NewServiceClient(cc)
	}
	return nil
}

// GovClient creates a new gov client.
func GovClient(r *http.Request, w http.ResponseWriter) govv1.QueryClient {
	if cc := grpcClientConn(r, w); cc != nil {
		return govv1.NewQueryClient(cc)
	}
	return nil
}

// IdentityClient creates a new identity client.
func IdentityClient(r *http.Request, w http.ResponseWriter) identityv1.QueryClient {
	if cc := grpcClientConn(r, w); cc != nil {
		return identityv1.NewQueryClient(cc)
	}
	return nil
}

// ServiceClient creates a new service client.
func ServiceClient(r *http.Request, w http.ResponseWriter) servicev1.QueryClient {
	if cc := grpcClientConn(r, w); cc != nil {
		return servicev1.NewQueryClient(cc)
	}
	return nil
}

// StakingClient creates a new staking client.
func StakingClient(r *http.Request, w http.ResponseWriter) stakingv1beta1.QueryClient {
	if cc := grpcClientConn(r, w); cc != nil {
		return stakingv1beta1.NewQueryClient(cc)
	}
	return nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                      Helper GRPC Client Wrapper Functions                      ||
// ! ||--------------------------------------------------------------------------------||

// grpcClientConn creates a gRPC client connection.
func grpcClientConn(r *http.Request, w http.ResponseWriter) *grpc.ClientConn {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		"127.0.0.1:9090", // your gRPC server address.
		grpc.WithTransportCredentials(insecure.NewCredentials()), // The Cosmos SDK doesn't support any transport security mechanism.
		// This instantiates a general gRPC codec which handles proto bytes. We pass in a nil interface registry
		// if the request/response types contain interface instead of 'nil' you should pass the application specific codec.
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
	)
	if err != nil {
		InternalServerError(w, fmt.Errorf("grpc client connection failed"))
		return nil
	}
	return grpcConn
}

// ipfsGateway returns the address of the user from the session cookie.
func ipfsGateway(r *http.Request) (nodeGrpcAddress string) {
	cookie, err := r.Cookie("ipfsGateway")
	if err != nil {
		return
	}
	return cookie.Value
}

// matrixConnection returns the address of the user from the session cookie.
func matrixConnection(r *http.Request) (nodeGrpcAddress string) {
	cookie, err := r.Cookie("matrixConnection")
	if err != nil {
		return
	}
	return cookie.Value
}

// nodeGrpcAddress returns the address of the user from the session cookie.
func nodeGrpcAddress(r *http.Request) (nodeGrpcAddress string) {
	cookie, err := r.Cookie("nodeGrpcAddress")
	if err != nil {
		return
	}
	return cookie.Value
}
