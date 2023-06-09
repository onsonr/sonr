package blocker

import (
	"context"
	"fmt"

	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/x/vault/internal/mpc"
	"github.com/sonrhq/core/x/vault/internal/sfs"
	"github.com/sonrhq/core/x/vault/types"
)

type Blocker interface {
	Pop() *types.ClaimableWallet
	Next()
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

func (s *blocker) Next() {
	if s.results != nil && len(s.results) > 0 {
		return
	}
	if s.jobsQueue.PendingJobs() < 10 {
		job := Job{
			Name:   "Build Claimable Wallet",
			Action: s.buildClaimableWallet,
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

func (s *blocker) buildClaimableWallet() error {
	// Call Handler for keygen
	confs, err := mpc.Keygen(crypto.PartyID("current"))
	if err != nil {
		s.errCh <- err
		return err
	}

	var kss []types.KeyShare
	for _, conf := range confs {
		ksb, err := conf.MarshalBinary()
		if err != nil {
			s.errCh <- err
			return err
		}
		ks, err := types.NewKeyshare(string(conf.ID), ksb, crypto.SONRCoinType)
		if err != nil {
			s.errCh <- err
			return err
		}

		go sfs.InsertKeyshare(ks)
		kss = append(kss, ks)
	}
	vaddr, _ := local.ValidatorAddress()
	cw, err := types.NewWalletClaims(vaddr, kss)
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
			fmt.Println(err)
		}
	}
}
