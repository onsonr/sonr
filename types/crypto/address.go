package crypto

import (
	"crypto/sha256"
	"encoding/hex"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/shengdoushi/base58"
	"golang.org/x/crypto/ripemd160"
)

func EthereumAddress(pk *PubKey) string {
	hash := ethcrypto.Keccak256(pk.Bytes()[1:])
	addressBytes := hash[len(hash)-20:]
	// Convert the address bytes to a hexadecimal string
	address := hex.EncodeToString(addressBytes)
	return "0x" + address
}

func BitcoinAddress(pk *PubKey) string {
	// Step 1: Compute the SHA256 hash of the public key
	pubKey, err := pk.Btcec()
	if err != nil {
		return ""
	}
	sha256Hash := sha256.New()
	sha256Hash.Write(pubKey.SerializeCompressed())
	sha256Digest := sha256Hash.Sum(nil)

	// Step 2: Compute the RIPEMD160 hash of the SHA256 hash
	ripemd160Hash := ripemd160.New()
	ripemd160Hash.Write(sha256Digest)
	ripemd160Digest := ripemd160Hash.Sum(nil)

	// Step 3: Add version byte (0x00 for Bitcoin mainnet)
	versionedPayload := append([]byte{0x00}, ripemd160Digest...)

	// Step 4: Compute the checksum (first 4 bytes of the double SHA256 hash of the versioned payload)
	firstSHA := sha256.Sum256(versionedPayload)
	secondSHA := sha256.Sum256(firstSHA[:])
	checksum := secondSHA[:4]

	// Step 5: Append the checksum to the versioned payload
	fullPayload := append(versionedPayload, checksum...)

	// Step 6: Encode the result using Base58
	address := base58.Encode(fullPayload, base58.BitcoinAlphabet)
	return address
}
