package ipfs

import (
	"fmt"

	"github.com/shengdoushi/base58"
	"github.com/sonrhq/core/pkg/node/config"
	"github.com/sonrhq/core/x/identity/types"
	"golang.org/x/crypto/nacl/box"
)

// `boxer` is a struct that contains a peer's public key, a node's public and private keys, and an IPFS
// node.
// @property {[]byte} peerPubKey - The public key of the peer we're communicating with.
// @property nodePubKey - The public key of the node.
// @property nodePrivKey - The private key of the node.
// @property node - This is the IPFS node that we'll be using to communicate with the IPFS network.
type boxer struct {
	peerPubKey  []byte
	nodePubKey  *[32]byte
	nodePrivKey *[32]byte
	node        config.IPFSNode
}

// NewBox creates a new box
func (l *localIpfs) newBoxer(peerPubKey []byte) (*boxer, error) {
	priv, pub, err := l.config.LoadEncKeys()
	if err != nil {
		return nil, err
	}
	return &boxer{
		peerPubKey:  peerPubKey,
		nodePubKey:  pub,
		nodePrivKey: priv,
		node:        l,
	}, nil
}

// Write encrypts a message using the box algorithm
func (b *boxer) Seal(msg []byte) (string, error) {
	// This encrypts msg and appends the result to the nonce.
	encrypted := box.Seal(nil, msg, b.Nonce(), b.nodePubKey, b.nodePrivKey)
	cid, err := b.node.Add(encrypted)
	if err != nil {
		return "", err
	}
	return cid, nil
}

// Read decrypts a message using the box algorithm
func (b *boxer) Open(cid string) ([]byte, error) {
	// Get the encrypted message from IPFS
	msg, err := b.node.Get(cid)
	if err != nil {
		return nil, err
	}
	// The recipient can decrypt the message using their private key and the
	// sender's public key. When you decrypt, you must use the same nonce you
	// used to encrypt the message. One way to achieve this is to store the
	// nonce alongside the encrypted message. Above, we stored the nonce in the
	// first 24 bytes of the encrypted text.
	decrypted, ok := box.Open(nil, msg, b.Nonce(), b.nodePubKey, b.nodePrivKey)
	if !ok {
		return nil, fmt.Errorf("decryption failed")
	}
	return decrypted, nil
}

// Nonce returns the nonce for the box
func (b *boxer) Nonce() *[24]byte {
	var nonce [24]byte
	copy(nonce[:], b.peerPubKey[:24])
	return &nonce
}

// GetService returns the service for the box
func (b *boxer) GetService(addr string, cid string) types.Service {
	return types.Service{
		ID:              types.ConvertAccAddressToDid(addr),
		Type:            types.ServiceType_ServiceType_ENCRYPTED_DATA_VAULT,
		ServiceEndpoint: cid,
	}
}

// GetCapabilityDelegation returns the capability delegation for the box
func (b *boxer) GetCapabilityDelegation(addr string, cid string) types.VerificationMethod {
	var pubKey []byte
	copy(pubKey, b.nodePubKey[:])
	return types.VerificationMethod{
		ID:                 types.ConvertAccAddressToDid(addr),
		Type:               types.KeyType_KeyType_ED25519_VERIFICATION_KEY_2018,
		PublicKeyMultibase: base58.Encode(pubKey, base58.BitcoinAlphabet),
		// sController: types.ConvertAccAddressToDid(b.node),
	}
}
