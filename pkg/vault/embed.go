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

// NewConfig uses the config template to generate the dwn config file
func NewConfig(keyshareJSON string, adddress string, chainID string, schema *dwn.Schema) *dwn.Config {
	dwnCfg := &dwn.Config{
		Motr:   createMotrConfig(keyshareJSON, adddress, "sonr.id"),
		Ipfs:   defaultIPFSConfig(),
		Sonr:   defaultSonrConfig(chainID),
		Schema: schema,
	}
	return dwnCfg
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

// Use DWNWasm template to generate the dwn wasm file
func DWNWasmFile() files.Node {
	return files.NewBytesFile(dwnWasmData)
}

// Use MotrMJS template to generate the dwn wasm file
func MotrMJSFile() files.Node {
	return files.NewBytesFile(motrMJSData)
}

// Use ServiceWorkerJS template to generate the service worker file
func SWJSFile() files.Node {
	return files.NewBytesFile(swJSData)
}

func createMotrConfig(keyshareJSON string, adddress string, origin string) *dwn.Motr {
	return &dwn.Motr{
		Keyshare: keyshareJSON,
		Address:  adddress,
		Origin:   origin,
	}
}

func defaultIPFSConfig() *dwn.IPFS {
	return &dwn.IPFS{
		ApiUrl:     "https://api.sonr-ipfs.land",
		GatewayUrl: "https://ipfs.sonr.land",
	}
}

func defaultSonrConfig(chainID string) *dwn.Sonr {
	return &dwn.Sonr{
		ApiUrl:       "https://api.sonr.land",
		GrpcUrl:      "https://grpc.sonr.land",
		RpcUrl:       "https://rpc.sonr.land",
		WebSocketUrl: "wss://rpc.sonr.land/ws",
		ChainId:      chainID,
	}
}
