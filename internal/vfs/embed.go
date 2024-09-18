package vfs

import (
	_ "embed"
	"encoding/json"

	"github.com/ipfs/boxo/files"
	"github.com/onsonr/sonr/config/dwn"
)

//go:embed app.wasm
var dwnWasmData []byte

//go:embed index.html
var indexData []byte

//go:embed sw.js
var swJSData []byte

// NewDWNConfigFile uses the config template to generate the dwn config file
func NewDWNConfigFile(keyshareJSON string, adddress string) (files.Node, error) {
	dwnCfg := &dwn.Config{
		Keyshare: &keyshareJSON,
		Address:  &adddress,
		Ipfs:     defaultIPFSConfig(),
		Sonr:     defaultSonrConfig(),
	}
	dwnConfigData, err := json.Marshal(dwnCfg)
	if err != nil {
		return nil, err
	}
	return files.NewBytesFile(dwnConfigData), nil
}

// Use DWNWasm template to generate the dwn wasm file
func DWNWasmFile() files.Node {
	return files.NewBytesFile(dwnWasmData)
}

// Use IndexHTML template to generate the index file
func IndexFile() files.Node {
	return files.NewBytesFile(indexData)
}

// Use ServiceWorkerJS template to generate the service worker file
func SWJSFile() files.Node {
	return files.NewBytesFile(swJSData)
}

func defaultIPFSConfig() *dwn.IPFS {
	return &dwn.IPFS{
		ApiUrl:     "https://api.sonr-ipfs.land",
		GatewayUrl: "https://ipfs.sonr.land",
	}
}

func defaultSonrConfig() *dwn.Sonr {
	return &dwn.Sonr{
		ApiUrl:       "https://api.sonr.land",
		GrpcUrl:      "https://grpc.sonr.land",
		RpcUrl:       "https://rpc.sonr.land",
		WebSocketUrl: "wss://rpc.sonr.land/ws",
	}
}
