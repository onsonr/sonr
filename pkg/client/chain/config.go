package chain

type APIEndpoint string

const (
	// List of known origin api endpoints.
	SonrLocalRpcOrigin  APIEndpoint = "localhost:9090"
	SonrPublicRpcOrigin APIEndpoint = "142.93.116.204:9090"
)
