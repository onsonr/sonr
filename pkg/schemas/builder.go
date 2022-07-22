package schemas

import (
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (as *appSchemaInternalImpl) BuildNodesFromDefinition(
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

	for _, t := range def.GetFields() {
		k := t.Name
		ma.AssembleKey().AssignString(k)
		if t.Field != st.SchemaKind_STRUCT && t.Field != st.SchemaKind_MAP {
			AssignValueToNode(t.Field, ma, object[k])
		}
	}

	buildErr := ma.Finish()

	if buildErr != nil {
		return nil, buildErr
	}
	node := nb.Build()

	return node, nil
}

func AssignValueToNode(field st.SchemaKind, ma datamodel.MapAssembler, value interface{}) error {
	switch field {
	case st.SchemaKind_STRING:
		val := value.(string)
		ma.AssembleValue().AssignString(val)
	case st.SchemaKind_INT:
		val := int64(value.(int))
		ma.AssembleValue().AssignInt(val)
	case st.SchemaKind_FLOAT:
		val := value.(float64)
		ma.AssembleValue().AssignFloat(val)
	case st.SchemaKind_BOOL:
		val := value.(bool)
		ma.AssembleValue().AssignBool(val)
	case st.SchemaKind_BYTES:
		val := value.([]byte)
		ma.AssembleValue().AssignBytes(val)
	default:
		return errSchemaFieldsInvalid
	}

	return nil
}
