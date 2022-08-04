package object

import (
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sonr-io/sonr/pkg/schemas"
)

type Config struct {
	storeImpl *shell.Shell
	schema    schemas.AppSchemaInternal
}

func (c *Config) WithStorage(store *shell.Shell) {
	c.storeImpl = store
}

func (c *Config) WithSchemaImpl(schema schemas.AppSchemaInternal) {
	c.schema = schema
}

type ObjectConfiguration = func(config *Config)
