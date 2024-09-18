//go:build js && wasm
// +build js,wasm

package jsexc

type LocalStorageAPI interface {
	Get(key string) string
	Set(key string, value string)
	Remove(key string)
}

type SessionStorageAPI interface{}
