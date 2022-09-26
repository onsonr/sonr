package schemas

import (
	"reflect"

	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/sonr-io/sonr/x/schema/types"
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
			err = as.AssignValueToNode(t.FieldKind, ma, object[k])
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

func (as *schemaImpl) AssignValueToNode(kind *st.SchemaFieldKind, ma datamodel.MapAssembler, value interface{}) error {
	switch kind.GetKind() {
	case st.Kind_STRING:
		val := value.(string)
		ma.AssembleValue().AssignString(val)
	case st.Kind_INT:
		switch value.(type) {
		case int:
			val := int64(value.(int))
			ma.AssembleValue().AssignInt(val)
		case int32:
			val := int64(value.(int32))
			ma.AssembleValue().AssignInt(val)
		case int64:
			val := int64(value.(int64))
			ma.AssembleValue().AssignInt(val)
		}
	case st.Kind_FLOAT:
		switch value.(type) {
		case float64:
			val := value.(float64)
			ma.AssembleValue().AssignFloat(val)
		case float32:
			val := value.(float32)
			ma.AssembleValue().AssignFloat(float64(val))
		}
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
		n, err := as.BuildNodeFromList(val, kind.ListKind)
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
			err := as.AssignValueToNode(f.FieldKind.ListKind, lma, value[f.Name])
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

func (as *schemaImpl) BuildNodeFromList(lst []interface{}, kind *types.SchemaFieldKind) (datamodel.Node, error) {
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

	err = as.VerifyList(lst, kind)

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
			lstItem := interface{}(val).(int64)
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
		case []byte:
			lstItem := interface{}(val).([]byte)
			la.AssembleValue().AssignBytes(lstItem)
		/*
			The below cases are for handling lists of up to 3 dimensions.
			Within each cases arrays are normalized to match a type of []interface{}
			each generic []interface{} array is then handed back to this function to further resolve types.
			depth is cut off at 3 dimensions due to having to implement explicit type cases here
		*/
		case []string, []int, []int32, []int64, []float32, []float64:
			value := make([]interface{}, 0)
			s := reflect.ValueOf(val)
			for i := 0; i < s.Len(); i++ {
				value = append(value, s.Index(i).Interface())
			}
			n, err := as.BuildNodeFromList(value, kind.ListKind)
			if err != nil {
				return nil, err
			}
			la.AssembleValue().AssignNode(n)
		default:
			value := make([]interface{}, 0)
			s := reflect.ValueOf(val)
			for i := 0; i < s.Len(); i++ {
				value = append(value, s.Index(i).Interface())
			}
			n, err := as.BuildNodeFromList(value, kind.ListKind)
			if err != nil {
				return nil, err
			}
			la.AssembleValue().AssignNode(n)
		}
	}
	err = la.Finish()

	if err != nil {
		return nil, err
	}

	return nb.Build(), nil
}
