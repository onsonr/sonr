package schemas

import (
	"context"
	"errors"

	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/sonr-io/sonr/pkg/client"
	st "github.com/sonr-io/sonr/x/schema/types"
)

var (
	errSchemaFieldsInvalid  = errors.New("supplied Schema is invalid")
	errSchemaFieldsNotFound = errors.New("no Schema Fields found")
	errNodeNotFound         = errors.New("no object definition built from schema")
)

type Encoding int

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
	Default initialization with a local client instance created in scope
*/
func New(fields []*st.SchemaKindDefinition, whatIs *st.WhatIs) *schemaImpl {
	asi := &schemaImpl{
		fields:     fields,
		subSchemas: make(map[string]*st.SchemaDefinition),
		whatIs:     whatIs,
		nodes:      nil,
		store: &readStoreImpl{
			client: client.NewClient(client.ConnEndpointType_LOCAL),
		},
	}

	asi.loadSubSchemas(context.Background(), fields)
	return asi
}

/*
	Initialize with a instance of pkg/client
*/
func NewWithShell(client *client.Client, fields []*st.SchemaKindDefinition, whatIs *st.WhatIs) *schemaImpl {
	asi := &schemaImpl{
		fields:     fields,
		subSchemas: make(map[string]*st.SchemaDefinition),
		whatIs:     whatIs,
		nodes:      nil,
		store: &readStoreImpl{
			client: client,
		},
	}

	asi.loadSubSchemas(context.Background(), fields)
	return asi
}
