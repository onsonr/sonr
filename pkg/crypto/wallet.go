package crypto

import (
	"bytes"
	"fmt"
	"strings"
	"sync"

	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/mr-tron/base58/base58"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
	"github.com/taurusgroup/multi-party-sig/pkg/ecdsa"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp/config"
)

type MPCWallet struct {
	pool *pool.Pool

	ID        party.ID
	Configs   map[party.ID]*cmp.Config
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
				return
			}
		}(id)
	}
	wg.Wait()
	return wallet, nil
}

// // Returns the cosmos compatible address of the given party.
// func (w *MPCWallet) AccountAddress(id ...party.ID) (types.AccAddress, error) {
// 	// c := types.GetConfig()
// 	// c.SetBech32PrefixForAccount("snr", "pub")
// 	bechAddr, err := w.Bech32Address(id...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	acc, err := types.AccAddressFromBech32(bechAddr)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return acc, nil
// }

// Returns the Bech32 representation of the given party.
func (w *MPCWallet) Bech32Address(id ...party.ID) (string, error) {
	// c := types.GetConfig()
	// c.SetBech32PrefixForAccount("snr", "pub")
	pub, err := w.PublicKey(id...)
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

// DID Returns the DID Address of the Wallet. When partyId is provided, it returns the DID of the given party. Only the first party in the wallet can create a DID.
func (w *MPCWallet) DID(party ...party.ID) (*did.DID, error) {
	if len(party) == 0 {
		addr, err := w.Bech32Address()
		if err != nil {
			return nil, err
		}
		return did.ParseDID(fmt.Sprintf("did:snr:%s", strings.TrimPrefix(addr, "snr")))
	} else if len(party) == 1 {
		id := party[0]
		if !w.Config.PartyIDs().Contains(id) {
			return nil, fmt.Errorf("party %s is not a member of the wallet", id)
		}

		baseDid, err := w.DID()
		if err != nil {
			return nil, err
		}
		return did.ParseDID(fmt.Sprintf("%s#%s", baseDid, id))
	}
	return nil, fmt.Errorf("invalid number of arguments")
}

// CreateDIDDocument creates a DID Document for the given party.
func (w *MPCWallet) DIDDocument() (did.Document, error) {
	// Get the DID of the wallet
	baseDid, err := w.DID()
	if err != nil {
		return nil, err
	}

	// Create the DID Document
	doc, err := did.NewDocument(baseDid.String())
	if err != nil {
		return nil, err
	}

	// Get ALL the VerificationMethods of the wallet.
	vmsAll := make([]*did.VerificationMethod, 0)
	for _, id := range w.Config.PartyIDs() {
		vm, err := w.GetVerificationMethod(id)
		if err != nil {
			return nil, err
		}
		vmsAll = append(vmsAll, vm)
	}

	for _, vm := range vmsAll {
		doc.AddAuthenticationMethod(vm)
	}

	if len(doc.GetAuthenticationMethods()) != len(vmsAll) {
		return nil, fmt.Errorf("failed to add all verification methods to DID Document")
	}
	return doc, nil
}

// GetVerificationMethod returns the VerificationMethod for the given party.
func (w *MPCWallet) GetVerificationMethod(id party.ID) (*did.VerificationMethod, error) {
	// Get the DID of the wallet
	baseDid, err := w.DID()
	if err != nil {
		return nil, err
	}

	// Get the DID of the party.
	vmdid, err := w.DID(id)
	if err != nil {
		return nil, err
	}

	// Get base58 encoded public key.
	pub, err := w.PublicKeyBase58(id)
	if err != nil {
		return nil, err
	}

	// Return the shares VerificationMethod
	return &did.VerificationMethod{
		ID:              *vmdid,
		Type:            ssi.ECDSASECP256K1VerificationKey2019,
		Controller:      *baseDid,
		PublicKeyBase58: pub,
	}, nil
}

// GetSigners returns the list of signers for the given message.
func (w *MPCWallet) GetSigners() party.IDSlice {
	signers := w.Configs[w.ID].PartyIDs()[:w.Threshold+1]
	if !signers.Contains(w.ID) {
		w.Network.Quit(w.ID)
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
	w.Configs[conf.ID] = conf
	return nil
}

// Returns the ECDSA public key of the given party.
func (w *MPCWallet) PublicKey(id ...party.ID) ([]byte, error) {
	var pub *config.Public
	if len(id) == 0 {
		pub = w.Config.Public[w.Config.ID]
	} else if len(id) == 1 {
		pub = w.Config.Public[id[0]]
	} else {
		return nil, fmt.Errorf("invalid number of arguments")
	}
	if pub == nil {
		return nil, fmt.Errorf("no public key found")
	}
	buffer := bytes.NewBuffer(nil)
	_, err := pub.WriteTo(buffer)
	if err != nil {
		return nil, err
	}
	buf := address.Hash("snr", buffer.Bytes())
	return buf, nil
}

// Returns the ECDSA public key of the given party.
func (w *MPCWallet) PublicKeyBase58(id ...party.ID) (string, error) {
	pub, err := w.PublicKey(id...)
	if err != nil {
		return "", err
	}
	return base58.Encode(pub), nil
}

// Refreshes all shares of an existing ECDSA private key.
func (w *MPCWallet) Refresh(pl *pool.Pool) (*cmp.Config, error) {
	hRefresh, err := protocol.NewMultiHandler(cmp.Refresh(w.Configs[w.ID], pl), nil)
	if err != nil {
		return nil, err
	}
	handlerLoop(w.ID, hRefresh, w.Network)

	r, err := hRefresh.Result()
	if err != nil {
		return nil, err
	}

	return r.(*cmp.Config), nil
}

// Generates an ECDSA signature for messageHash.
func sign(c *cmp.Config, m []byte, signers party.IDSlice, n *Network, pl *pool.Pool) (*ecdsa.Signature, error) {
	h, err := protocol.NewMultiHandler(cmp.Sign(c, signers, m, pl), nil)
	if err != nil {
		return nil, err
	}
	handlerLoop(c.ID, h, n)
	signResult, err := h.Result()
	if err != nil {
		return nil, err
	}
	signature := signResult.(*ecdsa.Signature)
	if !signature.Verify(c.PublicPoint(), m) {
		return nil, fmt.Errorf("failed to verify cmp signature")
	}
	return signature, nil
}

func (w *MPCWallet) Sign(m []byte /*, pl *pool.Pool*/) (*ecdsa.Signature, error) {
	var wg sync.WaitGroup
	signers := w.GetSigners()
	net := NewNetwork(signers)

	var (
		sig *ecdsa.Signature
		err error
	)

	for _, id := range signers {
		wg.Add(1)
		go func(id party.ID) {
			defer wg.Done()

			pl := pool.NewPool(0)
			defer pl.TearDown()

			if sig, err = sign(w.Configs[id], m, signers, net, pl); err != nil {
				return
			}
		}(id)
	}
	wg.Wait()

	return sig, err
}
