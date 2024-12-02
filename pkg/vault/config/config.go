package config

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/ipfs/boxo/files"
	"github.com/onsonr/sonr/pkg/common/models"
	"github.com/onsonr/sonr/pkg/vault/config/internal"
	"github.com/onsonr/sonr/pkg/vault/types"
)

const SchemaVersion = 1
const (
	AppManifestFileName   = "app.webmanifest"
	DWNConfigFileName     = "dwn.json"
	IndexHTMLFileName     = "index.html"
	MainJSFileName        = "main.js"
	ServiceWorkerFileName = "sw.js"
)

// spawnVaultDirectory creates a new directory with the default files
func NewFS(cfg *types.Config) (files.Directory, error) {
	manifestBz, err := newWebManifestBytes()
	if err != nil {
		return nil, err
	}
	cnfBz, err := json.Marshal(cfg)
	if err != nil {
		return nil, err
	}
	return files.NewMapDirectory(map[string]files.Node{
		AppManifestFileName:   files.NewBytesFile(manifestBz),
		DWNConfigFileName:     files.NewBytesFile(cnfBz),
		IndexHTMLFileName:     files.NewBytesFile(internal.IndexHTML),
		MainJSFileName:        files.NewBytesFile(internal.MainJS),
		ServiceWorkerFileName: files.NewBytesFile(internal.WorkerJS),
	}), nil
}

// DefaultSchema returns the default schema
func DefaultSchema() *types.Schema {
	return &types.Schema{
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
