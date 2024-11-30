package vault

import (
	"reflect"
	"strings"

	"github.com/onsonr/sonr/pkg/common/models"
	"github.com/onsonr/sonr/web/vault/types"
)

const (
	DWNConfigFileName     = "dwn.json"
	IndexHTMLFileName     = "index.html"
	ServiceWorkerFileName = "sw.js"
)

const SchemaVersion = 1

func toCamelCase(s string) string {
	if s == "" {
		return s
	}
	if len(s) == 1 {
		return strings.ToLower(s)
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func GetSchema(structType interface{}) string {
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

// DefaultSchema returns the default schema
func DefaultSchema() *types.Schema {
	return &types.Schema{
		Version:    SchemaVersion,
		Account:    GetSchema(&models.Account{}),
		Asset:      GetSchema(&models.Asset{}),
		Chain:      GetSchema(&models.Chain{}),
		Credential: GetSchema(&models.Credential{}),
		Grant:      GetSchema(&models.Grant{}),
		Keyshare:   GetSchema(&models.Keyshare{}),
		Profile:    GetSchema(&models.Profile{}),
	}
}
