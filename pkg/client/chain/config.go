package chain

type APIEndpoint string

const (
	// List of known origin api endpoints.
	SonrLocalRpcOrigin  APIEndpoint = "localhost:9090"
	SonrPublicRpcOrigin APIEndpoint = "159.65.236.177:9090"
)
