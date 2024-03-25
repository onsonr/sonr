package types

import grpc "google.golang.org/grpc"

const (
	// ModuleName defines the name of the middleware.
	ModuleName = "service"

	// StoreKey is the store key string for the middleware.
	StoreKey = ModuleName

	// RouterKey is the message route for the middleware.
	RouterKey = ModuleName

	// QuerierRoute is the querier route for the middleware.
	QuerierRoute = ModuleName
)

// getQueryServiceClient is a helper function to get a QueryClient
func getQueryServiceClient(conn *grpc.ClientConn) QueryClient {
	return NewQueryClient(conn)
}
