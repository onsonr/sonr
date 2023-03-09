package chain

import (
	"context"
	"testing"
)

func TestQueryLocalhostService(t *testing.T) {
	c := NewClient(SonrLocalRpcOrigin)
	s, err := c.GetService(context.Background(), "localhost")
	if err != nil {
		t.Error(err)
	}
	t.Log(s)
}

func TestQueryDevnetService(t *testing.T) {
	c := NewClient(SonrPublicRpcOrigin)
	s, err := c.GetAllServices(context.Background())
	if err != nil {
		t.Error(err)
	}
	t.Log(s)
}
