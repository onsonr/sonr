package schemas

import (
	"errors"
	"reflect"

	st "github.com/sonr-io/sonr/x/schema/types"
)

var (
	DocumentSpecialFields = []string{"@did"}
)

/*
	Top level verification of the given schema def
*/
func (as *schemaImpl) VerifyObject(doc map[string]interface{}) error {
	if as.fields == nil {
		return errSchemaFieldsNotFound
	}

	fields := make(map[string]st.Kind)
	for _, c := range as.fields {
		fields[c.Name] = c.FieldKind.Kind
	}

	for key, value := range doc {
		if _, ok := fields[key]; !ok {
			// check for special metadata fields, if found skip validation
			if !arrayContains(DocumentSpecialFields, key) {
				return errSchemaFieldsInvalid
			} else {
				continue
			}
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
func (as *schemaImpl) VerifySubObject(lst []*st.SchemaField, doc map[string]interface{}) error {
	if as.fields == nil {
		return errSchemaFieldsNotFound
	}

	fields := make(map[string]st.Kind)
	for _, c := range lst {
		fields[c.Name] = c.GetKind()
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

func (as *schemaImpl) VerifyList(lst []interface{}, itemType *st.SchemaFieldKind) error {
	if itemType == nil {
		for _, val := range lst {
			if reflect.TypeOf(val) != reflect.TypeOf(lst[0]) {
				return errors.New("array type is not of uniform values")
			}
		}
		return nil
	}

	for _, val := range lst {
		if !CheckValueOfField(val, itemType.GetKind()) {
			return errSchemaFieldsInvalid
		}
	}

	return nil
}

// Current supported IPLD types, will be adding more once supporting of Links and Complex types (Object)
func CheckValueOfField(value interface{}, fieldType st.Kind) bool {
	switch value.(type) {
	case int:
		return fieldType == st.Kind_INT
	case uint:
		return fieldType == st.Kind_INT
	case int32:
		return fieldType == st.Kind_INT
	case int64:
		return fieldType == st.Kind_INT
	case float64:
		return fieldType == st.Kind_FLOAT
	case float32:
		return fieldType == st.Kind_FLOAT
	case bool:
		return fieldType == st.Kind_BOOL
	case string:
		return fieldType == st.Kind_STRING
	case []byte:
		return fieldType == st.Kind_BYTES
	case []interface{}:
		return fieldType == st.Kind_LIST
	case []int:
		return fieldType == st.Kind_LIST
	case []int32:
		return fieldType == st.Kind_LIST
	case []int64:
		return fieldType == st.Kind_LIST
	case []bool:
		return fieldType == st.Kind_LIST
	case []float64:
		return fieldType == st.Kind_LIST
	case []float32:
		return fieldType == st.Kind_LIST
	case []string:
		return fieldType == st.Kind_LIST
	case [][]byte:
		return fieldType == st.Kind_LIST
	case [][]string:
		return fieldType == st.Kind_LIST
	case [][]int32:
		return fieldType == st.Kind_LIST
	case [][]int64:
		return fieldType == st.Kind_LIST
	case [][]float32:
		return fieldType == st.Kind_LIST
	case [][]float64:
		return fieldType == st.Kind_LIST
	case [][][]bool:
		return fieldType == st.Kind_LIST
	case [][][]byte:
		return fieldType == st.Kind_LIST
	case [][][]string:
		return fieldType == st.Kind_LIST
	case [][][]int32:
		return fieldType == st.Kind_LIST
	case [][][]int64:
		return fieldType == st.Kind_LIST
	case [][][]float32:
		return fieldType == st.Kind_LIST
	case [][][]float64:
		return fieldType == st.Kind_LIST
	case map[string]interface{}:
		return fieldType == st.Kind_LINK
	default:
		return false
	}
}
