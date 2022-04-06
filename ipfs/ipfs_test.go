package ipfs

import (
	"context"
	"fmt"
	"testing"

	iface "github.com/ipfs/interface-go-ipfs-core"
)

var nodeTemp iface.CoreAPI
var tempCID string

func TestCreateTemp(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	node, err := SpawnEphemeral(ctx)
	nodeTemp = node
	if err != nil {
		t.Error(err)
	}
	if node == nil {
		t.Errorf("SpawnDefault(ctx) resulted in nil result")
	}
}

func TestUploadTemp(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	data := []byte("Hello World!!!")
	resp, err := UploadData(ctx, data, nodeTemp)
	if err != nil {
		t.Errorf("UploadData([]byte, coreAPI) resulted in status %d", resp.Status)
		t.Error(err)
	}
	tempCID = resp.Cid
}

func TestCreatePerm(t *testing.T) {
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
