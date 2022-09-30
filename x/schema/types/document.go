package types

func (d *SchemaDocumentValue) GetValue() interface{} {
	switch d.Kind {
	case Kind_BOOL:
		if d.BoolValue != nil {
			return d.BoolValue.Value
		}
	case Kind_BYTES:
		if d.BytesValue != nil {
			return d.BytesValue.Value
		}
	case Kind_INT:
		if d.IntValue != nil {
			return int64(d.IntValue.Value)
		}
	case Kind_FLOAT:
		if d.FloatValue != nil {
			return d.FloatValue.Value
		}
	case Kind_STRING:

		if d.StringValue != nil {
			return d.StringValue.Value
		}
	case Kind_LINK:
		if d.LinkValue != nil {
			return d.LinkValue.Value
		}
	case Kind_LIST:
		if d.ArrayValue != nil {
			return resolveArrayValues(d.ArrayValue.Value)
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
		Fields:    fields,
		SchemaDid: schemaDid,
		Cid:       cid,
	}
}

func NewDocumentValueFromInterface(name string, value interface{}) *SchemaDocumentValue {
	switch v := value.(type) {
	case bool:
		return &SchemaDocumentValue{
			Name: name,
			Kind: Kind_BOOL,
			BoolValue: &BoolValue{
				Value: v,
			},
		}
	case []byte:
		// TODO: add schema DID to objects in IPFS so that the types don't need to be inferred.
		// if strings.Contains(v, "did:") && name != "@did" {
		// 	return &SchemaDocumentValue{
		// 		Name: name,
		// 		Kind: Kind_LINK,
		// 		LinkValue: &LinkValue{
		// 			Value: ,
		// 		},
		// 	}
		// }
		return &SchemaDocumentValue{
			Name: name,
			Kind: Kind_BYTES,
			BytesValue: &BytesValue{
				Value: v,
			},
		}
	case int:
		return &SchemaDocumentValue{
			Name: name,
			Kind: Kind_INT,
			IntValue: &IntValue{
				Value: int32(v),
			},
		}
	case float64:
		return &SchemaDocumentValue{
			Name: name,
			Kind: Kind_FLOAT,
			FloatValue: &FloatValue{
				Value: v,
			},
		}
	case string:
		return &SchemaDocumentValue{
			Name: name,
			Kind: Kind_STRING,
			StringValue: &StringValue{
				Value: v,
			},
		}
	case []*SchemaDocumentValue:
		return &SchemaDocumentValue{
			Name: name,
			Kind: Kind_LIST,
			ArrayValue: &ArrayValue{
				Value: v,
			},
		}
	default:
		return nil
	}
}

func resolveArrayValues(vals []*SchemaDocumentValue) []interface{} {
	arr := make([]interface{}, 0)
	for _, val := range vals {
		if val.BoolValue != nil {
			arr = append(arr, val.BoolValue.Value)
		} else if val.StringValue != nil {
			arr = append(arr, val.StringValue.Value)
		} else if val.IntValue != nil {
			arr = append(arr, val.IntValue.Value)
		} else if val.FloatValue != nil {
			arr = append(arr, val.FloatValue.Value)
		} else if val.BytesValue != nil {
			arr = append(arr, val.BytesValue.Value)
		}
	}

	return arr
}
