package schemas

import (
	"github.com/ipld/go-ipld-prime/datamodel"
)

func (as *appSchemaInternalImpl) GetTopLevelNodeById(id string) (datamodel.Node, error) {
	if _, ok := as.nodes[id]; ok {
		return as.nodes[id], nil
	}

	return nil, errIdNotFound
}

func (as *appSchemaInternalImpl) GetPath(id string) (datamodel.ListIterator, error) {
	if _, ok := as.nodes[id]; !ok {
		return nil, errIdNotFound
	}
	node := as.nodes[id]
	return node.ListIterator(), nil
}
