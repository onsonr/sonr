package types

import (
	"regexp"

	"github.com/cosmos/btcutil/bech32"
	"github.com/ethereum/go-ethereum/crypto"
)

// BitcoinAddress is a type for the BTC address
type BitcoinAddress string

// Bytes returns the bytes representation of the BTC address
func (a BitcoinAddress) Bytes() []byte {
	return []byte(a)
}

// DID returns the DID representation of the BTC address given a method
func (a BitcoinAddress) DID(method string) string {
	return "did:" + method + ":" + a.String()
}

// String returns the string representation of the BTC address
func (a BitcoinAddress) String() string {
	return string(a)
}

// Validate returns an error if the BTC address is invalid
func (a BitcoinAddress) Validate() error {
	re := regexp.MustCompile(`\b(bc(0([ac-hj-np-z02-9]{39}|[ac-hj-np-z02-9]{59})|1[ac-hj-np-z02-9]{8,87})|[13][a-km-zA-HJ-NP-Z1-9]{25,35})\b`)
	if !re.MatchString(string(a)) {
		return ErrInvalidBTCAddressFormat
	}
	return nil
}

// EthereumAddress is a type for the ETH address
type EthereumAddress string

// Bytes returns the bytes representation of the ETH address
func (a EthereumAddress) Bytes() []byte {
	return []byte(a)
}

// DID returns the DID representation of the ETH address given a method
func (a EthereumAddress) DID(method string) string {
	return "did:" + method + ":" + a.String()
}

// String returns the string representation of the ETH address
func (a EthereumAddress) String() string {
	return string(a)
}

// Validate returns an error if the ETH address is invalid
func (a EthereumAddress) Validate() error {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if !re.MatchString(string(a)) {
		return ErrInvalidETHAddressFormat
	}
	return nil
}

// SonrAddress is a type for the IDX address
type SonrAddress string

// Bytes returns the bytes representation of the IDX address
func (a SonrAddress) Bytes() []byte {
	return []byte(a)
}

// DID returns the DID representation of the IDX address given a method
func (a SonrAddress) DID(method string) string {
	return "did:" + method + ":" + a.String()
}

// String returns the string representation of the IDX address
func (a SonrAddress) String() string {
	return string(a)
}

// CreateBitcoinAddress returns the BTC address from the public key using bech32 encoding
func CreateBitcoinAddress(publicKey *PublicKey) (BitcoinAddress, error) {
	addr, err := bech32.Encode("bc", publicKey.Bytes())
	if err != nil {
		return "", err
	}
	return BitcoinAddress(addr), nil
}

// CreateEthereumAddress returns the ETH address from the public key using keccak256
func CreateEthereumAddress(publicKey *PublicKey) (EthereumAddress, error) {
	ecdsaPub, err := crypto.DecompressPubkey(publicKey.Bytes())
	if err != nil {
		return "", err
	}
	addr := crypto.PubkeyToAddress(*ecdsaPub).Hex()
	return EthereumAddress(addr), nil
}

// CreateSonrAddress returns the IDX address from the public key
func CreateSonrAddress(publicKey *PublicKey) (SonrAddress, error) {
	addr, err := bech32.Encode("idx", publicKey.Bytes())
	if err != nil {
		return "", err
	}
	return SonrAddress(addr), nil
}
