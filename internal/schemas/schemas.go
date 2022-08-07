package schemas

import (
	"context"
	"errors"

	"github.com/ipfs/go-cid"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/ipld/go-ipld-prime/datamodel"
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
	fields     []*st.SchemaKindDefinition
	subSchemas map[string]*st.SchemaDefinition
	whatIs     *st.WhatIs
	nodes      datamodel.Node
	store      *readStoreImpl
	next       *schemaImpl
}

/*
	Default initialization with a local shell for persistence
*/
func New(fields []*st.SchemaKindDefinition, whatIs *st.WhatIs) Schema {
	asi := &schemaImpl{
		fields:     fields,
		subSchemas: make(map[string]*st.SchemaDefinition),
		whatIs:     whatIs,
		nodes:      nil,
		store: &readStoreImpl{
			shell: shell.NewLocalShell(),
		},
	}

	asi.loadSubSchemas(context.Background(), fields)
	return asi
}

/*
	Initialize with a ipfs shell instance
*/
func NewWithShell(shell *shell.Shell, fields []*st.SchemaKindDefinition, whatIs *st.WhatIs) Schema {
	asi := &schemaImpl{
		fields:     fields,
		subSchemas: make(map[string]*st.SchemaDefinition),
		whatIs:     whatIs,
		nodes:      nil,
		store: &readStoreImpl{
			shell: shell,
		},
	}

	asi.loadSubSchemas(context.Background(), fields)
	return asi
}
