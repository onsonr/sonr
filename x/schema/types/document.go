package types

import "strings"

func (d *SchemaDocumentValue) GetValue() interface{} {
	switch d.Field {
	case SchemaKind_BOOL:
		if d.BoolValue != nil {
			return d.BoolValue
		}
	case SchemaKind_BYTES:
		if d.BytesValue != nil {
			return d.BytesValue
		}
	case SchemaKind_INT:
		if d.IntValue != nil {
			return d.IntValue
		}
	case SchemaKind_FLOAT:
		if d.FloatValue != nil {
			return d.FloatValue
		}
	case SchemaKind_STRING:

		if d.StringValue != nil {
			return d.StringValue
		}
	case SchemaKind_LINK:
		if d.LinkValue != nil {
			return d.LinkValue
		}
	case SchemaKind_LIST:
		if d.ArrayValue != nil {
			return d.ArrayValue
		}
	default:
		return nil
	}
	return nil
}

func NewDocumentFromMap(cid string, m map[string]interface{}) *SchemaDocument {
	var schemaDid string
	fields := make([]*SchemaDocumentValue, 0)
	for k, v := range m {
		if k == "@did" {
			schemaDid = v.(string)
			continue
		}
		fields = append(fields, NewDocumentValueFromInterface(k, v))
	}
	return &SchemaDocument{
		Fields: fields,
		Did:    schemaDid,
		Cid:    cid,
	}
}

func NewDocumentValueFromInterface(name string, value interface{}) *SchemaDocumentValue {
	switch v := value.(type) {
	case bool:
		return &SchemaDocumentValue{
			Name:  name,
			Field: SchemaKind_BOOL,
			BoolValue: &BoolValue{
				Value: v,
			},
		}
	case []byte:
		return &SchemaDocumentValue{
			Name:  name,
			Field: SchemaKind_BYTES,
			BytesValue: &BytesValue{
				Value: v,
			},
		}
	case int:
		return &SchemaDocumentValue{
			Name:  name,
			Field: SchemaKind_INT,
			IntValue: &IntValue{
				Value: int32(v),
			},
		}
	case float64:
		return &SchemaDocumentValue{
			Name:  name,
			Field: SchemaKind_FLOAT,
			FloatValue: &FloatValue{
				Value: v,
			},
		}
	case string:
		if strings.Contains(v, "did:") && name != "@did" {
			return &SchemaDocumentValue{
				Name:  name,
				Field: SchemaKind_LINK,
				LinkValue: &LinkValue{
					Value: v,
					Link:  LinkKind_UNKNOWN,
				},
			}
		}
		return &SchemaDocumentValue{
			Name:  name,
			Field: SchemaKind_STRING,
			StringValue: &StringValue{
				Value: v,
			},
		}
	case []*SchemaDocumentValue:
		return &SchemaDocumentValue{
			Name:  name,
			Field: SchemaKind_LIST,
			ArrayValue: &ArrayValue{
				Value: v,
			},
		}
	default:
		return nil
	}
}
