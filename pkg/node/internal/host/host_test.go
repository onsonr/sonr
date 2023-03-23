package host

import (
	"context"
	"testing"

	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonrhq/core/pkg/node/config"
)

func TestNewP2PHost(t *testing.T) {
	ctx, err := config.NewContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	config := config.DefaultConfig(ctx)
	h1, err := Initialize(config)
	if err != nil {
		t.Fatal(err)
	}
	h2, err := Initialize(config)
	if err != nil {
		t.Fatal(err)
	}

	if h1.PeerID() == h2.PeerID() {
		t.Fatal("Host IDs should be different")
	}
}

func TestPubsub(t *testing.T) {
	ctx, err := config.NewContext(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	config := config.DefaultConfig(ctx)
	h1, err := Initialize(config)
	if err != nil {
		t.Fatal(err)
	}
	h2, err := Initialize(config)
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
