package schemas

import (
	"errors"

	"github.com/ipfs/go-cid"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/linking"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/ipld/go-ipld-prime/storage/memstore"
	st "github.com/sonr-io/sonr/x/schema/types"
)

var (
	errSchemaFieldsInvalid  = errors.New("supplied Schema is invalid")
	errSchemaFieldsNotFound = errors.New("no Schema Fields found")
	errNodeNotFound         = errors.New("no object definition built from schema")
)

type Encoding int

type Event struct {
	name     string
	previous cid.Cid
}

const (
	EncType_DAG_CBOR Encoding = iota
	EncType_DAG_JSON
)

type schemaImpl struct {
	fields    []*st.SchemaKindDefinition
	whatIs    *st.WhatIs
	nodes     datamodel.Node
	linkProto cidlink.LinkPrototype
	linkSys   linking.LinkSystem
	store     *memstore.Store
	next      *schemaImpl
}

func New(fields []*st.SchemaKindDefinition, whatIs *st.WhatIs) *schemaImpl {
	asi := &schemaImpl{
		fields: fields,
		whatIs: whatIs,
		nodes:  nil,
		// TODO: replace this with the interface Daniel made
		store:   &memstore.Store{},
		linkSys: cidlink.DefaultLinkSystem(),
	}

	asi.linkSys.SetWriteStorage(asi.store)
	asi.linkSys.SetReadStorage(readStoreImpl{
		shell: shell.NewLocalShell(),
	})

	asi.linkProto = asi.CreateLinkPrototype()

	return asi
}
