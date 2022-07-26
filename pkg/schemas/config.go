package schemas

import "github.com/sonr-io/sonr/pkg/client"

type SchemasConfig struct {
	persistenceURI       string
	clientConnectionType client.ConnEndpointType
}
