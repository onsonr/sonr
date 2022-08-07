package object

import (
	"errors"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sonr-io/sonr/internal/schemas"
)

var (
	ErrObjectNotUploaded = errors.New("error while uploading object")
	ErrObjectEmpty       = errors.New("object cannot be empty")
)

type ObjectReference struct {
	Did   string
	Label string
	Cid   string
}

/*
	Object definition to be returned after object creation
*/
type ObjectUploadResult struct {
	Code      int32
	Reference *ObjectReference
	Message   string
}

type objectImpl struct {
	shell  *shell.Shell
	schema schemas.Schema
}

func New(schemaImpl schemas.Schema, shell *shell.Shell) *objectImpl {
	return &objectImpl{
		// TODO: replace with store interface that Daniel made
		shell:  shell,
		schema: schemaImpl,
	}
}

func NewWithConfig(c *Config) *objectImpl {
	return &objectImpl{
		shell:  c.storeImpl,
		schema: c.schema,
	}
}
