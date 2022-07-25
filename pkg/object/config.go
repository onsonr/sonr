package object

import "github.com/sonr-io/sonr/pkg/schemas"

type Config struct {
	schemaImpl schemas.AppSchemaInternal
}

func (c *Config) WithSchemaImplementation(schema schemas.AppSchemaInternal) {
	c.schemaImpl = schema
}

type ObjectConfiguration = func(config *Config)
