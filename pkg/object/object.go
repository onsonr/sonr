package object

import (
	shell "github.com/ipfs/go-ipfs-api"
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
	Code       int32
	Definition *ObjectDefinition
	Message    string
}

type AppObjectInternal interface {
	UploadObject(fields []*st.SchemaKindDefinition, object map[string]interface{}) (*ObjectUploadResult, error)
	GetObject(cid string) ([]byte, error)
}

type AppObjectInternalImpl struct {
	schemaInternal schemas.AppSchemaInternal
	rpcClient      *grpc.ClientConn
	shell          *shell.Shell
}

func New(config ObjectConfiguration) AppObjectInternal {
	c := Config{}
	config(&c)

	return &AppObjectInternalImpl{
		schemaInternal: c.schemaImpl,
		rpcClient:      c.rpcClient,
		shell:          shell.NewShell(c.storageEndpoint),
	}
}
