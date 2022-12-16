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

func TestPubSub(t *testing.T) {
	// Call Run method and check for panic (if any)
	node, err := New(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// Subscribe to a topic
	sub, err := node.Subscribe("test")
	if err != nil {
		t.Fatal(err)
	}

	// Publish a message to the topic
	err = sub.Publish([]byte("Hello World!"))
	if err != nil {
		t.Fatal(err)
	}

	// Get the message from the topic
	msg := <-sub.Messages()

	// Check if the message is the same as the one we published
	assert.Equal(t, []byte("Hello World!"), msg)
}

// TestMultiNodePubSub tests the pubsub functionality of multiple nodes
func TestMultiNodePubSub(t *testing.T) {
	// Call Run method and check for panic (if any)
	node1, err := New(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// Call Run method and check for panic (if any)
	node2, err := New(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// Connect the two nodes
	err = node1.Connect(node2.AddrInfo())
	if err != nil {
		t.Fatal(err)
	}

	// Subscribe to a topic
	_, err = node1.Subscribe("test")
	if err != nil {
		t.Fatal(err)
	}

	// Subscribe to a topic
	sub2, err := node2.Subscribe("test")
	if err != nil {
		t.Fatal(err)
	}

	// Publish a message to the topic
	err = node1.Publish("test", []byte("Hello World!"))
	if err != nil {
		t.Fatal(err)
	}

	// Get the message from the topic
	msg := <-sub2.Messages()

	// Check if the message is the same as the one we published
	assert.Equal(t, []byte("Hello World!"), msg)
}
