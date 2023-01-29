package dispatcher

import (
	"context"
	"sync"

	"github.com/sonrhq/core/pkg/crypto/wallet/accounts"
	"github.com/sonrhq/core/x/identity/protocol/vault/controller"
)

type Dispatcher struct {
	// n common.IPFSNode
	sync.Mutex
}

// NewDispatcher creates a new wallet dispatcher
func New() *Dispatcher {
	return &Dispatcher{
		// n: n,
	}
}

// BuildNewDIDController creates a new wallet
func (d *Dispatcher) BuildNewDIDController() (controller.DIDController, error) {
	// Lock the dispatcher
	d.Lock()
	defer d.Unlock()
	doneCh := make(chan controller.DIDController)
	errCh := make(chan error)

	// Create the wallet in a goroutine
	go func() {
		// The default shards that are added to the MPC wallet
		rootAcc, err := accounts.New()
		if err != nil {
			errCh <- err
		}
		control, err := controller.New(context.Background(), rootAcc)
		if err != nil {
			errCh <- err
		}
		doneCh <- control
	}()

	// Wait for the wallet to be created
	select {
	case w := <-doneCh:
		return w, nil
	case err := <-errCh:
		return nil, err
	}
}
