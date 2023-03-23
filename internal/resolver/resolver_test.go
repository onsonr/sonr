package resolver

import (
	"context"
	"testing"
)

func TestQueryLocalhostService(t *testing.T) {
	s, err := GetService(context.Background(), "localhost", SonrLocalRpcOrigin)
	if err != nil {
		t.Logf("error: %v. Its just a warning - probably offline", err)
	}
	t.Log(s)
}

func TestQueryDevnetService(t *testing.T) {
	s, err := GetService(context.Background(), "localhost", SonrPublicRpcOrigin)
	if err != nil {
		t.Error(err)
	}
	t.Log(s)
}
