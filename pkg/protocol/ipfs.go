package ipfs

import (
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/sonr-io/sonr/x/schema/types"
)

type Protocol interface {
	GetData(cid string) ([]byte, error)
	GetObjectSchema(cid *cid.Cid) (datamodel.Node, error)
	PutData(data []byte) (*cid.Cid, error)
	PutObjectSchema(doc *types.SchemaDefinition) (*cid.Cid, error)
	RemoveFile(cidstr string) error
}
