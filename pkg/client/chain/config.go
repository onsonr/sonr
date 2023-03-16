package chain

import "fmt"

type APIEndpoint string

const (
	// List of known origin api endpoints.
	SonrLocalRpcOrigin  APIEndpoint = "localhost:9090"
	SonrPublicRpcOrigin APIEndpoint = "142.93.116.204:9090"
)

func (e APIEndpoint) TcpAddress() string {
	return fmt.Sprintf("tcp://%s", string(e))
}
