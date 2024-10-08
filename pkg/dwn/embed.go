package dwn

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"

	"github.com/ipfs/boxo/files"
	"github.com/onsonr/sonr/pkg/nebula/global"
)

//go:embed app.wasm
var dwnWasmData []byte

//go:embed sw.js
var swJSData []byte

var (
	dwnWasmFile = files.NewBytesFile(dwnWasmData)
	swJSFile    = files.NewBytesFile(swJSData)
)

// NewVaultDirectory creates a new directory with the default files
func NewVaultDirectory(cnfg *Config) (files.Node, error) {
	dwnJSON, err := json.Marshal(cnfg)
	if err != nil {
		return nil, err
	}

	w := bytes.NewBuffer(nil)
	err = global.IndexFile().Render(context.Background(), w)
	if err != nil {
		return nil, err
	}
	fileMap := map[string]files.Node{
		"config.json": files.NewBytesFile(dwnJSON),
		"sw.js":       swJSFile,
		"app.wasm":    dwnWasmFile,
		"index.html":  files.NewBytesFile(w.Bytes()),
	}
	return files.NewMapDirectory(fileMap), nil
}

// Use IndexHTML template to generate the index file
func IndexHTMLFile() (files.Node, error) {
	w := bytes.NewBuffer(nil)
	err := global.IndexFile().Render(context.Background(), w)
	if err != nil {
		return nil, err
	}
	indexData := w.Bytes()
	return files.NewBytesFile(indexData), nil
}

// MarshalConfigFile uses the config template to generate the dwn config file
func MarshalConfigFile(c *Config) (files.Node, error) {
	dwnConfigData, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return files.NewBytesFile(dwnConfigData), nil
}
