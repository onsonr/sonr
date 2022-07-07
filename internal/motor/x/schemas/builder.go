package schemas

import (
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (as *appSchemaInternalImpl) BuildNodesFromDefinition(
	id string,
	def *st.SchemaDefinition,
	object map[string]interface{}) (datamodel.Node, error) {

	err := as.VerifyObject(object, def)

	if err != nil {

		return nil, errSchemaFieldsInvalid
	}

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
			val := object[k].(string)
			ma.AssembleValue().AssignString(val)
		case st.SchemaKind_INT:
			val := int64(object[k].(int))
			ma.AssembleValue().AssignInt(val)
		case st.SchemaKind_FLOAT:
			val := object[k].(float64)
			ma.AssembleValue().AssignFloat(val)
		case st.SchemaKind_BOOL:
			val := object[k].(bool)
			ma.AssembleValue().AssignBool(val)
		case st.SchemaKind_BYTES:
			val := object[k].([]byte)
			ma.AssembleValue().AssignBytes(val)
		case st.SchemaKind_LINK:
			ma.AssembleValue().AssignLink(nil)
		case st.SchemaKind_MAP:
			ma.AssembleValue().AssignNode(nil)
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
