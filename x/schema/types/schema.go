package types

func (skd *SchemaKindDefinition) GetKind() SchemaKind {
	return skd.Field
}

func (skd *SchemaKindDefinition) TryValue(val interface{}) bool {
	switch val.(type) {
	case int:
		return skd.Field == SchemaKind_INT
	case uint:
		return skd.Field == SchemaKind_INT
	case int32:
		return skd.Field == SchemaKind_INT
	case uint32:
		return skd.Field == SchemaKind_INT
	case int64:
		return skd.Field == SchemaKind_INT
	case uint64:
		return skd.Field == SchemaKind_INT
	case string:
		return skd.Field == SchemaKind_STRING
	case float64:
		return skd.Field == SchemaKind_FLOAT
	case bool:
		return skd.Field == SchemaKind_BOOL
	case []byte:
		return skd.Field == SchemaKind_BYTES
	case []interface{}:
		return skd.Field == SchemaKind_LIST
	default:
		return false
	}
}
