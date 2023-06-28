package blocker

import (
	"context"
	"fmt"

	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/internal/sfs"
	"github.com/sonrhq/core/pkg/mpc"
	"github.com/sonrhq/core/x/vault/types"
)

type Blocker interface {
	Pop() *types.ClaimableWallet
	Next(ctx context.Context)
}

func NewBlocker() Blocker {
	c := context.Background()
	s := &blocker{
		ctx:       c,
		jobsQueue: NewQueue(c, "WalletClaims"),
		results:   make([]*types.ClaimableWallet, 0),
		errCh:     make(chan error),
		doneCh:    make(chan *types.ClaimableWallet),
	}
	wk := NewWorker(s.jobsQueue)
	go s.run(wk)
	go s.handleProcess()
	return s
}

func (s *blocker) Pop() *types.ClaimableWallet {
	if len(s.results) == 0 {
		return nil
	}
	w := s.results[0]
	s.results = s.results[1:]
	return w
}

func (s *blocker) Next(ctx context.Context) {
	if s.results != nil && len(s.results) > 0 {
		return
	}
	if s.jobsQueue.PendingJobs() < 10 {
		job := Job{
			Name:   "Build Claimable Wallet",
			Action: s.buildClaimableWallet,
			Ctx:    ctx,
		}
		s.jobsQueue.AddJob(job)
	}
	return
}

type blocker struct {
	jobsQueue *Queue
	results   []*types.ClaimableWallet
	errCh     chan error
	doneCh    chan *types.ClaimableWallet
	ctx       context.Context
}

func (s *blocker) run(w *Worker) {
	for {
		done := w.DoWork(s.errCh)
		if done {
			fmt.Println("Worker done")
			break
		}
		continue
	}
}

func (s *blocker) buildClaimableWallet(ctx context.Context) error {
	// Call Handler for keygen
	confs, err := mpc.Keygen(crypto.PartyID("ucw-1"), mpc.WithPeers("ucw-2"), mpc.WithHandlers())
	if err != nil {
		s.errCh <- err
		return err
	}

	did, err := mpc.GetDIDAddress(confs[0], crypto.SONRCoinType)
	if err != nil {
		s.errCh <- err
		return err
	}

	for i, conf := range confs {
		ksb, err := conf.MarshalBinary()
		if err != nil {
			s.errCh <- err
			return err
		}
		ks, err := types.NewKeyshare(ksb, crypto.SONRCoinType, types.SetUnclaimed(i+1))
		if err != nil {
			s.errCh <- err
			return err
		}
		err = sfs.InsertUnclaimedKeyshare(ks, crypto.SONRCoinType, i+1)
		if err != nil {
			s.errCh <- err
			return err
		}
	}
	vaddr, _ := local.ValidatorAddress()
	cw, err := types.NewWalletClaims(vaddr, did)
	if err != nil {
		s.errCh <- err
		return err
	}
	s.doneCh <- cw
	return nil
}

func (s *blocker) handleProcess() {
	for {
		select {
		case cw := <-s.doneCh:
			s.results = append(s.results, cw)
		case err := <-s.errCh:
			panic(err)
		}
	}
}
