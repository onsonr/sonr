package motor

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/sonr-io/multi-party-sig/pkg/party"
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/crypto"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
)

type MotorNode struct {
	DeviceID    string
	Cosmos      *client.Client
	Wallet      *crypto.MPCWallet
	Address     string
	PubKey      *secp256k1.PubKey
	DID         did.DID
	DIDDocument did.Document

	// Sharding
	deviceShard   []byte
	sharedShard   []byte
	recoveryShard []byte
	unusedShards  [][]byte
}

func newMotor(id string, options ...crypto.WalletOption) (*MotorNode, error) {
	// Create Client instance
	c := client.NewClient(client.ConnEndpointType_BETA)

	// Generate wallet
	w, err := crypto.GenerateWallet(options...)
	if err != nil {
		return nil, err
	}

	// Get address
	bechAddr, err := w.Address()
	if err != nil {
		return nil, err
	}

	// Get public key
	pk, err := w.PublicKeyProto()
	if err != nil {
		return nil, err
	}

	// Set Base DID
	baseDid, err := did.ParseDID(fmt.Sprintf("did:snr:%s", strings.TrimPrefix(bechAddr, "snr")))
	if err != nil {
		return nil, err
	}

	// Create MotorNode
	m := &MotorNode{
		DeviceID: id,
		Cosmos:   c,
		Wallet:   w,
		Address:  bechAddr,
		PubKey:   pk,
		DID:      *baseDid,
	}

	return m, nil
}

func (m *MotorNode) Balance() int64 {
	cs, err := m.Cosmos.CheckBalance(m.Address)
	if err != nil {
		return 0
	}
	if len(cs) <= 0 {
		return 0
	}
	return cs[0].Amount.Int64()
}

// GetVerificationMethod returns the VerificationMethod for the given party.
func (w *MotorNode) GetVerificationMethod(id party.ID) (*did.VerificationMethod, error) {
	vmdid, err := did.ParseDID(fmt.Sprintf("did:snr:%s#%s", strings.TrimPrefix(w.Address, "snr"), id))
	if err != nil {
		return nil, err
	}

	// Get base58 encoded public key.
	pub, err := w.Wallet.PublicKeyBase58()
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
