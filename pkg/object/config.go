package object

import (
	"github.com/sonr-io/sonr/pkg/schemas"
)

type Config struct {
	schemaImpl      schemas.AppSchemaInternal
	storageEndpoint string
}

func (c *Config) WithSchemaImplementation(schema schemas.AppSchemaInternal) {
	c.schemaImpl = schema
}

func (c *Config) WithStorageEndpoint(uri string) {
	c.storageEndpoint = uri
}

type ObjectConfiguration = func(config *Config)
