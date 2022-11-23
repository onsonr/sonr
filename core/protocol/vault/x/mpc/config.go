package mpc

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sonr-io/multi-party-sig/pkg/party"
)

type WalletOption func(*Wallet)

const keygenDefaultThreshold = 1

func WithThreshold(threshold int) WalletOption {
	return func(w *Wallet) {
		if threshold > 0 {
			w.threshold = threshold
		}
	}
}

func (w *Wallet) Apply(opts ...WalletOption) {
	for _, opt := range opts {
		opt(w)
	}
	if w.threshold == 0 {
		w.threshold = keygenDefaultThreshold
	}
}

func makeWallet(id party.ID, options ...WalletOption) *Wallet {
	w := &Wallet{
		id: id,
	}
	w.Apply(options...)
	return w
}

func (w *Wallet) HasPartyId() bool {
	return w.id != ""
}

func idFromFilename(filename string) party.ID {
	base := filepath.Base(filename)
	ptrs := strings.Split(base, "_")
	return party.ID(ptrs[0])
}

func (w *Wallet) fileName() string {
	if w.id == "" {
		panic("id not set")
	}
	return fmt.Sprintf("%s_sonr_id.json", w.id)
}

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}
