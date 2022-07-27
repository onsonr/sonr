package object

import (
	"errors"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sonr-io/sonr/pkg/schemas"
)

var (
	ErrObjectNotUploaded = errors.New("error while uploading object")
	ErrObjectEmpty       = errors.New("object cannot be empty")
)

type ObjectDefinition struct {
	Did   string
	Label string
	Cid   string
}

type ObjectUploadResult struct {
	Code       int32
	Definition *ObjectDefinition
	Message    string
}

type AppObjectInternalImpl struct {
	schemaInternal schemas.AppSchemaInternal
	shell          *shell.Shell
}

func New(schemaImpl schemas.AppSchemaInternal, storageEndpoint string) AppObjectInternal {
	return &AppObjectInternalImpl{
		schemaInternal: schemaImpl,
		shell:          shell.NewShell(storageEndpoint),
	}
}

func NewWithConfig(c *Config) AppObjectInternal {

	return &AppObjectInternalImpl{
		schemaInternal: c.schemaImpl,
		shell:          shell.NewShell(c.storageEndpoint),
	}
}
