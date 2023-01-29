package ipfs

import (
	"context"
	"fmt"
	"testing"

	"github.com/sonrhq/core/pkg/common"
	"github.com/sonrhq/core/pkg/node/config"
	"github.com/stretchr/testify/assert"
)

func TestNewAddGet(t *testing.T) {
	// Call Run method and check for panic (if any)
	ctx, err := common.NewContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	cnfg := config.DefaultConfig(ctx)
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
	ctx, err := common.NewContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	cnfg := config.DefaultConfig(ctx)
	node, err := Initialize(cnfg)
	if err != nil {
		t.Fatal(err)
	}

	docsStore, err := node.LoadDocsStore("test")
	if err != nil {
		t.Fatal(err)
	}

	testData := map[string]interface{}{
		"_id":  "0",
		"test": "test",
	}
	op, err := docsStore.Put(context.Background(), testData)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("op: %v", op)

	// Get the file from the network
	rawVal, err := docsStore.Get(context.Background(), "0", nil)
	if err != nil {
		t.Fatal(err)
	}
	val := rawVal[0].(map[string]interface{})
	assert.Equal(t, "test", val["test"])
}
