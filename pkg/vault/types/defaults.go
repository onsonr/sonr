package types

import (
	"reflect"
	"strings"

	"github.com/onsonr/sonr/pkg/common/models"
)

const SchemaVersion = 1

// DefaultSchema returns the default schema
func DefaultSchema() *Schema {
	return &Schema{
		Version:    SchemaVersion,
		Account:    getSchema(&models.Account{}),
		Asset:      getSchema(&models.Asset{}),
		Chain:      getSchema(&models.Chain{}),
		Credential: getSchema(&models.Credential{}),
		Grant:      getSchema(&models.Grant{}),
		Keyshare:   getSchema(&models.Keyshare{}),
		Profile:    getSchema(&models.Profile{}),
	}
}

func getSchema(structType interface{}) string {
	t := reflect.TypeOf(structType)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return ""
	}

	var fields []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldName := toCamelCase(field.Name)
		fields = append(fields, fieldName)
	}

	// Add "++" at the beginning, separated by a comma
	return "++, " + strings.Join(fields, ", ")
}

func toCamelCase(s string) string {
	if s == "" {
		return s
	}
	if len(s) == 1 {
		return strings.ToLower(s)
	}
	return strings.ToLower(s[:1]) + s[1:]
}
