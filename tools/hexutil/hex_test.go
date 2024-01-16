package hexutil

import (
	"testing"
)

func TestDecodeHexString(t *testing.T) {
	bytes, err := DecodeHexString("0xe33ef3d7883cd3f6b9c2a72b916c36066cca8443c718fb53bc0a0607de9e4d9a")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(bytes)
}

func TestDecodeHexString2(t *testing.T) {
	bytes, err := DecodeHexString("e33ef3d7883cd3f6b9c2a72b916c36066cca8443c718fb53bc0a0607de9e4d9a")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(bytes)
}
