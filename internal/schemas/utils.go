package schemas

import (
	"github.com/ipld/go-ipld-prime/datamodel"
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (as *appSchemaInternalImpl) GetPath() (datamodel.ListIterator, error) {
	if as.nodes == nil {
		return nil, errNodeNotFound
	}
	return as.nodes.ListIterator(), nil
}

func (as *appSchemaInternalImpl) GetNode() (datamodel.Node, error) {
	if as.nodes == nil {
		return nil, errNodeNotFound
	}

	return as.nodes, nil
}

func (as *appSchemaInternalImpl) GetSchema() ([]*st.SchemaKindDefinition, error) {
	if as.fields == nil {
		return nil, errSchemaFieldsNotFound
	}
	return as.fields, nil
}
