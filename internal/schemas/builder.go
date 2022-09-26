package schemas

import (
	"reflect"

	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (as *schemaImpl) BuildNodesFromDefinition(
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
		if t.GetKind() != st.Kind_LINK {
			err = as.AssignValueToNode(t.GetKind(), ma, object[k])
			if err != nil {
				return err
			}
		} else if t.GetKind() == st.Kind_LINK {
			err := as.BuildSchemaFromLink(t.FieldKind.LinkDid, ma, object[t.Name].(map[string]interface{}))
			if err != nil {
				return err
			}
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

func (as *schemaImpl) AssignValueToNode(field st.Kind, ma datamodel.MapAssembler, value interface{}) error {
	switch field {
	case st.Kind_STRING:
		val := value.(string)
		ma.AssembleValue().AssignString(val)
	case st.Kind_INT:
		val := int64(value.(int64))
		ma.AssembleValue().AssignInt(val)
	case st.Kind_FLOAT:
		val := value.(float64)
		ma.AssembleValue().AssignFloat(val)
	case st.Kind_BOOL:
		val := value.(bool)
		ma.AssembleValue().AssignBool(val)
	case st.Kind_BYTES:
		val := value.([]byte)
		ma.AssembleValue().AssignBytes(val)
	case st.Kind_LIST:
		val := make([]interface{}, 0)
		s := reflect.ValueOf(value)
		for i := 0; i < s.Len(); i++ {
			val = append(val, s.Index(i).Interface())
		}
		n, err := as.BuildNodeFromList(val)
		if err != nil {
			return errSchemaFieldsInvalid
		}
		ma.AssembleValue().AssignNode(n)
	default:
		return errSchemaFieldsInvalid
	}

	return nil
}

func (as *schemaImpl) BuildSchemaFromLink(key string, ma datamodel.MapAssembler, value map[string]interface{}) error {
	if as.subSchemas[key] == nil {
		return errNodeNotFound
	}

	sd := as.subSchemas[key]

	err := as.VerifySubObject(sd.Fields, value)

	if err != nil {
		return err
	}

	// Create IPLD Node
	np := basicnode.Prototype.Any
	nb := np.NewBuilder() // Create a builder.
	lma, err := nb.BeginMap(int64(len(value)))

	if err != nil {
		return err
	}

	for _, f := range sd.Fields {
		lma.AssembleKey().AssignString(f.Name)
		if f.GetKind() != st.Kind_LINK {
			err := as.AssignValueToNode(f.GetKind(), lma, value[f.Name])
			if err != nil {
				return err
			}
		} else if f.GetKind() == st.Kind_LINK {
			err = as.BuildSchemaFromLink(f.FieldKind.LinkDid, lma, value[f.Name].(map[string]interface{}))
			if err != nil {
				return err
			}
		}

	}

	lma.Finish()
	n := nb.Build()
	ma.AssembleValue().AssignNode(n)

	return nil
}

func (as *schemaImpl) BuildNodeFromList(lst []interface{}) (datamodel.Node, error) {
	// Create IPLD Node
	np := basicnode.Prototype.Any
	nb := np.NewBuilder() // Create a builder.
	la, err := nb.BeginList(int64(len(lst)))

	if err != nil {
		return nil, err
	}

	if len(lst) < 1 {
		return nb.Build(), nil
	}

	err = as.VerifyList(lst)

	if err != nil {
		return nil, err
	}

	for _, val := range lst {
		switch lst[0].(type) {
		case int:
			lstItem := interface{}(val).(int)
			la.AssembleValue().AssignInt(int64(lstItem))
		case int32:
			lstItem := interface{}(val).(int32)
			la.AssembleValue().AssignInt(int64(lstItem))
		case int64:
			lstItem := interface{}(val).(int32)
			la.AssembleValue().AssignInt(int64(lstItem))
		case float64:
			lstItem := interface{}(val).(float64)
			la.AssembleValue().AssignFloat(lstItem)
		case float32:
			lstItem := interface{}(val).(float32)
			la.AssembleValue().AssignFloat(float64(lstItem))
		case bool:
			lstItem := interface{}(val).(bool)
			la.AssembleValue().AssignBool(lstItem)
		case string:
			lstItem := interface{}(val).(string)
			la.AssembleValue().AssignString(lstItem)
		}
	}
	err = la.Finish()

	if err != nil {
		return nil, err
	}

	return nb.Build(), nil
}
