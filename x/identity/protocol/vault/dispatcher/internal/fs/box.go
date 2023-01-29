package fs

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/shengdoushi/base58"
	"github.com/sonrhq/core/pkg/common"
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
	node        common.IPFSNode
}

// NewBox creates a new box
func (c *VaultConfig) newBoxer() (*boxer, error) {
	_, peerPubKey, err := bech32.DecodeAndConvert(c.address)
	if err != nil {
		return nil, err
	}
	privKey, pubKey, err := loadBoxKeys(c.cctx)
	if err != nil {
		return nil, err
	}
	return &boxer{
		peerPubKey:  peerPubKey,
		nodePubKey:  pubKey,
		nodePrivKey: privKey,
	}, nil
}

// Write encrypts a message using the box algorithm
func (b *boxer) Seal(msg []byte) ([]byte, error) {
	// This encrypts msg and appends the result to the nonce.
	encrypted := box.Seal(nil, msg, b.Nonce(), b.nodePubKey, b.nodePrivKey)
	return encrypted, nil
}

// Read decrypts a message using the box algorithm
func (b *boxer) Open(msg []byte) ([]byte, error) {
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
		Id:              types.ConvertAccAddressToDid(addr),
		Type:            types.ServiceType_ServiceType_ENCRYPTED_DATA_VAULT,
		ServiceEndpoint: cid,
	}
}

// GetCapabilityDelegation returns the capability delegation for the box
func (b *boxer) GetCapabilityDelegation(addr string, cid string) types.VerificationMethod {
	var pubKey []byte
	copy(pubKey, b.nodePubKey[:])
	return types.VerificationMethod{
		Id:                 types.ConvertAccAddressToDid(addr),
		Type:               types.KeyType_KeyType_ED25519_VERIFICATION_KEY_2018,
		PublicKeyMultibase: base58.Encode(pubKey, base58.BitcoinAlphabet),
		// sController: types.ConvertAccAddressToDid(b.node),
	}
}
