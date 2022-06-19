package crypto

import (
	"fmt"
	"strings"
	"sync"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
)

type MPCWallet struct {
	pool        *pool.Pool
	ID          party.ID
	DID         did.DID
	DIDDocument did.Document
	Configs     map[party.ID]*cmp.Config
	Network     *Network
	Threshold   int
}

// GenerateWallet a new ECDSA private key shared among all the given participants.
func GenerateWallet(options ...WalletOption) (*MPCWallet, error) {
	opt := defaultConfig()
	w := opt.Apply(options...)

	var wg sync.WaitGroup
	for _, id := range opt.participants {
		wg.Add(1)
		go func(id party.ID) {
			pl := pool.NewPool(0)
			defer pl.TearDown()
			defer wg.Done()
			conf, err := cmpKeygen(id, opt.participants, w.Network, opt.threshold, pl)
			if err != nil {
				return
			}
			w.Configs[conf.ID] = conf
		}(id)
	}
	wg.Wait()

	addr, err := w.Address()
	if err != nil {
		return nil, err
	}

	baseDid, err := did.ParseDID(fmt.Sprintf("did:snr:%s", strings.TrimPrefix(addr, "snr")))
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

	// Add the DID Document to the wallet.
	w.DID = *baseDid
	w.DIDDocument = doc
	return w, nil
}

// Returns the Bech32 representation of the given party.
func (w *MPCWallet) Address(id ...party.ID) (string, error) {
	pub, err := w.PublicKeyProto()
	if err != nil {
		return "", err
	}

	str, err := bech32.ConvertAndEncode("snr", pub.Address().Bytes())
	if err != nil {
		return "", err
	}
	return str, nil
}

// Config returns the configuration of this wallet.
func (w *MPCWallet) Config() *cmp.Config {
	return w.Configs[w.ID]
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
	addr, err := w.Address()
	if err != nil {
		return nil, err
	}
	vmdid, err := did.ParseDID(fmt.Sprintf("did:snr:%s#%s", strings.TrimPrefix(addr, "snr"), id))
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
		Controller:      w.DID,
		PublicKeyBase58: pub,
	}, nil
}

// Marshal returns the JSON representation of the entire wallet.
func (w *MPCWallet) Marshal() ([]byte, error) {
	return w.Config().MarshalBinary()
}

// Returns the ECDSA public key of the given party.
func (w *MPCWallet) PublicKey() ([]byte, error) {
	p := w.Config().PublicPoint().(*curve.Secp256k1Point)
	buf, err := p.MarshalBinary()
	if err != nil {
		return nil, err
	}
	// Check length of the public key.
	if len(buf) != 33 {
		return nil, fmt.Errorf("invalid public key length")
	}
	return buf, nil
}

// Returns the ECDSA public key of the given party.
func (w *MPCWallet) PublicKeyBase58() (string, error) {
	pub, err := w.PublicKey()
	if err != nil {
		return "", err
	}
	return base58.Encode(pub), nil
}

func (w *MPCWallet) PublicKeyProto() (*secp256k1.PubKey, error) {
	pubBz, err := w.PublicKey()
	if err != nil {
		return nil, err
	}
	return &secp256k1.PubKey{
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

// SignTx constructs a TxRaw from the given message and signs it.
func (w *MPCWallet) SignTx(msg sdk.Msg) ([]byte, error) {
	txb, err := BuildTx(w, msg)
	if err != nil {
		return nil, err
	}

	ai, err := w.GetAuthInfoSingle(2)
	if err != nil {
		return nil, err
	}

	sigDocBz, err := GetSignDocBytes(ai, txb)
	if err != nil {
		return nil, err
	}

	sig, err := w.Sign(sigDocBz)
	if err != nil {
		return nil, err
	}
	return SerializeSignature(sig)
}

// Unmarshal unmarshals the given JSON into the wallet.
func (w *MPCWallet) Unmarshal(buf []byte) error {
	c := &cmp.Config{}
	if err := c.UnmarshalBinary(buf); err != nil {
		return err
	}
	w.Configs[c.ID] = c
	w.ID = c.ID
	w.Threshold = c.Threshold
	return nil
}

// Verifies an ECDSA signature for messageHash.
func (w *MPCWallet) Verify(m []byte, sig []byte) bool {
	edsig, err := SignatureFromBytes(sig)
	if err != nil {
		return false
	}
	mpcVerif := edsig.Verify(w.Config().PublicPoint(), m)
	return mpcVerif
}
