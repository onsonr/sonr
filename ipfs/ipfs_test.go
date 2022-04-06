package ipfs

import (
	"context"
	"fmt"
	"os"
	"testing"

	iface "github.com/ipfs/interface-go-ipfs-core"
)

var nodeTemp iface.CoreAPI

func TestCreateTemp(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	node, err := SpawnEphemeral(ctx)
	nodeTemp = node
	if err != nil {
		t.Errorf("SpawnDefault(ctx) resulted in %d", err)
	}
	if node == nil {
		t.Errorf("SpawnDefault(ctx) resulted in nil result")
	}
}

func TestUploadTemp(t *testing.T) {
	//UploadData
}

func TestCreatePerm(t *testing.T) {
	os.Setenv("IPFS_PATH", "/Users/peytonthibodeaux/.ipfs")
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
