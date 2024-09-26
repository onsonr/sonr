package vault

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"

	"github.com/a-h/templ"
	"github.com/ipfs/boxo/files"
	"github.com/onsonr/sonr/config/dwn"
)

//go:embed app.wasm
var dwnWasmData []byte

//go:embed motr.mjs
var motrMJSData []byte

//go:embed sw.js
var swJSData []byte

var (
	dwnWasmFile = files.NewBytesFile(dwnWasmData)
	motrMJSFile = files.NewBytesFile(motrMJSData)
	swJSFile    = files.NewBytesFile(swJSData)
)

// NewVaultDirectory creates a new directory with the default files
func NewVaultDirectory(cnfg *dwn.Config) (files.Node, error) {
	dwnJSON, err := json.Marshal(cnfg)
	if err != nil {
		return nil, err
	}

	dwnStr, err := templ.JSONString(cnfg)
	if err != nil {
		return nil, err
	}
	w := bytes.NewBuffer(nil)
	err = indexFile(dwnStr).Render(context.Background(), w)
	if err != nil {
		return nil, err
	}
	fileMap := map[string]files.Node{
		"config.json": files.NewBytesFile(dwnJSON),
		"motr.mjs":    motrMJSFile,
		"sw.js":       swJSFile,
		"app.wasm":    dwnWasmFile,
		"index.html":  files.NewBytesFile(w.Bytes()),
	}
	return files.NewMapDirectory(fileMap), nil
}

// Use IndexHTML template to generate the index file
func IndexHTMLFile(c *dwn.Config) (files.Node, error) {
	str, err := templ.JSONString(c)
	if err != nil {
		return nil, err
	}
	w := bytes.NewBuffer(nil)
	err = indexFile(str).Render(context.Background(), w)
	if err != nil {
		return nil, err
	}
	indexData := w.Bytes()
	return files.NewBytesFile(indexData), nil
}

// MarshalConfigFile uses the config template to generate the dwn config file
func MarshalConfigFile(c *dwn.Config) (files.Node, error) {
	dwnConfigData, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return files.NewBytesFile(dwnConfigData), nil
}
