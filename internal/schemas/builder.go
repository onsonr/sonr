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
		if t.Field != st.SchemaKind_LINK {
			err = as.AssignValueToNode(t, ma, object[k])
			if err != nil {
				return err
			}
		} else if t.Field == st.SchemaKind_LINK {
			err := as.BuildSchemaFromLink(t.Link, ma, object[t.Name].(map[string]interface{}))
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

func (as *schemaImpl) AssignValueToNode(field *st.SchemaKindDefinition, ma datamodel.MapAssembler, value interface{}) error {
	switch field.Field {
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
	case st.SchemaKind_LIST:
		val := make([]interface{}, 0)
		s := reflect.ValueOf(value)
		for i := 0; i < s.Len(); i++ {
			val = append(val, s.Index(i).Interface())
		}
		n, err := as.BuildNodeFromList(val, field.Item)
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
		if f.Field != st.SchemaKind_LINK {
			err := as.AssignValueToNode(f, lma, value[f.Name])
			if err != nil {
				return err
			}
		} else if f.Field == st.SchemaKind_LINK {
			err = as.BuildSchemaFromLink(f.Link, lma, value[f.Name].(map[string]interface{}))
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

func (as *schemaImpl) BuildNodeFromList(lst []interface{}, itemType *types.SchemaItemKindDefinition) (datamodel.Node, error) {
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

	err = as.VerifyList(lst, itemType)

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
		case []string:
			value := make([]interface{}, 0)
			s := reflect.ValueOf(val)
			for i := 0; i < s.Len(); i++ {
				value = append(value, s.Index(i).Interface())
			}
			n, err := as.BuildNodeFromList(value, nil)
			if err != nil {
				return nil, err
			}
			la.AssembleValue().AssignNode(n)
		case []int:
			value := make([]interface{}, 0)
			s := reflect.ValueOf(val)
			for i := 0; i < s.Len(); i++ {
				value = append(value, s.Index(i).Interface())
			}
			n, err := as.BuildNodeFromList(value, nil)
			if err != nil {
				return nil, err
			}
			la.AssembleValue().AssignNode(n)
		case []int32:
			value := make([]interface{}, 0)
			s := reflect.ValueOf(val)
			for i := 0; i < s.Len(); i++ {
				value = append(value, s.Index(i).Interface())
			}
			n, err := as.BuildNodeFromList(value, nil)
			if err != nil {
				return nil, err
			}
			la.AssembleValue().AssignNode(n)
		case []int64:
			value := make([]interface{}, 0)
			s := reflect.ValueOf(val)
			for i := 0; i < s.Len(); i++ {
				value = append(value, s.Index(i).Interface())
			}
			n, err := as.BuildNodeFromList(value, nil)
			if err != nil {
				return nil, err
			}
			la.AssembleValue().AssignNode(n)
		case []float32:
			value := make([]interface{}, 0)
			s := reflect.ValueOf(val)
			for i := 0; i < s.Len(); i++ {
				value = append(value, s.Index(i).Interface())
			}
			n, err := as.BuildNodeFromList(value, nil)
			if err != nil {
				return nil, err
			}
			la.AssembleValue().AssignNode(n)
		case []float64:
			value := make([]interface{}, 0)
			s := reflect.ValueOf(val)
			for i := 0; i < s.Len(); i++ {
				value = append(value, s.Index(i).Interface())
			}
			n, err := as.BuildNodeFromList(value, nil)
			if err != nil {
				return nil, err
			}
			la.AssembleValue().AssignNode(n)
		case [][]string:
			value := make([]interface{}, 0)
			s := reflect.ValueOf(val)
			for i := 0; i < s.Len(); i++ {
				value = append(value, s.Index(i).Interface())
			}
			n, err := as.BuildNodeFromList(value, nil)
			if err != nil {
				return nil, err
			}
			la.AssembleValue().AssignNode(n)
		case [][]int:
			value := make([]interface{}, 0)
			s := reflect.ValueOf(val)
			for i := 0; i < s.Len(); i++ {
				value = append(value, s.Index(i).Interface())
			}
			n, err := as.BuildNodeFromList(value, nil)
			if err != nil {
				return nil, err
			}
			la.AssembleValue().AssignNode(n)
		case [][]int32:
			value := make([]interface{}, 0)
			s := reflect.ValueOf(val)
			for i := 0; i < s.Len(); i++ {
				value = append(value, s.Index(i).Interface())
			}
			n, err := as.BuildNodeFromList(value, nil)
			if err != nil {
				return nil, err
			}
			la.AssembleValue().AssignNode(n)
		case [][]int64:
			value := make([]interface{}, 0)
			s := reflect.ValueOf(val)
			for i := 0; i < s.Len(); i++ {
				value = append(value, s.Index(i).Interface())
			}
			n, err := as.BuildNodeFromList(value, nil)
			if err != nil {
				return nil, err
			}
			la.AssembleValue().AssignNode(n)
		case [][]float32:
			value := make([]interface{}, 0)
			s := reflect.ValueOf(val)
			for i := 0; i < s.Len(); i++ {
				value = append(value, s.Index(i).Interface())
			}
			n, err := as.BuildNodeFromList(value, nil)
			if err != nil {
				return nil, err
			}
			la.AssembleValue().AssignNode(n)
		case [][]float64:
			value := make([]interface{}, 0)
			s := reflect.ValueOf(val)
			for i := 0; i < s.Len(); i++ {
				value = append(value, s.Index(i).Interface())
			}
			n, err := as.BuildNodeFromList(value, nil)
			if err != nil {
				return nil, err
			}
			la.AssembleValue().AssignNode(n)
		case [][][]string:
			value := make([]interface{}, 0)
			s := reflect.ValueOf(val)
			for i := 0; i < s.Len(); i++ {
				value = append(value, s.Index(i).Interface())
			}
			n, err := as.BuildNodeFromList(value, nil)
			if err != nil {
				return nil, err
			}
			la.AssembleValue().AssignNode(n)
		case [][][]int:
			value := make([]interface{}, 0)
			s := reflect.ValueOf(val)
			for i := 0; i < s.Len(); i++ {
				value = append(value, s.Index(i).Interface())
			}
			n, err := as.BuildNodeFromList(value, nil)
			if err != nil {
				return nil, err
			}
			la.AssembleValue().AssignNode(n)
		case [][][]int32:
			value := make([]interface{}, 0)
			s := reflect.ValueOf(val)
			for i := 0; i < s.Len(); i++ {
				value = append(value, s.Index(i).Interface())
			}
			n, err := as.BuildNodeFromList(value, nil)
			if err != nil {
				return nil, err
			}
			la.AssembleValue().AssignNode(n)
		case [][][]int64:
			value := make([]interface{}, 0)
			s := reflect.ValueOf(val)
			for i := 0; i < s.Len(); i++ {
				value = append(value, s.Index(i).Interface())
			}
			n, err := as.BuildNodeFromList(value, nil)
			if err != nil {
				return nil, err
			}
			la.AssembleValue().AssignNode(n)
		case [][][]float32:
			value := make([]interface{}, 0)
			s := reflect.ValueOf(val)
			for i := 0; i < s.Len(); i++ {
				value = append(value, s.Index(i).Interface())
			}
			n, err := as.BuildNodeFromList(value, nil)
			if err != nil {
				return nil, err
			}
			la.AssembleValue().AssignNode(n)
		case [][][]float64:
			value := make([]interface{}, 0)
			s := reflect.ValueOf(val)
			for i := 0; i < s.Len(); i++ {
				value = append(value, s.Index(i).Interface())
			}
			n, err := as.BuildNodeFromList(value, nil)
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
