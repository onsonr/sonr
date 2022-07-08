package schemas

import (
	"github.com/ipld/go-ipld-prime/datamodel"
)

func (as *appSchemaInternalImpl) GetPath(node datamodel.Node) (datamodel.ListIterator, error) {
	return node.ListIterator(), nil
}
