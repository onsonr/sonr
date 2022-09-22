package mpc

import (
	"errors"
	"fmt"
	"sync"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/mr-tron/base58/base58"
	"github.com/sonr-io/multi-party-sig/pkg/ecdsa"
	"github.com/sonr-io/multi-party-sig/pkg/math/curve"
	"github.com/sonr-io/multi-party-sig/pkg/party"
	"github.com/sonr-io/multi-party-sig/pkg/pool"
	"github.com/sonr-io/multi-party-sig/pkg/protocol"
	"github.com/sonr-io/multi-party-sig/protocols/cmp"
	"github.com/sonr-io/sonr/third_party/types/common"
)

type Wallet struct {
	pool *pool.Pool
	ID   party.ID

	Configs   map[party.ID]*cmp.Config
	Network   *Network
	Threshold int
}

// GenerateWallet a new ECDSA private key shared among all the given participants.
func GenerateWallet(cb common.MotorCallback, options ...WalletOption) (*Wallet, error) {
	opt := defaultConfig()
	w := opt.Apply(options...)

	var wg sync.WaitGroup
	for _, id := range opt.participants {
		wg.Add(1)
		go func(id party.ID) {
			pl := pool.NewPool(0)
			defer pl.TearDown()
			conf, err := cmpKeygen(id, opt.participants, w.Network, opt.threshold, &wg, pl)
			if err != nil {
				return
			}
			w.Configs[conf.ID] = conf
		}(id)
	}
	wg.Wait()
	// Add the DID Document to the wallet.
	return w, nil
}

// Returns the Bech32 representation of the given party.
func (w *Wallet) Address(id ...party.ID) (string, error) {
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
func (w *Wallet) Config() *cmp.Config {
	return w.Configs[w.ID]
}

// GetSigners returns the list of signers for the given message.
func (w *Wallet) GetSigners() party.IDSlice {
	signers := party.IDSlice([]party.ID{"dsc", "psk"})
	// signers := w.Configs[w.ID].PartyIDs()[:w.Threshold+1]
	if !signers.Contains(w.ID) {
		w.Network.Quit(w.ID)
		return nil
	}
	return party.NewIDSlice(signers)
}

// Marshal returns the JSON representation of the entire wallet.
func (w *Wallet) Marshal() ([]byte, error) {
	return w.Config().MarshalBinary()
}

// Returns the ECDSA public key of the given party.
func (w *Wallet) PublicKey() ([]byte, error) {
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
func (w *Wallet) PublicKeyBase58() (string, error) {
	pub, err := w.PublicKey()
	if err != nil {
		return "", err
	}
	return base58.Encode(pub), nil
}

func (w *Wallet) PublicKeyProto() (*secp256k1.PubKey, error) {
	pubBz, err := w.PublicKey()
	if err != nil {
		return nil, err
	}
	return &secp256k1.PubKey{
		Key: pubBz,
	}, nil
}

// Refreshes all shares of an existing ECDSA private key.
func (w *Wallet) Refresh(pl *pool.Pool) (*cmp.Config, error) {
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
func (w *Wallet) Sign(m []byte) (*ecdsa.Signature, error) {
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
			pl := pool.NewPool(0)
			defer pl.TearDown()
			if sig, err = cmpSign(w.Configs[id], m, signers, net, &wg, pl); err != nil {
				return
			}
		}(id)
	}
	wg.Wait()
	return sig, err
}

// Unmarshal unmarshals the given JSON into the wallet.
func (w *Wallet) Unmarshal(buf []byte) error {
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
func (w *Wallet) Verify(m []byte, sig []byte) bool {
	edsig, err := SignatureFromBytes(sig)
	if err != nil {
		return false
	}
	mpcVerif := edsig.Verify(w.Config().PublicPoint(), m)
	return mpcVerif
}

func (w *Wallet) CreateInitialShards() (dscShard, pskShard, recShard []byte, unused [][]byte, err error) {
	ss, e := w.serializedShards()
	if e != nil {
		err = e
		return
	}
	if len(ss) < 6 {
		err = errors.New("not enough shards")
		return
	}

	var ok bool
	// assign dscShard
	dscShard, ok = ss["dsc"]
	if !ok {
		err = errors.New("could not find dsc shard")
		return
	}

	// assign pskShard
	pskShard, ok = ss["psk"]
	if !ok {
		err = errors.New("could not find psk shard")
		return
	}

	// assign recshard
	recShard, ok = ss["recovery"]
	if !ok {
		err = errors.New("could not find recovery shard")
		return
	}

	// sign unused shards using MPC
	// TODO: this is insecure! We're signing but should be encrypting.
	// Can't encrypt with ECDSA. Is that the only encryption supported by mpc?
	for i := 0; i < len(ss); i++ {
		u, ok := ss[fmt.Sprintf("bank%d", i+1)]
		if !ok {
			continue
		}
		sig, e := w.Sign([]byte(u))
		if e != nil {
			err = e
			return
		}
		sSig, e := SerializeSignature(sig)
		if e != nil {
			err = e
			return
		}
		unused = append(unused, sSig)
	}
	if len(unused) == 0 {
		err = errors.New("no backup shards")
	}
	return
}

func (w *Wallet) serializedShards() (map[string][]byte, error) {
	deviceShards := make(map[string][]byte)
	for k, c := range w.Configs {
		b, err := c.MarshalBinary()
		if err != nil {
			return nil, err
		}
		deviceShards[string(k)] = b
	}
	return deviceShards, nil
}

// - The R and S values must be in the valid range for secp256k1 scalars:
//   - Negative values are rejected
//   - Zero is rejected
//   - Values greater than or equal to the secp256k1 group order are rejected
func SignatureFromBytes(sigStr []byte) (*ecdsa.Signature, error) {
	rBytes := sigStr[:33]
	sBytes := sigStr[33:65]

	sig := ecdsa.EmptySignature(curve.Secp256k1{})
	if err := sig.R.UnmarshalBinary(rBytes); err != nil {
		return nil, errors.New("malformed signature: R is not in the range [1, N-1]")
	}

	// S must be in the range [1, N-1].  Notice the check for the maximum number
	// of bytes is required because SetByteSlice truncates as noted in its
	// comment so it could otherwise fail to detect the overflow.
	if err := sig.S.UnmarshalBinary(sBytes); err != nil {
		return nil, errors.New("malformed signature: S is not in the range [1, N-1]")
	}

	// Create and return the signature.
	return &sig, nil
}

// SerializeSignature marshals an ECDSA signature to DER format for use with the CMP protocol
func SerializeSignature(sig *ecdsa.Signature) ([]byte, error) {
	rBytes, err := sig.R.MarshalBinary()
	if err != nil {
		return nil, err
	}

	sBytes, err := sig.S.MarshalBinary()
	if err != nil {
		return nil, err
	}

	sigBytes := make([]byte, 65)
	// 0 pad the byte arrays from the left if they aren't big enough.
	copy(sigBytes[33-len(rBytes):33], rBytes)
	copy(sigBytes[65-len(sBytes):65], sBytes)
	return sigBytes, nil
}
