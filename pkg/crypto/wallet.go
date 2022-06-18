package crypto

import (
	"fmt"
	"strings"
	"sync"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	at "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/mr-tron/base58/base58"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
	rt "github.com/sonr-io/sonr/x/registry/types"
	"github.com/taurusgroup/multi-party-sig/pkg/ecdsa"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

type MPCWallet struct {
	pool      *pool.Pool
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

// Returns the Bech32 representation of the given party.
func (w *MPCWallet) Bech32Address(id ...party.ID) (string, error) {
	// c := types.GetConfig()
	// c.SetBech32PrefixForAccount("snr", "pub")
	pub, err := w.PublicKey()
	if err != nil {
		return "", err
	}
	str, err := bech32.ConvertAndEncode("snr", pub)
	if err != nil {
		return "", err
	}
	return str, nil
}

// Config returns the configuration of this wallet.
func (w *MPCWallet) Config() *cmp.Config {
	return w.Configs[w.ID]
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
		if !w.Config().PartyIDs().Contains(id) {
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
	for _, id := range w.Configs[w.ID].PartyIDs() {
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

// GetSigners returns the list of signers for the given message.
func (w *MPCWallet) GetSigners() party.IDSlice {
	signers := w.Configs[w.ID].PartyIDs()[:w.Threshold+1]
	if !signers.Contains(w.ID) {
		w.Network.Quit(w.ID)
		return nil
	}
	return party.NewIDSlice(signers)
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
	pub, err := w.PublicKeyBase58()
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
func (w *MPCWallet) PublicKey() ([]byte, error) {
	p := w.Config().PublicPoint().(*curve.Secp256k1Point)
	return p.MarshalBinary()
}

// Returns the ECDSA public key of the given party.
func (w *MPCWallet) PublicKeyBase58() (string, error) {
	pub, err := w.PublicKey()
	if err != nil {
		return "", err
	}
	return base58.Encode(pub), nil
}

func (w *MPCWallet) PublicKeyProto() (*rt.PubKey, error) {
	pubBz, err := w.PublicKey()
	if err != nil {
		return nil, err
	}
	return &rt.PubKey{
		Key: pubBz,
	}, nil
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
	handlerLoop(w.ID, hRefresh, w.Network)
	return r.(*cmp.Config), nil
}

// Generates an ECDSA signature for messageHash.
func (w *MPCWallet) Sign(m []byte) (*ecdsa.Signature, error) {
	var wg sync.WaitGroup
	signers := w.GetSigners()
	net := NewNetwork(signers)

	var (
		sig *ecdsa.Signature
		err error
	)

	cmpSign := func(c *cmp.Config, m []byte, signers party.IDSlice, n *Network, pl *pool.Pool) (*ecdsa.Signature, error) {
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

	for _, id := range signers {
		wg.Add(1)
		go func(id party.ID) {
			defer wg.Done()

			pl := pool.NewPool(0)
			defer pl.TearDown()

			if sig, err = cmpSign(w.Configs[id], m, signers, net, pl); err != nil {
				return
			}
		}(id)
	}
	wg.Wait()
	return sig, err
}

func (w *MPCWallet) SignTx(account *at.BaseAccount, authInfo *txtypes.AuthInfo, txBody *txtypes.TxBody) (*ecdsa.Signature, error) {
	// Serialize the transaction body.
	txBodyBz, err := txBody.Marshal()
	if err != nil {
		return nil, err
	}

	// Serialize the auth info.
	authInfoBz, err := authInfo.Marshal()
	if err != nil {
		return nil, err
	}

	// Create SignDoc
	signDoc := &txtypes.SignDoc{
		BodyBytes:     txBodyBz,
		AuthInfoBytes: authInfoBz,
		ChainId:       "sonr",
		AccountNumber: account.GetAccountNumber(),
	}

	// Serialize the sign doc.
	signDocBz, err := signDoc.Marshal()
	if err != nil {
		return nil, err
	}
	return w.Sign(signDocBz)
}
