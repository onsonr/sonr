package orm

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/onsonr/sonr/internal/orm/didmethod"
)

type ChainCode uint32

func GetChainCode(m didmethod.DIDMethod) ChainCode {
	switch m {
	case didmethod.Bitcoin:
		return 0 // 0
	case didmethod.Ethereum:
		return 64 // 60
	case didmethod.Ibc:
		return 118 // 118
	case didmethod.Sonr:
		return 703 // 703
	default:
		return 0
	}
}

func FormatAddress(pubKey *PublicKey, m didmethod.DIDMethod) (string, error) {
	hexPubKey, err := hex.DecodeString(pubKey.Raw)
	if err != nil {
		return "", err
	}

	// switch m {
	// case didmethod.Bitcoin:
	// 	return bech32.Encode("bc", pubKey.Bytes())
	//
	// case didmethod.Ethereum:
	// 	epk, err := pubKey.ECDSA()
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	return ComputeEthAddress(*epk), nil
	//
	// case didmethod.Sonr:
	// 	return bech32.Encode("idx", hexPubKey)
	//
	// case didmethod.Ibc:
	// 	return bech32.Encode("cosmos", hexPubKey)
	//
	// }
	return string(hexPubKey), nil
}

// ComputeAccountPublicKey computes the public key of a child key given the extended public key, chain code, and index.
func ComputeBip32AccountPublicKey(extPubKey PublicKey, chainCode ChainCode, index int) (*PublicKey, error) {
	// Check if the index is a hardened child key
	if chainCode&0x80000000 != 0 && index < 0 {
		return nil, errors.New("invalid index")
	}

	hexPubKey, err := hex.DecodeString(extPubKey.Raw)
	if err != nil {
		return nil, err
	}

	// Serialize the public key
	pubKey, err := btcec.ParsePubKey(hexPubKey)
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
	_ = btcec.NewPublicKey(lx, ly)
	return &PublicKey{}, nil
}

// newBigIntFieldVal creates a new field value from a big integer.
func newBigIntFieldVal(val *big.Int) *btcec.FieldVal {
	lx := new(btcec.FieldVal)
	lx.SetByteSlice(val.Bytes())
	return lx
}
