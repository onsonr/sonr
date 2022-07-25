package object

import (
	"github.com/sonr-io/sonr/pkg/schemas"
	"google.golang.org/grpc"
)

type Config struct {
	schemaImpl      schemas.AppSchemaInternal
	rpcClient       *grpc.ClientConn
	storageEndpoint string
}

func (c *Config) WithSchemaImplementation(schema schemas.AppSchemaInternal) {
	c.schemaImpl = schema
}

func (c *Config) WithRpcClient(client *grpc.ClientConn) {
	c.rpcClient = client
}

func (c *Config) WithStorageEndpoint(uri string) {
	c.storageEndpoint = uri
}

type ObjectConfiguration = func(config *Config)
