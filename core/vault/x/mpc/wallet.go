package mpc

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	crypto_pb "github.com/libp2p/go-libp2p/core/crypto/pb"
	"github.com/sonr-io/multi-party-sig/pkg/ecdsa"
	"github.com/sonr-io/multi-party-sig/pkg/math/curve"
	"github.com/sonr-io/multi-party-sig/pkg/party"
	"github.com/sonr-io/multi-party-sig/pkg/pool"
	"github.com/sonr-io/multi-party-sig/protocols/cmp"
	"github.com/sonr-hq/sonr/internal/node"
)

type Wallet struct {
	cmp.Config
	id        party.ID
	node      node.Node
	threshold int
}

func NewWallet(n node.Node, ids node.IDSlice, options ...WalletOption) (*Wallet, error) {
	id, err := n.ID()
	if err != nil {
		return nil, err
	}

	w := makeWallet(id.GetPartyID(), options...)
	w.node = n
	ch, err := n.Join("test")
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	pl := pool.NewPool(8)
	defer pl.TearDown()
	conf, err := cmpKeygen(id.GetPartyID(), ids, ch, w.threshold, &wg, pl)
	if err != nil {
		panic(err)
	}
	w.Config = *conf
	return w, nil
}

func NewWalletFromDisk(p string, password string) (*Wallet, error) {
	id := idFromFilename(p)
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return nil, err
	}
	lockedBz, err := os.ReadFile(p)
	if err != nil {
		fmt.Println("os.ReadFile:", err)
		return nil, err
	}
	bz, err := AesDecryptWithPassword(password, lockedBz)
	if err != nil {
		fmt.Println("AesDecryptWithPassword:", err)
		return nil, err
	}
	conf := cmp.EmptyConfig(curve.Secp256k1{})
	err = conf.UnmarshalBinary(bz)
	if err != nil {
		return nil, err
	}
	w := makeWallet(id)
	w.Config = *conf
	return w, nil
}

func (w *Wallet) Save(password string) (string, error) {
	p := filepath.Join(os.Getenv("HOME"), ".sonr", "wallet", w.fileName())
	cnfg, err := w.Config.MarshalBinary()
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return "", err
	}
	f, err := os.Create(p)
	if err != nil {
		return "", err
	}
	bz, err := AesEncryptWithPassword(password, cnfg)
	if _, err := f.Write(bz); err != nil {
		return "", err
	}
	return p, nil
}

// func (w *Wallet) Refresh() (*Wallet, error) {
// 	ch, err := w.node.Join("test")
// 	if err != nil {
// 		return nil, err
// 	}
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go func(channel *node.Channel) {
// 		pl := pool.NewPool(0)
// 		defer pl.TearDown()
// 		conf, err := cmpRefresh(&w.Config, channel, &wg, pl)
// 		if err != nil {
// 			return
// 		}
// 		w.Config = *conf
// 	}(ch)
// 	wg.Wait()
// 	return w, nil
// }

func (w *Wallet) GetPreSignature() (*ecdsa.PreSignature, error) {
	ch, err := w.node.Join("test")
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	var preSignature *ecdsa.PreSignature
	wg.Add(1)

	go func(channel *node.Channel) {
		pl := pool.NewPool(0)
		defer pl.TearDown()
		res, err := cmpPreSign(&w.Config, w.PartyIDs(), channel, &wg, pl)
		if err != nil {
			return
		}
		preSignature = res
	}(ch)
	wg.Wait()
	return preSignature, nil
}

func (w *Wallet) SignWithPreSignature(m []byte, preSignature *ecdsa.PreSignature) ([]byte, error) {
	ch, err := w.node.Join("test")
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	var signature *ecdsa.Signature
	wg.Add(1)
	go func(channel *node.Channel) {
		pl := pool.NewPool(0)
		defer pl.TearDown()
		res, err := cmpPreSignOnline(&w.Config, preSignature, m, channel, &wg, pl)
		if err != nil {
			return
		}
		signature = res
	}(ch)

	wg.Wait()
	return SerializeSignature(signature)
}

func (w *Wallet) Sign(m []byte) ([]byte, error) {
	ch, err := w.node.Join("test")
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	pl := pool.NewPool(0)
	defer pl.TearDown()
	wg.Add(1)
	res, err := cmpSign(&w.Config, m, w.PartyIDs(), ch, &wg, pl)
	if err != nil {
		return nil, err
	}
	wg.Wait()
	return SerializeSignature(res)
}

func (w *Wallet) Verify(data []byte, sig []byte) (bool, error) {
	signature, err := DeserializeSignature(sig)
	if err != nil {
		return false, err
	}
	return signature.Verify(w.PublicPoint(), data), nil
}

func (p *Wallet) Type() crypto_pb.KeyType {
	return crypto_pb.KeyType_Secp256k1
}

func (p *Wallet) Raw() ([]byte, error) {
	return p.Config.MarshalBinary()
}
