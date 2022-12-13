package node

import (
	"context"
	"testing"
	"time"

	"github.com/libp2p/go-libp2p-core/network"
	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/stretchr/testify/assert"
)

func TestNodeNew(t *testing.T) {
	_, err := New(context.Background())
	assert.NoError(t, err, "Error creating node")
}

func TestNodeStreamHandler(t *testing.T) {
	// Create two nodes
	n1, err := New(context.Background())
	assert.NoError(t, err, "Error creating node")
	n2, err := New(context.Background())
	assert.NoError(t, err, "Error creating node")

	// Set a stream handler on n1
	n2.SetStreamHandler("/test/1.0.0", func(s network.Stream) {
		// Write to the stream
		_, err := s.Write([]byte("hello"))
		assert.NoError(t, err, "Error writing to stream")
	})

	// Open a stream on n2
	s, err := n1.NewStream(n2.Host().ID(), "/test/1.0.0")
	assert.NoError(t, err, "Error opening stream")

	// Read from the stream
	buf := make([]byte, 1024)
	_, err = s.Read(buf)
	assert.NoError(t, err, "Error reading from stream")
	assert.Equal(t, "hello", string(buf))
}

func TestNodeChannelSendReceive(t *testing.T) {
	// Create two nodes
	n1, err := New(context.Background())
	assert.NoError(t, err, "Error creating node")
	n2, err := New(context.Background())
	assert.NoError(t, err, "Error creating node")
	receivedChan := make(chan bool)
	receiveCount := 0

	// Create a channel on n1
	c1, err := n1.Join("test", WithOnMessage(func(msg *ps.Message) {
		t.Log("Received message on c2:", string(msg.Data))
		receivedChan <- true
	}))
	assert.NoError(t, err, "NewChannel")

	// Create a channel on n2
	c2, err := n2.Join("test", WithOnMessage(func(msg *ps.Message) {
		t.Log("Received message on c2:", string(msg.Data))
		receivedChan <- true
	}))
	assert.NoError(t, err, "NewChannel")

	// Send a message from n1 to n2
	err = c1.Send([]byte("hello"))
	assert.NoError(t, err, "should be able to send a message")

	// Send a message from n2 to n1
	err = c2.Send([]byte("hello"))
	assert.NoError(t, err, "should be able to send a message")

	// Await for the next message in channel on n2 using loop
	for {
		select {
		case <-receivedChan:
			receiveCount++
			if receiveCount == 1 {
				return
			}

		// Timeout
		case <-time.After(15 * time.Second):
			t.Fatal("timeout")
			return
		}
	}

}
