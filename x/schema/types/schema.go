package types

const (
	IPLD_LABEL      = "label"
	IPLD_SCHEMA_DID = "schema_did"
	IPLD_DOCUMENT   = "document"
)

func (skd *SchemaField) GetKind() Kind {
	return skd.FieldKind.Kind
}

func (skd *SchemaField) TryValue(val interface{}) bool {
	switch val.(type) {
	case int:
		return skd.GetKind() == Kind_INT
	case uint:
		return skd.GetKind() == Kind_INT
	case int32:
		return skd.GetKind() == Kind_INT
	case uint32:
		return skd.GetKind() == Kind_INT
	case int64:
		return skd.GetKind() == Kind_INT
	case uint64:
		return skd.GetKind() == Kind_INT
	case string:
		return skd.GetKind() == Kind_STRING
	case float64:
		return skd.GetKind() == Kind_FLOAT
	case bool:
		return skd.GetKind() == Kind_BOOL
	case []byte:
		return skd.GetKind() == Kind_BYTES
	case []interface{}, []map[string]interface{}, []string, []int, []int32, []int64, []float32, []float64:
		return skd.GetKind() == Kind_LIST
	default:
		return false
	}
}
