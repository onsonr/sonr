package crypto

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/taurusgroup/multi-party-sig/pkg/ecdsa"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

type MPCWallet struct {
	Config    *cmp.Config
	Network   *Network
	Threshold int
}

// Generate a new ECDSA private key shared among all the given participants.
func Generate(options ...WalletOption) (*MPCWallet, error) {
	opt := defaultConfig()
	wallet := opt.Apply(options...)

	var wg sync.WaitGroup
	for _, id := range opt.participants {
		wg.Add(1)
		go func(id party.ID) {
			pl := pool.NewPool(0)
			defer pl.TearDown()
			defer wg.Done()
			if err := wallet.Keygen(id, opt.participants, pl); err != nil {
				fmt.Println(err)
				return
			}
			buf, err := wallet.Sign([]byte("Test"), pl)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(buf)
		}(id)
	}
	wg.Wait()
	return wallet, nil
}

// Returns the cosmos compatible address of the given party.
func (w *MPCWallet) AccountAddress() (types.AccAddress, error) {
	c := types.GetConfig()
	c.SetBech32PrefixForAccount("snr", "pub")
	bechAddr, err := w.Bech32Address()
	if err != nil {
		return nil, err
	}
	acc, err := types.AccAddressFromBech32(bechAddr)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

// Returns the Bech32 representation of the given party.
func (w *MPCWallet) Bech32Address() (string, error) {
	c := types.GetConfig()
	c.SetBech32PrefixForAccount("snr", "pub")
	pub, err := w.PublicKey()
	if err != nil {
		return "", err
	}
	str, err := bech32.ConvertAndEncode("snr", pub)
	if err != nil {
		return "", err
	}
	fmt.Println(str)
	return str, nil
}

// GetSigners returns the list of signers for the given message.
func (w *MPCWallet) GetSigners() party.IDSlice {
	signers := w.Config.PartyIDs()
	if !signers.Contains(w.Config.ID) {
		w.Network.Quit(w.Config.ID)
		return nil
	}
	return party.NewIDSlice(signers)
}

// Generate a new ECDSA private key shared among all the given participants.
func (w *MPCWallet) Keygen(id party.ID, ids party.IDSlice, pl *pool.Pool) error {
	h, err := protocol.NewMultiHandler(cmp.Keygen(curve.Secp256k1{}, id, ids, w.Threshold, pl), nil)
	if err != nil {
		return err
	}

	handlerLoop(id, h, w.Network)
	r, err := h.Result()
	if err != nil {
		return err
	}

	conf := r.(*cmp.Config)
	w.Config = conf
	return nil
}

// Returns the ECDSA public key of the given party.
func (w *MPCWallet) PublicKey() ([]byte, error) {
	pub := w.Config.Public[w.Config.ID]
	if pub == nil {
		return nil, fmt.Errorf("no public key found")
	}
	buffer := bytes.NewBuffer(nil)
	_, err := pub.WriteTo(buffer)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	buf := address.Hash("snr", buffer.Bytes())
	return buf, nil
}

// Refreshes all shares of an existing ECDSA private key.
func (w *MPCWallet) Refresh(pl *pool.Pool) (*cmp.Config, error) {
	hRefresh, err := protocol.NewMultiHandler(cmp.Refresh(w.Config, pl), nil)
	if err != nil {
		return nil, err
	}
	handlerLoop(w.Config.ID, hRefresh, w.Network)

	r, err := hRefresh.Result()
	if err != nil {
		return nil, err
	}

	return r.(*cmp.Config), nil
}

// Generates an ECDSA signature for messageHash.
func (w *MPCWallet) Sign(m []byte, pl *pool.Pool) (*ecdsa.Signature, error) {
	
	signers := w.GetSigners()
	if len(signers) == 0 {
		return nil, fmt.Errorf("no signers found")
	}
	h, err := protocol.NewMultiHandler(cmp.Sign(w.Config, signers, m, pl), nil)
	if err != nil {
		return nil, err
	}
	handlerLoop(w.Config.ID, h, w.Network)
	signResult, err := h.Result()
	if err != nil {
		return nil, err
	}
	signature := signResult.(*ecdsa.Signature)
	if !signature.Verify(w.Config.PublicPoint(), m) {
		return nil, fmt.Errorf("failed to verify cmp signature")
	}
	fmt.Println("cmp signature verified: ", signature)
	return signature, nil
}
