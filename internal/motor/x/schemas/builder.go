package schemas

import (
	"github.com/ipfs/go-cid"
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

func (as *appSchemaInternalImpl) LoadLink(val interface{}) (cidlink.Link, error) {
	value := val.(string)
	cid, err := cid.Decode(value)
	if err != nil {
		return cidlink.Link{}, err
	}

	lnk := cidlink.Link{Cid: cid}
	/*
		lsys := cidlink.DefaultLinkSystem()
		lsys.SetReadStorage(store)

		np := basicnode.Prototype.Any

		// Apply the LinkSystem loader for the given cid
		node, err := lsys.Load(
			linking.LinkContext{},
			lnk,
			np,
		)

		if err != nil {
			return nil, err
		}
	*/
	return lnk, nil
}
