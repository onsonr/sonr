package schemas

import (
	"github.com/ipld/go-ipld-prime/datamodel"
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (as *SchemaImpl) GetPath() (datamodel.ListIterator, error) {
	if as.nodes == nil {
		return nil, errNodeNotFound
	}
	return as.nodes.ListIterator(), nil
}

func (as *SchemaImpl) GetNode() (datamodel.Node, error) {
	if as.nodes == nil {
		return nil, errNodeNotFound
	}

	return as.nodes, nil
}

func (as *SchemaImpl) GetSchema() (*st.Schema, error) {
	if as.fields == nil {
		return nil, errSchemaNotFound
	}
	return as.whatIs.Schema, nil
}

func (as *SchemaImpl) GetSubSchema(did string) (*SchemaImpl, error) {
	subWhatIs, ok := as.subWhatIs[did]
	if !ok {
		return nil, errSchemaNotFound
	}

	return NewWithClient(as.store.Client, subWhatIs), nil
}

func arrayContains(arr []string, val interface{}) bool {

	for _, v := range arr {
		if v == val {
			return true
		}
	}

	return false
}
