package crypto

import (
	"bytes"
	"fmt"
	"strings"
	"sync"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	stdtx "github.com/cosmos/cosmos-sdk/types/tx"
	rt "github.com/sonr-io/sonr/x/registry/types"

	"github.com/mr-tron/base58/base58"
	"github.com/sonr-io/sonr/pkg/client"
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

// Balances returns the balances of the given party.
func (w *MPCWallet) Balances() sdk.Coins {
	addr, err := w.Bech32Address()
	if err != nil {
		return nil
	}

	resp, err := client.CheckBalance(addr)
	if err != nil {
		return nil
	}
	fmt.Println("-- Check Balance --\n", resp)
	return resp
}

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
		pub = w.Config().Public[w.ID]
	} else if len(id) == 1 {
		pub = w.Config().Public[id[0]]
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

func (w *MPCWallet) BroadcastCreateWhoIs() error {
	addr, err := w.Bech32Address()
	if err != nil {
		return err
	}
	doc, err := w.DIDDocument()
	if err != nil {
		return err
	}

	docJSON, err := doc.MarshalJSON()
	if err != nil {
		return err
	}

	msg := rt.NewMsgCreateWhoIs(addr, docJSON, rt.WhoIsType_USER)
	msgBytes, err := msg.Marshal()
	if err != nil {
		return err
	}

	// Sign the transaction.
	tx, err := w.SignTx(msgBytes, fmt.Sprintf("%s/%s", msg.Route(), msg.Type()))
	if err != nil {
		return err
	}

	// Generate a JSON string.
	txBytes, err := tx.Marshal()
	if err != nil {
		return err
	}

	resp, err := client.BroadcastTx(txBytes)
	if err != nil {
		return err
	}
	fmt.Println("-- TX Response --\n", resp)
	return nil

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

// Generates an ECDSA signature for messageHash.
func (w *MPCWallet) SignTx(m []byte, typeUrl string) (*stdtx.Tx, error) {
	sig, err := w.Sign(m)
	if err != nil {
		return nil, err
	}

	pubKey, err := w.PublicKey()
	if err != nil {
		return nil, err
	}

	tx := stdtx.Tx{
		Body: &stdtx.TxBody{
			Messages: []*codectypes.Any{
				{
					TypeUrl: typeUrl,
					Value:   m,
				},
			},
		},
		AuthInfo: w.StdTxAuthInfo(2, pubKey),
		Signatures: [][]byte{
			ECDSASignatureToBytes(sig),
		},
	}
	return &tx, err
}

// StdTxAuthInfo returns the auth info for the given public key and fee.
func (w *MPCWallet) StdTxAuthInfo(fee int64, pubKey []byte) *stdtx.AuthInfo {
	return &stdtx.AuthInfo{
		SignerInfos: []*stdtx.SignerInfo{
			{
				PublicKey: codectypes.UnsafePackAny(pubKey),
				ModeInfo: &stdtx.ModeInfo{
					Sum: &stdtx.ModeInfo_Single_{},
				},
				Sequence: 1,
			},
		},
		Fee: &stdtx.Fee{
			Amount: sdk.NewCoins(sdk.NewCoin("snr", sdk.NewInt(fee))),
		},
	}
}

// Helper function to cmpSign a message.
func cmpSign(c *cmp.Config, m []byte, signers party.IDSlice, n *Network, pl *pool.Pool) (*ecdsa.Signature, error) {
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

func (w *MPCWallet) getCreateWhoIsMsg() ([]byte, error) {
	addr, err := w.Bech32Address()
	if err != nil {
		return nil, err
	}
	doc, err := w.DIDDocument()
	if err != nil {
		return nil, err
	}

	docJSON, err := doc.MarshalJSON()
	if err != nil {
		return nil, err
	}

	msg := rt.NewMsgCreateWhoIs(addr, docJSON, rt.WhoIsType_USER)
	msgBytes, err := msg.Marshal()
	if err != nil {
		return nil, err
	}
	return msgBytes, nil
}
