package document

import (
	"fmt"
	"reflect"

	st "github.com/sonr-io/sonr/x/schema/types"
)

func NewDocumentFromDag(m map[string]interface{}, schema Schema) (*st.Document, error) {
	fields := make([]*st.DocumentValue, 0)

	label, ok := m[st.IPLD_LABEL].(string)
	if !ok {
		return nil, fmt.Errorf("could not find label")
	}

	document, ok := m[st.IPLD_DOCUMENT].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("could not find document for document %s", label)
	}

	for _, field := range schema.GetFields() {
		v, ok := document[field.GetName()]
		if !ok {
			return nil, fmt.Errorf("DAG does not contain key %s", field.GetName())
		}

		vals, err := newDocumentValueFromInterface(field.GetName(), schema, field.GetFieldKind(), v)
		if err != nil {
			return nil, err
		}
		fields = append(fields, vals)
	}

	s, err := schema.GetSchema()
	if err != nil {
		return nil, err
	}

	return &st.Document{
		Label:     label,
		Fields:    fields,
		SchemaDid: s.Did,
	}, nil
}

func newDocumentValueFromInterface(key string, schema Schema, keyType *st.SchemaFieldKind, value interface{}) (*st.DocumentValue, error) {
	switch keyType.Kind {
	case st.Kind_BOOL:
		v, ok := value.(bool)
		if !ok {
			return nil, fmt.Errorf("could not cast to bool")
		}
		return &st.DocumentValue{
			Key:  key,
			Kind: st.Kind_BOOL,
			BoolValue: &st.BoolValue{
				Value: v,
			},
		}, nil
	case st.Kind_BYTES:
		v, ok := value.([]byte)
		if !ok {
			return nil, fmt.Errorf("could not cast to []byte")
		}
		return &st.DocumentValue{
			Key:  key,
			Kind: st.Kind_BYTES,
			BytesValue: &st.BytesValue{
				Value: v,
			},
		}, nil
	case st.Kind_INT:
		switch v := value.(type) {
		case int:
			return &st.DocumentValue{
				Key:  key,
				Kind: st.Kind_INT,
				IntValue: &st.IntValue{
					Value: int32(v),
				},
			}, nil
		case int32:
			return &st.DocumentValue{
				Key:  key,
				Kind: st.Kind_INT,
				IntValue: &st.IntValue{
					Value: int32(v),
				},
			}, nil
		case int64:
			return &st.DocumentValue{
				Key:  key,
				Kind: st.Kind_INT,
				IntValue: &st.IntValue{
					Value: int32(v),
				},
			}, nil
		default:
			return nil, fmt.Errorf("could not cast to int")

		}
	case st.Kind_FLOAT:
		switch v := value.(type) {
		case float32:
			return &st.DocumentValue{
				Key:  key,
				Kind: st.Kind_FLOAT,
				FloatValue: &st.FloatValue{
					Value: float64(v),
				},
			}, nil
		case float64:
			return &st.DocumentValue{
				Key:  key,
				Kind: st.Kind_FLOAT,
				FloatValue: &st.FloatValue{
					Value: v,
				},
			}, nil
		default:
			return nil, fmt.Errorf("could not cast to float")
		}
	case st.Kind_STRING:
		v, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("could not cast to string")
		}
		return &st.DocumentValue{
			Key:  key,
			Kind: st.Kind_STRING,
			StringValue: &st.StringValue{
				Value: v,
			},
		}, nil
	case st.Kind_LIST:
		s := reflect.ValueOf(value)
		val := make([]*st.DocumentValue, s.Len())

		for i := 0; i < s.Len(); i++ {
			item, err := newDocumentValueFromInterface("", schema, keyType.ListKind, s.Index(i).Interface())
			if err != nil {
				return nil, fmt.Errorf("deciphering list value: %s", err)
			}
			val[i] = item
		}
		return &st.DocumentValue{
			Key:  key,
			Kind: st.Kind_LIST,
			ListValue: &st.ListValue{
				Value: val,
			},
		}, nil
	case st.Kind_LINK:
		v, ok := value.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("could not cast to string for list")
		}

		subSchemaDid, ok := v[st.IPLD_SCHEMA_DID].(string)
		if !ok {
			return nil, fmt.Errorf("sub schema has no did")
		}

		subSchema, err := schema.GetSubSchema(subSchemaDid)
		if err != nil {
			return nil, fmt.Errorf("get subschema: %s", err)
		}

		linkedValue, err := NewDocumentFromDag(v, subSchema)
		if err != nil {
			return nil, err
		}
		return &st.DocumentValue{
			Key:  key,
			Kind: st.Kind_LINK,
			LinkValue: &st.LinkValue{
				Value: linkedValue,
			},
		}, nil
	default:
		return nil, fmt.Errorf("unknown schema type %s", keyType)
	}
}
