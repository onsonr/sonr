package mpc

import (
	"context"
	"sync"
	"testing"

	"github.com/sonr-hq/sonr/pkg/node"
)

func TestCMPKeygen(t *testing.T) {
	ctx := context.Background()
	n1, err := node.New(ctx)
	if err != nil {
		t.Fatal(err)
	}

	n2, err := node.New(ctx)
	if err != nil {
		t.Fatal(err)
	}

	_, err = n1.CreateNetwork(ctx, n2.ID())
	if err != nil {
		t.Fatal(err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		// _, err = CMPKeygen(net, wg)
		// if err != nil {
		// 	return
		// }
	}()
	wg.Wait()
}
