package document

import (
	shell "github.com/ipfs/go-ipfs-api"
)

type Config struct {
	storeImpl *shell.Shell
	schema    Schema
}

func (c *Config) WithStorage(store *shell.Shell) {
	c.storeImpl = store
}

func (c *Config) WithSchemaImpl(schema Schema) {
	c.schema = schema
}

type DocumentConfiguration = func(config *Config)
