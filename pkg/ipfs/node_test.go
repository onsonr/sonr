package ipfs

import (
	"context"
	"fmt"
	"testing"
	"time"

	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/options"
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

func TestBasicPubSub(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Call Run method and check for panic (if any)
	node1, err := New(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Call Run method and check for panic (if any)
	node2, err := New(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// Connect the two nodes
	err = node2.Connect(node1.MultiAddr())
	if err != nil {
		t.Fatal(err)
	}

	sub, err := node2.PubSub().Subscribe(ctx, "testch")
	if err != nil {
		t.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		defer close(done)

		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		for {
			err := node2.PubSub().Publish(ctx, "testch", []byte("hello world"))
			switch err {
			case nil:
			case context.Canceled:
				return
			default:
				t.Error(err)
				cancel()
				return
			}
			select {
			case <-ticker.C:
			case <-ctx.Done():
				return
			}
		}
	}()

	// Wait for the sender to finish before we return.
	// Otherwise, we can get random errors as publish fails.
	defer func() {
		cancel()
		<-done
	}()

	m, err := sub.Next(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if string(m.Data()) != "hello world" {
		t.Errorf("got invalid data: %s", string(m.Data()))
	}

	self1, err := node2.Key().Self(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if m.From() != self1.ID() {
		t.Errorf("m.From didn't match")
	}

	peers, err := node1.PubSub().Peers(ctx, options.PubSub.Topic("testch"))
	if err != nil {
		t.Fatal(err)
	}

	if len(peers) != 1 {
		t.Fatalf("got incorrect number of peers: %d", len(peers))
	}

	self0, err := node2.Key().Self(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if peers[0] != self0.ID() {
		t.Errorf("peer didn't match")
	}

	peers, err = node2.PubSub().Peers(ctx, options.PubSub.Topic("nottestch"))
	if err != nil {
		t.Fatal(err)
	}

	if len(peers) != 0 {
		t.Fatalf("got incorrect number of peers: %d", len(peers))
	}

	topics, err := node2.PubSub().Ls(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if len(topics) != 1 {
		t.Fatalf("got incorrect number of topics: %d", len(peers))
	}

	if topics[0] != "testch" {
		t.Errorf("topic didn't match")
	}
}

func TestGroupPubSub(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	n1, err := New(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	n2, err := New(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	n3, err := New(ctx)
	if err != nil {
		t.Fatal(err)
	}

	fromKey, err := n3.Key().Self(ctx)
	if err != nil {
		t.Fatal(err)
	}

	err = n3.Connect(n1.MultiAddr(), n2.MultiAddr())
	if err != nil {
		t.Fatal(err)
	}

	sub1, err := n1.PubSub().Subscribe(ctx, "testch")
	if err != nil {
		t.Fatal(err)
	}

	sub2, err := n2.PubSub().Subscribe(ctx, "testch")
	if err != nil {
		t.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		defer close(done)

		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		for {
			err := n3.PubSub().Publish(ctx, "testch", []byte("hello world"))
			switch err {
			case nil:
			case context.Canceled:
				return
			default:
				t.Error(err)
				cancel()
				return
			}
			select {
			case <-ticker.C:
			case <-ctx.Done():
				return
			}
		}
	}()

	// Wait for the sender to finish before we return.
	// Otherwise, we can get random errors as publish fails.
	defer func() {
		cancel()
		<-done
	}()

	// Check message content and verify that it came from the correct peer
	m1, err := sub1.Next(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if string(m1.Data()) != "hello world" {
		t.Errorf("got invalid data: %s", string(m1.Data()))
	}
	if m1.From() != fromKey.ID() {
		t.Errorf("m.From didn't match for first node message")
	}

	// Check for the same message on the other node
	m2, err := sub2.Next(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if string(m2.Data()) != "hello world" {
		t.Errorf("got invalid data: %s", string(m2.Data()))
	}

	if m2.From() != fromKey.ID() {
		t.Errorf("m.From didn't match for second node message")
	}
}

func TestSubPubNode(t *testing.T) {
	ctx := context.Background()
	node1, err := New(ctx)
	if err != nil {
		t.Fatal(err)
	}

	node2, err := New(ctx)
	if err != nil {
		t.Fatal(err)
	}

	err = node2.Connect(node1.MultiAddr())
	if err != nil {
		t.Fatal(err)
	}
	inMsgCount := 0
	done := make(chan struct{})
	err = node1.Subscribe(ctx, "testch", func(topic string, msg icore.PubSubMessage) error {
		if topic != "testch" {
			t.Errorf("got invalid topic: %s", topic)
		}
		if string(msg.Data()) != "hello world" {
			t.Errorf("got invalid data: %s", string(msg.Data()))
		}
		inMsgCount++
		if inMsgCount == 3 {
			fmt.Println("got 3 messages, cancelling context")
			done <- struct{}{}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	err = node2.Publish("testch", []byte("hello world"))
	if err != nil {
		t.Fatal(err)
	}
	err = node2.Publish("testch", []byte("hello world"))
	if err != nil {
		t.Fatal(err)
	}
	err = node2.Publish("testch", []byte("hello world"))
	if err != nil {
		t.Fatal(err)
	}
	<-done
}
