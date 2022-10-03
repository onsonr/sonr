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
	store      *ReadStoreImpl
	next       *schemaImpl
}

/*
	Default initialization with a local client instance created in scope
*/
func New(store *ReadStoreImpl, whatIs *st.WhatIs) *schemaImpl {
	asi := &schemaImpl{
		fields:     whatIs.Schema.Fields,
		subSchemas: make(map[string]*st.SchemaDefinition),
		whatIs:     whatIs,
		nodes:      nil,
		store:      store,
	}

	asi.LoadSubSchemas(context.Background())
	return asi
}

/*
	Initialize with a instance of pkg/client
*/
func NewWithClient(client *client.Client, whatIs *st.WhatIs) *schemaImpl {
	asi := &schemaImpl{
		fields:     whatIs.Schema.Fields,
		subSchemas: make(map[string]*st.SchemaDefinition),
		whatIs:     whatIs,
		nodes:      nil,
		store: &ReadStoreImpl{
			Client: client,
		},
	}

	asi.LoadSubSchemas(context.Background())
	return asi
}
