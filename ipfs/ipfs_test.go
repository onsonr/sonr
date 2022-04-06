package ipfs

import (
	"context"
	"fmt"
	"os"
	"testing"
)

// Test methods start with `Test`
func TestCreate(t *testing.T) {
	os.Setenv("IPFS_PATH", "")
	var ctx context.Context
	node, err := SpawnEphemeral(ctx)
	fmt.Println(err)
	if err != nil {
		t.Errorf("SpawnDefault(ctx) resulted in %d", err)
	}
	if node == nil {
		t.Errorf("SpawnDefault(ctx) resulted in nil result")
	}
}
