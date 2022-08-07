package schemas

import (
	"errors"
	"reflect"

	st "github.com/sonr-io/sonr/x/schema/types"
)

/*
	Top level verification of the given schema def
*/
func (as *schemaImpl) VerifyObject(doc map[string]interface{}) error {
	if as.fields == nil {
		return errSchemaFieldsNotFound
	}

	fields := make(map[string]st.SchemaKind)
	for _, c := range as.fields {
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

/*
	Sub level verification of the given schema def
*/
func (as *schemaImpl) VerifySubObject(lst []*st.SchemaKindDefinition, doc map[string]interface{}) error {
	if as.fields == nil {
		return errSchemaFieldsNotFound
	}

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

func (as *schemaImpl) VerifyList(lst []interface{}) error {
	for _, val := range lst {
		if reflect.TypeOf(val) != reflect.TypeOf(lst[0]) {
			return errors.New("array type is not of uniform values")
		}
	}

	return nil
}

// Current supported IPLD types, will be adding more once supporting of Links and Complex types (Object)
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
	case []int:
		return fieldType == st.SchemaKind_LIST
	case []bool:
		return fieldType == st.SchemaKind_LIST
	case []float64:
		return fieldType == st.SchemaKind_LIST
	case []string:
		return fieldType == st.SchemaKind_LIST
	case map[string]interface{}:
		return fieldType == st.SchemaKind_LINK
	case interface{}:
		return fieldType == st.SchemaKind_ANY
	default:
		return false
	}
}
