package host

import (
	"context"
	"testing"

	ps "github.com/libp2p/go-libp2p-pubsub"
)

func TestNewP2PHost(t *testing.T) {
	h1, err := New(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	h2, err := New(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if h1.PeerID() == h2.PeerID() {
		t.Fatal("Host IDs should be different")
	}
}

func TestPubsub(t *testing.T) {
	h1, err := New(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	h2, err := New(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	done := make(chan struct{})
	_, err = h1.Subscribe("test", func(msg *ps.Message) {
		t.Log("Got message:", string(msg.Data))
		if string(msg.Data) != "Hello World" {
			t.Fatal("Got invalid message")
			close(done)
		}
	})
	if err != nil {
		t.Fatal(err)
	}

	err = h2.Publish("test", []byte("Hello World"))
	if err != nil {
		t.Fatal(err)
	}

	<-done
}
