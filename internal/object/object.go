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

/*
	Object definition to be returned after object creation
*/
type ObjectUploadResult struct {
	Code       int32
	Definition *ObjectDefinition
	Message    string
}

type AppObjectInternalImpl struct {
	shell  *shell.Shell
	schema schemas.AppSchemaInternal
}

func New(schemaImpl schemas.AppSchemaInternal, shell *shell.Shell) AppObjectInternal {
	return &AppObjectInternalImpl{
		shell:  shell,
		schema: schemaImpl,
	}
}

func NewWithConfig(c *Config) AppObjectInternal {

	return &AppObjectInternalImpl{
		shell:  c.storeImpl,
		schema: c.schema,
	}
}
