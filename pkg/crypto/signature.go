package crypto

import (
	"crypto/elliptic"
	"math/big"

	"github.com/taurusgroup/multi-party-sig/pkg/ecdsa"
)

// p256Order returns the curve order for the secp256r1 curve
// NOTE: this is specific to the secp256r1/P256 curve,
// and not taken from the domain params for the key itself
// (which would be a more generic approach for all EC).
var p256Order = elliptic.P256().Params().N

// p256HalfOrder returns half the curve order
// a bit shift of 1 to the right (Rsh) is equivalent
// to division by 2, only faster.
var p256HalfOrder = new(big.Int).Rsh(p256Order, 1)

// IsSNormalized returns true for the integer sigS if sigS falls in
// lower half of the curve order
func IsSNormalized(sigS *big.Int) bool {
	return sigS.Cmp(p256HalfOrder) != 1
}

// NormalizeS will invert the s value if not already in the lower half
// of curve order value
func NormalizeS(sigS *big.Int) *big.Int {

	if IsSNormalized(sigS) {
		return sigS
	}

	return new(big.Int).Sub(p256Order, sigS)
}

// signatureRaw will serialize signature to R || S.
// R, S are padded to 32 bytes respectively.
// code roughly copied from secp256k1_nocgo.go
func signatureRaw(r *big.Int, s *big.Int) []byte {

	rBytes := r.Bytes()
	sBytes := s.Bytes()
	sigBytes := make([]byte, 64)
	// 0 pad the byte arrays from the left if they aren't big enough.
	copy(sigBytes[32-len(rBytes):32], rBytes)
	copy(sigBytes[64-len(sBytes):64], sBytes)
	return sigBytes
}

// ECDSASignatureToBytes converts an ECDSA signature to bytes
func ECDSASignatureToBytes(sig *ecdsa.Signature) []byte {
	// Get normalized scalar values
	normS := NormalizeS(sig.S.Curve().Order().Big())
	r := sig.R.Curve().Order().Big()
	return signatureRaw(r, normS)
}
