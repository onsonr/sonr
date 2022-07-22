package schemas

import (
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (as *appSchemaInternalImpl) VerifyObject(doc map[string]interface{}, def *st.SchemaDefinition) error {
	lst := def.GetField()
	fields := make(map[string]st.SchemaKind)
	for _, c := range lst {
		fields[c.Name] = c.Field
	}

	for key, value := range doc {
		if _, ok := fields[key]; !ok {
			return errSchemaFieldsInvalid
		}
		if !CheckValueOfField(value, fields[key]) {
			return errSchemaFieldsInvalid
		}
	}

	return nil
}

func CheckValueOfField(value interface{}, fieldType st.SchemaKind) bool {
	switch value.(type) {
	case int:
		return fieldType == st.SchemaKind_INT
	case float64:
		return fieldType == st.SchemaKind_FLOAT
	case bool:
		return fieldType == st.SchemaKind_BOOL
	case string:
		return fieldType == st.SchemaKind_STRING
	case []byte:
		return fieldType == st.SchemaKind_BYTES
	case interface{}:
		return fieldType == st.SchemaKind_ANY
	default:
		return false
	}
}
