package document

import (
	"fmt"

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

		vals, err := NewDocumentValueFromInterface(field.GetName(), field.GetFieldKind(), v)
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

func NewDocumentValueFromInterface(key string, keyType *st.SchemaFieldKind, value interface{}) (*st.SchemaDocumentValue, error) {
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
		v, ok := value.(int)
		if !ok {
			return nil, fmt.Errorf("could not cast to int")
		}
		return &st.SchemaDocumentValue{
			Key:  key,
			Kind: st.Kind_INT,
			IntValue: &st.IntValue{
				Value: int32(v),
			},
		}, nil
	case st.Kind_FLOAT:
		v, ok := value.(float64)
		if !ok {
			return nil, fmt.Errorf("could not cast to float")
		}
		return &st.SchemaDocumentValue{
			Key:  key,
			Kind: st.Kind_FLOAT,
			FloatValue: &st.FloatValue{
				Value: v,
			},
		}, nil
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
		v, ok := value.([]*st.SchemaDocumentValue)
		if !ok {
			return nil, fmt.Errorf("could not cast to List")
		}
		return &st.SchemaDocumentValue{
			Key:  key,
			Kind: st.Kind_LIST,
			ListValue: &st.ListValue{
				Value: v,
			},
		}, nil
	case st.Kind_LINK:
		v, ok := value.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("could not cast to string for list")
		}
		linkedValue, err := NewDocumentFromDag(v, nil) // TODO
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
