package schemas

import (
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (as *appSchemaInternalImpl) BuildNodesFromDefinition(
	object map[string]interface{}) error {
	if as.fields == nil {
		return errSchemaFieldsInvalid
	}

	err := as.VerifyObject(object)

	if err != nil {

		return errSchemaFieldsInvalid
	}

	// Create IPLD Node
	np := basicnode.Prototype.Any
	nb := np.NewBuilder() // Create a builder.
	ma, err := nb.BeginMap(int64(len(as.fields)))

	if err != nil {
		return err
	}

	for _, t := range as.fields {
		k := t.Name
		ma.AssembleKey().AssignString(k)
		if t.Field != st.SchemaKind_STRUCT && t.Field != st.SchemaKind_MAP {
			AssignValueToNode(t.Field, ma, object[k])
		}
	}

	buildErr := ma.Finish()

	if buildErr != nil {
		return buildErr
	}
	node := nb.Build()

	as.nodes = node

	return nil
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
