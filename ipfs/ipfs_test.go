package ipfs

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestCreateTemp(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	node, err := SpawnEphemeral(ctx)
	fmt.Println(err)
	if err != nil {
		t.Errorf("SpawnDefault(ctx) resulted in %d", err)
	}
	if node == nil {
		t.Errorf("SpawnDefault(ctx) resulted in nil result")
	}
}

func TestCreatePerm(t *testing.T) {
	os.Setenv("IPFS_PATH", "")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	node, err := SpawnDefault(ctx)
	fmt.Println(err)
	if err != nil {
		t.Error(err)
	}
	if node == nil {
		t.Errorf("SpawnDefault(ctx) resulted in nil result")
	}
}
