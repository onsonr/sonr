package curves

import (
	"crypto/elliptic"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func SP256() elliptic.Curve {
	return secp256k1.S256()
}
