package document

import (
	"fmt"
	"reflect"

	st "github.com/sonr-io/sonr/x/schema/types"
)

func NewDocumentFromDag(m map[string]interface{}, schema Schema) (*st.SchemaDocument, error) {
	fields := make([]*st.SchemaDocumentValue, 0)

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

	return &st.SchemaDocument{
		Label:     label,
		Fields:    fields,
		SchemaDid: s.Did,
	}, nil
}

func newDocumentValueFromInterface(key string, schema Schema, keyType *st.SchemaFieldKind, value interface{}) (*st.SchemaDocumentValue, error) {
	switch keyType.Kind {
	case st.Kind_BOOL:
		v, ok := value.(bool)
		if !ok {
			return nil, fmt.Errorf("could not cast to bool")
		}
		return &st.SchemaDocumentValue{
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
		return &st.SchemaDocumentValue{
			Key:  key,
			Kind: st.Kind_BYTES,
			BytesValue: &st.BytesValue{
				Value: v,
			},
		}, nil
	case st.Kind_INT:
		switch v := value.(type) {
		case int:
			return &st.SchemaDocumentValue{
				Key:  key,
				Kind: st.Kind_INT,
				IntValue: &st.IntValue{
					Value: int32(v),
				},
			}, nil
		case int32:
			return &st.SchemaDocumentValue{
				Key:  key,
				Kind: st.Kind_INT,
				IntValue: &st.IntValue{
					Value: v,
				},
			}, nil
		case int64:
			return &st.SchemaDocumentValue{
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
			return &st.SchemaDocumentValue{
				Key:  key,
				Kind: st.Kind_FLOAT,
				FloatValue: &st.FloatValue{
					Value: float64(v),
				},
			}, nil
		case float64:
			return &st.SchemaDocumentValue{
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
		return &st.SchemaDocumentValue{
			Key:  key,
			Kind: st.Kind_STRING,
			StringValue: &st.StringValue{
				Value: v,
			},
		}, nil
	case st.Kind_LIST:
		s := reflect.ValueOf(value)
		val := make([]*st.SchemaDocumentValue, s.Len())

		for i := 0; i < s.Len(); i++ {
			item, err := newDocumentValueFromInterface("", schema, keyType.ListKind, s.Index(i).Interface())
			if err != nil {
				return nil, fmt.Errorf("deciphering list value: %s", err)
			}
			val[i] = item
		}
		return &st.SchemaDocumentValue{
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

		subSchema, err := schema.GetSubSchema(keyType.LinkDid)
		if err != nil {
			return nil, fmt.Errorf("get subschema: %s", err)
		}

		v = embedDocument(v, keyType)
		linkedValue, err := NewDocumentFromDag(v, subSchema)
		if err != nil {
			return nil, err
		}
		return &st.SchemaDocumentValue{
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

func embedDocument(doc map[string]interface{}, field *st.SchemaFieldKind) map[string]interface{} {
	if field.GetKind() == st.Kind_LINK {
		return map[string]interface{}{
			st.IPLD_LABEL:      "",
			st.IPLD_SCHEMA_DID: field.LinkDid,
			st.IPLD_DOCUMENT:   doc,
		}
	}
	return doc
}
