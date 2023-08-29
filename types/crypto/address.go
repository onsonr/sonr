package crypto

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"strings"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/shengdoushi/base58"
	"golang.org/x/crypto/ripemd160"
)

// EthereumAddress returns the Ethereum address of the public key.
func EthereumAddress(pk *PubKey) string {
	hash := ethcrypto.Keccak256(pk.Bytes()[1:])
	addressBytes := hash[len(hash)-20:]
	// Convert the address bytes to a hexadecimal string
	address := hex.EncodeToString(addressBytes)
	return "0x" + address
}

// BitcoinAddress returns the Bitcoin address of the public key.
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

// Base64UrlToBytes converts a base64url string to bytes.
func Base64UrlToBytes(base64Url string) ([]byte, error) {
	base64String := strings.ReplaceAll(strings.ReplaceAll(base64Url, "-", "+"), "_", "/")
	missingPadding := len(base64String) % 4
	if missingPadding > 0 {
		base64String += strings.Repeat("=", 4-missingPadding)
	}
	return base64.StdEncoding.DecodeString(base64String)
}

// ParseCredentialPublicKey parses a public key from a base64url string.
func ParseCredentialPublicKey(pubStr string) (interface{}, error) {
	derEncodedPublicKey, err := Base64UrlToBytes(pubStr)
	if err != nil {
		return nil, err
	}

	publicKey, err := x509.ParsePKIXPublicKey(derEncodedPublicKey)
	if err != nil {
		return nil, err
	}
	return publicKey, nil
}
