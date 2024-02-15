package service

import (
	"cosmossdk.io/collections"
	grpc "google.golang.org/grpc"

	modulev1 "github.com/sonrhq/sonr/api/sonr/service/module/v1"
	servicev1 "github.com/sonrhq/sonr/api/sonr/service/v1"
)

// ModuleName is the name of the module
const ModuleName = "service"

var (
	// ParamsKey is the key for the module parameters
	ParamsKey = collections.NewPrefix(0)

	// CounterKey is the key for the module counter
	CounterKey = collections.NewPrefix(1)
)

// getQueryServiceClient is a helper function to get a QueryClient
func getQueryServiceClient(conn *grpc.ClientConn) servicev1.QueryClient {
	return servicev1.NewQueryClient(conn)
}

// getStateServiceClient is a helper function to get a StateClient
func getStateServiceClient(conn *grpc.ClientConn) modulev1.StateQueryServiceClient {
	return modulev1.NewStateQueryServiceClient(conn)
}
