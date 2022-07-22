package schemas

import (
	"github.com/ipld/go-ipld-prime/datamodel"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
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
	ma, err := nb.BeginMap(int64(len(def.GetField())))

	if err != nil {
		return nil, err
	}

	for _, t := range def.GetField() {
		k := t.Name
		ma.AssembleKey().AssignString(k)
		switch t.Field {
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
			val := object[k].(string)
			link, err := as.LoadLink(val)
			if err != nil {
				return nil, err
			}
			ma.AssembleValue().AssignLink(link)
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

	return node, nil
}

func (as *appSchemaInternalImpl) LoadJOSE() {

}

func (as *appSchemaInternalImpl) BuildJoseLink() {

}

func (as *appSchemaInternalImpl) LoadLink(val interface{}) (cidlink.Link, error) {
	return cidlink.Link{}, nil
}
