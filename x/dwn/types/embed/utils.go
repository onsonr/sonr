package embed

import (
	_ "embed"
	"reflect"
	"strings"
)

//go:embed index.html
var IndexHTML []byte

//go:embed main.js
var MainJS []byte

//go:embed sw.js
var WorkerJS []byte

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
