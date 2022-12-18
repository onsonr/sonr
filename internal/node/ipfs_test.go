package node

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAddGet(t *testing.T) {
	// Call Run method and check for panic (if any)
	node, err := New(context.Background())
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

	// Check if the file is the same as the one we added
	assert.Equal(t, []byte("Hello World!"), file)
}
