package schemas

import (
	"context"
	"errors"

	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/sonr-io/sonr/pkg/client"
	st "github.com/sonr-io/sonr/x/schema/types"
)

var (
	errSchemaNotFound       = errors.New("schema not found")
	errSchemaFieldsInvalid  = errors.New("supplied Schema is invalid")
	errSchemaFieldsNotFound = errors.New("no Schema Fields found")
	errNodeNotFound         = errors.New("no object definition built from schema")
)

type Encoding int

const (
	EncType_DAG_CBOR Encoding = iota
	EncType_DAG_JSON
)

type SchemaImpl struct {
	fields    []*st.SchemaField
	subWhatIs map[string]*st.WhatIs
	whatIs    *st.WhatIs
	nodes     datamodel.Node
	store     *ReadStoreImpl
	next      *SchemaImpl
}

/*
	Default initialization with a local client instance created in scope
*/
func New(store *ReadStoreImpl, whatIs *st.WhatIs) *SchemaImpl {
	asi := &SchemaImpl{
		fields:    whatIs.Schema.Fields,
		subWhatIs: make(map[string]*st.WhatIs),
		whatIs:    whatIs,
		nodes:     nil,
		store:     store,
	}

	asi.LoadSubSchemas(context.Background())
	return asi
}

/*
	Initialize with a instance of pkg/client
*/
func NewWithClient(client *client.Client, whatIs *st.WhatIs) *SchemaImpl {
	asi := &SchemaImpl{
		fields:    whatIs.Schema.Fields,
		subWhatIs: make(map[string]*st.WhatIs),
		whatIs:    whatIs,
		nodes:     nil,
		store: &ReadStoreImpl{
			Client: client,
		},
	}

	asi.LoadSubSchemas(context.Background())
	return asi
}
