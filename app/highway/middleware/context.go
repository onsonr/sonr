package middleware

import (
	"net/http"

	bankv1beta1 "cosmossdk.io/api/cosmos/bank/v1beta1"
	govv1 "cosmossdk.io/api/cosmos/gov/v1"
	stakingv1beta1 "cosmossdk.io/api/cosmos/staking/v1beta1"
	"github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/ipfs/kubo/client/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	identityv1 "github.com/sonrhq/sonr/api/sonr/identity/v1"
	servicev1 "github.com/sonrhq/sonr/api/sonr/service/v1"
)

func Context(next http.Handler) http.Handler {
	mw := ControllerMiddleware{
		Next:     next,
		Secure:   true,
		HTTPOnly: true,
	}
	return mw
}

type ContextMiddleware struct {
	Next     http.Handler
	Secure   bool
	HTTPOnly bool
}

func GrpcClientConn(r *http.Request) *grpc.ClientConn {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		"127.0.0.1:9090", // your gRPC server address.
		grpc.WithTransportCredentials(insecure.NewCredentials()), // The Cosmos SDK doesn't support any transport security mechanism.
		// This instantiates a general gRPC codec which handles proto bytes. We pass in a nil interface registry
		// if the request/response types contain interface instead of 'nil' you should pass the application specific codec.
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
	)
	if err != nil {
		panic(err)
	}
	return grpcConn
}

func IPFSClient(r *http.Request) *rpc.HttpApi {
	// The `IPFSClient()` function is a method of the `context` struct that returns an instance of the `rpc.HttpApi` type.
	ipfsC, err := rpc.NewLocalApi()
	if err != nil {
		panic(err)
	}
	return ipfsC
}

func (mw ContextMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mw.Next.ServeHTTP(w, r)
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                      Helper GRPC Client Wrapper Functions                      ||
// ! ||--------------------------------------------------------------------------------||

type GRPCConn = *grpc.ClientConn

type CometClient = cmtservice.ServiceClient

type BankClient = bankv1beta1.QueryClient

type GovClient = govv1.QueryClient

type IdentityClient = identityv1.QueryClient

type ServiceClient = servicev1.QueryClient

type StakingClient = stakingv1beta1.QueryClient

func NewBankClient(conn GRPCConn) BankClient {
	return bankv1beta1.NewQueryClient(conn)
}

func NewGovClient(conn GRPCConn) GovClient {
	return govv1.NewQueryClient(conn)
}

func NewIdentityClient(conn GRPCConn) IdentityClient {
	return identityv1.NewQueryClient(conn)
}

func NewServiceClient(conn GRPCConn) ServiceClient {
	return servicev1.NewQueryClient(conn)
}

func NewStakingClient(conn GRPCConn) StakingClient {
	return stakingv1beta1.NewQueryClient(conn)
}

func NewCometClient(conn GRPCConn) CometClient {
	return cmtservice.NewServiceClient(conn)
}
