package service

import (
	"cosmossdk.io/collections"
	grpc "google.golang.org/grpc"

	servicev1 "github.com/sonrhq/sonr/api/sonr/service/v1"
)

// ModuleName is the name of the module
const ModuleName = "service"

var (
	// ParamsKey is the key for the module parameters
	ParamsKey = collections.NewPrefix(0)

	// ServiceKey is the key for the module counter
	ServiceKey = collections.NewPrefix(1)
)

// getQueryServiceClient is a helper function to get a QueryClient
func getQueryServiceClient(conn *grpc.ClientConn) servicev1.QueryClient {
	return servicev1.NewQueryClient(conn)
}
