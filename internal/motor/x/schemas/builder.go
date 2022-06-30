package schemas

import (
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	st "github.com/sonr-io/sonr/x/schema/types"
)

/*
	Takes an DID of the schema and the definition to create a Basic Node definition
	maps the id of the schema to the top level node in `nodes`
*/
func (as *appSchemaInternalImpl) BuildNodesFromDefinition(id string, def *st.SchemaDefinition) (*datamodel.Node, error) {
	// Create IPLD Node
	np := basicnode.Prototype.Any
	nb := np.NewBuilder() // Create a builder.
	ma, err := nb.BeginMap(int64(len(def.GetFields())))

	if err != nil {
		return nil, err
	}
	for k, _ := range def.GetFields() {
		ma.AssembleKey().AssignString(k)
	}

	buildErr := ma.Finish()

	if buildErr != nil {
		return nil, buildErr
	}
	node := nb.Build()

	as.nodes[id] = &node

	return &node, nil
}
