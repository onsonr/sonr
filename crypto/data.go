package crypto

import (
	"bytes"
	"crypto/elliptic"
	"encoding/hex"
	"errors"
	"math/big"
	"strconv"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
)

// SignCompact Calculate the signature string according to R and S
func SignCompact(curve *btcec.KoblitzCurve, r, s *big.Int, publicKey btcec.PublicKey, hash []byte, isCompressedKey bool) ([]byte, error) {
	// bitcoind checks the bit length of R and S here. The ecdsa signature
	// algorithm returns R and S mod N therefore they will be the bitsize of
	// the curve, and thus correctly sized.
	for i := 0; i < (1+1)*2; i++ {
		bitSize := curve.BitSize / 8
		result := make([]byte, 1, 2*bitSize+1)
		result[0] = 27 + byte(i)
		if isCompressedKey {
			result[0] += 4
		}
		// Not sure this needs rounding but safer to do so.
		curvelen := (curve.BitSize + 7) / 8

		// Pad R and S to curvelen if needed.
		bytelen := (r.BitLen() + 7) / 8
		if bytelen < curvelen {
			result = append(result,
				make([]byte, curvelen-bytelen)...)
		}
		result = append(result, r.Bytes()...)

		bytelen = (s.BitLen() + 7) / 8
		if bytelen < curvelen {
			result = append(result,
				make([]byte, curvelen-bytelen)...)
		}
		result = append(result, s.Bytes()...)
		// Verify this signature, if not recalculate
		recoverPublicKey, _, err := ecdsa.RecoverCompact(result, hash)
		if err == nil && recoverPublicKey.X().Cmp(publicKey.X()) == 0 && recoverPublicKey.Y().Cmp(publicKey.Y()) == 0 {
			return result, nil
		}
	}
	return nil, errors.New("no valid solution for pubkey found")
}

func NewSignatureData(msgHash []byte, publicKey string, r, s *big.Int) (string, error) {
	// Calculate v, r, and s
	pubBytes, err := hex.DecodeString(publicKey)
	if err != nil {
		return "", err
	}

	pubKey, err := btcec.ParsePubKey(pubBytes)
	if err != nil {
		return "", err
	}
	sig, err := SignCompact(btcec.S256(), r, s, *pubKey, msgHash, false)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(sig), nil

}

// https://tools.ietf.org/html/rfc6979#section-2.3.3
func int2octets(v *big.Int, rolen int) []byte {
	out := v.Bytes()

	// left pad with zeros if it's too short
	if len(out) < rolen {
		out2 := make([]byte, rolen)
		copy(out2[rolen-len(out):], out)
		return out2
	}

	// drop most significant bytes if it's too long
	if len(out) > rolen {
		out2 := make([]byte, rolen)
		copy(out2, out[len(out)-rolen:])
		return out2
	}

	return out
}

// https://tools.ietf.org/html/rfc6979#section-2.3.4
func bits2octets(in []byte, curve elliptic.Curve, rolen int) []byte {
	z1 := hashToInt(in, curve)
	z2 := new(big.Int).Sub(z1, curve.Params().N)
	if z2.Sign() < 0 {
		return int2octets(z1, rolen)
	}
	return int2octets(z2, rolen)
}

// hashToInt converts a hash value to an integer. There is some disagreement
// about how this is done. [NSA] suggests that this is done in the obvious
// manner, but [SECG] truncates the hash to the bit-length of the curve order
// first. We follow [SECG] because that's what OpenSSL does. Additionally,
// OpenSSL right shifts excess bits from the number if the hash is too large
// and we mirror that too.
// This is borrowed from crypto/ecdsa.
func hashToInt(hash []byte, c elliptic.Curve) *big.Int {
	orderBits := c.Params().N.BitLen()
	orderBytes := (orderBits + 7) / 8
	if len(hash) > orderBytes {
		hash = hash[:orderBytes]
	}

	ret := new(big.Int).SetBytes(hash)
	excess := len(hash)*8 - orderBits
	if excess > 0 {
		ret.Rsh(ret, uint(excess))
	}
	return ret
}

func SignToString(sign []byte) string {
	return bytesToHexString(sign)
}

func bytesToHexString(b []byte) string {
	var buf bytes.Buffer
	for _, v := range b {
		t := strconv.FormatInt(int64(v), 16)
		if len(t) > 1 {
			buf.WriteString(t)
		} else {
			buf.WriteString("0" + t)
		}
	}
	return buf.String()
}
