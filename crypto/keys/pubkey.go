package keys

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

	"github.com/onsonr/sonr/crypto/core/curves"
)

type PubKey interface {
	Bytes() []byte
	Method() string
	Type() string
	Hex() string
	Verify(msg []byte, sig []byte) bool
}

type pubKey struct {
	publicPoint curves.Point
	method      string
}

func (p pubKey) Bytes() []byte {
	return p.publicPoint.ToAffineCompressed()
}

func (p pubKey) Hex() string {
	return hex.EncodeToString(p.publicPoint.ToAffineCompressed())
}

func (p pubKey) Method() string {
	return fmt.Sprintf("did:%s", p.method)
}

func (p pubKey) Type() string {
	return "secp256k1"
}

func (p pubKey) Verify(msgBz []byte, sigBz []byte) bool {
	sig, err := deserializeSignature(sigBz)
	if err != nil {
		return false
	}
	pp, err := getEcdsaPoint(p.Bytes())
	if err != nil {
		return false
	}
	pk := &ecdsa.PublicKey{
		Curve: pp.Curve,
		X:     pp.X,
		Y:     pp.Y,
	}
	return ecdsa.Verify(pk, msgBz, sig.R, sig.S)
}
