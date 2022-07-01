package schemas

import (
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	dmt "github.com/ipld/go-ipld-prime/schema/dmt"
	dsl "github.com/ipld/go-ipld-prime/schema/dsl"
	st "github.com/sonr-io/sonr/x/schema/types"
)

/*
	Takes an DID of the schema and the definition to create a Basic Node definition
	maps the id of the schema to the top level node in `nodes`
*/
func (as *appSchemaInternalImpl) BuildNodesFromDefinition(id string, def *st.SchemaDefinition) (datamodel.Node, error) {
	// Create IPLD Node
	np := basicnode.Prototype.Any
	nb := np.NewBuilder() // Create a builder.
	ma, err := nb.BeginMap(int64(len(def.GetFields())))

	if err != nil {
		return nil, err
	}
	for k, t := range def.GetFields() {
		ma.AssembleKey().AssignString(k)
		switch t {
		case st.SchemaKind_STRING:
			ma.AssembleValue().AssignString("")
		case st.SchemaKind_INT:
			ma.AssembleValue().AssignInt(0)
		case st.SchemaKind_FLOAT:
			ma.AssembleValue().AssignFloat(0.0)
		case st.SchemaKind_BOOL:
			ma.AssembleValue().AssignBool(false)
		case st.SchemaKind_BYTES:
			ma.AssembleValue().AssignBytes([]byte{})
		case st.SchemaKind_LINK:
			ma.AssembleValue().AssignLink(nil)
		default:
			ma.AssembleValue().AssignNull()
		}
	}

	buildErr := ma.Finish()

	if buildErr != nil {
		return nil, buildErr
	}
	node := nb.Build()

	as.nodes[id] = node
	return node, nil
}

func (as *appSchemaInternalImpl) BuildSchemaFromNodes(id string) (*dmt.Schema, error) {
	if _, ok := as.nodes[id]; !ok {
		return nil, errIdNotFound
	}

	tln := as.nodes[id]

	bytes, err := tln.AsBytes()

	if err != nil {
		return nil, err
	}

	schema, err := dsl.ParseBytes(bytes)

	if err != nil {
		return nil, err
	}

	return schema, nil
}
