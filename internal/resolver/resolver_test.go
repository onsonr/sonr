package resolver

import (
	"context"
	"testing"
)

func TestQueryDevnetService(t *testing.T) {
	s, err := GetService(context.Background(), "localhost")
	if err != nil {
		t.Error(err)
	}
	t.Log(s)
}
