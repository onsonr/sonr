package ipfs

import (
	"context"
	"fmt"
	"testing"

	"github.com/sonrhq/core/x/vault/internal/node/config"
	"github.com/stretchr/testify/assert"
)

func TestNewAddGet(t *testing.T) {
	cnfg := config.DefaultConfig()
	node, err := Initialize(cnfg)
	if err != nil {
		t.Fatal(err)
	}

	// Add a file to the network
	cid, err := node.Add([]byte("Hello World!"))
	if err != nil {
		t.Fatal(err)
	}

	// Get the file from the network
	file, err := node.Get(cid)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("File: %s\n", file)
	fmt.Printf("CID: %s\n", cid)
	// Check if the file is the same as the one we added
	assert.Equal(t, []byte("Hello World!"), file)
}

func TestOrbitDB(t *testing.T) {
	cnfg := config.DefaultConfig()
	node, err := Initialize(cnfg)
	if err != nil {
		t.Fatal(err)
	}

	// Add a file to the network
	docsStore, err := node.LoadDocsStore("testDocStore", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Store Name: %v", docsStore.DBName())
	t.Logf("Store Identity: %v", docsStore.Identity().ID)
	t.Logf("Store Address: %v", docsStore.Address().String())
	t.Logf("Store Type: %v", docsStore.Type())

	testData := map[string]interface{}{
		"_id":  "0",
		"test": "test",
	}
	op, err := docsStore.Put(context.Background(), testData)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("op: %v", op.GetOperation())

	// Test KV Store
	kvStore, err := node.LoadKeyValueStore("testKVStore")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Store Name: %v", kvStore.DBName())
	t.Logf("Store Identity: %v", kvStore.Identity().ID)
	t.Logf("Store Address: %v", kvStore.Address().String())
	t.Logf("Store Type: %v", kvStore.Type())

	op, err = kvStore.Put(context.Background(), "test", []byte("test"))
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("op: %v", op.GetOperation())

	// Get the file from the network
	rawVal, err := docsStore.Get(context.Background(), "0", nil)
	if err != nil {
		t.Fatal(err)
	}
	val := rawVal[0].(map[string]interface{})
	assert.Equal(t, "test", val["test"])
}
