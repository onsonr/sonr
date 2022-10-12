package document

import (
	"errors"

	shell "github.com/ipfs/go-ipfs-api"
)

var (
	ErrDocumentNotUploaded = errors.New("error while uploading document")
	ErrDocumentEmpty       = errors.New("document cannot be empty")
)

type documentImpl struct {
	shell  *shell.Shell
	schema Schema
}

func New(schemaImpl Schema, shell *shell.Shell) *documentImpl {
	return &documentImpl{
		// TODO: replace with store interface that Daniel made
		shell:  shell,
		schema: schemaImpl,
	}
}

func NewWithConfig(c *Config) *documentImpl {
	return &documentImpl{
		shell:  c.storeImpl,
		schema: c.schema,
	}
}
