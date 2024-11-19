package types

import (
	"crypto/ecdsa"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

// ComputeSonrAddr computes the Sonr address from a public key
func ComputeSonrAddr(pk []byte) (string, error) {
	sonrAddr, err := bech32.ConvertAndEncode("idx", pk)
	if err != nil {
		return "", err
	}
	return sonrAddr, nil
}

// ComputeBitcoinAddr computes the Bitcoin address from a public key
func ComputeBitcoinAddr(pk []byte) (string, error) {
	btcAddr, err := bech32.ConvertAndEncode("bc", pk)
	if err != nil {
		return "", err
	}
	return btcAddr, nil
}

// ComputeEthereumAddr computes the Ethereum address from a public key
func ComputeEthereumAddr(pk *ecdsa.PublicKey) string {
	// Generate Ethereum address
	address := ethcrypto.PubkeyToAddress(*pk)

	// Apply ERC-55 checksum encoding
	addr := address.Hex()
	addr = strings.ToLower(addr)
	addr = strings.TrimPrefix(addr, "0x")
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(addr))
	hashBytes := hash.Sum(nil)

	result := "0x"
	for i, c := range addr {
		if c >= '0' && c <= '9' {
			result += string(c)
		} else {
			if hashBytes[i/2]>>(4-i%2*4)&0xf >= 8 {
				result += strings.ToUpper(string(c))
			} else {
				result += string(c)
			}
		}
	}
	return result
}
