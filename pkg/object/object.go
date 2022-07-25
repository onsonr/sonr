package object

import (
	"github.com/sonr-io/sonr/pkg/schemas"
	st "github.com/sonr-io/sonr/x/schema/types"
	"google.golang.org/grpc"
)

// Defines and object and relates it to a given Schema Definition
// Does not yet support encryption keys
type ObjectDefinition struct {
	Did       string
	Cid       string
	Creator   string
	SchemaCid string
}

type ObjectUploadResult struct {
	Code    int32
	Cid     string
	Message string
}

type AppObjectInternal interface {
	UploadObject(schemaDefinition *st.SchemaDefinition) error
	GetObject(cid string) ([]byte, error)
}

type AppObjectSchemaInternal struct {
	schemaInternal schemas.AppSchemaInternal
	rpcClient      *grpc.ClientConn
}

func New(config ObjectConfiguration) *AppObjectInternal {
	c := Config{}
	config(&c)

	return AppObjectSchemaInternal{
		schemaInternal: c.schemaImpl,
	}
}
