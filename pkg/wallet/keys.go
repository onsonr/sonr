package wallet

import (
	"github.com/libp2p/go-libp2p/core/crypto"
	crypto_pb "github.com/libp2p/go-libp2p/core/crypto/pb"

	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
)

type PublicKey struct {
	// contains filtered or unexported fields
	p curve.Point
}

func makePublicKey(p curve.Point) *PublicKey {
	return &PublicKey{p: p}
}

func (p *PublicKey) Equals(k crypto.Key) bool {
	pbz, err := p.Raw()
	if err != nil {
		return false
	}

	kbz, err := k.Raw()
	if err != nil {
		return false
	}
	return string(pbz) == string(kbz)
}

func (p *PublicKey) Raw() ([]byte, error) {
	return p.p.MarshalBinary()
}

func (p *PublicKey) Type() crypto_pb.KeyType {
	return crypto_pb.KeyType_Secp256k1
}

func (w *PublicKey) Verify(data []byte, sig []byte) (bool, error) {
	signature, err := DeserializeSignature(sig)
	if err != nil {
		return false, err
	}
	return signature.Verify(w.p, data), nil
}
