package object

import (
	shell "github.com/ipfs/go-ipfs-api"
)

type Config struct {
	storeImpl *shell.Shell
}

func (c *Config) WithStorage(store *shell.Shell) {
	c.storeImpl = store
}

type ObjectConfiguration = func(config *Config)
