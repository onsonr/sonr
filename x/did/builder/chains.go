package builder

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/binary"
	"errors"
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/onsonr/sonr/x/did/types"
)

// ComputeAccountPublicKey computes the public key of a child key given the extended public key, chain code, and index.
func computeBip32AccountPublicKey(extPubKey PublicKey, chainCode types.ChainCode, index int) (*types.PubKey, error) {
	// Check if the index is a hardened child key
	if chainCode&0x80000000 != 0 && index < 0 {
		return nil, errors.New("invalid index")
	}

	// Serialize the public key
	pubKey, err := btcec.ParsePubKey(extPubKey.GetRaw())
	if err != nil {
		return nil, err
	}
	pubKeyBytes := pubKey.SerializeCompressed()

	// Serialize the index
	indexBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(indexBytes, uint32(index))

	// Compute the HMAC-SHA512
	mac := hmac.New(sha512.New, []byte{byte(chainCode)})
	mac.Write(pubKeyBytes)
	mac.Write(indexBytes)
	I := mac.Sum(nil)

	// Split I into two 32-byte sequences
	IL := I[:32]

	// Convert IL to a big integer
	ilNum := new(big.Int).SetBytes(IL)

	// Check if parse256(IL) >= n
	curve := btcec.S256()
	if ilNum.Cmp(curve.N) >= 0 {
		return nil, errors.New("invalid child key")
	}

	// Compute the child public key
	ilx, ily := curve.ScalarBaseMult(IL)
	childX, childY := curve.Add(ilx, ily, pubKey.X(), pubKey.Y())
	lx := newBigIntFieldVal(childX)
	ly := newBigIntFieldVal(childY)

	// Create the child public key
	childPubKey := btcec.NewPublicKey(lx, ly)
	pk, err := types.NewPublicKey(childPubKey.SerializeCompressed(), types.ChainCodeKeyInfos[chainCode])
	if err != nil {
		return nil, err
	}
	return pk, nil
}

// newBigIntFieldVal creates a new field value from a big integer.
func newBigIntFieldVal(val *big.Int) *btcec.FieldVal {
	lx := new(btcec.FieldVal)
	lx.SetByteSlice(val.Bytes())
	return lx
}
