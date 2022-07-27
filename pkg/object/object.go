package object

import (
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sonr-io/sonr/pkg/schemas"
	"google.golang.org/grpc"
)

// Defines and object and relates it to a given Schema Definition
// Does not yet support encryption keys
type ObjectEncoding int

var (
	cbor ObjectEncoding = 1
	json ObjectEncoding = 2
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
	rpcClient      *grpc.ClientConn
	shell          *shell.Shell
}

func New(c *Config) AppObjectInternal {

	return &AppObjectInternalImpl{
		schemaInternal: c.schemaImpl,
		rpcClient:      c.rpcClient,
		shell:          shell.NewShell(c.storageEndpoint),
	}
}
